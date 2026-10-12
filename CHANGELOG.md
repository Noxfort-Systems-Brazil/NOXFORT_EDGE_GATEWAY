# Changelog

All notable changes to the **NOXFORT EDGE GATEWAY** project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [1.0.0] - 2026-10-12

### Added
- **Domain-Driven Design (DDD) Architecture:** Pure core domain models (`UpstreamService`, `DDNSTarget`, `RouteRule`) and contracts (`IDDNSSynchronizer`, `IHealthMonitor`, `IProxyRouter`).
- **Clean Architecture Use Cases:**
  - `SyncAllDDNSUseCase`: Dual-stack DuckDNS dynamic DNS synchronization engine.
  - `MonitorHealthUseCase`: Concurrent active TCP/HTTP upstream probing.
  - `GetGatewayStatusUseCase`: Telemetry status consolidation into `GatewayStatusDTO`.
- **Infrastructure Integrations:**
  - `DuckDNSClient`: Robust HTTPS client with timeout management and backoff.
  - `IPv4Detector`: External egress public IPv4 detection via `api.ipify.org`.
  - `GlobalIPv6Detector`: Local network interface scanner extracting public global unicast IPv6 addresses.
  - `TCPHTTPHealthChecker`: Two-phase connection validator (TCP handshake + HTTP status code).
- **Caddy Multi-Tenant Reverse Proxy:**
  - Custom Docker multi-stage build embedding `caddy-dns/duckdns`.
  - SNI reverse proxy routes for `carinanoxfort.duckdns.org`, `noxfortmonitor.duckdns.org` (including WebSocket `/mqtt` proxying to port 9001).
- **Automation & Operations:**
  - Full `Makefile` for building, starting, stopping, testing, and tailing logs.
  - Comprehensive unit test suite for domain models and use cases (`100% PASS`).
  - Standardized Apache 2.0 license headers across all source files.
  - Full technical documentation suite compatible with MkDocs and Obsidian Vault.

---

<div align="center">
  <img src="docs/assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licensed under the Apache License 2.0.</small>
</div>
