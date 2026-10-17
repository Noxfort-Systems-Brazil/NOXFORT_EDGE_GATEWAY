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
// File: internal/domain/models/ddns_target.go
// Author: Gabriel Moraes
// Date: October 2026

package models

import (
	"fmt"
	"strings"
	"time"
)

// DDNSTarget is a Value Object holding configuration and dynamic state for a DDNS domain.
type DDNSTarget struct {
	Subdomain   string    `json:"subdomain"`
	Token       string    `json:"-"`
	Provider    string    `json:"provider"`
	LastIPv4    string    `json:"last_ipv4"`
	LastIPv6    string    `json:"last_ipv6"`
	LastSync    time.Time `json:"last_sync"`
	SyncSuccess bool      `json:"sync_success"`
	LastError   string    `json:"last_error,omitempty"`
}

// NewDDNSTarget creates and validates a DDNSTarget.
func NewDDNSTarget(subdomain, token, provider string) (*DDNSTarget, error) {
	cleanSub := strings.TrimSpace(subdomain)
	cleanSub = strings.TrimPrefix(cleanSub, "https://")
	cleanSub = strings.TrimPrefix(cleanSub, "http://")
	cleanSub = strings.TrimSuffix(cleanSub, "/")
	cleanSub = strings.TrimSuffix(cleanSub, ".duckdns.org")

	if cleanSub == "" {
		return nil, fmt.Errorf("ddns subdomain cannot be empty")
	}
	if strings.TrimSpace(token) == "" {
		return nil, fmt.Errorf("ddns token cannot be empty for subdomain %s", cleanSub)
	}
	if provider == "" {
		provider = "duckdns"
	}

	return &DDNSTarget{
		Subdomain: cleanSub,
		Token:     strings.TrimSpace(token),
		Provider:  provider,
	}, nil
}

// FullDomain returns the FQDN representation.
func (d *DDNSTarget) FullDomain() string {
	if d.Provider == "duckdns" {
		return fmt.Sprintf("%s.duckdns.org", d.Subdomain)
	}
	return d.Subdomain
}

// RecordSync updates the sync status.
func (d *DDNSTarget) RecordSync(ipv4, ipv6 string, success bool, err string) {
	d.LastIPv4 = ipv4
	d.LastIPv6 = ipv6
	d.SyncSuccess = success
	d.LastError = err
	d.LastSync = time.Now()
}
