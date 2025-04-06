package routers

import (
	"context"
	"strings"
	"testing"

	mock_routers "github.com/famiphoto/famiphoto/api/testing/mocks/interfaces/http/routers"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func TestApiRouter_route(t *testing.T) {
	loadOpenAPISpec := func(t *testing.T, path string) *openapi3.T {
		loader := openapi3.NewLoader()
		doc, err := loader.LoadFromFile(path)
		if err != nil {
			t.Fatalf("failed to load OpenAPI spec: %v", err)
		}
		if err := doc.Validate(context.Background()); err != nil {
			t.Fatalf("failed to validate OpenAPI spec: %v", err)
		}
		return doc
	}

	getRegisteredRoutes := func(e *echo.Echo) map[string]string {
		routes := make(map[string]string) // key: METHOD PATH, value: name
		for _, r := range e.Routes() {
			key := r.Method + " " + r.Path
			routes[key] = r.Name
		}
		return routes
	}

	openAPIPathToEchoPath := func(path string) string {
		return strings.ReplaceAll(path, "{", ":")
	}

	t.Run("ルーティングのパスがOpenAPI定義と一致しているかテスト", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish(t)

		e := echo.New()
		handler := mock_routers.NewMockServerInterface(t)
		r := apiRouter{}
		r.route(e, handler)

		actual := getRegisteredRoutes(e)
		doc := loadOpenAPISpec(t, "openapi/openapi-bundle.yaml")

		for path, pathItem := range doc.Paths {
			for method, operation := range pathItem.Operations() {
				method = strings.ToUpper(method)
				echoPath := openAPIPathToEchoPath(path)
				key := method + " " + echoPath
				if _, ok := actual[key]; !ok {
					t.Errorf("Missing route implementation: %s %s (operationId: %s)", method, path, operation.OperationID)
				}
			}
		}

	})
}
