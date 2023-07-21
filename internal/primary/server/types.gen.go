// Package server provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package server

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Defines values for ChannelInputSource.
const (
	LIVE ChannelInputSource = "LIVE"
	VOD  ChannelInputSource = "VOD"
)

// Channel defines model for Channel.
type Channel struct {
	Id    *openapi_types.UUID `json:"id,omitempty"`
	Input ChannelInput        `json:"input"`

	// Name Channel configuration unique name. At creation, if not provided it will be generated.
	Name string `json:"name"`

	// OnAir If enabled, channel is available for end-users.
	OnAir *bool `json:"on_air,omitempty"`
}

// ChannelID defines model for ChannelID.
type ChannelID struct {
	Id *openapi_types.UUID `json:"id,omitempty"`
}

// ChannelInput defines model for ChannelInput.
type ChannelInput struct {
	// OriginUrl Origin URL of the input.
	OriginUrl string `json:"origin_url"`

	// Source Origin of the input.
	Source ChannelInputSource `json:"source"`
}

// ChannelInputSource Origin of the input.
type ChannelInputSource string

// Channels defines model for Channels.
type Channels = []Channel

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// PostChannelsJSONRequestBody defines body for PostChannels for application/json ContentType.
type PostChannelsJSONRequestBody = Channel
