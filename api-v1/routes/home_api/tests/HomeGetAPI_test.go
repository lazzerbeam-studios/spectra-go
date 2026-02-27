package home_tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"

	"api-go/routes/home_api"
)

type ResponseOutput struct {
	Message string `json:"message"`
}

func Test_HomeGetAPI(t *testing.T) {
	_, api := humatest.New(t)

	home_api.Register(api)
	response := api.Get("/")

	var responseOutput ResponseOutput
	err := json.Unmarshal(response.Body.Bytes(), &responseOutput)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %s", err.Error())
	}

	if !strings.EqualFold(responseOutput.Message, "Welcome to Photon") {
		t.Fatalf("Unexpected response: %s", response.Body.String())
	}
}
