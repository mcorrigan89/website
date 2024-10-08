// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: website.sql

package models

import (
	"context"

	"github.com/google/uuid"
)

const createPalette = `-- name: CreatePalette :one
INSERT INTO palette (website_styles_id, color_one, color_two, color_three, color_four, color_five, color_six) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, website_styles_id, color_one, color_two, color_three, color_four, color_five, color_six, created_at, updated_at, version
`

type CreatePaletteParams struct {
	WebsiteStylesID uuid.UUID `json:"website_styles_id"`
	ColorOne        string    `json:"color_one"`
	ColorTwo        string    `json:"color_two"`
	ColorThree      string    `json:"color_three"`
	ColorFour       string    `json:"color_four"`
	ColorFive       string    `json:"color_five"`
	ColorSix        string    `json:"color_six"`
}

func (q *Queries) CreatePalette(ctx context.Context, arg CreatePaletteParams) (Palette, error) {
	row := q.db.QueryRow(ctx, createPalette,
		arg.WebsiteStylesID,
		arg.ColorOne,
		arg.ColorTwo,
		arg.ColorThree,
		arg.ColorFour,
		arg.ColorFive,
		arg.ColorSix,
	)
	var i Palette
	err := row.Scan(
		&i.ID,
		&i.WebsiteStylesID,
		&i.ColorOne,
		&i.ColorTwo,
		&i.ColorThree,
		&i.ColorFour,
		&i.ColorFive,
		&i.ColorSix,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const createWebsite = `-- name: CreateWebsite :one
INSERT INTO website (handle, default_locale) VALUES ($1, $2) RETURNING id, handle, default_locale, created_at, updated_at, version
`

type CreateWebsiteParams struct {
	Handle        string `json:"handle"`
	DefaultLocale string `json:"default_locale"`
}

func (q *Queries) CreateWebsite(ctx context.Context, arg CreateWebsiteParams) (Website, error) {
	row := q.db.QueryRow(ctx, createWebsite, arg.Handle, arg.DefaultLocale)
	var i Website
	err := row.Scan(
		&i.ID,
		&i.Handle,
		&i.DefaultLocale,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const createWebsiteConfig = `-- name: CreateWebsiteConfig :one
INSERT INTO website_config (website_id, default_page_id) VALUES ($1, $2) RETURNING id, website_id, default_page_id, created_at, updated_at, version
`

type CreateWebsiteConfigParams struct {
	WebsiteID     uuid.UUID `json:"website_id"`
	DefaultPageID uuid.UUID `json:"default_page_id"`
}

func (q *Queries) CreateWebsiteConfig(ctx context.Context, arg CreateWebsiteConfigParams) (WebsiteConfig, error) {
	row := q.db.QueryRow(ctx, createWebsiteConfig, arg.WebsiteID, arg.DefaultPageID)
	var i WebsiteConfig
	err := row.Scan(
		&i.ID,
		&i.WebsiteID,
		&i.DefaultPageID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const createWebsiteContent = `-- name: CreateWebsiteContent :one
INSERT INTO website_content (website_id, locale, website_display_name, website_display_description) VALUES ($1, $2, $3, $4) RETURNING id, website_id, locale, website_display_name, website_display_description, created_at, updated_at, version
`

type CreateWebsiteContentParams struct {
	WebsiteID                 uuid.UUID `json:"website_id"`
	Locale                    string    `json:"locale"`
	WebsiteDisplayName        string    `json:"website_display_name"`
	WebsiteDisplayDescription *string   `json:"website_display_description"`
}

func (q *Queries) CreateWebsiteContent(ctx context.Context, arg CreateWebsiteContentParams) (WebsiteContent, error) {
	row := q.db.QueryRow(ctx, createWebsiteContent,
		arg.WebsiteID,
		arg.Locale,
		arg.WebsiteDisplayName,
		arg.WebsiteDisplayDescription,
	)
	var i WebsiteContent
	err := row.Scan(
		&i.ID,
		&i.WebsiteID,
		&i.Locale,
		&i.WebsiteDisplayName,
		&i.WebsiteDisplayDescription,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const createWebsiteStyles = `-- name: CreateWebsiteStyles :one
INSERT INTO website_styles (website_id) VALUES ($1) RETURNING id, website_id, created_at, updated_at, version
`

func (q *Queries) CreateWebsiteStyles(ctx context.Context, websiteID uuid.UUID) (WebsiteStyle, error) {
	row := q.db.QueryRow(ctx, createWebsiteStyles, websiteID)
	var i WebsiteStyle
	err := row.Scan(
		&i.ID,
		&i.WebsiteID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const getWebsiteByComponentID = `-- name: GetWebsiteByComponentID :one
SELECT website.id, website.handle, website.default_locale, website.created_at, website.updated_at, website.version FROM website 
LEFT JOIN website_component ON website.id = website_component.website_id
WHERE website_component.id = $1
`

type GetWebsiteByComponentIDRow struct {
	Website Website `json:"website"`
}

func (q *Queries) GetWebsiteByComponentID(ctx context.Context, id uuid.UUID) (GetWebsiteByComponentIDRow, error) {
	row := q.db.QueryRow(ctx, getWebsiteByComponentID, id)
	var i GetWebsiteByComponentIDRow
	err := row.Scan(
		&i.Website.ID,
		&i.Website.Handle,
		&i.Website.DefaultLocale,
		&i.Website.CreatedAt,
		&i.Website.UpdatedAt,
		&i.Website.Version,
	)
	return i, err
}

const getWebsiteByHandle = `-- name: GetWebsiteByHandle :one
SELECT website.id, website.handle, website.default_locale, website.created_at, website.updated_at, website.version, website_config.id, website_config.website_id, website_config.default_page_id, website_config.created_at, website_config.updated_at, website_config.version, website_styles.id, website_styles.website_id, website_styles.created_at, website_styles.updated_at, website_styles.version, palette.id, palette.website_styles_id, palette.color_one, palette.color_two, palette.color_three, palette.color_four, palette.color_five, palette.color_six, palette.created_at, palette.updated_at, palette.version FROM website 
LEFT JOIN website_config ON website.id = website_config.website_id
LEFT JOIN website_styles ON website.id = website_styles.website_id
LEFT JOIN palette ON website_styles.id = palette.website_styles_id
WHERE website.handle = $1
`

type GetWebsiteByHandleRow struct {
	Website       Website       `json:"website"`
	WebsiteConfig WebsiteConfig `json:"website_config"`
	WebsiteStyle  WebsiteStyle  `json:"website_style"`
	Palette       Palette       `json:"palette"`
}

func (q *Queries) GetWebsiteByHandle(ctx context.Context, handle string) (GetWebsiteByHandleRow, error) {
	row := q.db.QueryRow(ctx, getWebsiteByHandle, handle)
	var i GetWebsiteByHandleRow
	err := row.Scan(
		&i.Website.ID,
		&i.Website.Handle,
		&i.Website.DefaultLocale,
		&i.Website.CreatedAt,
		&i.Website.UpdatedAt,
		&i.Website.Version,
		&i.WebsiteConfig.ID,
		&i.WebsiteConfig.WebsiteID,
		&i.WebsiteConfig.DefaultPageID,
		&i.WebsiteConfig.CreatedAt,
		&i.WebsiteConfig.UpdatedAt,
		&i.WebsiteConfig.Version,
		&i.WebsiteStyle.ID,
		&i.WebsiteStyle.WebsiteID,
		&i.WebsiteStyle.CreatedAt,
		&i.WebsiteStyle.UpdatedAt,
		&i.WebsiteStyle.Version,
		&i.Palette.ID,
		&i.Palette.WebsiteStylesID,
		&i.Palette.ColorOne,
		&i.Palette.ColorTwo,
		&i.Palette.ColorThree,
		&i.Palette.ColorFour,
		&i.Palette.ColorFive,
		&i.Palette.ColorSix,
		&i.Palette.CreatedAt,
		&i.Palette.UpdatedAt,
		&i.Palette.Version,
	)
	return i, err
}

const getWebsiteByID = `-- name: GetWebsiteByID :one
SELECT website.id, website.handle, website.default_locale, website.created_at, website.updated_at, website.version, website_config.id, website_config.website_id, website_config.default_page_id, website_config.created_at, website_config.updated_at, website_config.version, website_styles.id, website_styles.website_id, website_styles.created_at, website_styles.updated_at, website_styles.version, palette.id, palette.website_styles_id, palette.color_one, palette.color_two, palette.color_three, palette.color_four, palette.color_five, palette.color_six, palette.created_at, palette.updated_at, palette.version FROM website 
LEFT JOIN website_config ON website.id = website_config.website_id
LEFT JOIN website_styles ON website.id = website_styles.website_id
LEFT JOIN palette ON website_styles.id = palette.website_styles_id
WHERE website.id = $1
`

type GetWebsiteByIDRow struct {
	Website       Website       `json:"website"`
	WebsiteConfig WebsiteConfig `json:"website_config"`
	WebsiteStyle  WebsiteStyle  `json:"website_style"`
	Palette       Palette       `json:"palette"`
}

func (q *Queries) GetWebsiteByID(ctx context.Context, id uuid.UUID) (GetWebsiteByIDRow, error) {
	row := q.db.QueryRow(ctx, getWebsiteByID, id)
	var i GetWebsiteByIDRow
	err := row.Scan(
		&i.Website.ID,
		&i.Website.Handle,
		&i.Website.DefaultLocale,
		&i.Website.CreatedAt,
		&i.Website.UpdatedAt,
		&i.Website.Version,
		&i.WebsiteConfig.ID,
		&i.WebsiteConfig.WebsiteID,
		&i.WebsiteConfig.DefaultPageID,
		&i.WebsiteConfig.CreatedAt,
		&i.WebsiteConfig.UpdatedAt,
		&i.WebsiteConfig.Version,
		&i.WebsiteStyle.ID,
		&i.WebsiteStyle.WebsiteID,
		&i.WebsiteStyle.CreatedAt,
		&i.WebsiteStyle.UpdatedAt,
		&i.WebsiteStyle.Version,
		&i.Palette.ID,
		&i.Palette.WebsiteStylesID,
		&i.Palette.ColorOne,
		&i.Palette.ColorTwo,
		&i.Palette.ColorThree,
		&i.Palette.ColorFour,
		&i.Palette.ColorFive,
		&i.Palette.ColorSix,
		&i.Palette.CreatedAt,
		&i.Palette.UpdatedAt,
		&i.Palette.Version,
	)
	return i, err
}

const getWebsiteByPageID = `-- name: GetWebsiteByPageID :one
SELECT website.id, website.handle, website.default_locale, website.created_at, website.updated_at, website.version FROM website 
LEFT JOIN website_page ON website.id = website_page.website_id
WHERE website_page.id = $1
`

type GetWebsiteByPageIDRow struct {
	Website Website `json:"website"`
}

func (q *Queries) GetWebsiteByPageID(ctx context.Context, id uuid.UUID) (GetWebsiteByPageIDRow, error) {
	row := q.db.QueryRow(ctx, getWebsiteByPageID, id)
	var i GetWebsiteByPageIDRow
	err := row.Scan(
		&i.Website.ID,
		&i.Website.Handle,
		&i.Website.DefaultLocale,
		&i.Website.CreatedAt,
		&i.Website.UpdatedAt,
		&i.Website.Version,
	)
	return i, err
}

const getWebsiteBySectionID = `-- name: GetWebsiteBySectionID :one
SELECT website.id, website.handle, website.default_locale, website.created_at, website.updated_at, website.version FROM website 
LEFT JOIN website_section ON website.id = website_section.website_id
WHERE website_section.id = $1
`

type GetWebsiteBySectionIDRow struct {
	Website Website `json:"website"`
}

func (q *Queries) GetWebsiteBySectionID(ctx context.Context, id uuid.UUID) (GetWebsiteBySectionIDRow, error) {
	row := q.db.QueryRow(ctx, getWebsiteBySectionID, id)
	var i GetWebsiteBySectionIDRow
	err := row.Scan(
		&i.Website.ID,
		&i.Website.Handle,
		&i.Website.DefaultLocale,
		&i.Website.CreatedAt,
		&i.Website.UpdatedAt,
		&i.Website.Version,
	)
	return i, err
}

const getWebsiteContentByWebsiteID = `-- name: GetWebsiteContentByWebsiteID :one
SELECT website_content.id, website_content.website_id, website_content.locale, website_content.website_display_name, website_content.website_display_description, website_content.created_at, website_content.updated_at, website_content.version FROM website_content WHERE website_id = $1 AND locale = $2
`

type GetWebsiteContentByWebsiteIDParams struct {
	WebsiteID uuid.UUID `json:"website_id"`
	Locale    string    `json:"locale"`
}

type GetWebsiteContentByWebsiteIDRow struct {
	WebsiteContent WebsiteContent `json:"website_content"`
}

func (q *Queries) GetWebsiteContentByWebsiteID(ctx context.Context, arg GetWebsiteContentByWebsiteIDParams) (GetWebsiteContentByWebsiteIDRow, error) {
	row := q.db.QueryRow(ctx, getWebsiteContentByWebsiteID, arg.WebsiteID, arg.Locale)
	var i GetWebsiteContentByWebsiteIDRow
	err := row.Scan(
		&i.WebsiteContent.ID,
		&i.WebsiteContent.WebsiteID,
		&i.WebsiteContent.Locale,
		&i.WebsiteContent.WebsiteDisplayName,
		&i.WebsiteContent.WebsiteDisplayDescription,
		&i.WebsiteContent.CreatedAt,
		&i.WebsiteContent.UpdatedAt,
		&i.WebsiteContent.Version,
	)
	return i, err
}

const getWebsitePagesByWebsiteID = `-- name: GetWebsitePagesByWebsiteID :many
SELECT website_page.id, website_page.website_id, website_page.url_slug, website_page.sort_key, website_page.created_at, website_page.updated_at, website_page.version, website_page_content.id, website_page_content.website_page_id, website_page_content.locale, website_page_content.title, website_page_content.subtitle, website_page_content.created_at, website_page_content.updated_at, website_page_content.version FROM website_page 
LEFT JOIN website_page_content ON website_page.id = website_page_content.website_page_id
WHERE website_page.website_id = $1 AND website_page_content.locale = $2
ORDER BY website_page.sort_key
`

type GetWebsitePagesByWebsiteIDParams struct {
	WebsiteID uuid.UUID `json:"website_id"`
	Locale    string    `json:"locale"`
}

type GetWebsitePagesByWebsiteIDRow struct {
	WebsitePage        WebsitePage        `json:"website_page"`
	WebsitePageContent WebsitePageContent `json:"website_page_content"`
}

func (q *Queries) GetWebsitePagesByWebsiteID(ctx context.Context, arg GetWebsitePagesByWebsiteIDParams) ([]GetWebsitePagesByWebsiteIDRow, error) {
	rows, err := q.db.Query(ctx, getWebsitePagesByWebsiteID, arg.WebsiteID, arg.Locale)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetWebsitePagesByWebsiteIDRow{}
	for rows.Next() {
		var i GetWebsitePagesByWebsiteIDRow
		if err := rows.Scan(
			&i.WebsitePage.ID,
			&i.WebsitePage.WebsiteID,
			&i.WebsitePage.UrlSlug,
			&i.WebsitePage.SortKey,
			&i.WebsitePage.CreatedAt,
			&i.WebsitePage.UpdatedAt,
			&i.WebsitePage.Version,
			&i.WebsitePageContent.ID,
			&i.WebsitePageContent.WebsitePageID,
			&i.WebsitePageContent.Locale,
			&i.WebsitePageContent.Title,
			&i.WebsitePageContent.Subtitle,
			&i.WebsitePageContent.CreatedAt,
			&i.WebsitePageContent.UpdatedAt,
			&i.WebsitePageContent.Version,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWebsiteSectionByID = `-- name: GetWebsiteSectionByID :one
SELECT website_section.id, website_section.website_id, website_section.website_page_id, website_section.sort_key, website_section.created_at, website_section.updated_at, website_section.version FROM website_section WHERE id = $1
`

type GetWebsiteSectionByIDRow struct {
	WebsiteSection WebsiteSection `json:"website_section"`
}

func (q *Queries) GetWebsiteSectionByID(ctx context.Context, id uuid.UUID) (GetWebsiteSectionByIDRow, error) {
	row := q.db.QueryRow(ctx, getWebsiteSectionByID, id)
	var i GetWebsiteSectionByIDRow
	err := row.Scan(
		&i.WebsiteSection.ID,
		&i.WebsiteSection.WebsiteID,
		&i.WebsiteSection.WebsitePageID,
		&i.WebsiteSection.SortKey,
		&i.WebsiteSection.CreatedAt,
		&i.WebsiteSection.UpdatedAt,
		&i.WebsiteSection.Version,
	)
	return i, err
}
