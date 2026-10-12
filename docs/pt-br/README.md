---
tags: [readme, home, gateway, edge, pt-br]
aliases: [Central de Documentação, Visão Geral PT-BR]
---

<div align="center">

# NOXFORT EDGE GATEWAY
### Gateway de Borda Unificado Multi-Tenant & Proxy Reverso de Alta Disponibilidade
*Noxfort Systems — A State Of Art Company*

[![Status](https://img.shields.io/badge/Status-Ativo-brightgreen?style=flat&logo=github)](https://github.com/Noxfort-Systems-Brazil)
[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Caddy](https://img.shields.io/badge/Caddy-2.8%2B-1F88C0?style=flat&logo=caddy&logoColor=white)](https://caddyserver.com/)
[![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![Licença](https://img.shields.io/badge/Licen%C3%A7a-Apache_2.0-blue?style=flat)](../../LICENSE)

---

🌐 **Traduções / Idiomas:** **[🇧🇷 Português](README.md)** • **[🇺🇸 English](../../README.md)** • **[🇪🇸 Español](../es/README.md)** • **[🇫🇷 Français](../fr/README.md)** • **[🇷🇺 Русский](../ru/README.md)** • **[🇨🇳 简体中文](../zh/README.md)** • **[📚 Portal Principal](../index.md)**

---

</div>

O **NOXFORT EDGE GATEWAY** é o ponto de entrada central de rede de alta disponibilidade da Noxfort Systems para plataformas de cidades inteligentes e mobilidade urbana. Atuando como o **porteiro exclusivo das portas públicas 80 (HTTP) e 443 (HTTPS)** na máquina servidora do CCO municipal, ele provê terminação TLS automática, roteamento multi-tenant por SNI, sincronização periódica de DNS dinâmico dual-stack (DuckDNS IPv4/IPv6) e sondas ativas de integridade para serviços em portas locais privadas:

- 🚦 **CARINA (CCO Videowall, Semáforos e Ambulâncias EVP):** `carinanoxfort.duckdns.org` ➔ `127.0.0.1:2201`
- 📊 **NOXFORT MONITOR (Telemetria Industrial & Incidentes):** `noxfortmonitor.duckdns.org` ➔ `127.0.0.1:22100` e WSS `9001` (MQTT)
- 🧠 **SYNAPSE (IA de Alta Frequência & Orquestração):** `synapse.duckdns.org` ➔ Porta upstream do Synapse

---

## 📚 Central de Conhecimento e Módulos Técnicos

| Cartão / Subsistema | Área de Foco | Link Direto |
| :--- | :--- | :---: |
| 📚 **Master MOC** | Índice Principal, Obsidian Vault e Mapa do Código | [Explorar Hub](noxfort_edge_gateway_moc.md) |
| 🏛️ **Arquitetura do Sistema** | Camadas DDD, Topologia de Ingress e Proxy Reverso Caddy | [Ver Arquitetura](architecture.md) |
| ⚡ **Referência de APIs e Contratos** | Caddy Admin API, Protocolo DuckDNS e DTOs de Status | [Ver Referência](api_reference.md) |
| 🌐 **Sincronização DDNS** | Mecanismo DuckDNS, Detecção de IPv4 Público e IPv6 Global | [Ver Guia DDNS](ddns_synchronization.md) |
| 🛡️ **Monitoramento de Saúde** | Sondas Periódicas TCP/HTTP e Máquina de Estados Upstream | [Ver Guia Saúde](health_monitoring.md) |
| ⚙️ **Guia de Configuração** | Variáveis de Ambiente, Caddyfile e Mapeamento de Portas | [Ver Configurações](configuration_guide.md) |
| 🛠️ **Guia do Desenvolvedor** | Toolchain Go, Adição de Upstreams e Padrões de Código | [Ver Guias](developer_guides.md) |
| 🧪 **Testes & Validação** | Testes Unitários Go, Contratos de Domínio e Mocks | [Ver Diretrizes](testing.md) |
| 🚀 **Implantação & Operação** | Docker Compose, Serviço Systemd e Operação em Produção | [Ver Implantação](deployment_and_packaging.md) |

---

## ⚡ Destaques Arquiteturais

- **Domain-Driven Design (DDD) & Clean Architecture:** Isolamento total das regras de negócio (Domínio e Casos de Uso) de frameworks externos e drivers de rede.
- **Portas Públicas Exclusivas (80/443):** Impede colisões de portas no servidor central ao multiplexar o tráfego via Server Name Indication (SNI).
- **Gestão Automática de Certificados SSL/TLS:** Emissão e renovação automática via Let's Encrypt / ZeroSSL usando desafio DNS-01 DuckDNS.
- **Sincronizador Dual-Stack:** Detecção autônoma de IPv4 público e IPv6 global com retentativa exponencial.
- **Monitoramento Ativo de Microsserviços:** Sondas TCP e HTTP contínuas a cada 15 segundos com métricas de latência em milissegundos.

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Logo Noxfort Systems" width="45" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <i>Controlled Artificial Road-traffic Intelligence Network Architecture • NOXFORT EDGE GATEWAY v1.0.0</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado sob Apache 2.0.</small>
</div>
