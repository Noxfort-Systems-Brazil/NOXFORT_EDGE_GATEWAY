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
// File: internal/application/usecases/monitor_health_usecase.go
// Author: Gabriel Moraes
// Date: October 2026

package usecases

import (
	"context"
	"sync"

	"noxfort/edge-gateway/internal/domain/contracts"
	"noxfort/edge-gateway/internal/domain/models"
)

// MonitorHealthUseCase executes concurrent health checks on all upstream services (SRP).
type MonitorHealthUseCase struct {
	mu        sync.RWMutex
	upstreams map[string]*models.UpstreamService
	checker   contracts.IHealthMonitor
}

// NewMonitorHealthUseCase instantiates the use case with injected monitor contract.
func NewMonitorHealthUseCase(
	upstreams []*models.UpstreamService,
	checker contracts.IHealthMonitor,
) *MonitorHealthUseCase {
	m := make(map[string]*models.UpstreamService, len(upstreams))
	for _, u := range upstreams {
		m[u.ID] = u
	}
	return &MonitorHealthUseCase{
		upstreams: m,
		checker:   checker,
	}
}

// Execute triggers health checks across all upstreams concurrently.
func (uc *MonitorHealthUseCase) Execute(ctx context.Context) []contracts.CheckResult {
	uc.mu.RLock()
	services := make([]*models.UpstreamService, 0, len(uc.upstreams))
	for _, s := range uc.upstreams {
		services = append(services, s)
	}
	uc.mu.RUnlock()

	results := make([]contracts.CheckResult, len(services))
	var wg sync.WaitGroup

	for i, svc := range services {
		wg.Add(1)
		go func(idx int, target *models.UpstreamService) {
			defer wg.Done()
			res := uc.checker.CheckHealth(ctx, *target)
			results[idx] = res

			uc.mu.Lock()
			if res.IsAlive {
				target.MarkOnline(res.LatencyMs)
			} else {
				target.MarkOffline()
			}
			uc.mu.Unlock()
		}(i, svc)
	}

	wg.Wait()
	return results
}

// GetUpstreams returns the current registered upstream services.
func (uc *MonitorHealthUseCase) GetUpstreams() []*models.UpstreamService {
	uc.mu.RLock()
	defer uc.mu.RUnlock()
	res := make([]*models.UpstreamService, 0, len(uc.upstreams))
	for _, u := range uc.upstreams {
		res = append(res, u)
	}
	return res
}
