package tests

import (
	"io"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func startServer(t *testing.T, port string) *exec.Cmd {
	cmd := exec.Command("go", "run", "../cmd/shinyvanity/main.go")
	cmd.Env = append(os.Environ(),
		"VANITY_DOMAIN=example.com",
		"VANITY_ORGANIZATION=testorg",
		"LISTEN_PORT="+port,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start server: %v", err)
	}
	// Wait for server to start
	for i := 0; i < 10; i++ {
		resp, err := http.Get("http://localhost:" + port + "/healthz")
		if err == nil {
			errClose := resp.Body.Close()
			if errClose != nil {
				t.Fatalf("failed to close response body: %v", errClose)
			}
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	return cmd
}

func TestIntegration_GoGet(t *testing.T) {
	port := "8090"
	cmd := startServer(t, port)
	defer func() {
		err := cmd.Process.Kill()
		if err != nil {
			t.Fatalf("failed to kill process: %v", err)
		}
	}()

	resp, err := http.Get("http://localhost:" + port + "/foo/bar?go-get=1")
	assert.NoError(t, err, "http get should not fail")
	if err != nil {
		t.FailNow()
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			t.Fatalf("failed to close response body: %v", err)
		}
	}()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "status code should be 200")
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err, "should read response body without error")
	if err != nil {
		t.FailNow()
	}
	assert.Contains(t, string(body), `meta name="go-import" content="example.com/foo/bar git https://github.com/testorg/foo/bar"`, "body should contain go-import meta tag")
}

func TestIntegration_Redirect(t *testing.T) {
	port := "8091"
	cmd := startServer(t, port)
	defer func() {
		err := cmd.Process.Kill()
		if err != nil {
			t.Fatalf("failed to kill process: %v", err)
		}
	}()

	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
	resp, err := client.Get("http://localhost:" + port + "/foo/bar")
	assert.NoError(t, err, "http get should not fail")
	if err != nil {
		t.FailNow()
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			t.Fatalf("failed to close response body: %v", err)
		}
	}()
	assert.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode, "status code should be 307")
	loc := resp.Header.Get("Location")
	assert.Equal(t, "https://github.com/testorg/foo/bar", loc, "should redirect to repo")
}
