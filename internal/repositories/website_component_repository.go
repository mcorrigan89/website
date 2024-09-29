package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/website/internal/entities"
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

type CreateTextComponentArgs struct {
	WebsiteSectionID uuid.UUID
	Locale           *string
	Json             []byte
	Html             *string
}

func (repo *WebsiteComponentRepository) CreateTextComponent(ctx context.Context, args CreateTextComponentArgs) (*entities.WebsiteComponentEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	tx, err := repo.DB.Begin(ctx)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with transaction updating simple text component")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := repo.queries.WithTx(tx)

	websiteRow, err := qtx.GetWebsiteBySectionID(ctx, args.WebsiteSectionID)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with GetWebsiteBySectionID")
		return nil, err
	}

	var locale string
	if args.Locale == nil {
		locale = websiteRow.Website.DefaultLocale
	} else {
		locale = *args.Locale
	}

	createdComponent, err := qtx.CreateWebsiteComponent(ctx, models.CreateWebsiteComponentParams{
		WebsiteID:        websiteRow.Website.ID,
		WebsiteSectionID: args.WebsiteSectionID,
	})
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with CreateWebsiteComponent")
		return nil, err
	}

	createdTextRow, err := qtx.UpsertWebsiteSimpleTextComponent(ctx, models.UpsertWebsiteSimpleTextComponentParams{
		Locale:             locale,
		WebsiteComponentID: createdComponent.ID,
		ContentJson:        args.Json,
		ContentHtml:        args.Html,
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
		WebsiteID: createdComponent.WebsiteID,
		SectionID: createdComponent.WebsiteSectionID,
		TextComponent: &entities.WebsiteTextComponentEntity{
			ID:          createdTextRow.ID,
			ComponentID: createdComponent.ID,
			Json:        createdTextRow.ContentJson,
			Html:        createdTextRow.ContentHtml,
		},
	}, nil

}

type UpdateTextComponent struct {
	ID     uuid.UUID
	Locale *string
	Json   []byte
	Html   *string
}

func (repo *WebsiteComponentRepository) UpdateTextComponent(ctx context.Context, args UpdateTextComponent) (*entities.WebsiteComponentEntity, error) {
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

	_, err = qtx.UpdateWebsiteTextComponent(ctx, models.UpdateWebsiteTextComponentParams{
		Locale:             locale,
		ContentJson:        args.Json,
		ContentHtml:        args.Html,
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
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with GetWebsiteTextComponent")
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with transaction commit")
		return nil, err
	}

	return &entities.WebsiteComponentEntity{
		ID:        row.WebsiteComponent.ID,
		SectionID: row.WebsiteComponent.WebsiteSectionID,
		WebsiteID: row.WebsiteComponent.WebsiteID,
		TextComponent: &entities.WebsiteTextComponentEntity{
			ID:          row.TextComponent.ID,
			ComponentID: row.TextComponent.WebsiteComponentID,
			Json:        row.TextComponent.ContentJson,
			Html:        row.TextComponent.ContentHtml,
		},
	}, nil
}
