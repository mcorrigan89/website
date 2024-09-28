// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: website_page.sql

package models

import (
	"context"

	"github.com/google/uuid"
)

const createWebsitePage = `-- name: CreateWebsitePage :one
INSERT INTO website_page (website_id, sort_key, url_slug)
VALUES ($1, $2, $3)
RETURNING id, website_id, url_slug, sort_key, created_at, updated_at, version
`

type CreateWebsitePageParams struct {
	WebsiteID uuid.UUID `json:"website_id"`
	SortKey   string    `json:"sort_key"`
	UrlSlug   string    `json:"url_slug"`
}

func (q *Queries) CreateWebsitePage(ctx context.Context, arg CreateWebsitePageParams) (WebsitePage, error) {
	row := q.db.QueryRow(ctx, createWebsitePage, arg.WebsiteID, arg.SortKey, arg.UrlSlug)
	var i WebsitePage
	err := row.Scan(
		&i.ID,
		&i.WebsiteID,
		&i.UrlSlug,
		&i.SortKey,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const getWebsitePageByID = `-- name: GetWebsitePageByID :one
SELECT website_page.id, website_page.website_id, website_page.url_slug, website_page.sort_key, website_page.created_at, website_page.updated_at, website_page.version, website_page_content.id, website_page_content.website_page_id, website_page_content.locale, website_page_content.title, website_page_content.subtitle, website_page_content.created_at, website_page_content.updated_at, website_page_content.version FROM website_page 
LEFT JOIN website_page_content ON website_page.id = website_page_content.website_page_id
WHERE website_page.id = $1
AND website_page_content.locale = $2
`

type GetWebsitePageByIDParams struct {
	ID     uuid.UUID `json:"id"`
	Locale string    `json:"locale"`
}

type GetWebsitePageByIDRow struct {
	WebsitePage        WebsitePage        `json:"website_page"`
	WebsitePageContent WebsitePageContent `json:"website_page_content"`
}

func (q *Queries) GetWebsitePageByID(ctx context.Context, arg GetWebsitePageByIDParams) (GetWebsitePageByIDRow, error) {
	row := q.db.QueryRow(ctx, getWebsitePageByID, arg.ID, arg.Locale)
	var i GetWebsitePageByIDRow
	err := row.Scan(
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
	)
	return i, err
}

const updateWebsitePage = `-- name: UpdateWebsitePage :one
UPDATE website_page SET 
    url_slug = coalesce($2, website_page.url_slug), 
    updated_at = now(), 
    version = website_page.version + 1
WHERE id = $1 RETURNING id, website_id, url_slug, sort_key, created_at, updated_at, version
`

type UpdateWebsitePageParams struct {
	ID      uuid.UUID `json:"id"`
	UrlSlug *string   `json:"url_slug"`
}

func (q *Queries) UpdateWebsitePage(ctx context.Context, arg UpdateWebsitePageParams) (WebsitePage, error) {
	row := q.db.QueryRow(ctx, updateWebsitePage, arg.ID, arg.UrlSlug)
	var i WebsitePage
	err := row.Scan(
		&i.ID,
		&i.WebsiteID,
		&i.UrlSlug,
		&i.SortKey,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}

const upsertWebsitePageContent = `-- name: UpsertWebsitePageContent :one
INSERT INTO website_page_content (id, website_page_id, locale, title, subtitle) VALUES ($1, $2, $3, $4, $5) 
ON CONFLICT (website_page_id, locale) DO UPDATE SET 
    website_page_id = $2, 
    locale = coalesce($3, website_page_content.locale),
    title = coalesce($4, website_page_content.title),
    subtitle = coalesce($5, website_page_content.subtitle),
    updated_at = now(), 
    version = website_page_content.version + 1 
RETURNING id, website_page_id, locale, title, subtitle, created_at, updated_at, version
`

type UpsertWebsitePageContentParams struct {
	ID            uuid.UUID `json:"id"`
	WebsitePageID uuid.UUID `json:"website_page_id"`
	Locale        string    `json:"locale"`
	Title         *string   `json:"title"`
	Subtitle      *string   `json:"subtitle"`
}

func (q *Queries) UpsertWebsitePageContent(ctx context.Context, arg UpsertWebsitePageContentParams) (WebsitePageContent, error) {
	row := q.db.QueryRow(ctx, upsertWebsitePageContent,
		arg.ID,
		arg.WebsitePageID,
		arg.Locale,
		arg.Title,
		arg.Subtitle,
	)
	var i WebsitePageContent
	err := row.Scan(
		&i.ID,
		&i.WebsitePageID,
		&i.Locale,
		&i.Title,
		&i.Subtitle,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Version,
	)
	return i, err
}
