# 🚀 Guía de Despliegue y Operaciones

Este documento describe el empaquetado para producción, la containerización con Docker Compose y la gestión del servicio Systemd para **NOXFORT EDGE GATEWAY**.

⬅️ Volver al [Centro de Documentación](README.md) | 🏛️ Ver [Arquitectura](architecture.md)

---

## 1. Operación con Docker Compose

```bash
# Iniciar el servicio en segundo plano
make start

# Inspeccionar el estado de los puertos y contenedor
make status

# Ver registros en tiempo real
make logs

# Detener el contenedor de forma segura
make stop
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado bajo Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
