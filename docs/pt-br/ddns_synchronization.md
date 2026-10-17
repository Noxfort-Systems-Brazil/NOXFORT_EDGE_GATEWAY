---
tags: [ddns, duckdns, ipv4, ipv6, rede, dual-stack, pt-br]
aliases: [Sincronização DDNS, Mecanismo DuckDNS, Guia DNS Dinâmico]
---

# 🌐 Subsistema de Sincronização de DNS Dinâmico (DuckDNS)

Este documento detalha o mecanismo de sincronização dual-stack de DNS Dinâmico (DuckDNS) implementado no **NOXFORT EDGE GATEWAY**.

⬅️ Voltar para a [Central de Documentação](noxfort_edge_gateway_moc.md) | 🏛️ Ver [Arquitetura](architecture.md) | ⚡ Ver [APIs e Contratos](api_reference.md)

---

## 1. Motivação e Desafios de Borda

Em servidores de tráfego e CCOs urbanos, as conexões de internet frequentemente utilizam:
1. **IPv4 público dinâmico** ou CGNAT (Carrier-Grade NAT).
2. **IPv6 público globalmente roteável** atribuído à máquina local.

O **NOXFORT EDGE GATEWAY** garante acessibilidade contínua ao atualizar dinamicamente tanto o IPv4 quanto o IPv6 nos subdomínios DuckDNS cadastrados a cada 300 segundos.

---

## 2. Fluxo de Descoberta Dual-Stack

1. **Descoberta de IPv4 Público (`IPv4Detector`):** Consulta endpoints remotos (`https://api.ipify.org`) com timeout estrito de 5 segundos.
2. **Descoberta de IPv6 Global (`GlobalIPv6Detector`):** Inspeciona localmente as interfaces de rede do sistema operacional (`net.Interfaces()`), filtrando e descartando endereços de loopback (`::1`), link-local (`fe80::/10`) e locais únicos (ULA). O primeiro endereço unicast global é selecionado.
3. **Atualização Concorrente (`SyncAllDDNSUseCase`):** Dispara goroutines concorrentes para cada domínio configurado, enviando requisições com os endereços descobertos para a API do DuckDNS.

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Logo Noxfort Systems" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado sob Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
