package services

import (
	"sync"

	"github.com/mcorrigan89/website/internal/config"
	"github.com/mcorrigan89/website/internal/repositories"
	"github.com/mcorrigan89/website/internal/serviceapis"
	"github.com/rs/zerolog"
)

type ServicesUtils struct {
	logger *zerolog.Logger
	wg     *sync.WaitGroup
	config *config.Config
}

type Services struct {
	utils                   ServicesUtils
	WebsiteService          *WebsiteService
	WebsitePageService      *WebsitePageService
	WebsiteComponentService *WebsiteComponentService
}

func NewServices(repositories *repositories.Repositories, serviceApiClients *serviceapis.ServiceApiClients, cfg *config.Config, logger *zerolog.Logger, wg *sync.WaitGroup) Services {
	utils := ServicesUtils{
		logger: logger,
		wg:     wg,
		config: cfg,
	}

	websiteService := NewWebsiteService(utils, repositories)
	websitePageService := NewWebsitePageService(utils, repositories)
	websiteComponentService := NewWebsiteComponentService(utils, repositories)

	return Services{
		utils:                   utils,
		WebsiteService:          &websiteService,
		WebsitePageService:      &websitePageService,
		WebsiteComponentService: &websiteComponentService,
	}
}
