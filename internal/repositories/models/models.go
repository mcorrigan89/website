// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type DrizzleDrizzleMigration struct {
	ID        int32  `json:"id"`
	Hash      string `json:"hash"`
	CreatedAt *int64 `json:"created_at"`
}

type ImageComponent struct {
	ID                 uuid.UUID          `json:"id"`
	WebsiteComponentID uuid.UUID          `json:"website_component_id"`
	ImageID            uuid.UUID          `json:"image_id"`
	CreatedAt          pgtype.Timestamptz `json:"created_at"`
	UpdatedAt          pgtype.Timestamptz `json:"updated_at"`
	Version            int32              `json:"version"`
}

type Palette struct {
	ID              uuid.UUID          `json:"id"`
	WebsiteStylesID uuid.UUID          `json:"website_styles_id"`
	ColorOne        string             `json:"color_one"`
	ColorTwo        string             `json:"color_two"`
	ColorThree      string             `json:"color_three"`
	ColorFour       string             `json:"color_four"`
	ColorFive       string             `json:"color_five"`
	ColorSix        string             `json:"color_six"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
	Version         int32              `json:"version"`
}

type SchemaMigration struct {
	Version int64 `json:"version"`
	Dirty   bool  `json:"dirty"`
}

type TextComponent struct {
	ID                 uuid.UUID          `json:"id"`
	WebsiteComponentID uuid.UUID          `json:"website_component_id"`
	Locale             string             `json:"locale"`
	ContentJson        []byte             `json:"content_json"`
	ContentHtml        *string            `json:"content_html"`
	CreatedAt          pgtype.Timestamptz `json:"created_at"`
	UpdatedAt          pgtype.Timestamptz `json:"updated_at"`
	Version            int32              `json:"version"`
}

type Website struct {
	ID            uuid.UUID          `json:"id"`
	Handle        string             `json:"handle"`
	DefaultLocale string             `json:"default_locale"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
	Version       int32              `json:"version"`
}

type WebsiteComponent struct {
	ID               uuid.UUID          `json:"id"`
	WebsiteID        uuid.UUID          `json:"website_id"`
	WebsiteSectionID uuid.UUID          `json:"website_section_id"`
	CreatedAt        pgtype.Timestamptz `json:"created_at"`
	UpdatedAt        pgtype.Timestamptz `json:"updated_at"`
	Version          int32              `json:"version"`
}

type WebsiteComponentDisplay struct {
	ID                 uuid.UUID          `json:"id"`
	WebsiteComponentID uuid.UUID          `json:"website_component_id"`
	Height             int32              `json:"height"`
	Width              int32              `json:"width"`
	XCoordinate        int32              `json:"x_coordinate"`
	YCoordinate        int32              `json:"y_coordinate"`
	MobileHeight       *int32             `json:"mobile_height"`
	MobileWidth        *int32             `json:"mobile_width"`
	MobileXCoordinate  *int32             `json:"mobile_x_coordinate"`
	MobileYCoordinate  *int32             `json:"mobile_y_coordinate"`
	CreatedAt          pgtype.Timestamptz `json:"created_at"`
	UpdatedAt          pgtype.Timestamptz `json:"updated_at"`
	Version            int32              `json:"version"`
}

type WebsiteConfig struct {
	ID            uuid.UUID          `json:"id"`
	WebsiteID     uuid.UUID          `json:"website_id"`
	DefaultPageID uuid.UUID          `json:"default_page_id"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
	Version       int32              `json:"version"`
}

type WebsiteContent struct {
	ID                        uuid.UUID          `json:"id"`
	WebsiteID                 uuid.UUID          `json:"website_id"`
	Locale                    string             `json:"locale"`
	WebsiteDisplayName        string             `json:"website_display_name"`
	WebsiteDisplayDescription *string            `json:"website_display_description"`
	CreatedAt                 pgtype.Timestamptz `json:"created_at"`
	UpdatedAt                 pgtype.Timestamptz `json:"updated_at"`
	Version                   int32              `json:"version"`
}

type WebsitePage struct {
	ID        uuid.UUID          `json:"id"`
	WebsiteID uuid.UUID          `json:"website_id"`
	UrlSlug   string             `json:"url_slug"`
	SortKey   string             `json:"sort_key"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	Version   int32              `json:"version"`
}

type WebsitePageContent struct {
	ID            uuid.UUID          `json:"id"`
	WebsitePageID uuid.UUID          `json:"website_page_id"`
	Locale        string             `json:"locale"`
	Title         string             `json:"title"`
	Subtitle      *string            `json:"subtitle"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
	Version       int32              `json:"version"`
}

type WebsiteSection struct {
	ID            uuid.UUID          `json:"id"`
	WebsiteID     uuid.UUID          `json:"website_id"`
	WebsitePageID uuid.UUID          `json:"website_page_id"`
	SortKey       string             `json:"sort_key"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
	Version       int32              `json:"version"`
}

type WebsiteSectionDisplay struct {
	ID               uuid.UUID          `json:"id"`
	WebsiteSectionID uuid.UUID          `json:"website_section_id"`
	RowCount         int32              `json:"row_count"`
	ImageID          *uuid.UUID         `json:"image_id"`
	CreatedAt        pgtype.Timestamptz `json:"created_at"`
	UpdatedAt        pgtype.Timestamptz `json:"updated_at"`
	Version          int32              `json:"version"`
}

type WebsiteStyle struct {
	ID        uuid.UUID          `json:"id"`
	WebsiteID uuid.UUID          `json:"website_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	Version   int32              `json:"version"`
}
