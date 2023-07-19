package validator

//go:generate mockgen -source=openapi.go -destination=mocks/openapi.gen.go -package=mocks

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
)

var (
	// ErrValidateRequest returned when http request is not valid.
	ErrValidateRequest = errors.New("http request is not valid")

	// ErrValidateResponse returned when http response is not valid.
	ErrValidateResponse = errors.New("http response is not valid")
)

// OpenAPI interfaces the Open API validation.
type OpenAPI interface {
	ValidateRequest(ctx context.Context, req *http.Request) error
	ValidateResponse(ctx context.Context, req *http.Request, resp *http.Response) error
}

// KinOpenAPI implements the Open API validation
// based on Kin-openapi library.
type KinOpenAPI struct {
	router routers.Router
}

// NewKinOpenAPI contructs a new Kin open API validator.
func NewKinOpenAPI(ctx context.Context, schemaContent []byte) *KinOpenAPI {
	loader := openapi3.NewLoader()

	schema, err := loader.LoadFromData(schemaContent)
	if err != nil {
		panic(err)
	}

	if err := schema.Validate(ctx); err != nil {
		panic(err)
	}

	router, err := gorillamux.NewRouter(schema)
	if err != nil {
		panic(err)
	}

	validator := &KinOpenAPI{
		router: router,
	}

	return validator
}

// ValidateRequest validates the http request.
func (kin *KinOpenAPI) ValidateRequest(ctx context.Context, req *http.Request) error {
	route, pathParams, err := kin.router.FindRoute(req)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrValidateRequest, err)
	}

	requestValidateInput := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: pathParams,
		Route:      route,
	}

	if err := openapi3filter.ValidateRequest(ctx, requestValidateInput); err != nil {
		return fmt.Errorf("%w: %s", ErrValidateRequest, err)
	}

	return nil
}

// ValidateResponse validates the http response.
func (kin *KinOpenAPI) ValidateResponse(ctx context.Context, req *http.Request, resp *http.Response) error {
	route, pathParams, err := kin.router.FindRoute(req)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrValidateResponse, err)
	}

	requestValidateInput := &openapi3filter.RequestValidationInput{
		Request:    req,
		PathParams: pathParams,
		Route:      route,
	}

	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidateInput,
		Status:                 resp.StatusCode,
		Header:                 resp.Header,
		Body:                   resp.Body,
	}

	if err := openapi3filter.ValidateResponse(ctx, responseValidationInput); err != nil {
		return fmt.Errorf("%w: %s", ErrValidateResponse, err)
	}

	return nil
}
