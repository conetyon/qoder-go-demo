# Qoder Go Demo

> ⚡ Built with [Qoder](https://qoder.ai) — AI-powered coding assistant

A demonstration project showcasing modern Go features with Qoder theming, including **generics**, **interfaces**, **concurrency**, and **CLI interaction**.

![Go Version](https://img.shields.io/badge/Go-1.22-00ADD8?style=flat-square&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)

---

## 🎯 Project Overview

This project demonstrates key Go language features in a practical CLI application:

| Feature | Description |
|---|---|
| **Generics** | `Map[T, U]` and `Filter[T]` generic functions |
| **Interfaces** | `Greeter` interface with multiple implementations |
| **Concurrency** | Worker pool pattern with goroutines and channels |
| **CLI** | Flag-based command-line interface |
| **Testing** | Table-driven tests for all packages |

---

## 📁 Project Structure

```
qoder-go-demo/
├── cmd/
│   └── main.go              # CLI entry point
├── pkg/
│   └── qoder/
│       ├── banner.go        # ASCII banner, version, features
│       ├── greeting.go      # Greeter interface, generics, concurrency
│       └── qoder_test.go    # Package tests
├── internal/
│   └── utils/
│       ├── utils.go         # String utilities, table formatter
│       └── utils_test.go    # Utils tests
├── go.mod
└── README.md
```

---

## 🚀 Quick Start

### Prerequisites

- Go 1.22 or later
- Git

### Run the Demo

```bash
# Clone the repository
git clone https://github.com/yty-dev/qoder-go-demo.git
cd qoder-go-demo

# Run directly
go run ./cmd/main.go

# With custom options
go run ./cmd/main.go --name "Your Name" --style formal --table

# Build binary
go build -o qoder-demo ./cmd/main.go
./qoder-demo --workers 8 --table
```

### CLI Flags

| Flag | Default | Description |
|---|---|---|
| `--name` | `Developer` | Your name for the greeting |
| `--style` | `casual` | Greeting style: `casual` or `formal` |
| `--workers` | `4` | Number of concurrent workers |
| `--table` | `false` | Show feature comparison table |

---

## 🧪 Testing

```bash
# Run all tests
go test ./... -v

# With coverage
go test ./... -v -cover
```

Expected output:

```
ok  github.com/yty-dev/qoder-go-demo/internal/utils   0.005s
ok  github.com/yty-dev/qoder-go-demo/pkg/qoder         0.011s
```

---

## 💡 Key Go Features Demonstrated

### 1. Generics (Go 1.18+)

```go
// Map applies a function to each element of a slice
func Map[T any, U any](items []T, fn func(T) U) []U {
    result := make([]U, len(items))
    for i, item := range items {
        result[i] = fn(item)
    }
    return result
}
```

### 2. Interfaces

```go
type Greeter interface {
    Greet(name string) string
}

// Multiple implementations: FormalGreeter, CasualGreeter
greeter := qoder.NewGreeter("formal")
fmt.Println(greeter.Greet("Developer"))
```

### 3. Concurrency with Goroutines

```go
// Worker pool pattern — process tasks in parallel
results := ConcurrentProcess(tasks, numWorkers, func(task string) string {
    return fmt.Sprintf("done: %s", task)
})
```

---

## 🤖 Qoder Elements

This project is themed around **Qoder**, the AI coding assistant:

- **ASCII Banner**: Qoder branding rendered in the terminal
- **Feature Matrix**: Showcases Qoder capabilities (code gen, review, debugging, etc.)
- **Session ID**: Generates a random hex session ID, simulating an AI session
- **Greeter**: Qoder-themed welcome messages in casual and formal styles

---

## 📄 License

MIT License — feel free to use and modify.
