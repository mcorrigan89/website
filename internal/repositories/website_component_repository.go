package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/website/internal/entities"
	"github.com/mcorrigan89/website/internal/helpers"
	"github.com/mcorrigan89/website/internal/repositories/models"
)

type WebsiteComponentRepository struct {
	utils   ServicesUtils
	DB      *pgxpool.Pool
	queries *models.Queries
}

func NewWebsiteComponentRepository(utils ServicesUtils, db *pgxpool.Pool, queries *models.Queries) *WebsiteComponentRepository {
	return &WebsiteComponentRepository{
		utils:   utils,
		DB:      db,
		queries: queries,
	}
}

type CreateSimpleTextComponentArgs struct {
	WebsitePageID uuid.UUID
	Locale        *string
	Content       *string
}

func (repo *WebsiteComponentRepository) CreateSimpleTextComponent(ctx context.Context, args CreateSimpleTextComponentArgs) (*entities.WebsiteComponentEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	tx, err := repo.DB.Begin(ctx)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with transaction updating simple text component")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := repo.queries.WithTx(tx)

	websiteRow, err := qtx.GetWebsiteByPageID(ctx, args.WebsitePageID)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with GetWebsiteByPageID")
		return nil, err
	}

	var locale string
	if args.Locale == nil {
		locale = websiteRow.Website.DefaultLocale
	} else {
		locale = *args.Locale
	}

	websitePageComponents, err := qtx.GetWebsiteComponentsByWebsitePageID(ctx, args.WebsitePageID)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with GetWebsiteComponentsByWebsitePageID")
		return nil, err
	}

	mostRecentSortKey := ""
	if len(websitePageComponents) > 0 {
		mostRecentSortKey = websitePageComponents[len(websitePageComponents)-1].WebsiteComponent.SortKey
	}

	sortKeyToAdd, err := helpers.KeyBetween(mostRecentSortKey, "")
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with KeyBetween")
		return nil, err
	}

	createdComponent, err := qtx.CreateWebsiteComponent(ctx, models.CreateWebsiteComponentParams{
		WebsiteID:     websiteRow.Website.ID,
		WebsitePageID: args.WebsitePageID,
		SortKey:       sortKeyToAdd,
	})
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with CreateWebsiteComponent")
		return nil, err
	}

	fmt.Println(createdComponent.ID)

	createdTextRow, err := qtx.UpsertWebsiteSimpleTextComponent(ctx, models.UpsertWebsiteSimpleTextComponentParams{
		Locale:             locale,
		WebsiteComponentID: createdComponent.ID,
		Content:            args.Content,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with UpsertWebsiteSimpleTextComponent")
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with transaction commit")
		return nil, err
	}

	return &entities.WebsiteComponentEntity{
		ID:        createdComponent.ID,
		SortKey:   createdComponent.SortKey,
		WebsiteID: createdComponent.WebsiteID,
		TextComponent: &entities.WebsiteTextComponentEntity{
			ID:   createdTextRow.ID,
			Text: createdTextRow.Content,
		},
	}, nil

}

type UpdateSimpleTextComponent struct {
	ID      uuid.UUID
	Locale  *string
	Content *string
}

func (repo *WebsiteComponentRepository) UpdateSimpleTextComponent(ctx context.Context, args UpdateSimpleTextComponent) (*entities.WebsiteComponentEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	tx, err := repo.DB.Begin(ctx)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with transaction updating simple text component")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := repo.queries.WithTx(tx)

	var locale string
	if args.Locale == nil {
		websiteRow, err := qtx.GetWebsiteByComponentID(ctx, args.ID)
		if err != nil {
			repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with GetWebsiteByPageID")
			return nil, err
		}

		locale = websiteRow.Website.DefaultLocale
	} else {
		locale = *args.Locale
	}

	updatedRow, err := qtx.UpdateWebsiteTextComponent(ctx, models.UpdateWebsiteTextComponentParams{
		Locale:             locale,
		Content:            args.Content,
		WebsiteComponentID: args.ID,
	})
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with UpdateWebsiteTextComponent")
		return nil, err
	}

	row, err := qtx.GetWebsiteTextComponent(ctx, models.GetWebsiteTextComponentParams{
		ID:     args.ID,
		Locale: locale,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with GetTextComponentsByWebsiteID")
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with transaction commit")
		return nil, err
	}

	return &entities.WebsiteComponentEntity{
		ID:        row.WebsiteComponent.ID,
		SortKey:   row.WebsiteComponent.SortKey,
		WebsiteID: row.WebsiteComponent.WebsiteID,
		TextComponent: &entities.WebsiteTextComponentEntity{
			Text: updatedRow.Content,
		},
	}, nil
}
