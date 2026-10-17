---
tags: [architecture, gateway, clean-architecture, ddd, caddy, reverse-proxy, solid]
aliases: [System Architecture, Gateway Blueprint, Clean Architecture Topology]
---

# 🏛️ NOXFORT EDGE GATEWAY: System Blueprint & Architecture

This document specifies the internal engineering architecture of **NOXFORT EDGE GATEWAY**. It details the Domain-Driven Design (DDD) Clean Architecture layers, the multi-tenant SNI ingress routing model, automated dual-stack Dynamic DNS (DuckDNS) synchronization, active health monitoring, and containerized deployment topology.

⬅️ Back to [Main Documentation Hub](docs/NOXFORT_EDGE_GATEWAY_MOC.md) | ⚡ See [API & Contracts](docs/API_REFERENCE.md) | 🌐 See [DDNS Synchronization](docs/DDNS_SYNCHRONIZATION.md) | 🛡️ See [Health Monitoring](docs/HEALTH_MONITORING.md)

---

## 1. Network Topology & Ingress Concurrency Model

In municipal traffic control centres (CCOs) and edge server nodes, multiple independent software platforms (CARINA Traffic AI, Noxfort Monitor, Synapse HFT) must coexist securely on the same host operating system without colliding on standard web ports.

The **NOXFORT EDGE GATEWAY** resolves this by seizing public TCP ports **80** (HTTP / ACME Challenge) and **443** (HTTPS / TLS SNI) through a custom Caddy container, multiplexing encrypted traffic directly to backend services running on loopback ports:

```mermaid
graph TD
    ClientInternet([Internet / CCO Clients / Field Devices]) -->|HTTP Port 80 / HTTPS Port 443| Ingress[NOXFORT EDGE GATEWAY Caddy]
    
    subgraph Edge Reverse Proxy Layer
        Ingress -->|SNI: carinanoxfort.duckdns.org| R1[Route Handler: CARINA]
        Ingress -->|SNI: noxfortmonitor.duckdns.org| R2[Route Handler: MONITOR]
        Ingress -->|SNI: synapse.duckdns.org| R3[Route Handler: SYNAPSE]
        Ingress -->|Path: /mqtt| R_WSS[WebSocket WSS Proxy]
    end

    subgraph Host Private Services
        R1 -->|host.docker.internal:2201| CarinaCore[CARINA Web Gateway & CCO API]
        R2 -->|host.docker.internal:22100| NoxfortMon[Noxfort Monitor Web Dashboard]
        R_WSS -->|host.docker.internal:9001| Mosquitto[Mosquitto MQTT Broker WSS]
        R3 -->|host.docker.internal:SynapsePort| SynapseAI[SYNAPSE HFT AI Engine]
    end

    subgraph Go Orchestration Daemon
        GoDaemon[noxfort-gateway Daemon Process] -->|Sync Every 300s| DDNS[DuckDNS API IPv4 & IPv6]
        GoDaemon -->|Health Probe Every 15s| HealthChecker[TCP/HTTP Active Checker]
        HealthChecker --> CarinaCore
        HealthChecker --> NoxfortMon
    end
```

---

## 2. Clean Architecture & DDD Layering (Go)

The native Go orchestrator (`cmd/gateway/main.go`) is structured strictly according to Domain-Driven Design and Clean Architecture principles, ensuring that high-level enterprise business rules remain entirely agnostic of third-party frameworks, network drivers, or DNS providers.

```mermaid
graph BT
    subgraph Domain Layer
        Entities[Entities & Value Objects: UpstreamService, DDNSTarget, RouteRule]
        Contracts[Domain Contracts: IDDNSSynchronizer, IHealthMonitor, IProxyRouter]
    end

    subgraph Application Layer
        U1[SyncAllDDNSUseCase] --> Contracts
        U2[MonitorHealthUseCase] --> Contracts
        U3[GetGatewayStatusUseCase] --> Entities
        DTOs[DTO: GatewayStatusDTO]
    end

    subgraph Infrastructure Layer
        DuckDNSClient[DuckDNSClient] -->|Implements| Contracts
        IPDetectors[IPv4Detector, GlobalIPv6Detector] --> DuckDNSClient
        HealthChecker[TCPHTTPHealthChecker] -->|Implements| Contracts
        CaddyProxy[Caddy Reverse Proxy] -->|Implements| Contracts
    end

    subgraph Presentation / Composition Root
        Main[cmd/gateway/main.go] --> ApplicationLayer
        Config[config/config.go] --> Main
    end
```

### 2.1 Domain Layer (`internal/domain`)
- **Entities & Value Objects:**
  - `UpstreamService`: Represents a private backend target with mutable availability state (`StatusHealthy`, `StatusUnhealthy`, `StatusDegraded`), check latency, and consecutive failure counters.
  - `DDNSTarget`: Immutable Value Object holding the dynamic DNS registration parameters, public FQDN, DuckDNS token, and resolved IPv4/IPv6 addresses.
  - `RouteRule`: Represents an SNI routing configuration linking an external domain to a target service ID.
- **Contracts (`internal/domain/contracts`):**
  - `IDDNSSynchronizer`: Contract defining idempotent DNS update capabilities over IPv4/IPv6.
  - `IHealthMonitor`: Contract defining active probing mechanisms (TCP connect + HTTP GET).
  - `IProxyRouter`: Contract abstracting dynamic route reloading and Caddy admin API interactions.

### 2.2 Application Layer (`internal/application`)
- **Use Cases:**
  - `SyncAllDDNSUseCase`: Iterates through all configured DDNS targets, resolves external public IP addresses (IPv4 via `api.ipify.org` and IPv6 via local network interface inspection), and pushes updates to the DuckDNS API concurrently.
  - `MonitorHealthUseCase`: Dispatches parallel health checks against all registered upstream microservices, calculating latency and updating domain status.
  - `GetGatewayStatusUseCase`: Aggregates the operational health of all routes, DDNS synchronizations, and upstream services into a unified `GatewayStatusDTO`.

### 2.3 Infrastructure Layer (`internal/infrastructure`)
- **`ddns/`:**
  - `DuckDNSClient`: Implements `IDDNSSynchronizer`, communicating over HTTPS with `https://www.duckdns.org/update`.
  - `IPv4Detector`: Fetches public egress IPv4 using redundant public endpoints with strict timeouts.
  - `GlobalIPv6Detector`: Inspects local OS network interfaces (`net.Interfaces`) to identify valid public global unicast IPv6 addresses (excluding link-local `fe80::` and loopback `::1`).
- **`health/`:**
  - `TCPHTTPHealthChecker`: Performs a dual-phase verification: first validates raw TCP handshake on host:port, followed by an HTTP GET against the designated health endpoint (e.g., `/health`).
- **`caddy/`:**
  - Custom multi-stage `Dockerfile` building Caddy with `github.com/caddy-dns/duckdns`.
  - Dynamic `Caddyfile` managing multi-tenant SSL/TLS certificates and reverse proxy blocks.

---

## 3. Dynamic DNS Synchronization Workflow

```mermaid
sequenceDiagram
    autonumber
    participant Daemon as Go Gateway Loop
    participant UseCase as SyncAllDDNSUseCase
    participant Det4 as IPv4Detector
    participant Det6 as GlobalIPv6Detector
    participant Duck as DuckDNS API

    Daemon->>UseCase: Execute(ctx, targets)
    par Discover Public IP
        UseCase->>Det4: DetectIPv4(ctx)
        Det4-->>UseCase: Return Public IPv4
    and Discover Public IPv6
        UseCase->>Det6: DetectIPv6(ctx)
        Det6-->>UseCase: Return Public IPv6
    end
    loop For each target in targets
        UseCase->>Duck: Update(domain, token, ipv4, ipv6)
        Duck-->>UseCase: "OK" or Error
    end
    UseCase-->>Daemon: Consolidated SyncResult
```

---

## 4. High-Availability & Self-Healing Safeguards

1. **Atomic Health State Transitions:** Upstream services transition from `HEALTHY` to `UNHEALTHY` only after crossing the configured failure threshold, preventing flap anomalies during brief microservice restarts.
2. **Exponential Backoff on DDNS Outages:** When DuckDNS or external IP detector endpoints return network errors, the synchronization engine logs the incident and retries without crashing the daemon process.
3. **Caddy Zero-Downtime Hot Reloads:** Route updates and certificate renewals are processed in-memory via Caddy's internal Admin API (`localhost:2019`) without dropping active TLS connections.
4. **Docker Container Liveness:** Container health is supervised by Docker daemon with `restart: unless-stopped` policy.

---

<div align="center">
  <img src="docs/assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licensed under Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
