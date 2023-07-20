package entity

import "github.com/google/uuid"

// Channel defines model for Channel.
type Channel struct {
	Id    uuid.UUID
	Input ChannelInput
	Name  string
	OnAir *bool
}

// ChannelInput defines model for ChannelInput.
type ChannelInput struct {
	OriginUrl string
	Source    string
}
