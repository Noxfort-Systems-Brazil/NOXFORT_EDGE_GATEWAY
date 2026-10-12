# 🧪 Тестирование и контроль качества

В данном документе описываются правила тестирования, архитектура моков и проверка гонок данных для **NOXFORT EDGE GATEWAY**.

⬅️ Назад в [Центр документации](README.md) | 🏛️ См. [Архитектуру](architecture.md)

---

## 1. Запуск тестов

```bash
# Стандартный запуск
make test

# Проверка на состояние гонки данных
go test -v -race ./...
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Лицензия Apache 2.0.</small>
</div>
