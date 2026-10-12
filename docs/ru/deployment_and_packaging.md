# 🚀 Развертывание и промышленная эксплуатация

Этот документ описывает работу шлюза в производственной среде с использованием Docker Compose и системного сервиса Systemd для **NOXFORT EDGE GATEWAY**.

⬅️ Назад в [Центр документации](README.md) | 🏛️ См. [Архитектуру](architecture.md)

---

## 1. Команды управления (Make)

```bash
make start   # Запуск контейнера в фоновом режиме
make status  # Проверка состояния контейнера и портов
make logs    # Просмотр журналов в реальном времени
make stop    # Корректная остановка контейнера
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Лицензия Apache 2.0.</small>
</div>
