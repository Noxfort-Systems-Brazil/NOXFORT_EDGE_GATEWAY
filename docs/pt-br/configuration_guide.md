---
tags: [configuration, environment, caddyfile, configuracao, pt-br]
aliases: [Guia de Configuração, Variáveis de Ambiente, Diretivas Caddyfile]
---

# ⚙️ Guia de Configuração e Diretivas do Caddyfile

Este guia descreve todas as opções de configuração, variáveis de ambiente (12-Factor) e blocos do Caddyfile no **NOXFORT EDGE GATEWAY**.

⬅️ Voltar para a [Central de Documentação](noxfort_edge_gateway_moc.md) | 🏛️ Ver [Arquitetura](architecture.md) | ⚡ Ver [APIs e Contratos](api_reference.md)

---

## 1. Variáveis de Ambiente

As configurações são lidas do ambiente do sistema operacional ou arquivo `.env`:

| Variável | Valor Padrão | Descrição |
| :--- | :--- | :--- |
| `GATEWAY_HTTP_PORT` | `80` | Porta pública HTTP vinculada pelo proxy reverso. |
| `GATEWAY_HTTPS_PORT` | `443` | Porta pública HTTPS vinculada pelo proxy reverso. |
| `GATEWAY_SYNC_INTERVAL_SEC` | `300` | Intervalo em segundos entre sincronizações de DDNS. |
| `GATEWAY_HEALTH_INTERVAL_SEC` | `15` | Intervalo em segundos entre checagens de saúde ativas. |
| `CARINA_DOMAIN` | `carinanoxfort` | Subdomínio DuckDNS para o CARINA Traffic AI. |
| `CARINA_PORT` | `2201` | Porta local do gateway Web do CARINA. |
| `MONITOR_DOMAIN` | `noxfortmonitor` | Subdomínio DuckDNS para o Noxfort Monitor. |
| `MONITOR_PORT` | `22100` | Porta local do painel Web do Monitor. |

---

## 2. Estrutura do `Caddyfile`

O arquivo [`internal/infrastructure/caddy/Caddyfile`](../../internal/infrastructure/caddy/Caddyfile) gerencia as rotas criptografadas por domínio e paths especializados (como `/amb*`, `/cco*` e WebSockets `/mqtt` para MQTT Mosquitto).

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Logo Noxfort Systems" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado sob Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
