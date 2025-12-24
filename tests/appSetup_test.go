package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"github.com/tmazitov/fiberplus/internal/app"
)

type Database struct{}

func (d *Database) Test() string {
	return "Hello world!"
}

type Services struct {
	Database Database
}

// Helper function to create a test app instance
func setupTestApp() *app.App[Services] {
	services := &Services{
		Database: Database{},
	}

	testApp := app.NewApp(&app.AppConfig[Services]{
		Services: services,
	})

	testApp.Group("/test").
		Add(&app.Endpoint[Services]{Method: "POST", Route: "/in", Handler: &ReadHandlerExample{}})

	return testApp
}

// Helper function to make JSON requests
func makeJSONRequest(t *testing.T, app *fiber.App, method, url string, body interface{}) *http.Response {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		require.NoError(t, err)
	}

	req := httptest.NewRequest(method, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1) // -1 disables timeout
	require.NoError(t, err)

	return resp
}
