---
tags: [api, reference, caddy, duckdns, contracts, dto, pt-br]
aliases: [Referência de APIs, Especificação de Protocolos, Contratos Go]
---

# ⚡ Referência de APIs, Protocolos e Contratos

Este documento fornece especificações técnicas para as APIs internas, protocolos de rede, contratos de domínio e esquemas JSON de telemetria utilizados no **NOXFORT EDGE GATEWAY**.

⬅️ Voltar para a [Central de Documentação](noxfort_edge_gateway_moc.md) | 🏛️ Ver [Arquitetura](architecture.md) | 🌐 Ver [Sincronização DDNS](ddns_synchronization.md)

---

## 1. API Administrativa do Caddy

O container Caddy expõe uma interface de administração HTTP restrita a `localhost:2019`.

| Endpoint | Método | Descrição | Payload |
| :--- | :---: | :--- | :--- |
| `http://localhost:2019/load` | `POST` | Recarrega atomicamente as rotas e certificados sem derrubar conexões ativas | `Content-Type: text/caddyfile` |
| `http://localhost:2019/config/` | `GET` | Retorna a configuração em execução no formato JSON | Nenhum |
| `http://localhost:2019/stop` | `POST` | Solicita encerramento gracioso do processo Caddy | Nenhum |

---

## 2. Protocolo de Sincronização DuckDNS

A sincronização de domínios ocorre via requisição HTTPS REST:

```text
GET https://www.duckdns.org/update?domains={DOMINIO}&token={TOKEN}&ip={IPV4}&ipv6={IPV6}&verbose=true
```

### Respostas Esperadas
- `OK`: Registro DNS atualizado com êxito.
- `KO`: Falha de autenticação ou subdomínio não encontrado.

---

## 3. DTO de Telemetria e Status do Gateway

O daemon Go consolida as métricas no objeto `GatewayStatusDTO`:

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
    }
  ]
}
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Logo Noxfort Systems" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado sob Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
