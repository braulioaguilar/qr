package cors

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	c "github.com/rs/cors"
)

// SetupHeaders func
func SetupHeaders() *c.Cors {
	envs := map[string]string{
		"production":  os.Getenv("ALLOW_ORIGINS"),
		"development": "*",
	}

	currentEnv := os.Getenv("ENVIRONMENT")

	if _, ok := envs[currentEnv]; !ok {
		fmt.Printf("Environment: %s not valid.", currentEnv)
	}

	allowedOrigins := strings.Split(envs[currentEnv], ",")

	opts := c.Options{
		AllowedOrigins:     allowedOrigins,
		AllowedMethods:     []string{"GET", "POST", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders:     []string{"Accept", "Content-Type", "Origin", "Content-Length", "Accept-Encoding", "Authorization"},
		OptionsPassthrough: true,
		Debug:              envs[currentEnv] == "development",
	}

	return c.New(opts)
}

// SetupOptions func
func SetupOptions(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		w.WriteHeader(http.StatusOK)
		return
	}

	next(w, r)
}
