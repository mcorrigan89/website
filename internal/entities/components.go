package entities

import "github.com/google/uuid"

type WebsiteComponentEntity struct {
	ID             uuid.UUID
	SectionID      uuid.UUID
	WebsiteID      uuid.UUID
	TextComponent  *WebsiteTextComponentEntity
	ImageComponent *WebsiteImageComponentEntity
}

type WebsiteTextComponentEntity struct {
	ID          uuid.UUID
	ComponentID uuid.UUID
	Json        []byte
	Html        *string
}

type WebsiteImageComponentEntity struct {
	ID          uuid.UUID
	ComponentID uuid.UUID
	PhotoURL    string
}
