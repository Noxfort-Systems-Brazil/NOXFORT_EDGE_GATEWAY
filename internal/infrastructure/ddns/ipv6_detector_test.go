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
// File: internal/infrastructure/ddns/ipv6_detector_test.go
// Author: Gabriel Moraes
// Date: October 2026

package ddns_test

import (
	"testing"

	"noxfort/edge-gateway/internal/infrastructure/ddns"
)

func TestDetectGlobalIPv6ReturnsStable(t *testing.T) {
	detector := ddns.NewGlobalIPv6Detector()
	ip := detector.DetectGlobalIPv6()
	if ip == "" {
		t.Skip("no global IPv6 present on test machine")
	}

	t.Logf("Detected IPv6: %s", ip)
	// Must not be the temporary privacy address ending in 170d
	if ip == "2804:84:8112:e301:62b8:4fc0:9f7f:170d" {
		t.Errorf("Detector returned temporary privacy address instead of stable SLAAC: %s", ip)
	}

	if ip != "2804:84:8112:e301:2e0:4cff:feaf:42" {
		t.Logf("Notice: Detected IP %s (expected stable SLAAC)", ip)
	}
}
