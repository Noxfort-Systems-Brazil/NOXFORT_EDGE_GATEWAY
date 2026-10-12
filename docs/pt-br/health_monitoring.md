---
tags: [health, monitoring, tcp, http, saude, sondas, pt-br]
aliases: [Monitoramento de Saúde, Sondas Ativas, Máquina de Estados]
---

# 🛡️ Subsistema de Monitoramento Ativo de Integridade

Este documento especifica a engenharia de sondagem ativa e a máquina de estados de disponibilidade dos serviços upstream do **NOXFORT EDGE GATEWAY**.

⬅️ Voltar para a [Central de Documentação](noxfort_edge_gateway_moc.md) | 🏛️ Ver [Arquitetura](architecture.md) | ⚡ Ver [APIs e Contratos](api_reference.md)

---

## 1. Verificação em Duas Fases

O gateway não presume que um microsserviço está saudável apenas por manter uma porta aberta. Ele executa uma verificação em duas etapas:

1. **Fase 1 (TCP Handshake):** Executa `net.DialTimeout` de 3 segundos no IP e porta privada do serviço (ex: `127.0.0.1:2201`). Se a porta estiver fechada ou recusar conexão, o serviço é imediatamente classificado como `UNHEALTHY`.
2. **Fase 2 (HTTP Probe):** Caso a Fase 1 tenha sucesso, emite uma requisição HTTP GET no endpoint de saúde configurado (ex: `/health` ou `/api/health`). Status 200..299 confirmam operação normal (`HEALTHY`). Respostas de erro HTTP ou falhas repetidas colocam o serviço em estado `DEGRADED` ou `UNHEALTHY`.

---

## 2. Máquina de Estados de Disponibilidade

| Estado | Significado Operacional | Ação do Gateway |
| :--- | :--- | :--- |
| `HEALTHY` | Porta conectada, HTTP 200 OK, latência dentro do limite. | Tráfego roteado normalmente sem intervenção. |
| `DEGRADED` | Falha pontual de sonda ou latência elevada (>500ms). | Alerta gerado em telemetria; tráfego mantido. |
| `UNHEALTHY` | 3 ou mais falhas consecutivas ou porta fechada. | Marcação crítica de incidente no painel. |

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Logo Noxfort Systems" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado sob Apache 2.0.</small>
</div>
