package entities

import "github.com/google/uuid"

type WebsiteEntity struct {
	ID                 uuid.UUID
	Handle             string
	DisplayName        string
	DisplayDescription *string
	Config             *WebsiteConfigEntity
	Styles             *WebsiteStylesEntity
	Pages              []*WebsitePageEntity
}

func (w *WebsiteEntity) GetDefaultPage() *WebsitePageEntity {
	for _, page := range w.Pages {
		if page.ID == w.Config.DefaultPageID {
			return page
		}
	}

	return nil
}

type WebsiteConfigEntity struct {
	ID            uuid.UUID
	WebsiteID     uuid.UUID
	DefaultPageID uuid.UUID
}

type PaletteEntity struct {
	ID        uuid.UUID
	WebsiteID uuid.UUID
	Color1    string
	Color2    string
	Color3    string
	Color4    string
	Color5    string
	Color6    string
}

type WebsiteStylesEntity struct {
	ID        uuid.UUID
	WebsiteID uuid.UUID
	Palette   *PaletteEntity
}

type WebsitePageEntity struct {
	ID        uuid.UUID
	WebsiteID uuid.UUID
	UrlSlug   string
	SortKey   string
	Title     string
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

type WebsiteSectionDisplayEntity struct {
	ID       uuid.UUID
	RowCount int32
	ImageID  uuid.UUID
}
