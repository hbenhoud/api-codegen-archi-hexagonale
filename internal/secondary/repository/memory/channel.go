package memory

import (
	"backend-api/internal/core/entity"
	"backend-api/internal/core/ports"
)

var _ ports.ChannelRepository = (*ChannelMemory)(nil)

type ChannelMemory struct {
	channels []entity.Channel
}

func NewChannel() *ChannelMemory {
	return &ChannelMemory{
		channels: []entity.Channel{},
	}
}

func (memory *ChannelMemory) Create(channel entity.Channel) error {
	memory.channels = append(memory.channels, channel)
	return nil
}

func (memory *ChannelMemory) FindAll() ([]entity.Channel, error) {
	return memory.channels, nil
}
