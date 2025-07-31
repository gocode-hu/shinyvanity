package main__test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var server *httptest.Server

func TestMain(m *testing.M) {
	server = httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	os.Exit(m.Run())
}

func TestHelloWorldHandler(t *testing.T) {
	// Create a test HTTP server with your handler
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	// Make an HTTP GET request
	resp, err := http.Get(server.URL)
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Read the body
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	// Assertions
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "Hello, World!\n", string(body))
}
