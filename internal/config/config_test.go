package config

import (
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("CONFIG_FILE", "../../cmd/config/local.yaml")
	os.Setenv("ENCRYPTION_SECRET_KEY", "12345678901234567890123456789012")
	cfg := GetConfig()
	if cfg == nil {
		t.Errorf("Expected config to be loaded")
	}
	assert.Equal(t, "8080", cfg.Server.Port)
}
