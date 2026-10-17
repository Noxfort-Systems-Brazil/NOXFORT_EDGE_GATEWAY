# ⚡ Referencia de APIs, Protocolos y Contratos

Este documento proporciona especificaciones técnicas para las APIs internas, protocolos de red y contratos de dominio del **NOXFORT EDGE GATEWAY**.

⬅️ Volver al [Centro de Documentación](README.md) | 🏛️ Ver [Arquitectura](architecture.md)

---

## 1. API Administrativa de Caddy

Caddy expone una interfaz de administración interna en `http://localhost:2019`.

| Endpoint | Método | Descripción | Carga Útil |
| :--- | :---: | :--- | :--- |
| `http://localhost:2019/load` | `POST` | Recarga atómicamente la configuración sin interrumpir conexiones | `Content-Type: text/caddyfile` |
| `http://localhost:2019/config/` | `GET` | Obtiene la configuración activa en formato JSON | Ninguno |

---

## 2. Protocolo de Sincronización DuckDNS

La actualización de dominios se realiza mediante peticiones HTTPS GET:

```text
GET https://www.duckdns.org/update?domains={DOMINIO}&token={TOKEN}&ip={IPV4}&ipv6={IPV6}&verbose=true
```

Respuestas:
- `OK`: Actualización completada con éxito.
- `KO`: Token no válido o subdominio inexistente.

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licenciado bajo Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
