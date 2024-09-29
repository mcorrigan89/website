package entities

import "github.com/google/uuid"

type WebsiteComponentEntity struct {
	ID             uuid.UUID
	SectionID      uuid.UUID
	WebsiteID      uuid.UUID
	Display        *WebsiteComponentDisplayEntity
	TextComponent  *WebsiteTextComponentEntity
	ImageComponent *WebsiteImageComponentEntity
}

type WebsiteComponentPositioningEntity struct {
	Xcoord *int32
	Ycoord *int32
	Height *int32
	Width  *int32
}

type WebsiteComponentDisplayEntity struct {
	ID                    uuid.UUID
	ComponentID           uuid.UUID
	MobilePositioning     *WebsiteComponentPositioningEntity
	FullScreenPositioning *WebsiteComponentPositioningEntity
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
