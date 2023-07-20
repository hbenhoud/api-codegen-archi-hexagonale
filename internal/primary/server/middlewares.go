package server

import (
	"backend-api/internal/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func requestValidatorMiddleware(opav validator.OpenAPI) MiddlewareFunc {
	return func(c *gin.Context) {
		if err := opav.ValidateRequest(c.Request.Context(), c.Request); err != nil {
			c.JSON(http.StatusBadRequest, Error{
				Message: err.Error(),
			})

			return
		}

		c.Next()
	}
}
