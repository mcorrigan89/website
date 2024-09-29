package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/mcorrigan89/website/internal/entities"
	"github.com/mcorrigan89/website/internal/repositories"
)

type WebsiteComponentService struct {
	utils                      ServicesUtils
	websiteRepository          *repositories.WebsiteRepository
	websitePageRepository      *repositories.WebsitePageRepository
	websiteComponentRepository *repositories.WebsiteComponentRepository
}

func NewWebsiteComponentService(utils ServicesUtils, repos *repositories.Repositories) WebsiteComponentService {
	return WebsiteComponentService{
		utils:                      utils,
		websiteRepository:          repos.WebsiteRepository,
		websitePageRepository:      repos.WebsitePageRepository,
		websiteComponentRepository: repos.WebsiteComponentRepository,
	}
}

type CreateTextComponentArgs struct {
	WebsiteSectionID uuid.UUID
	Locale           *string
	Json             []byte
	Html             *string
}

func (service *WebsiteComponentService) CreateTextComponent(ctx context.Context, args CreateTextComponentArgs) (*entities.WebsiteComponentEntity, error) {
	service.utils.logger.Info().Ctx(ctx).Interface("args", args).Msg("Creating smple text component")
	component, err := service.websiteComponentRepository.CreateTextComponent(ctx, repositories.CreateTextComponentArgs{
		WebsiteSectionID: args.WebsiteSectionID,
		Locale:           args.Locale,
		Json:             args.Json,
		Html:             args.Html,
	})

	if err != nil {
		service.utils.logger.Err(err).Ctx(ctx).Msg("Error creating simple text component")
		return nil, err
	}

	return component, nil

}

type UpdateTextComponentArgs struct {
	ID     uuid.UUID
	Locale *string
	Json   []byte
	Html   *string
}

func (service *WebsiteComponentService) UpdateTextComponent(ctx context.Context, args UpdateTextComponentArgs) (*entities.WebsiteComponentEntity, error) {
	service.utils.logger.Info().Ctx(ctx).Interface("args", args).Msg("Updating text component")
	component, err := service.websiteComponentRepository.UpdateTextComponent(ctx, repositories.UpdateTextComponent{
		ID:     args.ID,
		Locale: args.Locale,
		Json:   args.Json,
		Html:   args.Html,
	})

	if err != nil {
		service.utils.logger.Err(err).Ctx(ctx).Msg("Error updating text component")
		return nil, err
	}

	return component, nil

}
