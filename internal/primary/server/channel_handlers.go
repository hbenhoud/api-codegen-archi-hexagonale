package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handlers) GetChannels(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (h Handlers) PostChannels(c *gin.Context) {
	c.Status(http.StatusCreated)
}
