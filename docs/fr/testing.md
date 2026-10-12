# 🧪 Tests et Assurance Qualité

Ce document présente la stratégie de tests unitaires, l'utilisation de mocks pour isoler le réseau et les commandes d'exécution pour **NOXFORT EDGE GATEWAY**.

⬅️ Retour au [Hub de Documentation](README.md) | 🏛️ Voir [Architecture](architecture.md)

---

## 1. Exécution des Tests

```bash
# Exécution standard
make test

# Exécution avec détection des accès concurrents (race conditions)
go test -v -race ./...
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Sous licence Apache 2.0.</small>
</div>
