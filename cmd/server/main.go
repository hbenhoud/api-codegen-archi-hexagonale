package main

import (
	"backend-api/api"
	"backend-api/internal/primary/server"
	"backend-api/internal/validator"
	"context"
)

func main() {
	const addr = ":8080"
	var ctx = context.Background()

	validator := validator.NewKinOpenAPI(
		ctx, api.GetConfigSchema(),
	)

	server := server.New(addr, validator)

	server.ListenAndServe()
}
