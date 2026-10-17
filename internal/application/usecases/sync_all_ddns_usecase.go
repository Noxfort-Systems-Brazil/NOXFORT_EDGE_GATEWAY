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
// File: internal/application/usecases/sync_all_ddns_usecase.go
// Author: Gabriel Moraes
// Date: October 2026

package usecases

import (
	"context"
	"log"
	"sync"

	"noxfort/edge-gateway/internal/domain/contracts"
	"noxfort/edge-gateway/internal/domain/models"
)

// SyncAllDDNSUseCase orchestrates periodic IP detection and synchronization across all domains (SRP).
type SyncAllDDNSUseCase struct {
	mu           sync.RWMutex
	targets      []*models.DDNSTarget
	synchronizer contracts.IDDNSSynchronizer
	ipv6Detector contracts.IIPv6Detector
	ipv4Detector contracts.IIPv4Detector
	enableIPv4   bool
}

// NewSyncAllDDNSUseCase constructs the use case instance with injected dependencies (DIP).
func NewSyncAllDDNSUseCase(
	targets []*models.DDNSTarget,
	synchronizer contracts.IDDNSSynchronizer,
	ipv6Detector contracts.IIPv6Detector,
	ipv4Detector contracts.IIPv4Detector,
	enableIPv4 bool,
) *SyncAllDDNSUseCase {
	return &SyncAllDDNSUseCase{
		targets:      targets,
		synchronizer: synchronizer,
		ipv6Detector: ipv6Detector,
		ipv4Detector: ipv4Detector,
		enableIPv4:   enableIPv4,
	}
}

// Execute performs one synchronous detection and update round for all registered DDNS targets.
func (uc *SyncAllDDNSUseCase) Execute(ctx context.Context) []contracts.SyncResult {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	var ipv6 string
	if uc.ipv6Detector != nil {
		ipv6 = uc.ipv6Detector.DetectGlobalIPv6()
	}

	var ipv4 string
	if uc.enableIPv4 && uc.ipv4Detector != nil {
		ipv4 = uc.ipv4Detector.DetectPublicIPv4(ctx)
	}

	results := make([]contracts.SyncResult, 0, len(uc.targets))

	for _, target := range uc.targets {
		res := uc.synchronizer.Sync(ctx, *target, ipv4, ipv6)
		target.RecordSync(ipv4, ipv6, res.Success, res.Error)
		results = append(results, res)

		if res.Success {
			log.Printf("[SyncAllDDNS] Domain '%s' successfully synchronized (IPv4: %s, IPv6: %s)", target.FullDomain(), ipv4, ipv6)
		} else {
			log.Printf("[SyncAllDDNS] Failed updating domain '%s': %s", target.FullDomain(), res.Error)
		}
	}

	return results
}

// GetTargets returns the current list of targets.
func (uc *SyncAllDDNSUseCase) GetTargets() []*models.DDNSTarget {
	uc.mu.RLock()
	defer uc.mu.RUnlock()
	return uc.targets
}
