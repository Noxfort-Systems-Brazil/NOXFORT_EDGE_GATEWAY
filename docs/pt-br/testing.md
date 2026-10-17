---
tags: [testing, testes, validacao, go, race-detector, pt-br]
aliases: [Testes & Validação, Guia de Testes, QA]
---

# 🧪 Diretrizes de Testes e Validação

Este documento descreve a estratégia de testes unitários, mocks de interfaces e cobertura de código para o **NOXFORT EDGE GATEWAY**.

⬅️ Voltar para a [Central de Documentação](noxfort_edge_gateway_moc.md) | 🏛️ Ver [Arquitetura](architecture.md) | 🛠️ Ver [Guia do Desenvolvedor](developer_guides.md)

---

## 1. Filosofia de Testes

Os testes da camada de Domínio e Aplicação são estritamente isolados da rede externa:
- **Zero Dependências Externas:** Todas as chamadas de rede (API do DuckDNS, resoluções de IP, sockets TCP) utilizam test doubles/mocks.
- **Detecção de Condições de Corrida:** O flag `-race` do Go é utilizado para verificar a segurança da concorrência entre goroutines.

```bash
# Execução simples
make test

# Execução com race detector ativo
go test -v -race ./...

# Medição de cobertura de código
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Logo Noxfort Systems" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado sob Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
