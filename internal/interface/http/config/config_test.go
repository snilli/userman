package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig_MissingEnv(t *testing.T) {
	os.Clearenv()
	os.Setenv("GO_ENV", "test")
	defer os.Clearenv()

	_, err := NewConfig()

	assert.Error(t, err)
}

func TestNewConfig_WithEnv(t *testing.T) {
	os.Setenv("GO_ENV", "production")

	envVars := map[string]string{
		"DB_MONGO_HOST":     "localhost",
		"DB_MONGO_USERNAME": "user",
		"DB_MONGO_PASSWORD": "pass",
		"DB_MONGO_DATABASE": "testdb",
		"DB_MONGO_PORT":     "27017",
		"DB_MONGO_URI":      "mongodb://user:pass@localhost:27017/testdb",
		"HTTP_PORT":         "8080",
		"JWT_SECRET_KEY":    "secret",
		"JWT_EXPIRE_SEC":    "3600",
	}

	for key, value := range envVars {
		os.Setenv(key, value)
	}
	defer os.Clearenv()

	cfg, err := NewConfig()

	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "localhost", cfg.MongoDb.Host)
	assert.Equal(t, 8080, cfg.Http.Port)
	assert.Equal(t, "secret", cfg.Jwt.SecretKey)
}

func TestNewConfig_ProductionEnv(t *testing.T) {
	os.Setenv("GO_ENV", "production")
	defer os.Unsetenv("GO_ENV")

	_, err := NewConfig()

	assert.Error(t, err)
}
