# 🏛️ 系统架构核心设计与工程蓝图

本技术文档详述 **NOXFORT EDGE GATEWAY** 的内部工程架构，涵盖基于 Go 的整洁架构 (DDD) 分层、公网 80/443 端口 SNI 路由拓扑、DuckDNS IPv4/IPv6 双栈动态 DNS 自动同步机制以及主动式健康探测管道。

⬅️ 返回 [文档中心](README.md) | ⚡ 查看 [API 规范](api_reference.md)

---

## 1. 网络拓扑与端口多路复用

在市政交通管控中心（CCO）主机上，多个独立运行的微服务（CARINA 交通大脑、Noxfort Monitor、Synapse HFT）需在同一操作系统内安全共存。**NOXFORT EDGE GATEWAY** 通过独占外部 TCP **80** 和 **443** 端口，借助 Caddy 反向代理根据 SNI 域名直接将加密流量安全分流至各本地私有端口：

- `carinanoxfort.duckdns.org` ➔ `host.docker.internal:2201`
- `noxfortmonitor.duckdns.org` ➔ `host.docker.internal:22100`（以及 WebSocket `9001` 用于 MQTT）

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. 遵循 Apache-2.0 许可证 (SPDX: Apache-2.0)。</small>
</div>
