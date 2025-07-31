package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVanityHandler_GoGetRequest(t *testing.T) {
	err := os.Setenv("VANITY_DOMAIN", "example.com")
	if err != nil {
		t.Fatalf("failed to set VANITY_DOMAIN: %v", err)
	}
	err = os.Setenv("VANITY_ORGANIZATION", "testorg")
	if err != nil {
		t.Fatalf("failed to set VANITY_ORGANIZATION: %v", err)
	}

	req := httptest.NewRequest("GET", "/foo/bar?go-get=1", nil)
	rw := httptest.NewRecorder()

	VanityHandler(rw, req)

	resp := rw.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "status code should be 200")

	body := rw.Body.String()
	assert.Contains(t, body, `<meta name="go-import" content="example.com/foo/bar git https://github.com/testorg/foo/bar">`, "body should contain go-import meta tag")
}

func TestVanityHandler_NonGoGetRequest(t *testing.T) {
	err := os.Setenv("VANITY_DOMAIN", "example.com")
	if err != nil {
		t.Fatalf("failed to set VANITY_DOMAIN: %v", err)
	}
	err = os.Setenv("VANITY_ORGANIZATION", "testorg")
	if err != nil {
		t.Fatalf("failed to set VANITY_ORGANIZATION: %v", err)
	}

	req := httptest.NewRequest("GET", "/foo/bar", nil)
	rw := httptest.NewRecorder()

	VanityHandler(rw, req)

	resp := rw.Result()
	assert.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode, "status code should be 307")
	loc := resp.Header.Get("Location")
	assert.Equal(t, "https://github.com/testorg/foo/bar", loc, "should redirect to repo")
}
