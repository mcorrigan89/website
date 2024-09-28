-- name: GetWebsiteByID :one
SELECT sqlc.embed(website) FROM website WHERE id = $1;

-- name: GetWebsiteByHandle :one
SELECT sqlc.embed(website) FROM website WHERE handle = $1;

-- name: GetWebsiteByPageID :one
SELECT sqlc.embed(website) FROM website 
LEFT JOIN website_page ON website.id = website_page.website_id
WHERE website_page.id = $1;

-- name: GetWebsiteByComponentID :one
SELECT sqlc.embed(website) FROM website 
LEFT JOIN website_component ON website.id = website_component.website_id
WHERE website_component.id = $1;

-- name: GetWebsitePagesByWebsiteID :many
SELECT sqlc.embed(website_page), sqlc.embed(website_page_content) FROM website_page 
LEFT JOIN website_page_content ON website_page.id = website_page_content.website_page_id
WHERE website_page.website_id = $1 AND website_page_content.locale = $2
ORDER BY website_page.sort_key;

-- name: GetWebsiteContentByWebsiteID :one
SELECT sqlc.embed(website_content) FROM website_content WHERE website_id = $1 AND locale = $2;