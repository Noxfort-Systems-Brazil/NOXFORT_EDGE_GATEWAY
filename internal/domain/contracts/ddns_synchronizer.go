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
// File: internal/domain/contracts/ddns_synchronizer.go
// Author: Gabriel Moraes
// Date: 2026-10-12

package contracts

import (
	"context"

	"noxfort/edge-gateway/internal/domain/models"
)

// SyncResult defines the outcome of a dynamic DNS update operation.
type SyncResult struct {
	Subdomain string
	IPv4      string
	IPv6      string
	Success   bool
	Error     string
}

// IDDNSSynchronizer defines the contract for updating Dynamic DNS records (DIP).
type IDDNSSynchronizer interface {
	Sync(ctx context.Context, target models.DDNSTarget, ipv4, ipv6 string) SyncResult
	ProviderName() string
}

// IIPv6Detector defines the contract for discovering host public IPv6.
type IIPv6Detector interface {
	DetectGlobalIPv6() string
}

// IIPv4Detector defines the contract for discovering host public IPv4.
type IIPv4Detector interface {
	DetectPublicIPv4(ctx context.Context) string
}
