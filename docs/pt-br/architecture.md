---
tags: [architecture, gateway, clean-architecture, ddd, caddy, pt-br]
aliases: [Arquitetura do Sistema, Topologia do Gateway, Clean Architecture Blueprint]
---

# 🏛️ Arquitetura do Sistema e Blueprint de Engenharia

Este documento especifica a arquitetura interna do **NOXFORT EDGE GATEWAY**. Ele detalha as camadas de Clean Architecture e Domain-Driven Design (DDD), a topologia de ingress multi-tenant por SNI, o sincronizador periódico de DNS dinâmico (DuckDNS) dual-stack e a infraestrutura de sondas ativas de integridade.

⬅️ Voltar para a [Central de Documentação](noxfort_edge_gateway_moc.md) | ⚡ Ver [APIs e Contratos](api_reference.md) | 🌐 Ver [Sincronização DDNS](ddns_synchronization.md)

---

## 1. Topologia de Rede e Modelo de Concorrência

Em centros de controle operacional (CCO) municipais e servidores de borda, diversos sistemas independentes (IA Semafórica CARINA, Noxfort Monitor, Synapse HFT) precisam operar simultaneamente sem colidir nas portas de rede padrão.

O **NOXFORT EDGE GATEWAY** atua como porteiro exclusivo das portas públicas **80** (HTTP) e **443** (HTTPS) por meio de um container Docker Caddy, roteando conexões criptografadas por SNI diretamente para as portas privadas locais:

```mermaid
graph TD
    ClientInternet([Internet / Clientes CCO / Viaturas]) -->|Porta 80 HTTP / Porta 443 HTTPS| Ingress[NOXFORT EDGE GATEWAY Caddy]
    
    subgraph Camada de Proxy Reverso de Borda
        Ingress -->|SNI: carinanoxfort.duckdns.org| R1[Roteador: CARINA]
        Ingress -->|SNI: noxfortmonitor.duckdns.org| R2[Roteador: MONITOR]
        Ingress -->|SNI: synapse.duckdns.org| R3[Roteador: SYNAPSE]
        Ingress -->|Path: /mqtt| R_WSS[Proxy WebSocket WSS]
    end

    subgraph Serviços Privados no Host
        R1 -->|host.docker.internal:2201| CarinaCore[CARINA Web Gateway & CCO API]
        R2 -->|host.docker.internal:22100| NoxfortMon[Noxfort Monitor Web Dashboard]
        R_WSS -->|host.docker.internal:9001| Mosquitto[Broker MQTT Mosquitto WSS]
        R3 -->|host.docker.internal:SynapsePort| SynapseAI[SYNAPSE Motor IA]
    end

    subgraph Daemon Orquestrador Go
        GoDaemon[Processo noxfort-gateway] -->|Sincroniza a cada 300s| DDNS[API DuckDNS IPv4 e IPv6]
        GoDaemon -->|Sonda Saúde a cada 15s| HealthChecker[Verificador Ativo TCP/HTTP]
        HealthChecker --> CarinaCore
        HealthChecker --> NoxfortMon
    end
```

---

## 2. Camadas Clean Architecture e DDD em Go

O orquestrador nativo em Go é estruturado rigorosamente sob Clean Architecture e DDD:

```mermaid
graph BT
    subgraph Camada de Domínio
        Entities[Entidades e Value Objects: UpstreamService, DDNSTarget, RouteRule]
        Contracts[Contratos: IDDNSSynchronizer, IHealthMonitor, IProxyRouter]
    end

    subgraph Camada de Aplicação
        U1[SyncAllDDNSUseCase] --> Contracts
        U2[MonitorHealthUseCase] --> Contracts
        U3[GetGatewayStatusUseCase] --> Entities
        DTOs[DTO: GatewayStatusDTO]
    end

    subgraph Camada de Infraestrutura
        DuckDNSClient[DuckDNSClient] -->|Implementa| Contracts
        IPDetectors[IPv4Detector, GlobalIPv6Detector] --> DuckDNSClient
        HealthChecker[TCPHTTPHealthChecker] -->|Implementa| Contracts
        CaddyProxy[Caddy Reverse Proxy] -->|Implementa| Contracts
    end

    subgraph Apresentação / Composição
        Main[cmd/gateway/main.go] --> CamadaDeAplicacao
        Config[config/config.go] --> Main
    end
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Logo Noxfort Systems" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado sob Apache 2.0.</small>
</div>
