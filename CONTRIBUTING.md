# Contributing to NOXFORT EDGE GATEWAY

Thank you for contributing to **NOXFORT EDGE GATEWAY**! This project provides mission-critical edge ingress, reverse proxying, and multi-tenant security for urban mobility and smart city platforms.

---

## 🛠️ Development Setup

1. **Prerequisites:**
   - Go 1.22 or higher
   - Docker and Docker Compose
   - `make` utility
2. **Clone and Branch:**
   ```bash
   git clone https://github.com/Noxfort-Systems-Brazil/NOXFORT_EDGE_GATEWAY.git
   cd NOXFORT_EDGE_GATEWAY
   git checkout -b feature/your-feature-name
   ```
3. **Build the Gateway:**
   ```bash
   make build
   ```

---

## 📜 Coding Standards & Architectural Guidelines

1. **Domain-Driven Design (DDD) & Clean Architecture:**
   - **Domain Layer (`internal/domain`):** Pure entities and interfaces. No external framework imports (`caddy`, `http`, etc. unless pure contract interfaces).
   - **Application Layer (`internal/application`):** Use cases orchestrating domain entities via interfaces.
   - **Infrastructure Layer (`internal/infrastructure`):** Network drivers, HTTP clients, Caddy wrappers, DDNS synchronizers.
2. **License Headers:**
   Every new or modified source code file **MUST** include the standard **SPDX Apache 2.0 License Header** (`SPDX-FileCopyrightText: 2026 Noxfort Systems` and `SPDX-License-Identifier: Apache-2.0`) with proper author and date tags, as specified in `cabecalho.txt`.
3. **No Port Conflicts:**
   Never bind arbitrary public host ports in production containers. Only ports **80** and **443** are exposed publicly.

---

## 🧪 Testing Guidelines

Before opening a Pull Request, verify that all unit and integration tests succeed:

```bash
make test
# or directly:
go test -v -race ./...
```

See [docs/TESTING.md](docs/TESTING.md) for detailed guidelines.

---

## 🚀 Pull Request Checklist

- [ ] Code follows DDD and SOLID guidelines outlined in `ARCHITECTURE.md`.
- [ ] All unit tests pass with zero errors and race conditions.
- [ ] SPDX Apache-2.0 license header (`SPDX-License-Identifier: Apache-2.0`) is present in all newly created files.
- [ ] Relevant documentation in `docs/` is updated.
- [ ] Commit messages follow the conventional commit format (`feat:`, `fix:`, `docs:`, `refactor:`).

---

<div align="center">
  <img src="docs/assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licensed under Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
