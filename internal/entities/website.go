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
	ID         uuid.UUID
	WebsiteID  uuid.UUID
	UrlSlug    string
	SortKey    string
	Title      *string
	Subtitle   *string
	Components []*WebsiteComponentEntity
}
