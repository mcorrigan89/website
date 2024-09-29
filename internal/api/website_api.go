package api

import (
	"context"
	"errors"
	"sync"

	"connectrpc.com/connect"

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

	websitePages := []*websitev1.WebsitePage{}

	for _, page := range website.Pages {

		sectionsAndComponents := s.createSectionsAndComponents(page)

		p := websitev1.WebsitePage{
			Id:       page.ID.String(),
			Title:    page.Title,
			Subtitle: page.Subtitle,
			Sections: sectionsAndComponents,
		}

		websitePages = append(websitePages, &p)
	}

	res := connect.NewResponse(&websitev1.WebsiteByHandleResponse{
		Website: &websitev1.Website{
			Id:          website.ID.String(),
			Handle:      website.Handle,
			Name:        website.DisplayName,
			Description: website.DisplayDescription,
			Pages:       websitePages,
		},
	})
	res.Header().Set("Website-Version", "v1")
	return res, nil
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
