// NOXFORT EDGE GATEWAY - Unified Multi-Tenant Edge Gateway & Reverse Proxy
// SPDX-FileCopyrightText: 2026 Noxfort Systems
// SPDX-License-Identifier: Apache-2.0
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
// File: cmd/gateway/main.go
// Author: Gabriel Moraes
// Date: October 2026

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"noxfort/edge-gateway/config"
	"noxfort/edge-gateway/internal/application/usecases"
	"noxfort/edge-gateway/internal/infrastructure/ddns"
	"noxfort/edge-gateway/internal/infrastructure/health"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("==========================================================")
	log.Println("NOXFORT EDGE GATEWAY (Multi-Tenant Master Orchestrator)")
	log.Println("Architecture: DDD + SOLID + Clean Architecture")
	log.Println("==========================================================")

	startTime := time.Now()

	// 1. Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("[FATAL] Error loading gateway configuration: %v", err)
	}

	// 2. Initialize Infrastructure Layer (Adapters)
	ipv6Detector := ddns.NewGlobalIPv6Detector()
	ipv4Detector := ddns.NewPublicIPv4Detector()
	duckClient := ddns.NewDuckDNSClient()
	healthChecker := health.NewTCPAndHTTPChecker()

	// 3. Initialize Application Layer (Use Cases)
	syncDDNSUseCase := usecases.NewSyncAllDDNSUseCase(
		cfg.DDNSTargets,
		duckClient,
		ipv6Detector,
		ipv4Detector,
		cfg.EnableIPv4,
	)

	monitorHealthUseCase := usecases.NewMonitorHealthUseCase(
		cfg.Upstreams,
		healthChecker,
	)

	statusUseCase := usecases.NewGetGatewayStatusUseCase(
		startTime,
		syncDDNSUseCase,
		monitorHealthUseCase,
		cfg.Routes,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 4. Initial synchronous execution
	log.Println("[BOOT] Executing initial Dynamic DNS synchronization...")
	syncDDNSUseCase.Execute(ctx)

	log.Println("[BOOT] Executing initial upstream health checks...")
	monitorHealthUseCase.Execute(ctx)

	// 5. Start background periodic workers
	// 5.1 DDNS Sync Loop
	go func() {
		ticker := time.NewTicker(time.Duration(cfg.SyncIntervalSeconds) * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				log.Println("[HEARTBEAT] Refreshing DDNS records for all domains...")
				syncDDNSUseCase.Execute(ctx)
			}
		}
	}()

	// 5.2 Upstream Health Check Loop
	go func() {
		ticker := time.NewTicker(time.Duration(cfg.HealthIntervalSeconds) * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				monitorHealthUseCase.Execute(ctx)
			}
		}
	}()

	// 6. Expose local diagnostic and management endpoint on localhost:2026
	mux := http.NewServeMux()
	mux.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		status := statusUseCase.Execute()
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(status)
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"online","gateway":"noxfort-edge-gateway"}`))
	})

	srv := &http.Server{
		Addr:    "127.0.0.1:2026",
		Handler: mux,
	}

	go func() {
		log.Printf("[MANAGEMENT] Status & telemetry API listening on http://127.0.0.1:2026/status")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("[MANAGEMENT] Server error: %v", err)
		}
	}()

	log.Println("[READY] Noxfort Edge Gateway Orchestrator running. Press Ctrl+C to stop.")

	// 7. Graceful OS signal termination
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	sig := <-sigChan

	log.Printf("[SHUTDOWN] Signal %v received. Performing graceful shutdown...", sig)
	cancel()

	shutCtx, shutCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer shutCancel()
	_ = srv.Shutdown(shutCtx)

	log.Println("[SHUTDOWN] Noxfort Edge Gateway safely terminated.")
	fmt.Println("Gateway stopped.")
}
