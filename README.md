# Hermes-test

> Hermes Agent + Claude Code 协同工作演示项目
>
> **语言: Go** 🚀

## 文件说明

| 文件 | 说明 |
|------|------|
| `hello.go` | 入门演示 |
| `calculator.go` | 计算器实现（加减乘除） |
| `server.go` | TCP 服务端（端口 28080，多协程处理） |

## 运行

```bash
# 单文件运行
go run hello.go
go run calculator.go

# TCP 服务端
go run server.go
```

## TCP 服务端

监听 `:28080`，每接入一个客户端自动分配一个 goroutine 处理。

**内置命令：**
- `ping` → `pong`
- `time` → 当前服务器时间
- `exit` / `quit` → 断开连接
- 其他消息 → 回声返回

支持 `Ctrl+C` 优雅关闭。
