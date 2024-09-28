package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/website/internal/entities"
	"github.com/mcorrigan89/website/internal/repositories"
)

type WebsitePageService struct {
	utils                 ServicesUtils
	websiteRepository     *repositories.WebsiteRepository
	websitePageRepository *repositories.WebsitePageRepository
}

func NewWebsitePageService(utils ServicesUtils, repos *repositories.Repositories) WebsitePageService {
	return WebsitePageService{
		utils:                 utils,
		websiteRepository:     repos.WebsiteRepository,
		websitePageRepository: repos.WebsitePageRepository,
	}
}

type WebsitePageByIDArgs struct {
	ID     uuid.UUID
	Locale *string
}

func (service *WebsitePageService) WebsitePageByID(ctx context.Context, args WebsitePageByIDArgs) (*entities.WebsitePageEntity, error) {
	service.utils.logger.Info().Ctx(ctx).Interface("args", args).Msg("Getting website page by ID")
	website, err := service.websitePageRepository.WebsitePageByID(ctx, repositories.WebsitePageByIDArgs{
		ID:     args.ID,
		Locale: args.Locale,
	})
	if err != nil {
		service.utils.logger.Err(err).Ctx(ctx).Msg("Error getting schedule item by ID")
		return nil, err
	}

	return website, nil

}

type CreateWebsitePageArgs struct {
	WebsiteID  uuid.UUID
	UrlSlug    string
	Title      *string
	Subtitle   *string
	IsPrivate  *bool
	IsDisabled *bool
}

func (service *WebsitePageService) CreateWebsitePage(ctx context.Context, args CreateWebsitePageArgs) (*entities.WebsitePageEntity, error) {
	service.utils.logger.Info().Ctx(ctx).Interface("args", args).Msg("Creating website page")
	website, err := service.websitePageRepository.CreateWebsitePage(ctx, repositories.CreateWebsitePageArgs{
		WebsiteID: args.WebsiteID,
		UrlSlug:   args.UrlSlug,
		Title:     args.Title,
		Subtitle:  args.Subtitle,
	})
	if err != nil {
		service.utils.logger.Err(err).Ctx(ctx).Msg("Error getting schedule item by ID")
		return nil, err
	}

	return website, nil
}

type UpdateWebsitePageArgs struct {
	ID         uuid.UUID
	UrlSlug    *string
	Title      *string
	Subtitle   *string
	IsPrivate  *bool
	IsDisabled *bool
}

func (service *WebsitePageService) UpdateWebsitePage(ctx context.Context, args UpdateWebsitePageArgs) (*entities.WebsitePageEntity, error) {
	service.utils.logger.Info().Ctx(ctx).Interface("args", args).Msg("Updating website page")
	website, err := service.websitePageRepository.UpdateWebsitePage(ctx, repositories.UpdateWebsitePageArgs{
		ID:       args.ID,
		UrlSlug:  args.UrlSlug,
		Title:    args.Title,
		Subtitle: args.Subtitle,
	})
	if err != nil {
		service.utils.logger.Err(err).Ctx(ctx).Msg("Error getting schedule item by ID")
		return nil, err
	}

	return website, nil
}
