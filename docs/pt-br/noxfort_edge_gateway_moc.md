---
tags: [moc, hub, docs, obsidian, gateway, edge, pt-br]
aliases: [Master MOC PT-BR, Central de Conhecimento, Índice Principal]
---

<div align="center">

# Central Técnica Master de Documentação — NOXFORT EDGE GATEWAY
### Índice Geral e Vault de Conhecimento em Clean Architecture
*Noxfort Systems — A State Of Art Company*

</div>

Bem-vindo à biblioteca técnica do **NOXFORT EDGE GATEWAY**. Projetado como gateway de borda de alta disponibilidade e proxy reverso multi-tenant para plataformas de mobilidade inteligente da Noxfort Systems, ele gerencia a recepção de tráfego nas portas públicas 80 e 443, a emissão automática de certificados SSL/TLS via DNS-01 e o roteamento seguro para microsserviços locais.

Este índice master provê cobertura técnica para desenvolvedores backend, engenheiros DevOps e administradores de rede, totalmente compatível com **GitHub** e **[Obsidian](https://obsidian.md/)**.

---

## 🗺️ Mapa do Código-Fonte e Estrutura de Diretórios

```text
NOXFORT_EDGE_GATEWAY/
├── cmd/
│   └── gateway/
│       └── main.go                 # Ponto de Composição Master e Daemon Orquestrador
├── config/
│   └── config.go                   # Configurações 12-Factor e Upstreams Padrão
├── internal/
│   ├── domain/                     # Regras Puras de Negócio (Camada de Domínio DDD)
│   │   ├── models/                 # Entidades e Objetos de Valor
│   │   │   ├── upstream_service.go # Entidade UpstreamService (Estados, Latência, Falhas)
│   │   │   ├── ddns_target.go      # Objeto de Valor DDNSTarget (Domínio, Token, Provedor)
│   │   │   ├── route_rule.go       # Objeto de Valor RouteRule (Domínio, Serviço de Destino)
│   │   │   └── models_test.go      # Testes Unitários das Entidades de Domínio
│   │   └── contracts/              # Interfaces de Inversão de Controle (DIP)
│   │       ├── ddns_synchronizer.go# Interface IDDNSSynchronizer
│   │       ├── health_monitor.go   # Interface IHealthMonitor
│   │       └── reverse_proxy.go    # Interface IProxyRouter
│   ├── application/                # Regras de Negócio da Aplicação (Clean Architecture)
│   │   ├── usecases/               # Casos de Uso Orquestrados
│   │   │   ├── sync_all_ddns_usecase.go    # Sincronização Concorrente DuckDNS
│   │   │   ├── monitor_health_usecase.go  # Loop de Sondagem Ativa TCP/HTTP
│   │   │   ├── get_gateway_status_usecase.go # Agregação de Telemetria
│   │   │   └── usecases_test.go    # Testes Unitários de Casos de Uso com Mocks
│   │   └── dtos/                   # Objetos de Transferência de Dados
│   │       └── gateway_status_dto.go       # Representação JSON de Status em Tempo Real
│   └── infrastructure/             # Drivers, Frameworks e Infraestrutura de Rede
│       ├── ddns/                   # Implementações de DNS Dinâmico
│       │   ├── duckdns_client.go   # Cliente HTTP/HTTPS para API do DuckDNS
│       │   ├── ipv4_detector.go    # Detecção Externa de IPv4 Público (ipify)
│       │   └── ipv6_detector.go    # Varredor de IPv6 Global em Interfaces Locais
│       ├── health/                 # Drivers de Verificação de Saúde
│       │   └── tcp_http_checker.go # Verificador em Duas Fases (TCP Handshake + HTTP GET)
│       └── caddy/                  # Infraestrutura do Container Proxy Reverso
│           ├── Caddyfile           # Configuração de Rotas SNI Multi-Tenant
│           └── Dockerfile          # Build Customizado do Caddy com Plugin DuckDNS
├── docs/                           # Central de Documentação Técnica e Obsidian Vault
│   ├── pt-br/                      # Versão em Português do Brasil
│   └── en/                         # Versão em Inglês
├── bin/                            # Binários Nativos Compilados (noxfort-gateway)
├── docker-compose.yml              # Orquestração de Containers (Portas 80 e 443)
├── Makefile                        # Automação de Build, Testes e Operação
├── cabecalho.txt                   # Modelo Oficial de Cabeçalho SPDX Apache-2.0
├── LICENSE                         # Licença Oficial Apache 2.0 (SPDX: Apache-2.0)
├── pyproject.toml                  # Metadados do Projeto & Ferramentas MkDocs (SPDX: Apache-2.0)
└── README.md                       # Apresentação do Repositório
```

---

## 📚 Módulos de Documentação Técnica

| Cartão / Subsistema | Área de Foco | Link Direto |
| :--- | :--- | :---: |
| 🏛️ **Arquitetura do Sistema** | Camadas DDD, Topologia de Ingress e Proxy Reverso Caddy | [Ver Arquitetura](architecture.md) |
| ⚡ **Referência de APIs e Contratos** | Caddy Admin API, Protocolo DuckDNS e DTOs de Status | [Ver Referência](api_reference.md) |
| 🌐 **Sincronização DDNS** | Mecanismo DuckDNS, Detecção de IPv4 Público e IPv6 Global | [Ver Guia DDNS](ddns_synchronization.md) |
| 🛡️ **Monitoramento de Saúde** | Sondas Periódicas TCP/HTTP e Máquina de Estados Upstream | [Ver Guia Saúde](health_monitoring.md) |
| ⚙️ **Guia de Configuração** | Variáveis de Ambiente, Caddyfile e Mapeamento de Portas | [Ver Configurações](configuration_guide.md) |
| 🛠️ **Guia do Desenvolvedor** | Toolchain Go, Adição de Upstreams e Padrões de Código | [Ver Guias](developer_guides.md) |
| 🧪 **Testes & Validação** | Testes Unitários Go, Contratos de Domínio e Mocks | [Ver Diretrizes](testing.md) |
| 🚀 **Implantação & Operação** | Docker Compose, Serviço Systemd e Operação em Produção | [Ver Implantação](deployment_and_packaging.md) |

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Logo Noxfort Systems" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado sob Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
