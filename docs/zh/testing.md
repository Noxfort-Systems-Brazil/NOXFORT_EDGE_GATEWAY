# 🧪 测试规范与质量验证

本技术文档描述 **NOXFORT EDGE GATEWAY** 的测试隔离机制、Mock 接口契约与并发竞争检测。

⬅️ 返回 [文档中心](README.md) | 🏛️ 查看 [系统架构](architecture.md)

---

## 1. 测试运行命令

```bash
# 执行完整单元测试
make test

# 开启数据竞态检测运行
go test -v -race ./...
```

---

<div align="center">
  <img src="../assets/noxfort-logo.png" alt="Noxfort Systems Logo" width="40" /><br/>
  <b>Noxfort Systems</b> — <i>A State Of Art Company</i><br/>
  <small>© 2026 Noxfort Systems. 遵循 Apache-2.0 许可证 (SPDX: Apache-2.0)。</small>
</div>
