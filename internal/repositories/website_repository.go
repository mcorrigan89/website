package repositories

import (
	"context"
	"sort"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/website/internal/entities"
	"github.com/mcorrigan89/website/internal/repositories/models"
)

type WebsiteRepository struct {
	utils   ServicesUtils
	DB      *pgxpool.Pool
	queries *models.Queries
}

func NewWebsiteRepository(utils ServicesUtils, db *pgxpool.Pool, queries *models.Queries) *WebsiteRepository {
	return &WebsiteRepository{
		utils:   utils,
		DB:      db,
		queries: queries,
	}
}

func (repo *WebsiteRepository) GetWebsiteByHandle(ctx context.Context, handle string, locale *string) (*entities.WebsiteEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	tx, err := repo.DB.Begin(ctx)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with transaction creating person")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := repo.queries.WithTx(tx)

	row, err := qtx.GetWebsiteByHandle(ctx, handle)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		} else {
			repo.utils.logger.Err(err).Ctx(ctx).Str("handle", handle).Msg("Error getting website by handle")
			return nil, err
		}
	}

	websiteEntity, err := repo.gatherWebsiteData(ctx, qtx, &row.Website, locale)
	if err != nil {
		return nil, err
	}

	return websiteEntity, nil
}

func (repo *WebsiteRepository) gatherWebsiteData(ctx context.Context, qtx *models.Queries, website *models.Website, locale *string) (*entities.WebsiteEntity, error) {

	if locale == nil {
		locale = &website.DefaultLocale
	}

	websiteContentRow, err := qtx.GetWebsiteContentByWebsiteID(ctx, models.GetWebsiteContentByWebsiteIDParams{
		WebsiteID: website.ID,
		Locale:    *locale,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error getting website content by website id")
		return nil, err
	}

	pageRows, err := qtx.GetWebsitePagesByWebsiteID(ctx, models.GetWebsitePagesByWebsiteIDParams{
		WebsiteID: website.ID,
		Locale:    *locale,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error getting website pages by website id")
		return nil, err
	}

	textComponentRows, err := qtx.GetTextComponentsByWebsiteID(ctx, models.GetTextComponentsByWebsiteIDParams{
		WebsiteID: website.ID,
		Locale:    *locale,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error getting text components by website id")
		return nil, err
	}

	imageComponentRows, err := qtx.GetImageComponentsByWebsiteID(ctx, website.ID)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error getting image components by website id")
		return nil, err
	}

	websiteEntity := repo.modelToEntity(website, websiteContentRow, pageRows, componentArgs{
		textComponents:  textComponentRows,
		imageComponents: imageComponentRows,
	})

	return &websiteEntity, nil

}

type componentArgs struct {
	textComponents  []models.GetTextComponentsByWebsiteIDRow
	imageComponents []models.GetImageComponentsByWebsiteIDRow
}

func (repo *WebsiteRepository) modelToEntity(row *models.Website, content models.GetWebsiteContentByWebsiteIDRow, pages []models.GetWebsitePagesByWebsiteIDRow, components componentArgs) entities.WebsiteEntity {
	pageEntities := []*entities.WebsitePageEntity{}

	for _, pageRow := range pages {

		componentEntities := []*entities.WebsiteComponentEntity{}

		for _, component := range components.textComponents {
			if component.WebsiteComponent.WebsitePageID == pageRow.WebsitePage.ID {
				componentEntities = append(componentEntities, &entities.WebsiteComponentEntity{
					ID:        component.WebsiteComponent.ID,
					SortKey:   component.WebsiteComponent.SortKey,
					PageID:    component.WebsiteComponent.WebsitePageID,
					WebsiteID: component.WebsiteComponent.WebsiteID,
					TextComponent: &entities.WebsiteTextComponentEntity{
						ID:   component.SimpleTextComponent.ID,
						Text: component.SimpleTextComponent.Content,
					},
				})
			}
		}

		for _, component := range components.imageComponents {
			if component.WebsiteComponent.WebsitePageID == pageRow.WebsitePage.ID {
				componentEntities = append(componentEntities, &entities.WebsiteComponentEntity{
					ID:        component.WebsiteComponent.ID,
					SortKey:   component.WebsiteComponent.SortKey,
					PageID:    component.WebsiteComponent.WebsitePageID,
					WebsiteID: component.WebsiteComponent.WebsiteID,
					ImageComponent: &entities.WebsiteImageComponentEntity{
						ID:       component.ImageComponent.ID,
						PhotoURL: component.ImageComponent.ImageID.String(),
					},
				})
			}
		}

		sort.Slice(componentEntities, func(i, j int) bool {
			return componentEntities[i].SortKey < componentEntities[j].SortKey
		})

		pageEntities = append(pageEntities, &entities.WebsitePageEntity{
			ID:         pageRow.WebsitePage.ID,
			WebsiteID:  pageRow.WebsitePage.WebsiteID,
			SortKey:    pageRow.WebsitePage.SortKey,
			UrlSlug:    pageRow.WebsitePage.UrlSlug,
			Title:      pageRow.WebsitePageContent.Title,
			Subtitle:   pageRow.WebsitePageContent.Subtitle,
			Components: componentEntities,
		})
	}

	return entities.WebsiteEntity{
		ID:                 row.ID,
		Handle:             row.Handle,
		DisplayName:        content.WebsiteContent.WebsiteDisplayName,
		DisplayDescription: content.WebsiteContent.WebsiteDisplayDescription,
		Pages:              pageEntities,
	}
}
