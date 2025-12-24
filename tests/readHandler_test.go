package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tmazitov/fiberplus"
	"github.com/tmazitov/fiberplus/internal/behavior"
)

type ReadHandlerRequest struct {
	Field string `json:"field" validate:"required,min=5,max=20"`
}

type ReadHandlerExample struct {
	behavior.ReadHandler[Services, *ReadHandlerRequest]
}

type ReadHandlerResponse struct {
	Request *ReadHandlerRequest `json:"request"`
	Message string              `json:"message"`
}

func (h *ReadHandlerExample) Handle(app *fiberplus.App[Services]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		input := ctx.Locals("Input").(*ReadHandlerRequest)

		return ctx.JSON(&ReadHandlerResponse{Request: input, Message: "ok!"})
	}
}

func TestReadHandlerExample_TableDriven(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    ReadHandlerRequest
		expectedStatus int
		checkResponse  func(t *testing.T, resp *http.Response)
	}{
		{
			name: "valid request",
			requestBody: ReadHandlerRequest{
				Field: "validfield",
			},
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, resp *http.Response) {
				var response ReadHandlerResponse
				err := json.NewDecoder(resp.Body).Decode(&response)
				require.NoError(t, err)

				assert.Equal(t, "validfield", response.Request.Field)
				assert.Equal(t, "ok!", response.Message)
			},
		},
		{
			name: "field too short",
			requestBody: ReadHandlerRequest{
				Field: "abc",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "field too long",
			requestBody: ReadHandlerRequest{
				Field: "thisisfieldthatiswaytoolongforvalidation",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "empty field",
			requestBody: ReadHandlerRequest{
				Field: "",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "minimum length boundary",
			requestBody: ReadHandlerRequest{
				Field: "12345", // Exactly 5 characters
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "maximum length boundary",
			requestBody: ReadHandlerRequest{
				Field: "12345678901234567890", // Exactly 20 characters
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testApp := setupTestApp()

			resp := makeJSONRequest(t, testApp.Core(), "POST", "/test/in", tt.requestBody)
			defer resp.Body.Close()

			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			if tt.checkResponse != nil {
				tt.checkResponse(t, resp)
			}
		})
	}
}
