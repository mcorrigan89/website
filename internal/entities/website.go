package entities

import "github.com/google/uuid"

type WebsiteEntity struct {
	ID                 uuid.UUID
	Handle             string
	DisplayName        *string
	DisplayDescription *string
	Pages              []*WebsitePageEntity
}

type WebsitePageEntity struct {
	ID        uuid.UUID
	WebsiteID uuid.UUID
	UrlSlug   string
	SortKey   string
	Title     *string
	Subtitle  *string
	Sections  []*WebsitePageSectionEntity
}

type WebsitePageSectionEntity struct {
	ID        uuid.UUID
	WebsiteID uuid.UUID
	PageID    uuid.UUID
	SortKey   string

	RowCount int32

	Components []*WebsiteComponentEntity
}
