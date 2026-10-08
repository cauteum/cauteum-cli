# Changelog

## [v0.1.0-beta.2] - 2026-10-08

### Changed

- Update the driver dependency to v0.1.0-beta.1, including the platform-specific Docker isolation and executable checks.

## [v0.1.0-beta.1] - 2026-10-07

### Added

- Add shell completion scripts and structured command error output.

### Changed

- Use the beta core, proxy, and runtime modules and update Bubble Tea to v1.3.10.

### Fixed

- Protect CLI private files with Windows ACLs and cover Docker cleanup on every CI platform.

## [v0.1.0-alpha.3] - 2026-10-07

### Changed

- Align slogx, display, driver, proxy, and runtime dependencies with their published v0.1.0-alpha.2 modules; update `golang.org/x/crypto` to v0.57.0.
- Refresh the module graph so GoReleaser and cross-platform builds use the same compatible Docker and Moby dependency set as the released driver.

## [v0.1.0-alpha.2] - 2026-09-28

### Added

- Import, validate, update, and resolve OpenShell-compatible provider profiles from local files or the gateway catalog.
- Carry profile-declared OAuth2 refresh configuration and output-token mappings into gateway-managed providers.

### Changed

- Use current GoReleaser archive IDs so the pinned release tooling accepts the CLI and gateway archive configuration.

## [v0.0.2-alpha.1] - 2026-09-28

### Security

- `rule approve-all` skips security-flagged proposals unless `--include-security-flagged` is set.

### Changed

- Prefer `WHALESHELL_*` env over legacy `OPENSHELL_*` when both are set.
- Doctor warns on KEK migration need and binary-scoped identity limits on Docker Desktop.

### Fixed

- GoReleaser archive helpers use `strip_parent` so paths land at `libexec/whaleshell/linux-<arch>/`.
