// NOXFORT EDGE GATEWAY - Unified Multi-Tenant Edge Gateway & Reverse Proxy
// Copyright (C) 2026 Noxfort Systems
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// File: config/config.go
// Author: Gabriel Moraes
// Date: 2026-10-12

package config

import (
	"os"
	"strconv"
	"strings"

	"noxfort/edge-gateway/internal/domain/models"
)

// Config represents full edge gateway settings.
type Config struct {
	HTTPPort               int
	HTTPSPort              int
	SyncIntervalSeconds    int
	HealthIntervalSeconds  int
	Upstreams              []*models.UpstreamService
	DDNSTargets            []*models.DDNSTarget
	Routes                 []models.RouteRule
}

// LoadConfig loads configuration from environment variables with sensible defaults.
func LoadConfig() (*Config, error) {
	httpPort := getEnvInt("GATEWAY_HTTP_PORT", 80)
	httpsPort := getEnvInt("GATEWAY_HTTPS_PORT", 443)
	syncInterval := getEnvInt("GATEWAY_SYNC_INTERVAL_SEC", 300)
	healthInterval := getEnvInt("GATEWAY_HEALTH_INTERVAL_SEC", 15)

	carinaDomain := getEnv("CARINA_DOMAIN", "carinanoxfort")
	carinaToken := getEnv("CARINA_TOKEN", "ff528f24-1dd3-4c84-9c12-02b6d9805440")
	carinaPort := getEnvInt("CARINA_PORT", 2201)

	monitorDomain := getEnv("MONITOR_DOMAIN", "noxfortmonitor")
	monitorToken := getEnv("MONITOR_TOKEN", "2c42c374-2143-4292-93c3-5baccb6446c4")
	monitorPort := getEnvInt("MONITOR_PORT", 22100)

	// Build default domain entities
	uCarina, _ := models.NewUpstreamService("carina", "CARINA Traffic AI", "127.0.0.1", carinaPort, "/health")
	uMonitor, _ := models.NewUpstreamService("monitor", "Noxfort Monitor", "127.0.0.1", monitorPort, "/api/health")

	upstreams := []*models.UpstreamService{uCarina, uMonitor}

	tCarina, _ := models.NewDDNSTarget(carinaDomain, carinaToken, "duckdns")
	tMonitor, _ := models.NewDDNSTarget(monitorDomain, monitorToken, "duckdns")

	ddnsTargets := []*models.DDNSTarget{tCarina, tMonitor}

	rCarina, _ := models.NewRouteRule(tCarina.FullDomain(), "carina", true, true)
	rMonitor, _ := models.NewRouteRule(tMonitor.FullDomain(), "monitor", true, true)

	routes := []models.RouteRule{*rCarina, *rMonitor}

	return &Config{
		HTTPPort:              httpPort,
		HTTPSPort:             httpsPort,
		SyncIntervalSeconds:   syncInterval,
		HealthIntervalSeconds: healthInterval,
		Upstreams:             upstreams,
		DDNSTargets:           ddnsTargets,
		Routes:                routes,
	}, nil
}

func getEnv(key, fallback string) string {
	if val := strings.TrimSpace(os.Getenv(key)); val != "" {
		return val
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if val := strings.TrimSpace(os.Getenv(key)); val != "" {
		if n, err := strconv.Atoi(val); err == nil {
			return n
		}
	}
	return fallback
}
