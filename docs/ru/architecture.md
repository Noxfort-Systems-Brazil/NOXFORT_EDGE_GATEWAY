# 🏛️ Архитектура системы и технический проект

В данном документе приведена внутренняя инженерная архитектура **NOXFORT EDGE GATEWAY**. Он охватывает слои Clean Architecture и Domain-Driven Design (DDD), модель входящего трафика по SNI, синхронизацию DNS двух стеков (DuckDNS IPv4/IPv6) и активные проверки состояния.

⬅️ Назад в [Центр документации](README.md) | ⚡ См. [Спецификацию API](api_reference.md)

---

## 1. Топология сети и модель изоляции

На серверах ЦОУ несколько платформ (CARINA, Monitor, Synapse) сосуществуют на одном хосте. **NOXFORT EDGE GATEWAY** занимает общедоступные порты **80** и **443**, перенаправляя шифрованный трафик на приватные локальные порты:

- `carinanoxfort.duckdns.org` ➔ `host.docker.internal:2201`
- `noxfortmonitor.duckdns.org` ➔ `host.docker.internal:22100` (и WSS `9001` для MQTT)

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Лицензия Apache 2.0.</small>
</div>
