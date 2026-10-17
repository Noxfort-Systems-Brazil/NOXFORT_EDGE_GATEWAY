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
// Date: October 2026

package ddns

import (
	"bufio"
	"net"
	"os"
	"strconv"
	"strings"
)

// GlobalIPv6Detector discovers the stable public global unicast IPv6 on local interfaces.
type GlobalIPv6Detector struct{}

// NewGlobalIPv6Detector constructs a new detector.
func NewGlobalIPv6Detector() *GlobalIPv6Detector {
	return &GlobalIPv6Detector{}
}

// detectStableLinuxIPv6 reads /proc/net/if_inet6 and returns the first stable (non-temporary) global unicast IPv6.
func detectStableLinuxIPv6() string {
	file, err := os.Open("/proc/net/if_inet6")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)
		if len(parts) < 6 {
			continue
		}

		rawIP, scopeHex, flagsHex := parts[0], parts[3], parts[4]
		scope, err1 := strconv.ParseInt(scopeHex, 16, 64)
		flags, err2 := strconv.ParseInt(flagsHex, 16, 64)
		if err1 != nil || err2 != nil {
			continue
		}

		// scope == 0 means global unicast
		// IFA_F_TEMPORARY (0x01) means temporary privacy address
		if scope == 0 && (flags&0x01) == 0 {
			if len(rawIP) == 32 {
				var sb strings.Builder
				for i := 0; i < 32; i += 4 {
					if i > 0 {
						sb.WriteByte(':')
					}
					sb.WriteString(rawIP[i : i+4])
				}
				parsed := net.ParseIP(sb.String())
				if parsed != nil && parsed.IsGlobalUnicast() && !parsed.IsPrivate() && !parsed.IsLinkLocalUnicast() {
					return parsed.String()
				}
			}
		}
	}
	return ""
}

// DetectGlobalIPv6 scans network interfaces for a stable public global unicast IPv6.
func (d *GlobalIPv6Detector) DetectGlobalIPv6() string {
	// 1. Prioritize stable permanent SLAAC IPv6 from Linux kernel /proc/net/if_inet6
	if stable := detectStableLinuxIPv6(); stable != "" {
		return stable
	}

	// 2. Fallback to scanning net.InterfaceAddrs() with EUI-64 preference
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	var fallback string
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
				ipBytes := ip.To16()
				// Prefer SLAAC with standard EUI-64 marker (0xff, 0xfe at bytes 11, 12)
				if len(ipBytes) == 16 && ipBytes[11] == 0xff && ipBytes[12] == 0xfe {
					return ip.String()
				}
				if fallback == "" {
					fallback = ip.String()
				}
			}
		}
	}

	return fallback
}

