---
tags: [testing, quality-assurance, unit-tests, coverage, race-detector]
aliases: [Testing Guide, QA Standards, Test Suite Reference]
---

# 🧪 Testing & Validation Standards

This document outlines the testing architecture, mock strategies, test execution commands, and coverage targets for **NOXFORT EDGE GATEWAY**.

⬅️ Back to [Master Documentation Hub](NOXFORT_EDGE_GATEWAY_MOC.md) | 🏛️ See [Architecture](ARCHITECTURE.md) | 🛠️ See [Developer Guides](DEVELOPER_GUIDES.md)

---

## 1. Testing Philosophy

Because the Edge Gateway acts as the perimeter ingress for municipal traffic systems, reliability is paramount:
1. **Zero External Flakiness:** Unit tests for Domain and Application layers must execute with **zero network dependencies**. Network interactions (DuckDNS API, IP detection, HTTP/TCP sockets) are strictly mocked using Go interface test doubles.
2. **Deterministic Assertions:** All domain entity state transitions, latency bounds, and failure counters are validated deterministically.
3. **Data Race Free:** All concurrent routines (goroutines in `SyncAllDDNSUseCase` and `MonitorHealthUseCase`) are verified with Go's `-race` detector.

---

## 2. Test Execution Commands

```bash
# Run standard test suite with verbose reporting
make test

# Run tests with active race detector
go test -v -race ./...

# Generate coverage profile
go test -coverprofile=coverage.out ./...

# View coverage statistics in terminal
go tool cover -func=coverage.out

# View interactive HTML coverage report
go tool cover -html=coverage.out -o coverage.html
```

---

## 3. Test Suites Breakdown

### 3.1 Domain Layer Tests (`internal/domain/models/models_test.go`)
- **`TestNewUpstreamService`**: Verifies valid entity construction, default initialization of latency, and healthy initial status.
- **`TestNewUpstreamServiceValidation`**: Validates boundary rules (invalid ports `<= 0` or `> 65535`, empty host, empty ID).
- **`TestNewRouteRule`**: Tests immutable route rule creation and parameter validation.
- **`TestNewDDNSTarget`**: Validates FQDN resolution, provider assignment, and subdomain formatting.

### 3.2 Application Layer Tests (`internal/application/usecases/usecases_test.go`)
- **`TestSyncAllDDNSUseCase`**: Injects `mockDDNSSync` and `mockIPDetectors` to test parallel synchronization, IP propagation, and target state mutation without internet access.
- **`TestMonitorHealthUseCase`**: Injects `mockHealthChecker` to simulate healthy and degraded upstream responses, verifying latency tracking.
- **`TestGetGatewayStatusUseCase`**: Tests aggregation of live upstreams and DDNS targets into `GatewayStatusDTO`.

---

## 4. Mocking Pattern Example

```go
type mockDDNSSync struct {
    syncFunc func(ctx context.Context, target *models.DDNSTarget, ipv4, ipv6 string) (*contracts.SyncResult, error)
}

func (m *mockDDNSSync) SyncTarget(ctx context.Context, target *models.DDNSTarget, ipv4, ipv6 string) (*contracts.SyncResult, error) {
    return m.syncFunc(ctx, target, ipv4, ipv6)
}
```

This pattern isolates domain logic from infrastructure variability, enabling 100% reproducible test runs in CI/CD pipelines.

---

<div align="center">
  <img src="assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. Licensed under Apache-2.0 (SPDX: Apache-2.0).</small>
</div>
