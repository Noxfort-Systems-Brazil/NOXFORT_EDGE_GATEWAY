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
// File: internal/infrastructure/ddns/ipv6_detector.go
// Author: Gabriel Moraes
// Date: 2026-10-12

package ddns

import (
	"net"
)

// GlobalIPv6Detector discovers the first public global unicast IPv6 on local interfaces.
type GlobalIPv6Detector struct{}

// NewGlobalIPv6Detector constructs a new detector.
func NewGlobalIPv6Detector() *GlobalIPv6Detector {
	return &GlobalIPv6Detector{}
}

// DetectGlobalIPv6 scans network interfaces for a public global unicast IPv6.
func (d *GlobalIPv6Detector) DetectGlobalIPv6() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet.IP == nil {
			continue
		}

		ip := ipNet.IP
		// Check if it is an IPv6 address and not an IPv4-mapped address
		if ip.To4() == nil && ip.To16() != nil {
			// Must be global unicast, not link-local, loopback, or private ULA (fc00::/7)
			if ip.IsGlobalUnicast() && !ip.IsPrivate() && !ip.IsLinkLocalUnicast() {
				return ip.String()
			}
		}
	}

	return ""
}
