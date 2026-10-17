---
tags: [home, index, gateway, portal, hub]
aliases: [Documentation Index, Overview, Portal]
---

<div align="center">

# NOXFORT EDGE GATEWAY Documentation Hub
### Technical Master Portal & Architecture Index
*Noxfort Systems — A State Of Art Company*

</div>

Welcome to the technical documentation library for **NOXFORT EDGE GATEWAY** — the unified edge ingress reverse proxy and high-availability multi-tenant gatekeeper for Noxfort Systems smart city platforms.

Designed to operate seamlessly on **GitHub** and as an **[Obsidian](https://obsidian.md/) Knowledge Vault**, this documentation suite covers Clean Architecture in Go, Caddy SNI reverse proxying, Let's Encrypt / ZeroSSL TLS automation, dual-stack DuckDNS auto-synchronization, and active health monitoring.

---

## 🗺️ Master Navigation & Modules

| Subsystem / Dimension | Focus Area | Direct Link |
| :--- | :--- | :---: |
| 📚 **Master Map of Content** | Primary Obsidian Hub & Codebase Directory Map | [Explore Hub](NOXFORT_EDGE_GATEWAY_MOC.md) |
| 🏛️ **System Architecture** | Clean Architecture (DDD), Port 80/443 Ingress & Reverse Proxy Topology | [View Blueprint](ARCHITECTURE.md) |
| ⚡ **API & Contracts Reference** | Caddy Admin API, DuckDNS Protocol, DTOs & Health Contract Specs | [View API Reference](API_REFERENCE.md) |
| 🌐 **DDNS Synchronization** | DuckDNS Engine, Public IPv4 & Global IPv6 Interface Discovery | [View DDNS Guide](DDNS_SYNCHRONIZATION.md) |
| 🛡️ **Active Health Monitoring** | Periodic TCP/HTTP Health Probes & Upstream Availability Machine | [View Health Guide](HEALTH_MONITORING.md) |
| ⚙️ **Configuration Guide** | Environment Variables, Caddyfile Customization & Port Mappings | [View Configuration](CONFIGURATION_GUIDE.md) |
| 🛠️ **Developer Guides** | Go Toolchain, Adding Upstream Tenants & Contribution Workflow | [View Guides](DEVELOPER_GUIDES.md) |
| 🧪 **Testing & Validation** | Go Unit Tests, Race Detector, Domain Contracts & Mock Verification | [View Testing](TESTING.md) |
| 🚀 **Deployment & Operations** | Docker Compose, Systemd Units & Production Hardening | [View Deployment](DEPLOYMENT_AND_PACKAGING.md) |

---

## ⚡ Quick Architecture Summary

```text
       ┌────────────────────────────────────────────────────────┐
       │             NOXFORT EDGE GATEWAY (Caddy Container)     │
       │       (Ports 80 & 443 — Automatic TLS via DuckDNS)     │
       └──────────────────────────┬─────────────────────────────┘
                                  │
      ┌───────────────────────────┼───────────────────────────┐
      ▼                           ▼                           ▼
┌──────────────┐          ┌──────────────┐          ┌──────────────────┐
│CARINA Traffic│          │Noxfort Monitor│          │   Synapse AI     │
│(Port 2201)   │          │(Ports 22100 &│          │ (Upstream Port)  │
│              │          │ 9001 WSS)    │          │                  │
└──────────────┘          └──────────────┘          └──────────────────┘
```

---

<div align="center">
  <img src="assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="48" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <i>Controlled Artificial Road-traffic Intelligence Network Architecture • NOXFORT EDGE GATEWAY v1.0.0</i><br/>
  <small>Licensed under the <a href="../LICENSE">Apache License 2.0</a> (SPDX: Apache-2.0). © 2026 Noxfort Systems.</small>
</div>
