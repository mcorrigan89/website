// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package models

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreatePalette(ctx context.Context, arg CreatePaletteParams) (Palette, error)
	CreateWebsite(ctx context.Context, arg CreateWebsiteParams) (Website, error)
	CreateWebsiteComponent(ctx context.Context, arg CreateWebsiteComponentParams) (WebsiteComponent, error)
	CreateWebsiteConfig(ctx context.Context, arg CreateWebsiteConfigParams) (WebsiteConfig, error)
	CreateWebsiteContent(ctx context.Context, arg CreateWebsiteContentParams) (WebsiteContent, error)
	CreateWebsitePage(ctx context.Context, arg CreateWebsitePageParams) (WebsitePage, error)
	CreateWebsiteStyles(ctx context.Context, websiteID uuid.UUID) (WebsiteStyle, error)
	GetImageComponentsByWebsiteID(ctx context.Context, websiteID uuid.UUID) ([]GetImageComponentsByWebsiteIDRow, error)
	GetTextComponentsByWebsiteID(ctx context.Context, arg GetTextComponentsByWebsiteIDParams) ([]GetTextComponentsByWebsiteIDRow, error)
	GetWebsiteByComponentID(ctx context.Context, id uuid.UUID) (GetWebsiteByComponentIDRow, error)
	GetWebsiteByHandle(ctx context.Context, handle string) (GetWebsiteByHandleRow, error)
	GetWebsiteByID(ctx context.Context, id uuid.UUID) (GetWebsiteByIDRow, error)
	GetWebsiteByPageID(ctx context.Context, id uuid.UUID) (GetWebsiteByPageIDRow, error)
	GetWebsiteBySectionID(ctx context.Context, id uuid.UUID) (GetWebsiteBySectionIDRow, error)
	GetWebsiteComponentsByWebsiteSectionID(ctx context.Context, websiteSectionID uuid.UUID) ([]GetWebsiteComponentsByWebsiteSectionIDRow, error)
	GetWebsiteContentByWebsiteID(ctx context.Context, arg GetWebsiteContentByWebsiteIDParams) (GetWebsiteContentByWebsiteIDRow, error)
	GetWebsiteImageComponent(ctx context.Context, id uuid.UUID) (GetWebsiteImageComponentRow, error)
	GetWebsitePageByID(ctx context.Context, arg GetWebsitePageByIDParams) (GetWebsitePageByIDRow, error)
	GetWebsitePagesByWebsiteID(ctx context.Context, arg GetWebsitePagesByWebsiteIDParams) ([]GetWebsitePagesByWebsiteIDRow, error)
	GetWebsiteSectionByID(ctx context.Context, id uuid.UUID) (GetWebsiteSectionByIDRow, error)
	GetWebsiteSectionsByPageID(ctx context.Context, websitePageID uuid.UUID) ([]GetWebsiteSectionsByPageIDRow, error)
	GetWebsiteSectionsByWebsiteID(ctx context.Context, websiteID uuid.UUID) ([]GetWebsiteSectionsByWebsiteIDRow, error)
	GetWebsiteTextComponent(ctx context.Context, arg GetWebsiteTextComponentParams) (GetWebsiteTextComponentRow, error)
	UpdateWebsiteComponent(ctx context.Context, arg UpdateWebsiteComponentParams) (WebsiteComponent, error)
	UpdateWebsitePage(ctx context.Context, arg UpdateWebsitePageParams) (WebsitePage, error)
	UpdateWebsiteTextComponent(ctx context.Context, arg UpdateWebsiteTextComponentParams) (TextComponent, error)
	UpsertWebsitePageContent(ctx context.Context, arg UpsertWebsitePageContentParams) (WebsitePageContent, error)
	UpsertWebsiteQandAComponent(ctx context.Context, arg UpsertWebsiteQandAComponentParams) (ImageComponent, error)
	UpsertWebsiteSimpleTextComponent(ctx context.Context, arg UpsertWebsiteSimpleTextComponentParams) (TextComponent, error)
}

var _ Querier = (*Queries)(nil)
