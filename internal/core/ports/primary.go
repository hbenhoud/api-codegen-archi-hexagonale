package ports

import (
	"backend-api/internal/core/entity"

	"github.com/google/uuid"
)

//go:generate mockgen -source=primary.go -destination=mocks/primary.gen.go -package=mocks

type ChannelService interface {
	Create(channel entity.Channel) (uuid.UUID, error)
	FindAll() ([]entity.Channel, error)
}
