# gokitx

Go 泛型基础数据结构工具库 — 零依赖、可跨项目复用的通用组件。

## 项目概要

- **仓库**: [code-koan/gokitx](https://github.com/code-koan/gokitx)
- **语言**: Go 1.24+
- **模块**: `github.com/code-koan/gokitx`

## 文档输出规范（CRITICAL）

AI 产出的所有方案、文档、沉淀总结必须简洁：

- **步骤用列表** — 流程类内容用列表，不用大段叙述
- **多维用矩阵** — 多方案/多维度对比用表格
- **一句话高密度** — 一句说清链路：做什么 → 为什么 → 接下来几步

## Go 编写规范

遵循 [Go Proverbs](https://go-proverbs.github.io/) 和 [Effective Go](https://go.dev/doc/effective_go)。

- 扁平控制流，尽早返回
- 函数小而专注，单职责
- 优先声明式而非命令式
- 永不修改接收器状态或参数

## 命令

```bash
make test       # 运行所有测试
make lint       # 静态检查
make build      # 编译验证
```

## 目录结构

```
gokitx/
├── slice/          # 泛型切片操作
│   ├── add.go      #   Add — 合并两个 slice
│   ├── contain.go  #   Contain — 判断元素是否存在
│   ├── exist.go    #   Exist — 存在性检查
│   └── map.go      #   Map — 泛型 map 操作
├── errs/           # 通用错误辅助
│   └── errs.go     #   NewErrIndexOutOfRange
├── Makefile
└── go.mod
```

## 在 monorepo 中的位置

monorepo 最底层的基础设施，被 `aix`、`llm-sdk-go` 等 Go 仓库依赖。

规则：
- 任何 Go 项目遇到基础算法/数据结构需求时，优先查此仓库是否已有实现
- 如果没有所需功能，先在此仓库实现并通过测试，再在业务仓库中使用
- 只放零依赖、可跨项目复用的基础工具，不放业务逻辑

## codegraph

本仓库使用 [codegraph](https://github.com/colbymchenry/codegraph) 构建本地代码知识图谱。

```bash
codegraph init -i     # 首次构建图谱
codegraph status      # 查看索引状态
```

当 `.codegraph/codegraph.db` 存在时，Claude Code 自动通过 MCP Server 查询代码结构。

## issue → PR 规范

- issue → PR：使用 `.github/ISSUE_TEMPLATE/` 模板创建 issue 并打 label，变更通过 PR 提交，PR 描述关联 issue
