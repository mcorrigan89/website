package api

import (
	"context"
	"errors"
	"sync"

	"connectrpc.com/connect"

	"github.com/google/uuid"
	websitev1 "github.com/mcorrigan89/website/gen/serviceapis/website/v1"
	"github.com/mcorrigan89/website/internal/config"
	"github.com/mcorrigan89/website/internal/entities"
	"github.com/mcorrigan89/website/internal/services"

	"github.com/rs/zerolog"
)

type WebsiteServerV1 struct {
	config   *config.Config
	wg       *sync.WaitGroup
	logger   *zerolog.Logger
	services *services.Services
}

func newWebsiteProtoUrlServer(cfg *config.Config, logger *zerolog.Logger, wg *sync.WaitGroup, services *services.Services) *WebsiteServerV1 {
	return &WebsiteServerV1{
		config:   cfg,
		wg:       wg,
		logger:   logger,
		services: services,
	}
}

func (s *WebsiteServerV1) WebsiteByHandle(ctx context.Context, req *connect.Request[websitev1.WebsiteByHandleRequest]) (*connect.Response[websitev1.WebsiteByHandleResponse], error) {
	handle := req.Msg.Handle
	locale := req.Msg.Locale

	if handle == "" {
		err := errors.New("handle is empty")
		s.logger.Err(err).Ctx(ctx).Msg("Handle is empty")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	website, err := s.services.WebsiteService.GetWebsiteByHandle(ctx, handle, locale)
	if err != nil {
		s.logger.Err(err).Ctx(ctx).Msg("Error sending verification email")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	websiteMessage := s.serializeWebsite(website)

	res := connect.NewResponse(&websitev1.WebsiteByHandleResponse{
		Website: websiteMessage,
	})
	res.Header().Set("Website-Version", "v1")
	return res, nil
}

func (s *WebsiteServerV1) CreateWebsite(ctx context.Context, req *connect.Request[websitev1.CreateWebsiteRequest]) (*connect.Response[websitev1.CreateWebsiteResponse], error) {
	handle := req.Msg.Handle
	locale := req.Msg.Locale

	if handle == "" {
		err := errors.New("handle is empty")
		s.logger.Err(err).Ctx(ctx).Msg("Handle is empty")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	website, err := s.services.WebsiteService.CreateWebsite(ctx, services.CreateWebsiteArgs{
		Handle: handle,
		Locale: locale,
	})
	if err != nil {
		s.logger.Err(err).Ctx(ctx).Msg("Error sending verification email")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	websiteMessage := s.serializeWebsite(website)

	res := connect.NewResponse(&websitev1.CreateWebsiteResponse{
		Website: websiteMessage,
	})
	res.Header().Set("Website-Version", "v1")
	return res, nil
}

func (s *WebsiteServerV1) serializeWebsite(websiteEntity *entities.WebsiteEntity) *websitev1.Website {
	websitePages := []*websitev1.WebsitePage{}

	for _, page := range websiteEntity.Pages {

		sectionsAndComponents := s.createSectionsAndComponents(page)

		p := websitev1.WebsitePage{
			Id:       page.ID.String(),
			Title:    page.Title,
			Subtitle: page.Subtitle,
			UrlSlug:  page.UrlSlug,
			Sections: sectionsAndComponents,
		}

		websitePages = append(websitePages, &p)
	}

	website := websitev1.Website{
		Id:          websiteEntity.ID.String(),
		Handle:      websiteEntity.Handle,
		Name:        websiteEntity.DisplayName,
		Description: websiteEntity.DisplayDescription,
		Pages:       websitePages,
		Config: &websitev1.WebsiteConfig{
			Id:                 websiteEntity.Config.ID.String(),
			DefaultPageId:      websiteEntity.Config.DefaultPageID.String(),
			DefaultPageUrlSlug: websiteEntity.GetDefaultPage().UrlSlug,
		},
		Styles: &websitev1.WebsiteStyles{
			Id: websiteEntity.Styles.ID.String(),
			Palette: &websitev1.Palette{
				Id:         websiteEntity.Styles.Palette.ID.String(),
				ColorOne:   websiteEntity.Styles.Palette.Color1,
				ColorTwo:   websiteEntity.Styles.Palette.Color2,
				ColorThree: websiteEntity.Styles.Palette.Color3,
				ColorFour:  websiteEntity.Styles.Palette.Color4,
				ColorFive:  websiteEntity.Styles.Palette.Color5,
				ColorSix:   websiteEntity.Styles.Palette.Color6,
			},
		},
	}

	return &website
}

func (s *WebsiteServerV1) createSectionsAndComponents(page *entities.WebsitePageEntity) []*websitev1.WebsiteSection {

	sections := []*websitev1.WebsiteSection{}
	for _, section := range page.Sections {

		components := []*websitev1.WebsiteComponent{}

		for _, component := range section.Components {
			if component.TextComponent != nil {
				var jsonString string
				if component.TextComponent.Json != nil {
					jsonString = string(component.TextComponent.Json)
				}
				c := websitev1.WebsiteComponent{
					Id: component.ID.String(),
					ComponentContent: &websitev1.WebsiteComponent_TextComponent{
						TextComponent: &websitev1.TextComponent{
							Json: &jsonString,
							Html: component.TextComponent.Html,
						},
					},
				}

				components = append(components, &c)
			}

			if component.ImageComponent != nil {
				c := websitev1.WebsiteComponent{
					Id: component.ID.String(),
					Display: &websitev1.WebsiteComponentDisplay{
						Id: component.Display.ID.String(),
						FullScreen: &websitev1.WebsiteComponentPositioning{
							XCoordinate: component.Display.FullScreenPositioning.Xcoord,
							YCoordinate: component.Display.FullScreenPositioning.Ycoord,
							Width:       component.Display.FullScreenPositioning.Width,
							Height:      component.Display.FullScreenPositioning.Height,
						},
						MobileScreen: &websitev1.WebsiteComponentPositioning{
							XCoordinate: component.Display.FullScreenPositioning.Xcoord,
							YCoordinate: component.Display.FullScreenPositioning.Ycoord,
							Width:       component.Display.FullScreenPositioning.Width,
							Height:      component.Display.FullScreenPositioning.Height,
						},
					},
					ComponentContent: &websitev1.WebsiteComponent_ImageComponent{
						ImageComponent: &websitev1.ImageComponent{
							Url: component.ImageComponent.PhotoURL,
						},
					},
				}

				components = append(components, &c)
			}
		}

		s := websitev1.WebsiteSection{
			Id:         section.ID.String(),
			Components: components,
		}

		sections = append(sections, &s)
	}

	return sections
}

func (s *WebsiteServerV1) CreateWebsitePage(ctx context.Context, req *connect.Request[websitev1.CreateWebsitePageRequest]) (*connect.Response[websitev1.CreateWebsitePageResponse], error) {
	websiteId := req.Msg.WebsiteId
	title := req.Msg.Title
	subtitle := req.Msg.Subtitle
	urlSlug := req.Msg.UrlSlug

	if websiteId == "" {
		err := errors.New("website_id is empty")
		s.logger.Err(err).Ctx(ctx).Msg("WebsiteId is empty")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	websiteUuid, err := uuid.Parse(websiteId)
	if err != nil {
		s.logger.Err(err).Ctx(ctx).Msg("Error parsing website_id")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	if title == "" {
		err := errors.New("title is empty")
		s.logger.Err(err).Ctx(ctx).Msg("Title is empty")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	websitePage, err := s.services.WebsitePageService.CreateWebsitePage(ctx, services.CreateWebsitePageArgs{
		WebsiteID: websiteUuid,
		Title:     title,
		UrlSlug:   urlSlug,
		Subtitle:  subtitle,
	})
	if err != nil {
		s.logger.Err(err).Ctx(ctx).Msg("Error creating website page")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&websitev1.CreateWebsitePageResponse{
		Page: &websitev1.WebsitePage{
			Id:       websitePage.ID.String(),
			Title:    websitePage.Title,
			Subtitle: websitePage.Subtitle,
			UrlSlug:  websitePage.UrlSlug,
		},
	})
	res.Header().Set("Website-Version", "v1")
	return res, nil
}

func (s *WebsiteServerV1) UpdateWebsitePage(ctx context.Context, req *connect.Request[websitev1.UpdateWebsitePageRequest]) (*connect.Response[websitev1.UpdateWebsitePageResponse], error) {
	id := req.Msg.Id
	title := req.Msg.Title
	subtitle := req.Msg.Subtitle
	urlSlug := req.Msg.UrlSlug

	if id == "" {
		err := errors.New("id is empty")
		s.logger.Err(err).Ctx(ctx).Msg("ID is empty")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	Uuid, err := uuid.Parse(id)
	if err != nil {
		s.logger.Err(err).Ctx(ctx).Msg("Error parsing website_id")
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	websitePage, err := s.services.WebsitePageService.UpdateWebsitePage(ctx, services.UpdateWebsitePageArgs{
		ID:       Uuid,
		Title:    &title,
		UrlSlug:  &urlSlug,
		Subtitle: subtitle,
	})
	if err != nil {
		s.logger.Err(err).Ctx(ctx).Msg("Error updating website page")
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	res := connect.NewResponse(&websitev1.UpdateWebsitePageResponse{
		Page: &websitev1.WebsitePage{
			Id:       websitePage.ID.String(),
			Title:    websitePage.Title,
			Subtitle: websitePage.Subtitle,
			UrlSlug:  websitePage.UrlSlug,
		},
	})
	res.Header().Set("Website-Version", "v1")
	return res, nil
}
