package ports

import "backend-api/internal/core/entity"

//go:generate mockgen -source=secondary.go -destination=mocks/secondary.gen.go -package=mocks

type ChannelRepository interface {
	Create(channel entity.Channel) error
	FindAll() ([]entity.Channel, error)
}
