<div align="center">

# NOXFORT EDGE GATEWAY — Suite de Documentación Técnica
### Gateway de Borde Unificado Multi-Tenant y Proxy Inverso de Alta Disponibilidad
*Noxfort Systems — A State Of Art Company*

[![Status](https://img.shields.io/badge/Status-Activo-brightgreen?style=flat&logo=github)](https://github.com/Noxfort-Systems-Brazil)
[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Caddy](https://img.shields.io/badge/Caddy-2.8%2B-1F88C0?style=flat&logo=caddy&logoColor=white)](https://caddyserver.com/)
[![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![Licencia](https://img.shields.io/badge/Licencia-Apache_2.0-blue?style=flat)](../../LICENSE)

---

🌐 **Traducciones / Idiomas:** **[🇺🇸 English](../../README.md)** • **[🇧🇷 Português](../pt-br/README.md)** • **[🇪🇸 Español](README.md)** • **[🇫🇷 Français](../fr/README.md)** • **[🇷🇺 Русский](../ru/README.md)** • **[🇨🇳 简体中文](../zh/README.md)** • **[📚 Centro de Documentación](../index.md)**

---

</div>

## Bienvenido a la Documentación Oficial de NOXFORT EDGE GATEWAY

**NOXFORT EDGE GATEWAY** es el punto central de entrada a la red y proxy inverso unificado de alta disponibilidad para las plataformas de movilidad urbana inteligente de Noxfort Systems. Actuando como el **guardián exclusivo de los puertos públicos 80 (HTTP) y 443 (HTTPS)** en los servidores centrales de los Centros de Control Operativo (CCO), proporciona terminación TLS automatizada, enrutamiento multi-tenant por SNI, sincronización periódica de DNS dinámico dual-stack (DuckDNS IPv4/IPv6) y sondeo activo de salud para microservicios locales:

- 🚦 **CARINA (CCO Videowall, Semáforos e IA Adaptativa):** `carinanoxfort.duckdns.org` ➔ `127.0.0.1:2201`
- 📊 **NOXFORT MONITOR (Telemetría Industrial e Incidentes):** `noxfortmonitor.duckdns.org` ➔ `127.0.0.1:22100` y WSS `9001` (MQTT)
- 🧠 **SYNAPSE (IA de Alta Frecuencia y Orquestación):** `synapse.duckdns.org` ➔ Puerto privado de Synapse

---

## Directorio de Guías Técnicas

| Documento | Tema | Contenido Principal |
| :--- | :--- | :--- |
| 📖 **[Arquitectura del Sistema](architecture.md)** | Arquitectura y Topología | Clean Architecture en Go (DDD), multiplexación por SNI, puertos 80/443, Caddy y aislamiento de servicios. |
| ⚡ **[Referencia de APIs y Contratos](api_reference.md)** | Especificaciones Técnicas | API de administración de Caddy, protocolo REST de DuckDNS y esquemas JSON del DTO `GatewayStatusDTO`. |
| 🧪 **[Pruebas y Control de Calidad](testing.md)** | Aseguramiento de Calidad | Pruebas unitarias puras en Go (`go test`), simulación con mocks sin dependencias de red y detector de carreras `-race`. |
| 🚀 **[Implementación y Operación](deployment_and_packaging.md)** | Despliegue en Producción | Orquestación con Docker Compose (`restart: unless-stopped`), volumen persistente de certificados y servicio daemon Systemd. |

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="45" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <i>Ingeniería de Movilidad Inteligente • NOXFORT EDGE GATEWAY v1.0.0</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado bajo Apache 2.0.</small>
</div>
