---
tags: [readme, home, gateway, edge, reverse-proxy, noxfort]
aliases: [NOXFORT Edge Gateway, Gateway Root, Master Ingress]
---

<div align="center">

# NOXFORT EDGE GATEWAY
### Unified Multi-Tenant Edge Gateway & High-Availability Reverse Proxy
*Noxfort Systems — A State Of Art Company*

[![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=flat&logo=github)](https://github.com/Noxfort-Systems-Brazil)
[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Caddy](https://img.shields.io/badge/Caddy-2.8%2B-1F88C0?style=flat&logo=caddy&logoColor=white)](https://caddyserver.com/)
[![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![Architecture](https://img.shields.io/badge/Architecture-DDD%20%7C%20Clean-orange?style=flat)](#-arquitetura-e-destaques)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue?style=flat)](LICENSE)

---

🌐 **Translations / Idiomas:** **[🇺🇸 English](README.md)** • **[🇧🇷 Português](docs/pt-br/README.md)** • **[🇪🇸 Español](docs/es/README.md)** • **[🇫🇷 Français](docs/fr/README.md)** • **[🇷🇺 Русский](docs/ru/README.md)** • **[🇨🇳 简体中文](docs/zh/README.md)** • **[📚 Documentation Hub](docs/index.md)**

---

</div>

**NOXFORT EDGE GATEWAY** is the high-availability central network ingress and unified reverse proxy for Noxfort Systems smart city infrastructure. Operating as the exclusive gatekeeper for public HTTP/HTTPS ports (80 and 443), it provides automated TLS termination, multi-tenant SNI routing, periodic dual-stack dynamic DNS (DuckDNS IPv4/IPv6) synchronization, and active health probing for microservices running across private host ports:

- 🚦 **CARINA (CCO Videowall, Adaptive AI & EVP Ambulances):** `carinanoxfort.duckdns.org` ➔ `127.0.0.1:2201`
- 📊 **NOXFORT MONITOR (Industrial Telemetry & Incident Ingestion):** `noxfortmonitor.duckdns.org` ➔ `127.0.0.1:22100` and WSS `9001` (MQTT)
- 🧠 **SYNAPSE (High-Frequency AI & Edge Orchestration):** `synapse.duckdns.org` ➔ Synapse upstream port

---

## 📚 Documentation Hub & Knowledge Vault

Explore the full architecture, API specifications, DDNS mechanisms, and operations guides:

| Card / Subsystem | Focus Area | Direct Link |
| :--- | :--- | :---: |
| 📚 **Master Documentation Hub** | Central Index, Obsidian Vault & Repository Map | [Explore Hub](docs/NOXFORT_EDGE_GATEWAY_MOC.md) |
| 🏛️ **System Architecture** | DDD Layers, Ingress Topology & Caddy Reverse Proxy | [View Blueprint](ARCHITECTURE.md) |
| ⚡ **API & Contracts Reference** | Caddy Admin API, DuckDNS Protocol & Status DTOs | [View API Specs](docs/API_REFERENCE.md) |
| 🌐 **DDNS Synchronization** | DuckDNS Engine, Public IPv4 & Global IPv6 Detection | [View DDNS Guide](docs/DDNS_SYNCHRONIZATION.md) |
| 🛡️ **Active Health Monitoring** | Periodic TCP/HTTP Health Probes & Upstream States | [View Health Guide](docs/HEALTH_MONITORING.md) |
| ⚙️ **Configuration Guide** | Environment Variables, Caddyfile & Port Mappings | [View Config Guide](docs/CONFIGURATION_GUIDE.md) |
| 🛠️ **Developer Guides** | Go Toolchain, Adding Upstreams & Code Guidelines | [View Guides](docs/DEVELOPER_GUIDES.md) |
| 🧪 **Testing & Validation** | Go Unit Tests, Domain Contracts & Mock Verification | [View Guidelines](docs/TESTING.md) |
| 🚀 **Deployment & Operations** | Docker Compose, Systemd Service & Production Run | [View Deployment](docs/DEPLOYMENT_AND_PACKAGING.md) |

---

## ⚡ Key Highlights & Architecture

- **Domain-Driven Design (DDD) & Clean Architecture:** Pure separation between Domain Models/Value Objects, Domain Contracts (Interfaces), Application Use Cases, and Infrastructure Implementations (DuckDNS, HTTP Probers, Caddy).
- **Single Public Port Binding (80/443):** Eliminates port conflict hazards across multiple services on municipal CCO servers by multiplexing all inbound traffic via Server Name Indication (SNI).
- **Automated TLS Lifecycle:** Automatic issuance, staging, and renewal of Let's Encrypt / ZeroSSL wildcard and domain SSL certificates via Caddy with DuckDNS DNS-01 ACME challenge.
- **Dual-Stack DDNS Synchronizer:** Periodic automated synchronization daemon detecting public IPv4 (`api.ipify.org`) and global unicast IPv6 interfaces, updating DuckDNS endpoints with backoff resilience.
- **Active Health Probing Engine:** Concurrent non-blocking TCP socket and HTTP `/health` probes monitoring the health of all registered upstreams with millisecond latency metrics.
- **Industrial Zero-Downtime Operation:** Containerized via Docker Compose with `restart: unless-stopped` and host networking integration (`host.docker.internal`).

---

## 🚀 Quick Start

### 1. Requirements
- **Go 1.22+**
- **Docker** and **Docker Compose**
- Ports **80** and **443** free on the host interface

### 2. Build and Test
```bash
# Build the Go orchestrator and test all DDD use cases
make build
make test
```

### 3. Running the Gateway
```bash
# Start Caddy reverse proxy and daemon services
make start
```

### 4. Monitoring & Telemetry
```bash
# Check running container status
make status

# Stream live access and reverse proxy logs
make logs
```

### 5. Stopping the Gateway
```bash
make stop
```

---

<div align="center">
  <img src="docs/assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="48" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <i>Unified Multi-Tenant Edge Gateway & Reverse Proxy • NOXFORT EDGE GATEWAY v1.0.0</i><br/>
  <small>Licensed under the <a href="LICENSE">Apache License 2.0</a>. © 2026 Noxfort Systems.</small>
</div>
