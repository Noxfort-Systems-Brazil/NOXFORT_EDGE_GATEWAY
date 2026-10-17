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
// File: internal/domain/models/models_test.go
// Author: Gabriel Moraes
// Date: October 2026

package models_test

import (
	"testing"

	"noxfort/edge-gateway/internal/domain/models"
)

func TestNewUpstreamService(t *testing.T) {
	svc, err := models.NewUpstreamService("carina", "CARINA Traffic", "127.0.0.1", 2201, "/health")
	if err != nil {
		t.Fatalf("unexpected error creating upstream: %v", err)
	}

	if svc.TargetAddress() != "127.0.0.1:2201" {
		t.Errorf("expected target address 127.0.0.1:2201, got %s", svc.TargetAddress())
	}
	if svc.HealthURL() != "http://127.0.0.1:2201/health" {
		t.Errorf("expected health URL http://127.0.0.1:2201/health, got %s", svc.HealthURL())
	}

	svc.MarkOnline(15)
	if svc.Status != models.StatusOnline || svc.ResponseTime != 15 {
		t.Errorf("expected online status with 15ms latency")
	}

	svc.MarkOffline()
	if svc.Status != models.StatusOffline {
		t.Errorf("expected offline status")
	}
}

func TestNewUpstreamServiceValidation(t *testing.T) {
	_, err := models.NewUpstreamService("", "Name", "127.0.0.1", 2201, "")
	if err == nil {
		t.Error("expected error for empty id")
	}

	_, err = models.NewUpstreamService("id", "Name", "127.0.0.1", -1, "")
	if err == nil {
		t.Error("expected error for invalid port")
	}
}

func TestNewRouteRule(t *testing.T) {
	rule, err := models.NewRouteRule("carinanoxfort.duckdns.org", "carina", true, true)
	if err != nil {
		t.Fatalf("unexpected error creating route: %v", err)
	}
	if rule.Domain != "carinanoxfort.duckdns.org" || rule.UpstreamID != "carina" {
		t.Errorf("unexpected rule values: %+v", rule)
	}

	_, err = models.NewRouteRule("", "carina", true, true)
	if err == nil {
		t.Error("expected error for empty domain")
	}
}

func TestNewDDNSTarget(t *testing.T) {
	target, err := models.NewDDNSTarget("carinanoxfort", "token123", "duckdns")
	if err != nil {
		t.Fatalf("unexpected error creating target: %v", err)
	}
	if target.FullDomain() != "carinanoxfort.duckdns.org" {
		t.Errorf("expected carinanoxfort.duckdns.org, got %s", target.FullDomain())
	}

	target.RecordSync("1.2.3.4", "2804::1", true, "")
	if target.LastIPv4 != "1.2.3.4" || target.LastIPv6 != "2804::1" || !target.SyncSuccess {
		t.Errorf("record sync failed to store state")
	}
}
