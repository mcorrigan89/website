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

type CreateSimpleTextComponentArgs struct {
	WebsitePageID uuid.UUID
	Locale        *string
	Content       *string
}

func (service *WebsiteComponentService) CreateSimpleTextComponent(ctx context.Context, args CreateSimpleTextComponentArgs) (*entities.WebsiteComponentEntity, error) {
	service.utils.logger.Info().Ctx(ctx).Interface("args", args).Msg("Creating smple text component")
	component, err := service.websiteComponentRepository.CreateSimpleTextComponent(ctx, repositories.CreateSimpleTextComponentArgs{
		WebsitePageID: args.WebsitePageID,
		Locale:        args.Locale,
		Content:       args.Content,
	})

	if err != nil {
		service.utils.logger.Err(err).Ctx(ctx).Msg("Error creating simple text component")
		return nil, err
	}

	return component, nil

}

type UpdateTextComponentArgs struct {
	ID      uuid.UUID
	Locale  *string
	Content *string
}

func (service *WebsiteComponentService) UpdateSimpleTextComponent(ctx context.Context, args UpdateTextComponentArgs) (*entities.WebsiteComponentEntity, error) {
	service.utils.logger.Info().Ctx(ctx).Interface("args", args).Msg("Updating text component")
	component, err := service.websiteComponentRepository.UpdateSimpleTextComponent(ctx, repositories.UpdateSimpleTextComponent{
		ID:      args.ID,
		Locale:  args.Locale,
		Content: args.Content,
	})

	if err != nil {
		service.utils.logger.Err(err).Ctx(ctx).Msg("Error updating text component")
		return nil, err
	}

	return component, nil

}
