# gokitx

Go 泛型基础数据结构工具库 — 零依赖、可跨项目复用的通用组件。

## 项目概要

- **仓库**: [code-koan/gokitx](https://github.com/code-koan/gokitx)
- **语言**: Go 1.24+
- **模块**: `github.com/code-koan/gokitx`

## 当前模块

| 包 | 说明 |
|----|------|
| `slice` | 泛型切片操作（Add、Contain、Exist、Map 等） |
| `errs` | 通用错误辅助（如 `NewErrIndexOutOfRange`） |

## 在 monorepo 中的位置

monorepo 最底层的基础设施，被 `aix`、`llm-sdk-go` 等 Go 仓库依赖。

## 文件

- [architecture.md](architecture.md) — 包结构、设计哲学、依赖关系
