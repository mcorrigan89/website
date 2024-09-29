-- name: GetWebsiteComponentsByWebsiteSectionID :many
SELECT sqlc.embed(website_component), sqlc.embed(website_component_display), sqlc.embed(website_page) FROM website_component
LEFT JOIN website_component_display ON website_component.id = website_component_display.website_component_id
LEFT JOIN website_page ON website_component.website_section_id = website_page.id
WHERE website_section_id = $1;

-- name: GetTextComponentsByWebsiteID :many
SELECT sqlc.embed(website_component), sqlc.embed(website_component_display), sqlc.embed(text_component) FROM website_component
LEFT JOIN website_component_display ON website_component.id = website_component_display.website_component_id
JOIN text_component ON website_component.id = text_component.website_component_id
WHERE website_component.website_id = $1
AND text_component.locale = $2;

-- name: GetImageComponentsByWebsiteID :many
SELECT sqlc.embed(website_component), sqlc.embed(website_component_display), sqlc.embed(image_component) FROM website_component
LEFT JOIN website_component_display ON website_component.id = website_component_display.website_component_id
JOIN image_component ON website_component.id = image_component.website_component_id
WHERE website_component.website_id = $1;

-- name: GetWebsiteTextComponent :one
SELECT sqlc.embed(website_component), sqlc.embed(website_component_display), sqlc.embed(text_component) FROM website_component
LEFT JOIN website_component_display ON website_component.id = website_component_display.website_component_id
LEFT JOIN text_component ON website_component.id = text_component.website_component_id
WHERE website_component.id = $1
AND text_component.locale = $2;

-- name: GetWebsiteImageComponent :one
SELECT sqlc.embed(website_component), sqlc.embed(website_component_display), sqlc.embed(image_component) FROM website_component 
LEFT JOIN website_component_display ON website_component.id = website_component_display.website_component_id
LEFT JOIN image_component ON website_component.id = image_component.website_component_id
WHERE website_component.id = $1;

-- name: CreateWebsiteComponent :one
INSERT INTO website_component (website_id, website_section_id) VALUES ($1, $2) RETURNING *;

-- name: UpdateWebsiteComponent :one
UPDATE website_component SET
    website_id = coalesce(sqlc.arg(website_id), website_component.website_id),
    website_section_id = coalesce(sqlc.arg(website_section_id), website_component.website_section_id),
    updated_at = now(), 
    version = website_component.version + 1
WHERE id = sqlc.arg(id) RETURNING *;

-- name: UpdateWebsiteTextComponent :one
UPDATE text_component SET
    content_json = coalesce(sqlc.narg(content_json), text_component.content_json),
    content_html = coalesce(sqlc.narg(content_html), text_component.content_html),
    updated_at = now(), 
    version = text_component.version + 1
WHERE website_component_id = sqlc.arg(website_component_id)
AND locale = sqlc.arg(locale) RETURNING *;

-- name: UpsertWebsiteSimpleTextComponent :one
INSERT INTO text_component (website_component_id, locale, content_json, content_html) VALUES ($1, sqlc.arg(locale), sqlc.narg(content_json), sqlc.narg(content_html)) 
ON CONFLICT (website_component_id, locale) DO UPDATE SET 
    content_json = coalesce(sqlc.narg(content_json), text_component.content_json),
    content_html = coalesce(sqlc.narg(content_html), text_component.content_html),
    updated_at = now(), 
    version = text_component.version + 1 RETURNING *;

-- name: UpsertWebsiteQandAComponent :one
INSERT INTO image_component (website_component_id, image_id) VALUES (sqlc.arg(website_component_id), sqlc.arg(image_id)) 
ON CONFLICT (website_component_id) DO UPDATE SET 
    image_id = coalesce(sqlc.narg(image_id), image_component.image_id),
    updated_at = now(), 
    version = image_component.version + 1 RETURNING *;

