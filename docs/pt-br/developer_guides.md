---
tags: [developer, guides, go, desenvolvimento, contribuicao, pt-br]
aliases: [Guia do Desenvolvedor, Manual de Contribuição, Setup Go]
---

# 🛠️ Guia do Desenvolvedor e Boas Práticas

Este manual orienta desenvolvedores na preparação do ambiente de desenvolvimento, adição de novos upstreams, implementação de testes e adesão aos padrões arquiteturais do **NOXFORT EDGE GATEWAY**.

⬅️ Voltar para a [Central de Documentação](noxfort_edge_gateway_moc.md) | 🏛️ Ver [Arquitetura](architecture.md) | 🧪 Ver [Testes](testing.md)

---

## 1. Ambiente Local e Comandos `make`

```bash
# Compilar o binário em Go
make build

# Executar todos os testes unitários
make test

# Iniciar o container proxy em segundo plano
make start

# Consultar status do container
make status

# Parar o container
make stop

# Ver logs em tempo real
make logs
```

---

## 2. Padrões de Código e Cabeçalho de Licença SPDX

Todos os arquivos de código-fonte no repositório **devem** seguir o padrão de licença SPDX Apache-2.0 contido em `cabecalho.txt` (`SPDX-FileCopyrightText: 2026 Noxfort Systems` e `SPDX-License-Identifier: Apache-2.0`). Ao criar novos arquivos Go, certifique-se de incluir as tags SPDX, metadados do arquivo, autor e data.

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Logo Noxfort Systems" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado sob Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
