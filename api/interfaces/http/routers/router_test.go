package routers

import (
	"context"
	mock_middlewares "github.com/famiphoto/famiphoto/api/testing/mocks/interfaces/http/middlewares"
	mock_schema "github.com/famiphoto/famiphoto/api/testing/mocks/interfaces/http/schema"
	"regexp"
	"strings"
	"testing"

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

	// OpenAPIのパスパラメータ`{photoId}`形式をechoのパスパラメータ`:photoId`に変換する。
	openAPIPathToEchoPath := func(path string) string {
		re := regexp.MustCompile(`\{([^}]+)\}`)
		return re.ReplaceAllString(path, `:$1`)
	}

	t.Run("ルーティングのパスがOpenAPI定義と一致しているかテスト", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		e := echo.New()
		h := mock_schema.NewMockServerInterface(ctrl)
		r := &apiRouter{
			authMiddleware: mock_middlewares.NewMockAuthMiddleware(ctrl),
		}
		r.route(e, h)

		actual := getRegisteredRoutes(e)
		doc := loadOpenAPISpec(t, "../../../openapi/openapi-bundle.yaml")

		for path, pathItem := range doc.Paths.Map() {
			for method, operation := range pathItem.Operations() {
				method = strings.ToUpper(method)
				echoPath := openAPIPathToEchoPath(path)
				key := method + " " + echoPath
				if _, ok := actual[key]; !ok {
					t.Errorf("Missing route implementation: %s %s (operationId: %s)", method, path, operation.OperationID)
					t.Errorf("actual: %#v, %s", actual, key)
				}
			}
		}

	})
}
