-- name: GetWebsitePageByID :one
SELECT sqlc.embed(website_page), sqlc.embed(website_page_content) FROM website_page 
LEFT JOIN website_page_content ON website_page.id = website_page_content.website_page_id
WHERE website_page.id = $1
AND website_page_content.locale = $2;

-- name: CreateWebsitePage :one
INSERT INTO website_page (website_id, sort_key, url_slug)
VALUES ($1, $2, sqlc.arg(url_slug))
RETURNING *;

-- name: UpdateWebsitePage :one
UPDATE website_page SET 
    url_slug = coalesce(sqlc.narg(url_slug), website_page.url_slug), 
    updated_at = now(), 
    version = website_page.version + 1
WHERE id = $1 RETURNING *;

-- name: UpsertWebsitePageContent :one
INSERT INTO website_page_content (id, website_page_id, locale, title, subtitle) VALUES ($1, $2, sqlc.arg(locale), sqlc.narg(title), sqlc.narg(subtitle)) 
ON CONFLICT (website_page_id, locale) DO UPDATE SET 
    website_page_id = $2, 
    locale = coalesce(sqlc.arg(locale), website_page_content.locale),
    title = coalesce(sqlc.narg(title), website_page_content.title),
    subtitle = coalesce(sqlc.narg(subtitle), website_page_content.subtitle),
    updated_at = now(), 
    version = website_page_content.version + 1 
RETURNING *;

-- name: GetWebsiteSectionsByWebsiteID :many
SELECT sqlc.embed(website_section), sqlc.embed(website_section_display) FROM website_section 
LEFT JOIN website_section_display ON website_section.id = website_section_display.website_section_id
WHERE website_section.website_id = $1;

-- name: GetWebsiteSectionsByPageID :many
SELECT sqlc.embed(website_section), sqlc.embed(website_section_display) FROM website_section
LEFT JOIN website_section_display ON website_section.id = website_section_display.website_section_id
WHERE website_section.website_page_id = $1 
ORDER BY sort_key;