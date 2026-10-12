# 🧪 Pruebas y Validación de Calidad

Este documento describe la arquitectura de pruebas, el aislamiento de red mediante contratos mock y las metas de cobertura para **NOXFORT EDGE GATEWAY**.

⬅️ Volver al [Centro de Documentación](README.md) | 🏛️ Ver [Arquitectura](architecture.md)

---

## 1. Filosofía de Pruebas

- **Cero Dependencias de Red:** Las pruebas de las capas de Dominio y Aplicación no realizan peticiones reales a Internet ni dependen de que Caddy esté en ejecución.
- **Detector de Carreras:** Todas las rutinas concurrentes en Go se validan con `-race`.

```bash
# Ejecutar todas las pruebas unitarias
make test

# Ejecución con detección de carreras
go test -v -race ./...
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado bajo Apache 2.0.</small>
</div>
