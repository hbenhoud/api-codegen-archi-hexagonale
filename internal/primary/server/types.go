package server

import "backend-api/internal/core/entity"

func toEntityChannel(channel Channel) entity.Channel {
	return entity.Channel{
		Name: channel.Name,
		Input: entity.ChannelInput{
			Source:    string(channel.Input.Source),
			OriginUrl: channel.Input.OriginUrl,
		},
		OnAir: channel.OnAir,
	}
}
