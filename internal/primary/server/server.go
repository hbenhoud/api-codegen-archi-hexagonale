package server

import (
	"backend-api/internal/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:generate oapi-codegen -generate types -package server -o ./types.gen.go ../../../api/configuration.swagger.yaml
//go:generate oapi-codegen -generate gin -package server -o ./server.gen.go ../../../api/configuration.swagger.yaml

var (
	_      ServerInterface = (*Handlers)(nil)
	router                 = gin.New()
)

type Handlers struct{}

func newHandlers() *Handlers {
	return &Handlers{}
}

func New(addr string, opaValidator validator.OpenAPI) http.Server {
	return http.Server{
		Addr: addr,
		Handler: RegisterHandlersWithOptions(router, newHandlers(), GinServerOptions{
			Middlewares: []MiddlewareFunc{
				requestValidatorMiddleware(opaValidator),
			},
		}),
	}
}
