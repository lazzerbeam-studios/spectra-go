# AGENTS

This repo is a Go API service. The Go module lives in `api-v1`.

## Build, Lint, Test

Go version:
- `api-v1/go.mod` declares `go 1.25.0`.
- CI uses Go 1.23 in `.github/workflows/go-testing.yml`.

Common commands (run from repo root unless noted):
- Build (module): `cd api-v1 && go build api.go`
- Run API (dev): `cd api-v1 && go run api.go --env=dev --server=api`
- Run cron: `cd api-v1 && go run api.go --env=dev --server=cron`
- Download modules: `cd api-v1 && go mod download`
- Tidy modules: `cd api-v1 && go mod tidy`

Tests:
- All tests: `cd api-v1 && go test -v ./...`
- Single package: `cd api-v1 && go test -v ./routes/home_api/tests`
- Single test name: `cd api-v1 && go test -v ./routes/home_api/tests -run Test_HomeGetAPI`
- Single file (package): `cd api-v1 && go test -v ./routes/home_api/tests -run HomeGetAPI`

Lint/format:
- No repo-configured linter found. Use `gofmt` on changed files.

Codegen/migrations:
- Ent generate: `cd api-v1 && go generate ./ent`
- Ent new schema: `cd api-v1 && go run entgo.io/ent/cmd/ent new <Name>`
- Atlas diff/apply (dev): `cd api-v1 && atlas migrate diff --env dev` then `cd api-v1 && atlas migrate apply --env dev`

Docker/infra (root Makefile):
- `make docker-run` / `make docker-stop`
- `make docker-postgres` / `make docker-valkey`
- `make docker-api` (runs compose)

## Repo Structure

- `api-v1/` is the Go module and service entrypoint.
- `api-v1/api.go` wires Huma + Echo and registers routes.
- `api-v1/routes/*` contains API handlers and route registration.
- `api-v1/utils/*` contains helpers (auth, cache, db, etc.).
- `api-v1/ent/` is generated Ent ORM code.

## Code Style Guidelines

General:
- Follow idiomatic Go and keep functions focused and small.
- Prefer explicit error handling; avoid silent failures.
- Use `context.Context` in handlers; pass `ctx` to DB calls.
- Return `(output, error)` and use Huma error helpers for HTTP errors.

Imports:
- Use Go standard grouping:
  1) standard library
  2) third-party
  3) local module (`api-go/...`)
- Separate groups with a blank line.
- Keep imports sorted by `gofmt`.

Formatting:
- Run `gofmt` on all edited Go files.
- Use tabs for indentation as enforced by `gofmt`.
- Keep struct tags aligned by `gofmt` only.

Packages and files:
- Package names are lowercase with underscores for route groups (e.g., `auth_api`, `home_api`).
- Handler files often use `op-<Name>API.go` naming.
- Test packages can use a different name (e.g., `home_tests`), keep that consistent within a folder.

Types and naming:
- Exported types and functions use PascalCase (e.g., `SignInInput`, `SignInAPI`).
- Unexported helpers use camelCase.
- Request/response types typically include a `Body` struct with JSON tags.
- Operation IDs are PascalCase and match handler names.

Error handling:
- Prefer returning Huma error helpers: `huma.Error400BadRequest`, `huma.Error404NotFound`, etc.
- Wrap or log unexpected errors where appropriate; avoid `panic` in request handlers.
- `panic` is currently used in setup code (e.g., DB init); follow existing patterns unless changing behavior.

HTTP and API conventions:
- Routes are registered in `routes.go` files using `huma.Register`.
- Use `huma.Operation` fields consistently: `OperationID`, `Method`, `Path`, `Tags`.
- Keep handler signatures `func(ctx context.Context, input *InputType) (*OutputType, error)`.

DB and ORM:
- Ent client is initialized in `utils/db/ent.go` and stored in `db.EntDB`.
- Use Ent query builders from `api-go/ent/...` in handlers.
- Keep DB errors mapped to proper Huma errors for API responses.

Auth and security:
- JWT logic lives in `utils/auth` and uses `SecretKey` configured at startup.
- Avoid logging tokens or secrets.

Testing:
- Use `humatest.New` for API handler tests where appropriate.
- Prefer table-driven tests for multiple cases.
- Fail with `t.Fatalf` when responses are unexpected.

## Cursor/Copilot Rules

- No `.cursor/rules/`, `.cursorrules`, or `.github/copilot-instructions.md` found in this repo.
