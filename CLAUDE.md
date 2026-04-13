# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Build
go build ./cmd/rl-spectra-assure

# Run tests
go test ./...

# Run a single test
go test ./internal/app/... -run TestFunctionName

# Vet and format
go vet ./...
gofmt -w .
```

## Architecture

This is a **hexagonal architecture** Go CLI application using Cobra. The layers are:

- **`internal/domain/`** — Pure business logic with no external dependencies. `ports.go` defines the `CommunityRepository` interface; `models.go` holds all API response structs.
- **`internal/app/`** — Thin application services (`SearchService`, `VersionReportService`, `PackageDetailsService`) that orchestrate between CLI and repository.
- **`internal/adapters/primary/cli/`** — Cobra command tree. `root.go` wires commands using a lazy service factory closure that initializes the HTTP client only when a command runs.
- **`internal/adapters/secondary/rlclient/`** — HTTP adapter implementing `CommunityRepository`. `client.go` handles auth/transport; `repository.go` maps port methods to API endpoints.

The dependency flow is: CLI → App Service → Repository Port ← HTTP Client (adapter).

## API and Config

- Default endpoint: `https://data.reversinglabs.com/api/oss/community/v2/free`
- Auth: Bearer token via `--token` flag or `RL_API_TOKEN` env var
- Supported package communities: `npm`, `pypi`, `gem`, `nuget`, `vsx`, `psgallery`

## Key Patterns

- **Service factory**: In `root.go`, the factory closure is called lazily on `PersistentPreRunE` so flags are parsed before the HTTP client is constructed.
- **Context propagation**: All methods accept `context.Context` from `main.go`'s signal handler through to the HTTP client.
- **Output**: Commands support `--json` for raw JSON output; default is pretty-printed. The `--compact` flag on search collapses output.
- **Scoped packages**: The `search` command handles `@scope/name` PURL format for npm packages.
