# ⚡ API 接口规范与领域契约

本文件提供 **NOXFORT EDGE GATEWAY** 的 Caddy 内部管理端点、DuckDNS REST 协议以及 `GatewayStatusDTO` 遥测状态数据模型的规范说明。

⬅️ 返回 [文档中心](README.md) | 🏛️ 查看 [系统架构](architecture.md)

---

## 1. Caddy 内部管理接口 (`localhost:2019`)

- `POST /load` : 原子化热加载 Caddy 路由与证书配置，零中断现有 TLS 连接。
- `GET /config/` : 以 JSON 格式检索当前运行态完整配置。

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. 遵循 Apache-2.0 许可证 (SPDX: Apache-2.0)。</small>
</div>
