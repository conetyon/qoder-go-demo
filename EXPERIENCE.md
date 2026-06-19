# Qoder Go Demo — 构建经验总结

> 本文记录了在无 Go 预装、无 sudo 权限的 Linux 沙箱环境中，从零构建 Qoder Go Demo 并推送至 GitHub 的完整实践经验。

---

## 一、环境配置

### 1.1 问题：系统未安装 Go

```bash
$ go version
bash: go: command not found
```

### 1.2 问题：无全局目录写权限

尝试安装到 `/usr/local` 时报 `Permission denied`，沙箱环境限制了系统级写入。

### 1.3 解决方案：工作区本地安装

将 Go SDK 安装到工作区目录下，完全规避权限问题：

```bash
# 创建工作区 SDK 目录
mkdir -p ~/qoder_test/go_sdk

# 下载并解压 Go 1.22.4
curl -OL https://go.dev/dl/go1.22.4.linux-amd64.tar.gz
tar -C ~/qoder_test/go_sdk -xzf go1.22.4.linux-amd64.tar.gz
rm go1.22.4.linux-amd64.tar.gz

# 设置环境变量（每次终端需重新执行）
export GOROOT=~/qoder_test/go_sdk/go
export PATH=$GOROOT/bin:$PATH

# 验证
go version  # go version go1.22.4 linux/amd64
```

### 1.4 关键经验

| 问题 | 原因 | 解决方式 |
|---|---|---|
| `/usr/local` 写权限 | 沙箱限制 | 安装到工作区 `go_sdk/` |
| `bwrap: Permission denied` | 沙箱 bubblewrap 隔离 | 使用 `required_permissions='all'` |
| 环境变量不持久 | 非全局安装 | 每次命令前 `export GOROOT=...` |
| `.gitignore` 需排除 SDK | 避免提交 200MB+ SDK | 添加 `go_sdk/` 到 `.gitignore` |

---

## 二、项目设计与编程

### 2.1 项目结构（遵循 Go 标准布局）

```
qoder-go-demo/
├── cmd/main.go              # CLI 入口（flag 参数解析）
├── pkg/qoder/               # 公开包
│   ├── banner.go            # ASCII Banner、版本、Feature 列表
│   ├── greeting.go          # Greeter 接口、泛型函数、并发处理
│   └── qoder_test.go        # 13 个单元测试
├── internal/utils/          # 内部工具包
│   ├── utils.go             # 字符串填充、截断、随机数、表格格式化
│   └── utils_test.go        # 6 个单元测试
├── go.mod                   # 模块定义
└── README.md                # 项目文档
```

### 2.2 展示的 Go 核心特性

#### 泛型（Go 1.18+）

```go
func Map[T any, U any](items []T, fn func(T) U) []U {
    result := make([]U, len(items))
    for i, item := range items {
        result[i] = fn(item)
    }
    return result
}
```

- `Map[int, int]`：数字翻倍
- `Map[string, string]`：转大写
- `Filter[int]`：过滤偶数

#### 接口多态

```go
type Greeter interface {
    Greet(name string) string
}
// FormalGreeter / CasualGreeter 两种实现
// 通过 NewGreeter(style) 工厂函数创建
```

#### Goroutine 并发 Worker Pool

- 使用 `chan` + `sync.WaitGroup` 实现多 Worker 并发处理
- 带索引的结果收集，保证输出顺序与输入一致
- 支持自定义 Worker 数量（CLI `--workers` 参数）

#### 安全随机数

- 使用 `crypto/rand` 而非 `math/rand`，保证密码学安全
- `RandomHex(8)` 生成 16 字符 hex 串作为 Session ID

### 2.3 Qoder 品牌元素融入

| 元素 | 实现方式 |
|---|---|
| ASCII Art Banner | `Banner()` 函数返回终端 Logo |
| Feature Matrix | 6 大 Qoder 能力展示（代码生成、审查、调试等） |
| Session ID | `crypto/rand` 生成，模拟 AI 会话标识 |
| 欢迎语 | Casual/Formal 两种风格，含 Qoder 和火箭 emoji |
| 版本号 | `v1.0.0 — Built with Qoder + Go 1.22` |

### 2.4 测试策略

- **19 个测试，全部通过**
- 采用表驱动测试（Table-Driven Tests）风格
- 覆盖：接口多态、泛型类型推断、空切片边界、并发正确性、随机数范围

```
ok  github.com/yty-dev/qoder-go-demo/internal/utils   0.005s
ok  github.com/yty-dev/qoder-go-demo/pkg/qoder         0.011s
```

---

## 三、Git 管理与 GitHub 推送

### 3.1 Git 基础配置

```bash
git config --global user.name "your-username"
git config --global user.email "your-email@example.com"
```

### 3.2 仓库初始化

```bash
cd qoder-demo
git init
git branch -m master main   # 使用 main 作为默认分支
git add .
git commit -m "feat: Qoder Go Demo — generics, interfaces, concurrency"
```

### 3.3 Fine-grained PAT 推送踩坑

**问题**：Fine-grained Personal Access Token（`github_pat_xxx`）推送时返回 403：

```
remote: Permission to conetyon/qoder-go-demo.git denied to conetyon.
fatal: unable to access '...': The requested URL returned error: 403
```

**原因分析**：

| 踩坑点 | 说明 |
|---|---|
| Token 权限不足 | Fine-grained PAT 需要显式开启 **Contents: Read and write** |
| URL 内嵌认证失败 | `https://user:token@github.com/...` 格式对 Fine-grained PAT 不生效 |
| `x-access-token` 格式 | 同样返回 403 |

**最终解决方案**：使用 `credential.helper` 方式传递认证信息：

```bash
git -c credential.helper='!f() { echo "username=YOUR_USER"; echo "password=YOUR_TOKEN"; }; f' \
    push -u origin main
```

### 3.4 推送最佳实践建议

1. **优先使用 Classic Token**（勾选 `repo` 权限），比 Fine-grained PAT 更简单
2. 如使用 Fine-grained PAT，确保：
   - Repository access 包含目标仓库
   - **Contents** 权限设为 **Read and write**
   - **Metadata** 权限设为 **Read-only**（通常默认开启）
3. 推送后清理 remote URL 中的 token：
   ```bash
   git remote set-url origin https://github.com/user/repo.git
   ```

---

## 四、完整操作流程图

```
检查环境 (go version)
    ↓ 未安装
检测系统架构 (uname -m → x86_64)
    ↓
尝试全局安装 → 权限不足
    ↓
工作区本地安装 Go SDK
    ↓
设置 GOROOT + PATH
    ↓
go mod init 初始化模块
    ↓
分层创建代码 (pkg → internal → cmd)
    ↓
编写单元测试
    ↓
go test ./... 验证 (19 PASS)
    ↓
go run ./cmd/main.go 运行验证
    ↓
创建 README.md + .gitignore
    ↓
git init + commit
    ↓
配置 GitHub 认证 → credential.helper 推送
    ↓
✅ 成功推送到 GitHub
```

---

## 五、总结

| 维度 | 关键收获 |
|---|---|
| **环境** | 沙箱环境可通过工作区本地安装 SDK 规避权限问题 |
| **编程** | Go 泛型 + 接口 + 并发三大特性可在一个小项目中完整展示 |
| **测试** | 表驱动测试 + 边界用例，确保代码可靠性 |
| **Git** | Fine-grained PAT 推送需用 `credential.helper`，Classic Token 更省心 |
| **品牌** | 通过 Banner、Feature 列表、Session ID 等自然融入 Qoder 元素 |
