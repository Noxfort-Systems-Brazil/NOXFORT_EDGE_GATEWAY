---
tags: [configuration, environment, caddyfile, env, settings]
aliases: [Configuration Guide, Settings Reference, Environment Specs]
---

# ⚙️ Configuration Guide & Caddyfile Specifications

This guide details all configuration options, 12-Factor environment variables, and Caddy reverse proxy directives available in **NOXFORT EDGE GATEWAY**.

⬅️ Back to [Master Documentation Hub](NOXFORT_EDGE_GATEWAY_MOC.md) | 🏛️ See [Architecture](ARCHITECTURE.md) | ⚡ See [API & Contracts](API_REFERENCE.md)

---

## 1. Environment Variables (12-Factor Setup)

The Go orchestrator loads settings from environment variables with safe fallback defaults defined in [`config/config.go`](../config/config.go).

| Variable | Default Value | Description |
| :--- | :--- | :--- |
| `GATEWAY_HTTP_PORT` | `80` | Public HTTP port bound by the reverse proxy. |
| `GATEWAY_HTTPS_PORT` | `443` | Public HTTPS / TLS port bound by the reverse proxy. |
| `GATEWAY_SYNC_INTERVAL_SEC` | `300` | Period between dual-stack Dynamic DNS sync runs (seconds). |
| `GATEWAY_HEALTH_INTERVAL_SEC` | `15` | Period between active TCP/HTTP upstream health checks (seconds). |
| `CARINA_DOMAIN` | `carinanoxfort` | DuckDNS subdomain for CARINA Traffic AI platform. |
| `CARINA_TOKEN` | `ff528f...` | DuckDNS authentication token for CARINA. |
| `CARINA_PORT` | `2201` | Private host port where CARINA's web gateway listens. |
| `MONITOR_DOMAIN` | `noxfortmonitor` | DuckDNS subdomain for Noxfort Monitor platform. |
| `MONITOR_TOKEN` | `2c42c3...` | DuckDNS authentication token for Noxfort Monitor. |
| `MONITOR_PORT` | `22100` | Private host port where Noxfort Monitor listens. |

---

## 2. Reverse Proxy Directives (`Caddyfile`)

The reverse proxy configuration is maintained in [`internal/infrastructure/caddy/Caddyfile`](../internal/infrastructure/caddy/Caddyfile).

### 2.1 Domain Block: CARINA CCO & Ambulances
```caddy
{$CARINA_DOMAIN:carinanoxfort}.duckdns.org {
    tls {
        dns duckdns {$CARINA_TOKEN:ff528f24-1dd3-4c84-9c12-02b6d9805440}
    }

    log {
        output stdout
        format console
    }

    # Emergency Vehicle Preemption (EVP) Route
    handle /amb* {
        reverse_proxy host.docker.internal:2201 {
            header_up Host {host}
            header_up X-Real-IP {remote_host}
        }
    }

    # CCO Videowall Dashboard Route
    handle /cco* {
        reverse_proxy host.docker.internal:2201 {
            header_up Host {host}
            header_up X-Real-IP {remote_host}
        }
    }

    # Default catch-all
    handle {
        reverse_proxy host.docker.internal:2201 {
            header_up Host {host}
            header_up X-Real-IP {remote_host}
        }
    }
}
```

### 2.2 Domain Block: Noxfort Monitor & MQTT WebSocket
```caddy
{$MONITOR_DOMAIN:noxfortmonitor}.duckdns.org {
    tls {
        dns duckdns {$MONITOR_TOKEN:2c42c374-2143-4292-93c3-5baccb6446c4}
    }

    # WebSocket bridge for Mosquitto MQTT over WSS
    @mqtt_ws {
        path /mqtt
    }
    reverse_proxy @mqtt_ws host.docker.internal:9001

    # Ingestion & UI Web Endpoint
    reverse_proxy host.docker.internal:22100 {
        header_up Host {host}
        header_up X-Real-IP {remote_host}
    }
}
```

---

## 3. Adding a New Tenant Service

To register a new upstream tenant (e.g. `synapse`):
1. **Declare Environment Variables:**
   ```bash
   SYNAPSE_DOMAIN=synapsenoxfort
   SYNAPSE_TOKEN=<duckdns-token>
   SYNAPSE_PORT=22200
   ```
2. **Add Domain Block to `Caddyfile`:**
   ```caddy
   {$SYNAPSE_DOMAIN:synapsenoxfort}.duckdns.org {
       tls {
           dns duckdns {$SYNAPSE_TOKEN}
       }
       reverse_proxy host.docker.internal:22200
   }
   ```
3. **Register in `config/config.go`:**
   Instantiate `models.NewUpstreamService` and `models.NewDDNSTarget` inside `LoadConfig()`.

---

<div align="center">
  <img src="assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licensed under the Apache License 2.0.</small>
</div>
