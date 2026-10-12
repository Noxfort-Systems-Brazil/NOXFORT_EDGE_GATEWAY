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
// File: internal/application/usecases/usecases_test.go
// Author: Gabriel Moraes
// Date: 2026-10-12

package usecases_test

import (
	"context"
	"testing"
	"time"

	"noxfort/edge-gateway/internal/application/usecases"
	"noxfort/edge-gateway/internal/domain/contracts"
	"noxfort/edge-gateway/internal/domain/models"
)

// MockDDNSSynchronizer for testing.
type MockDDNSSynchronizer struct {
	ShouldFail bool
}

func (m *MockDDNSSynchronizer) Sync(ctx context.Context, target models.DDNSTarget, ipv4, ipv6 string) contracts.SyncResult {
	if m.ShouldFail {
		return contracts.SyncResult{
			Subdomain: target.Subdomain,
			Success:   false,
			Error:     "mock failure",
		}
	}
	return contracts.SyncResult{
		Subdomain: target.Subdomain,
		IPv4:      ipv4,
		IPv6:      ipv6,
		Success:   true,
	}
}

func (m *MockDDNSSynchronizer) ProviderName() string {
	return "mock"
}

// MockIPv6Detector for testing.
type MockIPv6Detector struct {
	IP string
}

func (m *MockIPv6Detector) DetectGlobalIPv6() string {
	return m.IP
}

// MockIPv4Detector for testing.
type MockIPv4Detector struct {
	IP string
}

func (m *MockIPv4Detector) DetectPublicIPv4(ctx context.Context) string {
	return m.IP
}

// MockHealthChecker for testing.
type MockHealthChecker struct {
	IsAlive bool
}

func (m *MockHealthChecker) CheckHealth(ctx context.Context, service models.UpstreamService) contracts.CheckResult {
	return contracts.CheckResult{
		ServiceID: service.ID,
		IsAlive:   m.IsAlive,
		LatencyMs: 10,
	}
}

func TestSyncAllDDNSUseCase(t *testing.T) {
	t1, _ := models.NewDDNSTarget("carinanoxfort", "tok1", "duckdns")
	t2, _ := models.NewDDNSTarget("noxfortmonitor", "tok2", "duckdns")

	mockSync := &MockDDNSSynchronizer{ShouldFail: false}
	mockV6 := &MockIPv6Detector{IP: "2804::1"}
	mockV4 := &MockIPv4Detector{IP: "189.1.1.1"}

	uc := usecases.NewSyncAllDDNSUseCase([]*models.DDNSTarget{t1, t2}, mockSync, mockV6, mockV4)
	results := uc.Execute(context.Background())

	if len(results) != 2 {
		t.Fatalf("expected 2 sync results, got %d", len(results))
	}
	for _, res := range results {
		if !res.Success {
			t.Errorf("expected success for target %s, got error: %s", res.Subdomain, res.Error)
		}
	}
}

func TestMonitorHealthUseCase(t *testing.T) {
	u1, _ := models.NewUpstreamService("carina", "Carina", "127.0.0.1", 2201, "")
	u2, _ := models.NewUpstreamService("monitor", "Monitor", "127.0.0.1", 22100, "")

	mockHealth := &MockHealthChecker{IsAlive: true}
	uc := usecases.NewMonitorHealthUseCase([]*models.UpstreamService{u1, u2}, mockHealth)

	results := uc.Execute(context.Background())
	if len(results) != 2 {
		t.Fatalf("expected 2 health results, got %d", len(results))
	}

	for _, res := range results {
		if !res.IsAlive {
			t.Errorf("expected service %s to be alive", res.ServiceID)
		}
	}
}

func TestGetGatewayStatusUseCase(t *testing.T) {
	t1, _ := models.NewDDNSTarget("carinanoxfort", "tok1", "duckdns")
	u1, _ := models.NewUpstreamService("carina", "Carina", "127.0.0.1", 2201, "")
	rule, _ := models.NewRouteRule("carinanoxfort.duckdns.org", "carina", true, true)

	syncUC := usecases.NewSyncAllDDNSUseCase(
		[]*models.DDNSTarget{t1},
		&MockDDNSSynchronizer{},
		&MockIPv6Detector{IP: "2804::1"},
		&MockIPv4Detector{IP: "189.1.1.1"},
	)
	healthUC := usecases.NewMonitorHealthUseCase(
		[]*models.UpstreamService{u1},
		&MockHealthChecker{IsAlive: true},
	)

	syncUC.Execute(context.Background())
	healthUC.Execute(context.Background())

	statusUC := usecases.NewGetGatewayStatusUseCase(time.Now(), syncUC, healthUC, []models.RouteRule{*rule})
	dto := statusUC.Execute()

	if dto.TotalRoutes != 1 || dto.HealthyServices != 1 {
		t.Errorf("unexpected status dto: %+v", dto)
	}
}
