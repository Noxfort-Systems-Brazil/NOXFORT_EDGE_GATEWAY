---
tags: [api, reference, caddy, duckdns, contracts, dto]
aliases: [API Reference, Protocol Specification, Contracts Reference]
---

# ⚡ API & Contracts Reference

This document provides technical specifications for the internal APIs, networking protocols, domain contracts, and JSON telemetry schemas used across the **NOXFORT EDGE GATEWAY**.

⬅️ Back to [Master Documentation Hub](NOXFORT_EDGE_GATEWAY_MOC.md) | 🏛️ See [Architecture](ARCHITECTURE.md) | 🌐 See [DDNS Synchronization](DDNS_SYNCHRONIZATION.md)

---

## 1. Caddy Reverse Proxy Admin API

The Caddy reverse proxy container exposes an internal HTTP administration interface bound strictly to `localhost:2019` inside the bridge network.

| Endpoint | Method | Description | Payload |
| :--- | :---: | :--- | :--- |
| `http://localhost:2019/load` | `POST` | Atomically reloads the running configuration without dropping connections | `Content-Type: text/caddyfile` |
| `http://localhost:2019/config/` | `GET` | Fetches active running configuration in JSON format | None |
| `http://localhost:2019/stop` | `POST` | Initiates graceful shutdown of the Caddy process | None |

### Dynamic Reload Example (cURL)
```bash
curl "http://localhost:2019/load" \
  -H "Content-Type: text/caddyfile" \
  --data-binary @internal/infrastructure/caddy/Caddyfile
```

---

## 2. DuckDNS Synchronization Protocol

The gateway synchronizes registered domains using DuckDNS's REST API over HTTPS:

```text
GET https://www.duckdns.org/update?domains={DOMAIN}&token={TOKEN}&ip={IPV4}&ipv6={IPV6}&verbose=true
```

### Query Parameters
- `domains` *(string, required)*: Subdomain identifier registered under `duckdns.org`.
- `token` *(UUID, required)*: Account authentication token.
- `ip` *(string, optional)*: Public IPv4 address.
- `ipv6` *(string, optional)*: Public global unicast IPv6 address.
- `verbose` *(string, optional)*: Enables verbose plain-text response status.

### Expected Responses
- `OK` / `OK\n{IP}\n{IPV6}\nUPDATED`: Successful domain record update.
- `KO`: Authentication failed or domain not found.

---

## 3. Telemetry & Gateway Status DTO

The Go orchestrator models runtime operational status using the Data Transfer Object `GatewayStatusDTO` (`internal/application/dtos/gateway_status_dto.go`):

```json
{
  "timestamp": "2026-10-12T13:45:00Z",
  "uptime_seconds": 86400,
  "total_routes": 2,
  "healthy_routes": 2,
  "upstreams": [
    {
      "id": "carina",
      "name": "CARINA Traffic AI",
      "host": "127.0.0.1",
      "port": 2201,
      "health_path": "/health",
      "status": "HEALTHY",
      "last_latency_ms": 1.25,
      "consecutive_fails": 0,
      "last_checked_at": "2026-10-12T13:44:50Z"
    },
    {
      "id": "monitor",
      "name": "Noxfort Monitor",
      "host": "127.0.0.1",
      "port": 22100,
      "health_path": "/api/health",
      "status": "HEALTHY",
      "last_latency_ms": 0.82,
      "consecutive_fails": 0,
      "last_checked_at": "2026-10-12T13:44:50Z"
    }
  ],
  "ddns_targets": [
    {
      "subdomain": "carinanoxfort",
      "full_domain": "carinanoxfort.duckdns.org",
      "provider": "duckdns",
      "last_ipv4": "189.1.2.3",
      "last_ipv6": "2804:14d:5c82::1",
      "last_sync_at": "2026-10-12T13:40:00Z",
      "is_synced": true
    }
  ]
}
```

---

## 4. Domain Contracts (Go Interfaces)

### 4.1 `contracts.IDDNSSynchronizer`
```go
type IDDNSSynchronizer interface {
    SyncTarget(ctx context.Context, target *models.DDNSTarget, ipv4, ipv6 string) (*SyncResult, error)
}
```

### 4.2 `contracts.IHealthMonitor`
```go
type IHealthMonitor interface {
    CheckHealth(ctx context.Context, service *models.UpstreamService) (*CheckResult, error)
}
```

### 4.3 `contracts.IProxyRouter`
```go
type IProxyRouter interface {
    ApplyRoutes(ctx context.Context, routes []models.RouteRule) error
    Reload(ctx context.Context) error
}
```

---

<div align="center">
  <img src="assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licensed under the Apache License 2.0.</small>
</div>
