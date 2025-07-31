package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_EnvPort(t *testing.T) {
	err := os.Setenv("LISTEN_PORT", "8080")
	if err != nil {
		t.Fatalf("failed to set env: %v", err)
	}
	port := os.Getenv("LISTEN_PORT")
	assert.Equal(t, "8080", port, "LISTEN_PORT should be '8080'")
}

func TestMain_NoEnvFile(t *testing.T) {
	err := os.Unsetenv("LISTEN_PORT")
	if err != nil {
		t.Fatalf("failed to unset env: %v", err)
	}
	port := os.Getenv("LISTEN_PORT")
	assert.Empty(t, port, "LISTEN_PORT should be empty")
}
