package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handlers) GetChannels(c *gin.Context) {
	channels, err := h.services.ChannelService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	c.JSON(
		http.StatusOK,
		toServerChannels(channels),
	)
}

func (h Handlers) PostChannels(c *gin.Context) {
	var body Channel

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})
		return
	}

	uid, err := h.services.ChannelService.Create(
		toEntityChannel(body),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Error{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, ChannelID{
		Id: &uid,
	})
}
