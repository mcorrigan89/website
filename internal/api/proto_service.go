package api

import (
	"net/http"
	"sync"

	"connectrpc.com/grpcreflect"
	websitev1connect "github.com/mcorrigan89/website/gen/serviceapis/website/v1/websitev1connect"
	"github.com/mcorrigan89/website/internal/config"
	"github.com/mcorrigan89/website/internal/services"
	"github.com/rs/zerolog"
)

type ProtoServer struct {
	config          *config.Config
	wg              *sync.WaitGroup
	logger          *zerolog.Logger
	services        *services.Services
	websiteV1Server *WebsiteServerV1
}

func NewProtoServer(cfg *config.Config, logger *zerolog.Logger, wg *sync.WaitGroup, services *services.Services) *ProtoServer {

	websiteV1Server := newWebsiteProtoUrlServer(cfg, logger, wg, services)

	return &ProtoServer{
		config:          cfg,
		wg:              wg,
		logger:          logger,
		services:        services,
		websiteV1Server: websiteV1Server,
	}
}

func (s *ProtoServer) Handle(r *http.ServeMux) {

	reflector := grpcreflect.NewStaticReflector(
		"serviceapis.website.v1.WebsiteService",
	)

	reflectPath, reflectHandler := grpcreflect.NewHandlerV1(reflector)
	r.Handle(reflectPath, reflectHandler)
	reflectPathAlpha, reflectHandlerAlpha := grpcreflect.NewHandlerV1Alpha(reflector)
	r.Handle(reflectPathAlpha, reflectHandlerAlpha)

	websiteV1Path, websiteV1Handle := websitev1connect.NewWebsiteServiceHandler(s.websiteV1Server)
	r.Handle(websiteV1Path, websiteV1Handle)
}
