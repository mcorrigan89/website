package repositories

import (
	"errors"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mcorrigan89/website/internal/repositories/models"
	"github.com/rs/zerolog"
)

var (
	ErrNotFound = errors.New("not found")
)

const defaultTimeout = 10 * time.Second

type ServicesUtils struct {
	logger *zerolog.Logger
	wg     *sync.WaitGroup
	db     *pgxpool.Pool
}

type Repositories struct {
	utils                      ServicesUtils
	WebsiteRepository          *WebsiteRepository
	WebsitePageRepository      *WebsitePageRepository
	WebsiteComponentRepository *WebsiteComponentRepository
}

func NewRepositories(db *pgxpool.Pool, logger *zerolog.Logger, wg *sync.WaitGroup) Repositories {
	queries := models.New(db)
	utils := ServicesUtils{
		logger: logger,
		wg:     wg,
		db:     db,
	}

	websiteRepo := NewWebsiteRepository(utils, db, queries)
	websitePageRepo := NewWebsitePageRepository(utils, db, queries)
	websiteComponentRepo := NewWebsiteComponentRepository(utils, db, queries)

	return Repositories{
		utils:                      utils,
		WebsiteRepository:          websiteRepo,
		WebsitePageRepository:      websitePageRepo,
		WebsiteComponentRepository: websiteComponentRepo,
	}
}
