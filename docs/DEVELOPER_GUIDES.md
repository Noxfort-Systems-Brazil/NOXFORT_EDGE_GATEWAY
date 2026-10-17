---
tags: [developer, guides, go, toolchain, clean-architecture, contribution]
aliases: [Developer Guides, Development Setup, Contributor Manual]
---

# 🛠️ Developer & Contributor Guides

This guide walks software engineers through setting up a local development environment, navigating the Clean Architecture layers, implementing new features, and adhering to Noxfort Systems coding standards.

⬅️ Back to [Master Documentation Hub](NOXFORT_EDGE_GATEWAY_MOC.md) | 🏛️ See [Architecture](ARCHITECTURE.md) | 🧪 See [Testing](TESTING.md)

---

## 1. Local Development Setup

### Prerequisites
- **Go 1.22+** installed (`go version`)
- **Docker 24+** & Docker Compose plugin
- GNU `make` installed

### Initial Setup
```bash
# Clone the repository
git clone https://github.com/Noxfort-Systems-Brazil/NOXFORT_EDGE_GATEWAY.git
cd NOXFORT_EDGE_GATEWAY

# Compile the Go orchestrator
make build

# Run all test suites
make test
```

---

## 2. Working with Automation (`Makefile`)

The project Makefile automates build, test, and lifecycle workflows:

```bash
make all      # Builds binaries and executes test suite
make build    # Compiles cmd/gateway/main.go into bin/noxfort-gateway
make test     # Runs Go tests with verbose output
make start    # Builds container and launches Caddy in background
make stop     # Gracefully terminates the running gateway container
make status   # Displays running container health and port bindings
make logs     # Streams live logs from the reverse proxy
```

---

## 3. Extending the Codebase

### 3.1 Adding a New Dynamic DNS Provider
1. Create a new adapter file in `internal/infrastructure/ddns/`:
   ```go
   // File: internal/infrastructure/ddns/cloudflare_client.go
   package ddns

   type CloudflareClient struct {
       apiToken string
   }

   func (c *CloudflareClient) SyncTarget(ctx context.Context, target *models.DDNSTarget, ipv4, ipv6 string) (*contracts.SyncResult, error) {
       // Implement Cloudflare DNS record update
   }
   ```
2. Ensure it satisfies the `contracts.IDDNSSynchronizer` interface.
3. Wire the new adapter into the composition root (`cmd/gateway/main.go`).

### 3.2 Adding a New Health Checker
Implement `contracts.IHealthMonitor` in `internal/infrastructure/health/` (e.g. gRPC Health Checking Protocol or UDP Ping).

---

## 4. Coding Standards & SPDX Apache 2.0 License Header

All Go files must adhere to standard `gofmt` and include the SPDX Apache 2.0 license header:

```go
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
// File: <file-relative-path>
// Author: Gabriel Moraes
// Date: October 2026
```

Use `cabecalho.txt` as a reference when creating new files.

---

<div align="center">
  <img src="assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licensed under Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
