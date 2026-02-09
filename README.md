# Toppira

Backend service for the Toppira platform.

---

## 🚀 Developer Commands

### 🏃 Start Dev Server

```sh
./scripts/start.sh
```

### 📄 Generate API Documentation

Generate Swagger/OpenAPI docs from annotations.

```sh
swag init -o ./docs -g ./cmd/http/main.go --pd
```

**Notes:**

- Output is written to `./docs`
- Entry point: `./cmd/http/main.go`
- `--pd` enables parsing of dependency packages

---

### 🧱 Generate Repositories

Run the internal code generator to scaffold repositories.

```sh
go run ./cmd/gen
```

---

### 🧹 Run Linters

Execute all configured linters via `golangci-lint`.

```sh
golangci-lint run
```

**Tips:**

- Ensure the correct Go version is installed
- Run before committing to avoid CI failures

---

## 📦 Requirements

- Go (matching `go.mod`)
- `swag` CLI (`go install github.com/swaggo/swag/cmd/swag@latest`)
- `golangci-lint`

---

## 🧠 Conventions

- Commands assume execution from the project root
- Generated files should not be edited manually

---

Happy hacking ✨
