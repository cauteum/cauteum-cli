# Changelog

## [Unreleased]

### Added

- Import, validate, update, and resolve OpenShell-compatible provider profiles from local files or the gateway catalog.
- Carry profile-declared OAuth2 refresh configuration and output-token mappings into gateway-managed providers.

## [v0.0.2-alpha.1] - 2026-09-28

### Security

- `rule approve-all` skips security-flagged proposals unless `--include-security-flagged` is set.

### Changed

- Prefer `WHALESHELL_*` env over legacy `OPENSHELL_*` when both are set.
- Doctor warns on KEK migration need and binary-scoped identity limits on Docker Desktop.

### Fixed

- GoReleaser archive helpers use `strip_parent` so paths land at `libexec/whaleshell/linux-<arch>/`.
