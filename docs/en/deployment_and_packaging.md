---
tags: [deployment, packaging, docker, systemd, production, devops]
aliases: [Deployment Guide, Operations Manual, Packaging Specs]
---

# 🚀 Deployment & Operations Guide

This document describes production packaging, containerization, systemd daemonization, and operational management for **NOXFORT EDGE GATEWAY**.

⬅️ Back to [Master Documentation Hub](NOXFORT_EDGE_GATEWAY_MOC.md) | 🏛️ See [Architecture](ARCHITECTURE.md) | ⚙️ See [Configuration](CONFIGURATION_GUIDE.md)

---

## 1. Production Architecture Overview

In a typical production deployment (municipal server or industrial edge PC), the gateway deployment consists of:
1. **Containerized Ingress (Docker):** Caddy compiled with the `caddy-dns/duckdns` plugin, listening on public host ports 80 & 443.
2. **Native Go Orchestrator Daemon:** Native binary running under systemd or cron, periodically performing dynamic DNS sync and health monitoring.

---

## 2. Containerized Ingress (`docker-compose.yml`)

The reverse proxy is orchestrated via Docker Compose:

```yaml
version: '3.8'

services:
  caddy:
    build:
      context: .
      dockerfile: internal/infrastructure/caddy/Dockerfile
    container_name: noxfort-edge-gateway
    restart: unless-stopped
    ports:
      - "80:80"
      - "443:443"
    environment:
      - CARINA_DOMAIN=${CARINA_DOMAIN:-carinanoxfort}
      - CARINA_TOKEN=${CARINA_TOKEN:-ff528f24-1dd3-4c84-9c12-02b6d9805440}
      - MONITOR_DOMAIN=${MONITOR_DOMAIN:-noxfortmonitor}
      - MONITOR_TOKEN=${MONITOR_TOKEN:-2c42c374-2143-4292-93c3-5baccb6446c4}
    volumes:
      - ./internal/infrastructure/caddy/Caddyfile:/etc/caddy/Caddyfile:ro
      - caddy_data:/data
      - caddy_config:/config
      - /etc/localtime:/etc/localtime:ro
    extra_hosts:
      - "host.docker.internal:host-gateway"

volumes:
  caddy_data:
  caddy_config:
```

### Key Production Settings
- `caddy_data` volume: Persists TLS certificates issued by Let's Encrypt / ZeroSSL, preventing ACME rate-limiting upon container restarts.
- `extra_hosts`: Maps `host.docker.internal` to the host's bridge IP, allowing the container to proxy traffic to private loopback services on the host machine.
- `restart: unless-stopped`: Ensures the gateway restarts automatically on server reboot or daemon crash.

---

## 3. Running as a Systemd Service (Go Daemon)

To keep the Go orchestrator daemon running continuously as a native background service:

1. **Compile Native Binary:**
   ```bash
   make build
   sudo cp bin/noxfort-gateway /usr/local/bin/noxfort-gateway
   ```
2. **Create Systemd Unit File (`/etc/systemd/system/noxfort-gateway.service`):**
   ```ini
   [Unit]
   Description=Noxfort Edge Gateway Orchestrator
   After=network-online.target docker.service
   Wants=network-online.target

   [Service]
   Type=simple
   User=root
   WorkingDirectory=/home/gabriel-moraes/Documentos/NOXFORT_EDGE_GATEWAY
   ExecStart=/usr/local/bin/noxfort-gateway
   Restart=always
   RestartSec=5s
   Environment=GATEWAY_SYNC_INTERVAL_SEC=300
   Environment=GATEWAY_HEALTH_INTERVAL_SEC=15

   [Install]
   WantedBy=multi-user.target
   ```
3. **Enable and Start:**
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable --now noxfort-gateway
   sudo systemctl status noxfort-gateway
   ```

---

## 4. Operational Runbook

### Checking Live Logs
```bash
make logs
```

### Rotating SSL/TLS Certificates
Caddy manages ACME certificate renewals automatically 30 days prior to expiration. If manual renewal is forced:
```bash
docker exec -it noxfort-edge-gateway caddy reload --config /etc/caddy/Caddyfile --force
```

### Graceful Restart
```bash
make stop
make start
```

---

<div align="center">
  <img src="assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licensed under the Apache License 2.0.</small>
</div>
