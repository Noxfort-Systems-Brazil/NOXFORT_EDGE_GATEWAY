# 🚀 Déploiement et Opérations en Production

Ce document décrit le déploiement conteneurisé avec Docker Compose et la gestion du service d'arrière-plan sous Systemd pour **NOXFORT EDGE GATEWAY**.

⬅️ Retour au [Hub de Documentation](README.md) | 🏛️ Voir [Architecture](architecture.md)

---

## 1. Cycle de Vie avec Make

```bash
make start   # Démarre le conteneur en arrière-plan
make status  # Affiche l'état des ports et du conteneur
make logs    # Affiche les logs en continu
make stop    # Arrête le conteneur proprement
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Sous licence Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
