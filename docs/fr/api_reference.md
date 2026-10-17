# ⚡ Référence des API et Contrats de Domaine

Ce document détaille l'API d'administration interne de Caddy, le protocole DuckDNS REST et le format du DTO `GatewayStatusDTO` pour **NOXFORT EDGE GATEWAY**.

⬅️ Retour au [Hub de Documentation](README.md) | 🏛️ Voir [Architecture](architecture.md)

---

## 1. API Caddy Admin (`localhost:2019`)

- `POST /load` : Recharge à chaud la configuration sans couper les flux TLS existants.
- `GET /config/` : Renvoie la configuration courante au format JSON.

## 2. Protocole de Synchronisation DuckDNS

```text
GET https://www.duckdns.org/update?domains={DOMAINE}&token={TOKEN}&ip={IPV4}&ipv6={IPV6}&verbose=true
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Sous licence Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
