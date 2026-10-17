---
tags: [deployment, operacao, docker, systemd, producao, pt-br]
aliases: [Implantação em Produção, Guia de Operações, Empacotamento]
---

# 🚀 Guia de Implantação e Operação em Produção

Este documento aborda a execução em ambiente produtivo, empacotamento com Docker Compose, configuração do daemon systemd e procedimentos operacionais para o **NOXFORT EDGE GATEWAY**.

⬅️ Voltar para a [Central de Documentação](noxfort_edge_gateway_moc.md) | 🏛️ Ver [Arquitetura](architecture.md) | ⚙️ Ver [Configurações](configuration_guide.md)

---

## 1. Operação via Docker Compose

O arquivo [`docker-compose.yml`](../../docker-compose.yml) provê reinício automático (`restart: unless-stopped`) e persistência segura dos certificados TLS emitidos:

```bash
# Iniciar o container
make start

# Acompanhar logs
make logs

# Parar com segurança
make stop
```

---

## 2. Serviço Systemd para o Daemon Go

Para manter o orquestrador Go em execução contínua no servidor do CCO municipal:

```bash
make build
sudo cp bin/noxfort-gateway /usr/local/bin/noxfort-gateway
```

Crie a unidade `/etc/systemd/system/noxfort-gateway.service`:
```ini
[Unit]
Description=Noxfort Edge Gateway Orchestrator
After=network-online.target docker.service

[Service]
Type=simple
User=root
WorkingDirectory=/home/gabriel-moraes/Documentos/NOXFORT_EDGE_GATEWAY
ExecStart=/usr/local/bin/noxfort-gateway
Restart=always
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

Ative e inicie com:
```bash
sudo systemctl daemon-reload
sudo systemctl enable --now noxfort-gateway
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Logo Noxfort Systems" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado sob Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
