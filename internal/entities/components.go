package entities

import "github.com/google/uuid"

type WebsiteComponentEntity struct {
	ID             uuid.UUID
	PageID         uuid.UUID
	SortKey        string
	WebsiteID      uuid.UUID
	TextComponent  *WebsiteTextComponentEntity
	ImageComponent *WebsiteImageComponentEntity
}

type WebsiteTextComponentEntity struct {
	ID          uuid.UUID
	ComponentID uuid.UUID
	Text        string
}

type WebsiteImageComponentEntity struct {
	ID          uuid.UUID
	ComponentID uuid.UUID
	PhotoURL    string
}
