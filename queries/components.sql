-- name: GetWebsiteComponentsByWebsitePageID :many
SELECT sqlc.embed(website_component), sqlc.embed(website_page) FROM website_component
LEFT JOIN website_page ON website_component.website_page_id = website_page.id
WHERE website_page_id = $1;

-- name: GetTextComponentsByWebsiteID :many
SELECT sqlc.embed(website_component), sqlc.embed(simple_text_component) FROM website_component
JOIN simple_text_component ON website_component.id = simple_text_component.website_component_id
WHERE website_component.website_id = $1
AND simple_text_component.locale = $2;

-- name: GetImageComponentsByWebsiteID :many
SELECT sqlc.embed(website_component), sqlc.embed(image_component) FROM website_component
JOIN image_component ON website_component.id = image_component.website_component_id
WHERE website_component.website_id = $1;

-- name: GetWebsiteTextComponent :one
SELECT sqlc.embed(website_component), sqlc.embed(simple_text_component) FROM website_component 
LEFT JOIN simple_text_component ON website_component.id = simple_text_component.website_component_id
WHERE website_component.id = $1
AND simple_text_component.locale = $2;

-- name: GetWebsiteImageComponent :one
SELECT sqlc.embed(website_component), sqlc.embed(image_component) FROM website_component 
LEFT JOIN image_component ON website_component.id = image_component.website_component_id
WHERE website_component.id = $1;

-- name: CreateWebsiteComponent :one
INSERT INTO website_component (website_id, website_page_id, sort_key) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateWebsiteComponent :one
UPDATE website_component SET
    website_id = coalesce(sqlc.arg(website_id), website_component.website_id),
    website_page_id = coalesce(sqlc.arg(website_page_id), website_component.website_page_id),
    sort_key = coalesce(sqlc.arg(sort_key), website_component.sort_key),
    updated_at = now(), 
    version = website_component.version + 1
WHERE id = sqlc.arg(id) RETURNING *;

-- name: UpdateWebsiteTextComponent :one
UPDATE simple_text_component SET
    content = coalesce(sqlc.narg(content), simple_text_component.content),
    updated_at = now(), 
    version = simple_text_component.version + 1
WHERE website_component_id = sqlc.arg(website_component_id)
AND locale = sqlc.arg(locale) RETURNING *;

-- -- name: CreateWebsiteQandAComponent :one
-- INSERT INTO simple_qanda_component (website_component_id, locale, question, answer, firebase_key, firebase_ref) 
-- VALUES (sqlc.arg(website_component_id), sqlc.arg(locale), sqlc.narg(question), sqlc.narg(answer), sqlc.narg(firebase_key), sqlc.narg(firebase_ref)) RETURNING *;

-- -- name: UpdateWebsiteQandAComponent :one
-- UPDATE simple_qanda_component SET 
--     question = coalesce(sqlc.narg(question), simple_qanda_component.question),
--     answer = coalesce(sqlc.narg(answer), simple_qanda_component.answer),
--     firebase_key = coalesce(sqlc.narg(firebase_key), simple_qanda_component.firebase_key),
--     firebase_ref = coalesce(sqlc.narg(firebase_ref), simple_qanda_component.firebase_ref),
--     updated_at = now(), 
--     version = simple_qanda_component.version + 1
-- WHERE website_component_id = sqlc.arg(website_component_id)
-- AND locale = sqlc.arg(locale) RETURNING *;

-- name: UpsertWebsiteSimpleTextComponent :one
INSERT INTO simple_text_component (website_component_id, locale, content) VALUES ($1, sqlc.arg(locale), sqlc.narg(content)) 
ON CONFLICT (website_component_id, locale) DO UPDATE SET 
    content = coalesce(sqlc.narg(content), simple_text_component.content),
    updated_at = now(), 
    version = simple_text_component.version + 1 RETURNING *;

-- name: UpsertWebsiteQandAComponent :one
INSERT INTO image_component (website_component_id, image_id) VALUES (sqlc.arg(website_component_id), sqlc.arg(image_id)) 
ON CONFLICT (website_component_id) DO UPDATE SET 
    image_id = coalesce(sqlc.narg(image_id), image_component.image_id),
    updated_at = now(), 
    version = image_component.version + 1 RETURNING *;

