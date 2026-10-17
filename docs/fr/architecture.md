# 🏛️ Architecture Système et Blueprint d'Ingénierie

Ce document spécifie l'architecture interne de **NOXFORT EDGE GATEWAY**. Il détaille les couches Clean Architecture et Domain-Driven Design (DDD), le modèle d'entrée multi-tenant par SNI, la synchronisation automatisée DNS double-pile (DuckDNS IPv4/IPv6) et les sondes actives d'intégrité.

⬅️ Retour au [Hub de Documentation](README.md) | ⚡ Voir [API & Contrats](api_reference.md)

---

## 1. Topologie Réseau et Concurrence

Sur les serveurs des centres de contrôle municipaux (CCO), plusieurs plateformes indépendantes (CARINA Traffic AI, Noxfort Monitor, Synapse HFT) doivent coexister en toute sécurité sans entrer en conflit sur les ports Web conventionnels.

**NOXFORT EDGE GATEWAY** assure cette fonction en monopolisant les ports publics **80** (HTTP) et **443** (HTTPS) via un conteneur Caddy personnalisé, routant le trafic chiffré vers les ports locaux privés :

- `carinanoxfort.duckdns.org` ➔ `host.docker.internal:2201`
- `noxfortmonitor.duckdns.org` ➔ `host.docker.internal:22100` (et WSS `9001` pour MQTT)

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Sous licence Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
