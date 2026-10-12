<div align="center">

# NOXFORT EDGE GATEWAY — Комплект технической документации
### Мультитенантный пограничный шлюз и высокодоступный обратный прокси
*Noxfort Systems — A State Of Art Company*

[![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=flat&logo=github)](https://github.com/Noxfort-Systems-Brazil)
[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Caddy](https://img.shields.io/badge/Caddy-2.8%2B-1F88C0?style=flat&logo=caddy&logoColor=white)](https://caddyserver.com/)
[![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![Лицензия](https://img.shields.io/badge/License-Apache_2.0-blue?style=flat)](../../LICENSE)

---

🌐 **Языки / Переводы:** **[🇺🇸 English](../../README.md)** • **[🇧🇷 Português](../pt-br/README.md)** • **[🇪🇸 Español](../es/README.md)** • **[🇫🇷 Français](../fr/README.md)** • **[🇷🇺 Русский](README.md)** • **[🇨🇳 简体中文](../zh/README.md)** • **[📚 Главный Хаб](../index.md)**

---

</div>

## Добро пожаловать в официальную документацию NOXFORT EDGE GATEWAY

**NOXFORT EDGE GATEWAY** — это высокодоступный пограничный сетевой шлюз и мультитенантный обратный прокси-сервер, разработанный компанией Noxfort Systems для платформ умного города и интеллектуальных транспортных систем. Выступая в качестве **эксклюзивного шлюза для общедоступных портов 80 (HTTP) и 443 (HTTPS)** на серверах центров оперативного управления (ЦОУ), он обеспечивает автоматическое завершение TLS, маршрутизацию по SNI, периодическую синхронизацию динамического DNS с поддержкой двух стеков (DuckDNS IPv4/IPv6) и активную проверку состояния локальных микросервисов:

- 🚦 **CARINA (Видеостена ЦОУ, Светофоры и Адаптивный ИИ):** `carinanoxfort.duckdns.org` ➔ `127.0.0.1:2201`
- 📊 **NOXFORT MONITOR (Промышленная телеметрия и инциденты):** `noxfortmonitor.duckdns.org` ➔ `127.0.0.1:22100` и WSS `9001` (MQTT)
- 🧠 **SYNAPSE (Высокочастотный ИИ и оркестрация):** `synapse.duckdns.org` ➔ порт микросервиса Synapse

---

## Разделы технической документации

| Документ | Раздел | Основные темы |
| :--- | :--- | :--- |
| 📖 **[Архитектура системы](architecture.md)** | Базовая архитектура | Clean Architecture на Go (DDD), мультиплексирование по SNI, порты 80/443, Caddy и изоляция сервисов. |
| ⚡ **[Спецификация API и контрактов](api_reference.md)** | Протоколы и схемы | Административный API Caddy, протокол REST DuckDNS и схемы DTO `GatewayStatusDTO`. |
| 🧪 **[Тестирование и валидация](testing.md)** | Контроль качества | Чистые модульные тесты на Go (`go test`), моки сетевых интерфейсов и детектор гонок данных `-race`. |
| 🚀 **[Развертывание и эксплуатация](deployment_and_packaging.md)** | Промышленный запуск | Docker Compose (`restart: unless-stopped`), постоянные тома сертификатов и служба Systemd. |

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="45" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <i>Интеллектуальные транспортные системы • NOXFORT EDGE GATEWAY v1.0.0</i><br/>
  <small>© 2026 Noxfort Systems. Лицензия Apache 2.0.</small>
</div>
