# gokitx 架构

## 目录结构

```
gokitx/
├── slice/          # 泛型切片操作
│   ├── add.go      # Add — 合并两个 slice
│   ├── contain.go  # Contain — 判断元素是否存在
│   ├── exist.go    # Exist — 存在性检查
│   └── map.go      # Map — 泛型 map 操作
├── errs/           # 通用错误辅助
│   └── errs.go     # NewErrIndexOutOfRange
├── Makefile
└── go.mod
```

## 设计哲学

- 零外部依赖 — 只依赖 Go 标准库
- 泛型优先 — 充分利用 Go 1.24+ 泛型能力
- 每个包聚焦一个职责（如 slice 只做切片操作）
- 函数小而专注，单一返回值

## monorepo 依赖关系

```
gokitx (最底层)
  ↑
aix, llm-sdk-go (Go 业务仓库)
```

任何 Go 项目遇到基础算法/数据结构需求时，先查 gokitx 是否已有实现。如果没有，先在此实现并通过测试，再在业务仓库使用。
