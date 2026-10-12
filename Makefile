# NOXFORT EDGE GATEWAY - Unified Multi-Tenant Edge Gateway & Reverse Proxy
# Copyright (C) 2026 Noxfort Systems
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# File: Makefile
# Author: Gabriel Moraes
# Date: 2026-10-12

.PHONY: all build start stop status logs test

all: build test

build:
	@echo "==> Compilando o orquestrador Go do Gateway..."
	go build -o bin/noxfort-gateway cmd/gateway/main.go
	@echo "==> Compilando o container Caddy com plugin DuckDNS..."
	docker compose build

start:
	@echo "==> Iniciando o Gateway de Borda Unificado..."
	docker compose up -d
	@echo "==> Gateway de Borda ativo nas portas 80 e 443!"

stop:
	@echo "==> Parando o Gateway de Borda..."
	docker compose down

status:
	@echo "==> Status dos containers:"
	docker compose ps
	@echo "==> Telemetria local:"
	@curl -s http://127.0.0.1:2026/status || echo "Orquestrador local não ativo."

logs:
	docker compose logs -f

test:
	@echo "==> Executando testes unitários Go (DDD / Clean Architecture)..."
	go test -v ./...
