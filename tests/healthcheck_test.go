package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jakottelaar/gobookreviewapp/api"
	"github.com/jakottelaar/gobookreviewapp/config"
	"github.com/jakottelaar/gobookreviewapp/pkg/database"
	"github.com/stretchr/testify/assert"
)

var (
	testServer           *httptest.Server
	baseAuthEndpointURL  string
	baseBooksEndpointURL string
	baseUserEndpointURL  string
)

func TestMain(m *testing.M) {

	cfg, err := config.Load()

	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	err = database.Initialize(cfg)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer database.Close()

	routes := api.SetupRoutes()

	testServer = httptest.NewServer(routes)

	baseAuthEndpointURL = testServer.URL + "/v1/api/auth/"
	baseBooksEndpointURL = testServer.URL + "/v1/api/books/"
	baseUserEndpointURL = testServer.URL + "/v1/api/users/"

	m.Run()

}

func TestHealthCheck(t *testing.T) {

	req, err := http.NewRequest("GET", testServer.URL+"/health", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Could not send request: %v", err)
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var response map[string]string
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Could not decode response: %v", err)
	}

	assert.Equal(t, "Health Check OK", response["message"])
}
