package handlers

import (
	"math/rand"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/gin-gonic/gin"
)

type TrafficSplitConfig struct {
	MigrationPercent int    `env:"MOVIES_MIGRATION_PERCENT"`
	MonolithURL      string `env:"MONOLITH_URL"`
	MoviesServiceURL string `env:"MOVIES_SERVICE_URL"`
	GradualMigration bool   `env:"GRADUAL_MIGRATION"`
}

func NewTrafficSplitConfigFromEnv() (TrafficSplitConfig, error) {
	var cfg TrafficSplitConfig
	err := env.Parse(&cfg)
	return cfg, err
}

func NewTrafficSplitHandler(config TrafficSplitConfig) gin.HandlerFunc {
	monolithURL, _ := url.Parse(config.MonolithURL)
	microserviceURL, _ := url.Parse(config.MoviesServiceURL)
	monolithProxy := httputil.NewSingleHostReverseProxy(monolithURL)
	microserviceProxy := httputil.NewSingleHostReverseProxy(microserviceURL)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/movies") {
			if config.GradualMigration && rnd.Intn(100) < config.MigrationPercent {
				microserviceProxy.ServeHTTP(c.Writer, c.Request)
			} else {
				monolithProxy.ServeHTTP(c.Writer, c.Request)
			}
			return
		}

		monolithProxy.ServeHTTP(c.Writer, c.Request)
	}
}
