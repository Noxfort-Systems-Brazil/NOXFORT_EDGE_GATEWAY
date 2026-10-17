<div align="center">

# NOXFORT EDGE GATEWAY — Suite de Documentation Technique
### Passerelle de Bordure Multi-Tenant & Proxy Inverse Haute Disponibilité
*Noxfort Systems — A State Of Art Company*

[![Status](https://img.shields.io/badge/Status-Actif-brightgreen?style=flat&logo=github)](https://github.com/Noxfort-Systems-Brazil)
[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Caddy](https://img.shields.io/badge/Caddy-2.8%2B-1F88C0?style=flat&logo=caddy&logoColor=white)](https://caddyserver.com/)
[![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![Licence: Apache-2.0](https://img.shields.io/badge/License-Apache--2.0-blue.svg)](../../LICENSE)

---

🌐 **Traductions / Langues :** **[🇺🇸 English](../../README.md)** • **[🇧🇷 Português](../pt-br/README.md)** • **[🇪🇸 Español](../es/README.md)** • **[🇫🇷 Français](README.md)** • **[🇷🇺 Русский](../ru/README.md)** • **[🇨🇳 简体中文](../zh/README.md)** • **[📚 Hub Central](../index.md)**

---

</div>

## Bienvenue dans la Documentation Officielle de NOXFORT EDGE GATEWAY

**NOXFORT EDGE GATEWAY** constitue le point d'entrée réseau et le proxy inverse haute disponibilité pour les infrastructures de mobilité urbaine intelligente de Noxfort Systems. Agissant comme le **gardien exclusif des ports publics 80 (HTTP) et 443 (HTTPS)** sur les serveurs de centres de contrôle opérationnel (CCO), il assure la terminaison TLS automatique, le routage multi-tenant par SNI, la synchronisation périodique de DNS dynamique double-pile (DuckDNS IPv4/IPv6) et la surveillance active des microservices locaux :

- 🚦 **CARINA (CCO Videowall, Feux Intelligents & IA Adaptative) :** `carinanoxfort.duckdns.org` ➔ `127.0.0.1:2201`
- 📊 **NOXFORT MONITOR (Télémétrie Industrielle & Incidents) :** `noxfortmonitor.duckdns.org` ➔ `127.0.0.1:22100` et WSS `9001` (MQTT)
- 🧠 **SYNAPSE (IA Haute Fréquence & Orchestration) :** `synapse.duckdns.org` ➔ Port privé Synapse

---

## Répertoire des Guides Techniques

| Document | Sujet | Contenu Principal |
| :--- | :--- | :--- |
| 📖 **[Architecture Système](architecture.md)** | Architecture et Topologie | Clean Architecture en Go (DDD), multiplexage SNI, ports 80/443, Caddy et isolation des services. |
| ⚡ **[Référence API et Contrats](api_reference.md)** | Spécifications Techniques | API d'administration Caddy, protocole REST DuckDNS et schémas JSON du DTO `GatewayStatusDTO`. |
| 🧪 **[Tests et Assurance Qualité](testing.md)** | Validation et Tests | Tests unitaires en Go (`go test`), isolation réseau par mocks et détection des accès concurrents `-race`. |
| 🚀 **[Déploiement et Opérations](deployment_and_packaging.md)** | Exploitation Production | Orchestration Docker Compose, volumes persistants de certificats et service Systemd. |

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="45" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <i>Ingénierie de la Mobilité Intelligente • NOXFORT EDGE GATEWAY v1.0.0</i><br/>
  <small>© 2026 Noxfort Systems. Sous licence Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
