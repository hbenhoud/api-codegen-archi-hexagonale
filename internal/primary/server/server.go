package server

import (
	"backend-api/internal/core/ports"
	"backend-api/internal/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:generate oapi-codegen -generate types -package server -o ./types.gen.go ../../../api/configuration.swagger.yaml
//go:generate oapi-codegen -generate gin -package server -o ./server.gen.go ../../../api/configuration.swagger.yaml

var (
	_      ServerInterface = (*Handlers)(nil)
	router                 = gin.New()
)

type Handlers struct {
	services CoreServices
}

type CoreServices struct {
	ChannelService ports.ChannelService
}

func New(services CoreServices, addr string, opaValidator validator.OpenAPI) http.Server {

	router.Use(
		requestValidatorMiddleware(opaValidator),
		responseValidatorMiddleware(opaValidator),
	)

	return http.Server{
		Addr: addr,
		Handler: RegisterHandlersWithOptions(
			router, &Handlers{
				services: services,
			},
			GinServerOptions{},
		),
	}
}
