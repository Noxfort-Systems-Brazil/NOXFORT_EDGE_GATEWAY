<div align="center">

# NOXFORT EDGE GATEWAY — 官方技术文档套件
### 多租户统一边缘网关与高可用反向代理
*Noxfort Systems — A State Of Art Company*

[![Status](https://img.shields.io/badge/Status-Active-brightgreen?style=flat&logo=github)](https://github.com/Noxfort-Systems-Brazil)
[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Caddy](https://img.shields.io/badge/Caddy-2.8%2B-1F88C0?style=flat&logo=caddy&logoColor=white)](https://caddyserver.com/)
[![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![许可证: Apache-2.0](https://img.shields.io/badge/License-Apache--2.0-blue.svg)](../../LICENSE)

---

🌐 **语言 / Translations:** **[🇺🇸 English](../../README.md)** • **[🇧🇷 Português](../pt-br/README.md)** • **[🇪🇸 Español](../es/README.md)** • **[🇫🇷 Français](../fr/README.md)** • **[🇷🇺 Русский](../ru/README.md)** • **[🇨🇳 简体中文](README.md)** • **[📚 文档中心](../index.md)**

---

</div>

## 欢迎访问 NOXFORT EDGE GATEWAY 官方技术文档中心

**NOXFORT EDGE GATEWAY** 是由 Noxfort Systems 开发的高可用中央网络入口与统一反向代理平台，专为智能城市交通控制中心（CCO）与边缘服务器节点设计。作为主机上**公网 80（HTTP）与 443（HTTPS）端口的独占守护进程**，它提供全自动 TLS 证书签发、基于 SNI 的多租户分流路由、双栈动态 DNS（DuckDNS IPv4/IPv6）定期同步以及对本地后端微服务的实时健康探测：

- 🚦 **CARINA（CCO 大屏显示、交通信号灯自适应控制与急救车辆优先）：** `carinanoxfort.duckdns.org` ➔ `127.0.0.1:2201`
- 📊 **NOXFORT MONITOR（工业遥测与事件监控）：** `noxfortmonitor.duckdns.org` ➔ `127.0.0.1:22100` 以及 WSS `9001` (MQTT)
- 🧠 **SYNAPSE（高频 AI 推理与集群编排）：** `synapse.duckdns.org` ➔ Synapse 专用上游端口

---

## 专业技术指南目录

| 文档 | 领域与范围 | 核心内容 |
| :--- | :--- | :--- |
| 📖 **[系统架构核心设计](architecture.md)** | 系统蓝图 | 基于 Go 的整洁架构 (DDD)、SNI 端口多路复用 (80/443)、Caddy 容器化与网络隔离。 |
| ⚡ **[API 规范与领域契约](api_reference.md)** | 接口与数据契约 | Caddy 管理端 API、DuckDNS REST 更新协议以及 `GatewayStatusDTO` 遥测状态数据模型。 |
| 🧪 **[测试规范与质量验证](testing.md)** | 质量保证与测试 | 纯内存无网络依赖的 Go 单元测试 (`go test`)、Mock 接口契约与 `-race` 竞态检测。 |
| 🚀 **[生产部署与运维指南](deployment_and_packaging.md)** | 运维与系统打包 | Docker Compose 编排 (`restart: unless-stopped`)、证书持久化卷与 Systemd 守护进程配置。 |

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="45" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <i>智能交通工程与移动性科技 • NOXFORT EDGE GATEWAY v1.0.0</i><br/>
  <small>© 2026 Noxfort Systems. 遵循 Apache-2.0 许可证 (SPDX: Apache-2.0)。</small>
</div>
