package service

import (
	"backend-api/internal/core/entity"
	"backend-api/internal/core/ports"

	"github.com/google/uuid"
)

var _ ports.ChannelService = (*ChannelCore)(nil)

type ChannelCore struct {
	channelRepo ports.ChannelRepository
}

func NewChannelCore(channelRepo ports.ChannelRepository) *ChannelCore {
	return &ChannelCore{
		channelRepo: channelRepo,
	}
}

func (core *ChannelCore) Create(channel entity.Channel) (uuid.UUID, error) {
	channel.Id = uuid.New()

	if err := core.channelRepo.Create(channel); err != nil {
		return uuid.Nil, err
	}

	return channel.Id, nil
}

func (core *ChannelCore) FindAll() ([]entity.Channel, error) {
	return core.channelRepo.FindAll()
}
