package repositories

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/website/internal/entities"
	"github.com/mcorrigan89/website/internal/helpers"
	"github.com/mcorrigan89/website/internal/repositories/models"
)

type WebsitePageRepository struct {
	utils   ServicesUtils
	DB      *pgxpool.Pool
	queries *models.Queries
}

func NewWebsitePageRepository(utils ServicesUtils, db *pgxpool.Pool, queries *models.Queries) *WebsitePageRepository {
	return &WebsitePageRepository{
		utils:   utils,
		DB:      db,
		queries: queries,
	}
}

type WebsitePageByIDArgs struct {
	ID     uuid.UUID
	Locale *string
}

func (repo *WebsitePageRepository) WebsitePageByID(ctx context.Context, args WebsitePageByIDArgs) (*entities.WebsitePageEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	tx, err := repo.DB.Begin(ctx)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with transaction creating person")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := repo.queries.WithTx(tx)

	var locale string
	if args.Locale == nil {
		websiteRow, err := qtx.GetWebsiteByPageID(ctx, args.ID)
		if err != nil {
			repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with GetWebsiteByPageID")
			return nil, err
		}

		locale = websiteRow.Website.DefaultLocale
	} else {
		locale = *args.Locale
	}

	row, err := qtx.GetWebsitePageByID(ctx, models.GetWebsitePageByIDParams{
		ID:     args.ID,
		Locale: locale,
	})
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with GetWebsiteByPageID")
		return nil, err
	}

	return &entities.WebsitePageEntity{
		ID:        row.WebsitePage.ID,
		WebsiteID: row.WebsitePage.WebsiteID,
		Title:     row.WebsitePageContent.Title,
		Subtitle:  row.WebsitePageContent.Subtitle,
		UrlSlug:   row.WebsitePage.UrlSlug,
		SortKey:   row.WebsitePage.SortKey,
	}, nil

}

type CreateWebsitePageArgs struct {
	WebsiteID uuid.UUID
	Locale    *string
	UrlSlug   string
	Title     string
	Subtitle  *string
}

func (repo *WebsitePageRepository) CreateWebsitePage(ctx context.Context, args CreateWebsitePageArgs) (*entities.WebsitePageEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	tx, err := repo.DB.Begin(ctx)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with transaction creating page")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := repo.queries.WithTx(tx)

	var locale string
	if args.Locale == nil {
		websiteRow, err := qtx.GetWebsiteByID(ctx, args.WebsiteID)
		if err != nil {
			repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with GetWebsiteByPageID")
			return nil, err
		}

		locale = websiteRow.Website.DefaultLocale
	} else {
		locale = *args.Locale
	}

	previousPages, err := qtx.GetWebsitePagesByWebsiteID(ctx, models.GetWebsitePagesByWebsiteIDParams{
		WebsiteID: args.WebsiteID,
		Locale:    locale,
	})
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with GetWebsitePagesByWebsiteID")
		return nil, err
	}

	for _, page := range previousPages {
		if page.WebsitePage.UrlSlug == args.UrlSlug {
			slugExistsError := errors.New("UrlSlug already exists")
			repo.utils.logger.Err(slugExistsError).Ctx(ctx).Msg("UrlSlug already exists")
			return nil, slugExistsError
		}
	}

	lastSortKey := ""
	if len(previousPages) > 0 {
		lastSortKey = previousPages[len(previousPages)-1].WebsitePage.SortKey
	}

	pageSortKey, err := helpers.KeyBetween(lastSortKey, "")
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with KeyBetween")
		return nil, err
	}

	row, err := qtx.CreateWebsitePage(ctx, models.CreateWebsitePageParams{
		WebsiteID: args.WebsiteID,
		UrlSlug:   args.UrlSlug,
		SortKey:   pageSortKey,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with UpdateWebsitePage")
		return nil, err
	}

	contentRow, err := qtx.UpsertWebsitePageContent(ctx, models.UpsertWebsitePageContentParams{
		ID:            uuid.New(),
		WebsitePageID: row.ID,
		Locale:        locale,
		Title:         &args.Title,
		Subtitle:      args.Subtitle,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with UpsertWebsitePageContent")
		return nil, err
	}

	err = tx.Commit(ctx)

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error committing transaction")
		return nil, err
	}

	return &entities.WebsitePageEntity{
		ID:        row.ID,
		WebsiteID: row.WebsiteID,
		Title:     contentRow.Title,
		Subtitle:  contentRow.Subtitle,
		UrlSlug:   row.UrlSlug,
		SortKey:   row.SortKey,
	}, nil

}

type UpdateWebsitePageArgs struct {
	ID       uuid.UUID
	Locale   *string
	UrlSlug  *string
	Title    *string
	Subtitle *string
}

func (repo *WebsitePageRepository) UpdateWebsitePage(ctx context.Context, args UpdateWebsitePageArgs) (*entities.WebsitePageEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	tx, err := repo.DB.Begin(ctx)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with transaction updating page")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := repo.queries.WithTx(tx)

	var locale string
	if args.Locale == nil {
		websiteRow, err := qtx.GetWebsiteByPageID(ctx, args.ID)
		if err != nil {
			repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with GetWebsiteByPageID")
			return nil, err
		}

		locale = websiteRow.Website.DefaultLocale
	} else {
		locale = *args.Locale
	}

	row, err := qtx.UpdateWebsitePage(ctx, models.UpdateWebsitePageParams{
		ID:      args.ID,
		UrlSlug: args.UrlSlug,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with UpdateWebsitePage")
		return nil, err
	}

	contentRow, err := qtx.UpsertWebsitePageContent(ctx, models.UpsertWebsitePageContentParams{
		WebsitePageID: row.ID,
		Locale:        locale,
		Title:         args.Title,
		Subtitle:      args.Subtitle,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with UpsertWebsitePageContent")
		return nil, err
	}

	err = tx.Commit(ctx)

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error committing transaction")
		return nil, err
	}

	return &entities.WebsitePageEntity{
		ID:        row.ID,
		WebsiteID: row.WebsiteID,
		Title:     contentRow.Title,
		Subtitle:  contentRow.Subtitle,
		UrlSlug:   row.UrlSlug,
		SortKey:   row.SortKey,
	}, nil

}
