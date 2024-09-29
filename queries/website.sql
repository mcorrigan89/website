-- name: GetWebsiteByID :one
SELECT sqlc.embed(website), sqlc.embed(website_config), sqlc.embed(website_styles), sqlc.embed(palette) FROM website 
LEFT JOIN website_config ON website.id = website_config.website_id
LEFT JOIN website_styles ON website.id = website_styles.website_id
LEFT JOIN palette ON website_styles.id = palette.website_styles_id
WHERE website.id = $1;

-- name: GetWebsiteByHandle :one
SELECT sqlc.embed(website), sqlc.embed(website_config), sqlc.embed(website_styles), sqlc.embed(palette) FROM website 
LEFT JOIN website_config ON website.id = website_config.website_id
LEFT JOIN website_styles ON website.id = website_styles.website_id
LEFT JOIN palette ON website_styles.id = palette.website_styles_id
WHERE website.handle = $1;

-- name: GetWebsiteByPageID :one
SELECT sqlc.embed(website) FROM website 
LEFT JOIN website_page ON website.id = website_page.website_id
WHERE website_page.id = $1;

-- name: GetWebsiteBySectionID :one
SELECT sqlc.embed(website) FROM website 
LEFT JOIN website_section ON website.id = website_section.website_id
WHERE website_section.id = $1;

-- name: GetWebsiteByComponentID :one
SELECT sqlc.embed(website) FROM website 
LEFT JOIN website_component ON website.id = website_component.website_id
WHERE website_component.id = $1;

-- name: GetWebsiteSectionByID :one
SELECT sqlc.embed(website_section) FROM website_section WHERE id = $1;

-- name: GetWebsitePagesByWebsiteID :many
SELECT sqlc.embed(website_page), sqlc.embed(website_page_content) FROM website_page 
LEFT JOIN website_page_content ON website_page.id = website_page_content.website_page_id
WHERE website_page.website_id = $1 AND website_page_content.locale = $2
ORDER BY website_page.sort_key;

-- name: GetWebsiteContentByWebsiteID :one
SELECT sqlc.embed(website_content) FROM website_content WHERE website_id = $1 AND locale = $2;

-- name: CreateWebsite :one
INSERT INTO website (handle, default_locale) VALUES ($1, $2) RETURNING *;

-- name: CreateWebsiteContent :one
INSERT INTO website_content (website_id, locale, website_display_name, website_display_description) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: CreateWebsiteConfig :one
INSERT INTO website_config (website_id, default_page_id) VALUES ($1, $2) RETURNING *;

-- name: CreateWebsiteStyles :one
INSERT INTO website_styles (website_id) VALUES ($1) RETURNING *;

-- name: CreatePalette :one
INSERT INTO palette (website_styles_id, color_one, color_two, color_three, color_four, color_five, color_six) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;