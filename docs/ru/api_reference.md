# ⚡ Спецификация API и системных контрактов

Этот документ описывает административный API Caddy, протокол DuckDNS REST и формат JSON-телеметрии `GatewayStatusDTO` для **NOXFORT EDGE GATEWAY**.

⬅️ Назад в [Центр документации](README.md) | 🏛️ См. [Архитектуру](architecture.md)

---

## 1. Административный интерфейс Caddy (`localhost:2019`)

- `POST /load` : Атомарное обновление конфигурации Caddy بدون разрыва существующих соединений.
- `GET /config/` : Получение текущей конфигурации в формате JSON.

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Лицензия Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
