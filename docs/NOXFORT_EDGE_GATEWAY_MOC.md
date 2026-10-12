---
tags: [moc, hub, docs, obsidian, gateway, edge, index]
aliases: [NOXFORT Edge Gateway MOC, Master Documentation Hub, Documentation Index, Knowledge Vault]
---

<div align="center">

# NOXFORT EDGE GATEWAY Technical Master Documentation Hub
### Central Index & Clean Architecture Knowledge Vault
*Noxfort Systems — A State Of Art Company*

</div>

Welcome to the **NOXFORT EDGE GATEWAY** technical documentation library. Designed as a high-availability, multi-tenant reverse proxy and perimeter security gateway for municipal smart city platforms, NOXFORT EDGE GATEWAY bridges public Internet ingress (ports 80 & 443), automated DNS-01 ACME TLS issuance, and local microservices.

This master documentation index provides deep technical coverage for core backend developers, DevOps engineers, and network administrators. It is fully compatible with both **GitHub** and **[Obsidian](https://obsidian.md/)**.

---

## 🗺️ Codebase Map & Directory Hierarchy

```text
NOXFORT_EDGE_GATEWAY/
├── cmd/
│   └── gateway/
│       └── main.go                 # Master Composition Root & Orchestrator Daemon
├── config/
│   └── config.go                   # 12-Factor Environment Configuration & Default Upstreams
├── internal/
│   ├── domain/                     # Pure Enterprise Business Rules (DDD Domain Layer)
│   │   ├── models/                 # Pure Entities & Value Objects
│   │   │   ├── upstream_service.go # UpstreamService Entity (State, Latency, Failures)
│   │   │   ├── ddns_target.go      # DDNSTarget Value Object (Domain, Token, Provider)
│   │   │   ├── route_rule.go       # RouteRule Value Object (Domain, TargetService)
│   │   │   └── models_test.go      # Pure Domain Entity Unit Tests
│   │   └── contracts/              # Inversion of Control Interfaces (DIP)
│   │       ├── ddns_synchronizer.go# IDDNSSynchronizer Interface
│   │       ├── health_monitor.go   # IHealthMonitor Interface
│   │       └── reverse_proxy.go    # IProxyRouter Interface
│   ├── application/                # Application Business Rules (Clean Architecture)
│   │   ├── usecases/               # Orchestrated Use Cases
│   │   │   ├── sync_all_ddns_usecase.go    # Concurrent DuckDNS Dynamic DNS Sync
│   │   │   ├── monitor_health_usecase.go  # Active TCP/HTTP Health Probing Loop
│   │   │   ├── get_gateway_status_usecase.go # Telemetry Aggregation
│   │   │   └── usecases_test.go    # Application Layer Unit Tests with Mocks
│   │   └── dtos/                   # Boundary Data Transfer Objects
│   │       └── gateway_status_dto.go       # JSON Real-Time Status Representation
│   └── infrastructure/             # Frameworks, Drivers & Network Infrastructure
│       ├── ddns/                   # Dynamic DNS Implementations
│       │   ├── duckdns_client.go   # DuckDNS HTTP/HTTPS API Integration
│       │   ├── ipv4_detector.go    # External Public IPv4 Discovery (ipify)
│       │   └── ipv6_detector.go    # Local Network Interface Global IPv6 Scanner
│       ├── health/                 # Health Checking Drivers
│       │   └── tcp_http_checker.go # Two-phase TCP Handshake + HTTP Prober
│       └── caddy/                  # Reverse Proxy Container Infrastructure
│           ├── Caddyfile           # Multi-Tenant SNI Routing Configuration
│           └── Dockerfile          # Custom Caddy Build with DuckDNS DNS-01 Plugin
├── docs/                           # Comprehensive Knowledge Vault & Obsidian Docs
│   ├── index.md                    # MkDocs Entry Point & Portal
│   ├── NOXFORT_EDGE_GATEWAY_MOC.md # Master Documentation Hub (This File)
│   ├── ARCHITECTURE.md             # Complete Architecture Blueprint & Mermaid Graphs
│   ├── API_REFERENCE.md            # Caddy Admin API, DuckDNS Protocol & Status DTOs
│   ├── DDNS_SYNCHRONIZATION.md     # DuckDNS Engine & Dual-Stack IP Resolution
│   ├── HEALTH_MONITORING.md        # Active Health Probing & Upstream Availability
│   ├── CONFIGURATION_GUIDE.md      # Environment Variables, Caddyfile & Port Maps
│   ├── DEVELOPER_GUIDES.md         # Go Toolchain, Adding Upstreams & Contribution
│   ├── TESTING.md                  # Unit Testing, Coverage & Race Condition Checks
│   ├── DEPLOYMENT_AND_PACKAGING.md # Production Deployment, Docker & Systemd
│   ├── assets/                     # Corporate Assets (noxfort-logo.png)
│   ├── stylesheets/                # Custom Theme Stylesheets
│   ├── pt-br/                      # Brazilian Portuguese Documentation Mirror
│   └── en/                         # English Documentation Mirror
├── bin/                            # Compiled Native Binaries (noxfort-gateway)
├── docker-compose.yml              # Container Orchestration (Host Ports 80 & 443)
├── Makefile                        # Build, Test, Start, Stop & Log Automation
├── cabecalho.txt                   # Standard Apache 2.0 License Header Template
├── LICENSE                         # Official Apache License 2.0
└── README.md                       # Repository Master Overview & Quick Start
```

---

## 📚 Technical Knowledge Vault Index

| Card / Subsystem | Primary Focus | Direct Link |
| :--- | :--- | :---: |
| 🏛️ **System Architecture** | Domain-Driven Design, Ingress Topology & Caddy Reverse Proxy | [View Blueprint](ARCHITECTURE.md) |
| ⚡ **API & Contracts Reference** | Caddy Admin API, DuckDNS Protocol & Status DTOs | [View API Specs](API_REFERENCE.md) |
| 🌐 **DDNS Synchronization** | DuckDNS Engine, Public IPv4 & Global IPv6 Detection | [View DDNS Guide](DDNS_SYNCHRONIZATION.md) |
| 🛡️ **Active Health Monitoring** | Periodic TCP/HTTP Health Probes & Upstream States | [View Health Guide](HEALTH_MONITORING.md) |
| ⚙️ **Configuration Guide** | Environment Variables, Caddyfile & Port Mappings | [View Config Guide](CONFIGURATION_GUIDE.md) |
| 🛠️ **Developer Guides** | Go Toolchain, Adding Upstreams & Code Guidelines | [View Guides](DEVELOPER_GUIDES.md) |
| 🧪 **Testing & Validation** | Go Unit Tests, Domain Contracts & Mock Verification | [View Guidelines](TESTING.md) |
| 🚀 **Deployment & Operations** | Docker Compose, Systemd Service & Production Run | [View Deployment](DEPLOYMENT_AND_PACKAGING.md) |

---

<div align="center">
  <img src="assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licensed under the Apache License 2.0.</small>
</div>
