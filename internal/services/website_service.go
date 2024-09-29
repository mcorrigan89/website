package services

import (
	"context"

	"github.com/mcorrigan89/website/internal/entities"
	"github.com/mcorrigan89/website/internal/repositories"
)

type WebsiteService struct {
	utils             ServicesUtils
	websiteRepository *repositories.WebsiteRepository
}

func NewWebsiteService(utils ServicesUtils, repos *repositories.Repositories) WebsiteService {
	return WebsiteService{
		utils:             utils,
		websiteRepository: repos.WebsiteRepository,
	}
}

func (service *WebsiteService) GetWebsiteByHandle(ctx context.Context, handle string, locale *string) (*entities.WebsiteEntity, error) {
	service.utils.logger.Info().Ctx(ctx).Str("handle", handle).Msg("Getting website by handle")
	website, err := service.websiteRepository.GetWebsiteByHandle(ctx, handle, locale)
	if err != nil {
		service.utils.logger.Err(err).Ctx(ctx).Msg("Error getting website by handle")
		return nil, err
	}

	return website, nil
}

type CreateWebsiteArgs struct {
	Handle string
	Locale *string
}

func (service *WebsiteService) CreateWebsite(ctx context.Context, args CreateWebsiteArgs) (*entities.WebsiteEntity, error) {
	service.utils.logger.Info().Ctx(ctx).Interface("args", args).Msg("Creating website")
	var locale string
	if args.Locale == nil {
		locale = "en"
	} else {
		locale = *args.Locale
	}

	website, err := service.websiteRepository.CreateWebsite(ctx, repositories.CreateWebsiteArgs{
		Handle: args.Handle,
		Locale: locale,
	})
	if err != nil {
		service.utils.logger.Err(err).Ctx(ctx).Msg("Error creating website")
		return nil, err
	}

	return website, nil
}
