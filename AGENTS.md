# whaleshell-cli agent guide

## Scope
Standalone `github.com/whaleshell/whaleshell-cli` module: user-facing CLI, command wiring, policies, and agent container definitions.

## Read first
- `README.md`, `providers/README.md`, `docker/README.md`, and `policies/README.md` as relevant.
- Workspace plan: [Cleanup and compatibility](../whaleshell-docs/WORKSPACE_MIGRATION_PLAN.md); published docs source: `../whaleshell-docs/docs/en/`.
- `go.mod` for dependencies; `../Taskfile.yml` for workspace tasks.

## Change guidance
- Treat CLI flags, output, exit codes, config files, and policy files as compatibility surfaces.
- Keep policy validation fail-closed where required; never print credential values.
- Keep provider-specific behavior in provider packages and shared behavior in the owning modules.
- Update English and Russian published docs in `../whaleshell-docs/` when user-visible behavior changes.

## Verification
Run `go test ./...`, `go vet ./...`, and `gofmt` on changed Go files here. From the workspace root, use `task workspace:check`; use `task check` for a full Go workspace gate.
