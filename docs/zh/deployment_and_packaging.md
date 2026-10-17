# 🚀 生产部署与运维指南

本技术文档提供 **NOXFORT EDGE GATEWAY** 的 Docker Compose 容器化编排与 Systemd 系统守护进程操作说明。

⬅️ 返回 [文档中心](README.md) | 🏛️ 查看 [系统架构](architecture.md)

---

## 1. 常用运维管理指令

```bash
make start   # 后台启动网关容器
make status  # 查看容器运行状态与端口映射
make logs    # 实时查看反向代理日志
make stop    # 安全停止网关容器
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. 遵循 Apache-2.0 许可证 (SPDX: Apache-2.0)。</small>
</div>
