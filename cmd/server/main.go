package main

import (
	"backend-api/api"
	"backend-api/internal/core/service"
	"backend-api/internal/pkg/validator"
	"backend-api/internal/primary/server"
	"backend-api/internal/secondary/repository/memory"
	"context"
)

func main() {
	const addr = ":8080"
	var ctx = context.Background()

	// init validator with config schema.
	validator := validator.NewKinOpenAPI(
		ctx, api.GetConfigSchema(),
	)

	// init memory repo.
	memoryRepo := memory.NewChannelMemory()

	// init core services.
	services := server.CoreServices{
		ChannelService: service.NewChannelCore(memoryRepo),
	}

	// init http server.
	server := server.New(services, addr, validator)
	defer server.Close()

	server.ListenAndServe()
}
