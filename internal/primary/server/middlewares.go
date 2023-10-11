package server

import (
	"backend-api/internal/pkg/validator"
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func requestValidatorMiddleware(opav validator.OpenAPI) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := opav.ValidateRequest(c.Request.Context(), c.Request); err != nil {
			c.JSON(http.StatusBadRequest, Error{
				Message: err.Error(),
			})

			c.Abort()
		}

		c.Next()
	}
}

func responseValidatorMiddleware(opav validator.OpenAPI) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer = &cutomWriter{
			ResponseWriter: c.Writer,
			opav:           opav,
			ctx:            c,
		}

		c.Next()
	}
}

type cutomWriter struct {
	gin.ResponseWriter
	opav validator.OpenAPI
	ctx  *gin.Context
}

func (w *cutomWriter) Write(b []byte) (int, error) {
	if err := w.opav.ValidateResponse(w.ctx.Request.Context(), w.ctx.Request, &http.Response{
		StatusCode: w.Status(),
		Body:       io.NopCloser(bytes.NewReader(b)),
		Header:     w.ctx.Writer.Header(),
	}); err != nil {
		w.ResponseWriter.WriteHeader(http.StatusInternalServerError)

		b, _ := json.Marshal(Error{
			Message: err.Error(),
		})

		return w.ResponseWriter.Write(b)
	}

	return w.ResponseWriter.Write(b)
}
