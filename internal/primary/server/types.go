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

func toServerChannel(channel entity.Channel) Channel {
	return Channel{
		Id:    &channel.Id,
		Name:  channel.Name,
		Input: toServerChannelInput(channel.Input),
		OnAir: channel.OnAir,
	}
}

func toServerChannelInput(input entity.ChannelInput) ChannelInput {
	return ChannelInput{
		Source:    ChannelInputSource(input.Source),
		OriginUrl: input.OriginUrl,
	}
}

func toServerChannels(channels []entity.Channel) []Channel {
	result := make([]Channel, 0, len(channels))

	for _, channel := range channels {
		result = append(result, toServerChannel(channel))
	}

	return result
}
