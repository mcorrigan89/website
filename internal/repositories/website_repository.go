package repositories

import (
	"context"
	"sort"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/website/internal/entities"
	"github.com/mcorrigan89/website/internal/helpers"
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

	row, err := repo.queries.GetWebsiteByHandle(ctx, handle)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		} else {
			repo.utils.logger.Err(err).Ctx(ctx).Str("handle", handle).Msg("Error getting website by handle")
			return nil, err
		}
	}

	websiteEntity, err := repo.gatherWebsiteData(ctx, repo.queries, websiteData{
		website: row.Website,
		config:  row.WebsiteConfig,
		styles:  row.WebsiteStyle,
		palette: row.Palette,
	}, locale)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error gathering website data")
		return nil, err
	}

	return websiteEntity, nil
}

type CreateWebsiteArgs struct {
	Handle string
	Locale string
}

func (repo *WebsiteRepository) CreateWebsite(ctx context.Context, args CreateWebsiteArgs) (*entities.WebsiteEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	tx, err := repo.DB.Begin(ctx)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with transaction creating person")
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := repo.queries.WithTx(tx)

	createdWebsite, err := qtx.CreateWebsite(ctx, models.CreateWebsiteParams{
		Handle:        args.Handle,
		DefaultLocale: args.Locale,
	})
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error creating website")
		return nil, err
	}

	_, err = qtx.CreateWebsiteContent(ctx, models.CreateWebsiteContentParams{
		WebsiteID:          createdWebsite.ID,
		Locale:             args.Locale,
		WebsiteDisplayName: "Website Title",
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error creating website content")
		return nil, err
	}

	sortKey, err := helpers.KeyBetween("", "")
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error generating sort key")
		return nil, err
	}

	defaultHomePage, err := qtx.CreateWebsitePage(ctx, models.CreateWebsitePageParams{
		WebsiteID: createdWebsite.ID,
		SortKey:   sortKey,
		UrlSlug:   "home",
	})
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error creating default home page")
		return nil, err
	}

	defaultPageTitle := "Home"

	_, err = qtx.UpsertWebsitePageContent(ctx, models.UpsertWebsitePageContentParams{
		WebsitePageID: defaultHomePage.ID,
		Locale:        args.Locale,
		Title:         &defaultPageTitle,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error creating default home page content")
		return nil, err
	}

	createdConfig, err := qtx.CreateWebsiteConfig(ctx, models.CreateWebsiteConfigParams{
		WebsiteID:     createdWebsite.ID,
		DefaultPageID: defaultHomePage.ID,
	})
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error creating website config")
		return nil, err
	}

	createWebsiteStyles, err := qtx.CreateWebsiteStyles(ctx, createdWebsite.ID)

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error creating website styles")
		return nil, err
	}

	createdPalette, err := qtx.CreatePalette(ctx, models.CreatePaletteParams{
		WebsiteStylesID: createWebsiteStyles.ID,
		ColorOne:        "#000000",
		ColorTwo:        "#000000",
		ColorThree:      "#000000",
		ColorFour:       "#000000",
		ColorFive:       "#000000",
		ColorSix:        "#000000",
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error creating palette")
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error with transaction commit")
		return nil, err
	}

	websiteEntity, err := repo.gatherWebsiteData(ctx, repo.queries, websiteData{
		website: createdWebsite,
		config:  createdConfig,
		styles:  createWebsiteStyles,
		palette: createdPalette,
	}, &args.Locale)

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error gathering website data")
		return nil, err
	}

	return websiteEntity, nil
}

type websiteData struct {
	website models.Website
	config  models.WebsiteConfig
	styles  models.WebsiteStyle
	palette models.Palette
}

func (repo *WebsiteRepository) gatherWebsiteData(ctx context.Context, queries *models.Queries, websiteData websiteData, locale *string) (*entities.WebsiteEntity, error) {

	if locale == nil {
		locale = &websiteData.website.DefaultLocale
	}

	websiteContentRow, err := queries.GetWebsiteContentByWebsiteID(ctx, models.GetWebsiteContentByWebsiteIDParams{
		WebsiteID: websiteData.website.ID,
		Locale:    *locale,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error getting website content by website id")
		return nil, err
	}

	pageRows, err := queries.GetWebsitePagesByWebsiteID(ctx, models.GetWebsitePagesByWebsiteIDParams{
		WebsiteID: websiteData.website.ID,
		Locale:    *locale,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error getting website pages by website id")
		return nil, err
	}

	sectionRows, err := queries.GetWebsiteSectionsByWebsiteID(ctx, websiteData.website.ID)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error getting website sections by page id")
		return nil, err
	}

	textComponentRows, err := queries.GetTextComponentsByWebsiteID(ctx, models.GetTextComponentsByWebsiteIDParams{
		WebsiteID: websiteData.website.ID,
		Locale:    *locale,
	})

	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error getting text components by website id")
		return nil, err
	}

	imageComponentRows, err := queries.GetImageComponentsByWebsiteID(ctx, websiteData.website.ID)
	if err != nil {
		repo.utils.logger.Err(err).Ctx(ctx).Msg("Error getting image components by website id")
		return nil, err
	}

	websiteEntity := repo.modelToEntity(websiteArgs{
		website:        websiteData.website,
		websiteContent: websiteContentRow.WebsiteContent,
		websiteConfig:  websiteData.config,
		websiteStyle:   websiteData.styles,
		palette:        websiteData.palette,
	}, pageRows, componentArgs{
		sections:        sectionRows,
		textComponents:  textComponentRows,
		imageComponents: imageComponentRows,
	})

	return &websiteEntity, nil

}

type websiteArgs struct {
	website        models.Website
	websiteContent models.WebsiteContent
	websiteStyle   models.WebsiteStyle
	websiteConfig  models.WebsiteConfig
	palette        models.Palette
}
type componentArgs struct {
	sections        []models.GetWebsiteSectionsByWebsiteIDRow
	textComponents  []models.GetTextComponentsByWebsiteIDRow
	imageComponents []models.GetImageComponentsByWebsiteIDRow
}

func (repo *WebsiteRepository) modelToEntity(websiteData websiteArgs, pages []models.GetWebsitePagesByWebsiteIDRow, components componentArgs) entities.WebsiteEntity {
	pageEntities := []*entities.WebsitePageEntity{}

	for _, pageRow := range pages {

		sectionEntities := []*entities.WebsitePageSectionEntity{}

		for _, section := range components.sections {
			if section.WebsiteSection.WebsitePageID == pageRow.WebsitePage.ID {

				componentEntities := []*entities.WebsiteComponentEntity{}

				for _, component := range components.textComponents {
					if component.WebsiteComponent.WebsiteSectionID == section.WebsiteSection.ID {

						componentDisplay := repo.createComponentDisplay(component.WebsiteComponentDisplay)

						componentEntities = append(componentEntities, &entities.WebsiteComponentEntity{
							ID:        component.WebsiteComponent.ID,
							SectionID: component.WebsiteComponent.WebsiteSectionID,
							WebsiteID: component.WebsiteComponent.WebsiteID,
							Display:   componentDisplay,
							TextComponent: &entities.WebsiteTextComponentEntity{
								ID:   component.TextComponent.ID,
								Json: component.TextComponent.ContentJson,
								Html: component.TextComponent.ContentHtml,
							},
						})
					}
				}

				for _, component := range components.imageComponents {
					if component.WebsiteComponent.WebsiteSectionID == section.WebsiteSection.ID {

						componentDisplay := repo.createComponentDisplay(component.WebsiteComponentDisplay)

						componentEntities = append(componentEntities, &entities.WebsiteComponentEntity{
							ID:        component.WebsiteComponent.ID,
							SectionID: component.WebsiteComponent.WebsiteSectionID,
							WebsiteID: component.WebsiteComponent.WebsiteID,
							Display:   componentDisplay,
							ImageComponent: &entities.WebsiteImageComponentEntity{
								ID:       component.ImageComponent.ID,
								PhotoURL: component.ImageComponent.ImageID.String(),
							},
						})
					}
				}

				sectionEntity := &entities.WebsitePageSectionEntity{
					ID:         section.WebsiteSection.ID,
					WebsiteID:  section.WebsiteSection.WebsiteID,
					PageID:     section.WebsiteSection.WebsitePageID,
					SortKey:    section.WebsiteSection.SortKey,
					RowCount:   section.WebsiteSectionDisplay.RowCount,
					Components: componentEntities,
				}

				sectionEntities = append(sectionEntities, sectionEntity)
			}
		}

		sort.Slice(sectionEntities, func(i, j int) bool {
			return sectionEntities[i].SortKey < sectionEntities[j].SortKey
		})

		pageEntities = append(pageEntities, &entities.WebsitePageEntity{
			ID:        pageRow.WebsitePage.ID,
			WebsiteID: pageRow.WebsitePage.WebsiteID,
			SortKey:   pageRow.WebsitePage.SortKey,
			UrlSlug:   pageRow.WebsitePage.UrlSlug,
			Title:     pageRow.WebsitePageContent.Title,
			Subtitle:  pageRow.WebsitePageContent.Subtitle,
			Sections:  sectionEntities,
		})
	}

	return entities.WebsiteEntity{
		ID:                 websiteData.website.ID,
		Handle:             websiteData.website.Handle,
		DisplayName:        websiteData.websiteContent.WebsiteDisplayName,
		DisplayDescription: websiteData.websiteContent.WebsiteDisplayDescription,
		Config: &entities.WebsiteConfigEntity{
			ID:            websiteData.websiteConfig.ID,
			WebsiteID:     websiteData.websiteConfig.WebsiteID,
			DefaultPageID: websiteData.websiteConfig.DefaultPageID,
		},
		Styles: &entities.WebsiteStylesEntity{
			ID:        websiteData.websiteStyle.ID,
			WebsiteID: websiteData.websiteStyle.WebsiteID,
			Palette: &entities.PaletteEntity{
				ID:        websiteData.palette.ID,
				WebsiteID: websiteData.website.ID,
				Color1:    websiteData.palette.ColorOne,
				Color2:    websiteData.palette.ColorTwo,
				Color3:    websiteData.palette.ColorThree,
				Color4:    websiteData.palette.ColorFour,
				Color5:    websiteData.palette.ColorFive,
				Color6:    websiteData.palette.ColorSix,
			},
		},

		Pages: pageEntities,
	}
}

func (s *WebsiteRepository) createComponentDisplay(componentDisplayModel models.WebsiteComponentDisplay) *entities.WebsiteComponentDisplayEntity {
	componentDisplay := &entities.WebsiteComponentDisplayEntity{
		ID:          componentDisplayModel.ID,
		ComponentID: componentDisplayModel.WebsiteComponentID,
		MobilePositioning: &entities.WebsiteComponentPositioningEntity{
			Xcoord: componentDisplayModel.MobileXCoordinate,
			Ycoord: componentDisplayModel.MobileYCoordinate,
			Height: componentDisplayModel.MobileHeight,
			Width:  componentDisplayModel.MobileWidth,
		},
		FullScreenPositioning: &entities.WebsiteComponentPositioningEntity{
			Xcoord: &componentDisplayModel.XCoordinate,
			Ycoord: &componentDisplayModel.YCoordinate,
			Height: &componentDisplayModel.Height,
			Width:  &componentDisplayModel.Width,
		},
	}

	return componentDisplay
}
