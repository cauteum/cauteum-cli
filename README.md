<h1 align="center">whaleshell-cli</h1>

<p align="center">
  <strong>Agent sandbox CLI</strong><br>
  Create, harden, and operate policy-bound sandboxes — Cursor, Claude, Codex, and BYOC.
</p>
<p align="center">
  <a href="https://github.com/whaleshell/whaleshell-cli/actions/workflows/ci.yml"><img src="https://github.com/whaleshell/whaleshell-cli/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
  <a href="https://github.com/whaleshell/whaleshell-cli/releases"><img src="https://img.shields.io/github/v/release/whaleshell/whaleshell-cli?include_prereleases&sort=semver&label=release" alt="release"></a>
  <a href="https://img.shields.io/badge/status-beta-blueviolet"><img src="https://img.shields.io/badge/status-beta-blueviolet" alt="beta"></a>
  <a href="https://www.apache.org/licenses/LICENSE-2.0"><img src="https://img.shields.io/badge/License-Apache--2.0-blue.svg" alt="License"></a>
  <a href="https://github.com/whaleshell/whaleshell-cli"><img src="https://img.shields.io/badge/Go-1.27+-00ADD8?logo=go" alt="Go Version"></a>
</p>
<p align="center">
  <sub>Part of the <a href="https://github.com/whaleshell">whaleshell / whaleshell</a> ecosystem</sub>
</p>

---

## Overview

The [quick start](https://whaleshell.github.io/get-started/) contains the supported installation and first-sandbox flow.

**whaleshell-cli** is the user-facing `whaleshell` binary for the whaleshell ecosystem: sandbox lifecycle, policy checks, provider attach, gateway selection, live logs, and agent images.

### Key Features

| Category | Capabilities |
|----------|--------------|
| **Sandboxes** | create / exec / connect / rm — Docker-first via `whaleshell-driver` |
| **Policy** | YAML check/compose with `whaleshell-core` + provider presets |
| **Agents** | `--from cursor\|claude\|codex\|base\|gui\|gpu` (GHCR or local) |
| **Providers** | Cursor, GitHub, NVIDIA, local inference (`host.whaleshell.internal`) |
| **Observability** | `whaleshell logs --tail`, TUI (`whaleshell term`), OCSF audit lines |
| **Gateway** | register with `whaleshell-gateway`, relay exec, policy proposals |

---

## Installation

Install the current beta (`v0.1.0-beta.2`) into `~/.local/bin`:

```bash
curl -LsSf https://raw.githubusercontent.com/whaleshell/whaleshell-cli/main/install.sh \
  | WHALESHELL_VERSION=v0.1.0-beta.2 sh
```

From source in the [multi-repo workspace](https://github.com/whaleshell):

```bash
go build -C whaleshell-cli -o ../whaleshell ./cmd/whaleshell
./whaleshell install
```

**Requirements:** Docker or Podman. Go 1.27+ only if building from source.

`v0.1.0-beta.2` is a GitHub prerelease. It builds from published module dependencies; OpenShell behavioral compatibility remains partial. See [development and releases](https://whaleshell.github.io/reference/development/) and the [compatibility status](https://whaleshell.github.io/reference/openshell-compatibility/).

---

## Quick Start

```bash
./whaleshell status
./whaleshell policy check ./policies/default.yaml

./whaleshell sandbox create --name demo --workspace . --policy ./policies/default.yaml
./whaleshell sandbox exec demo -- uname -a
./whaleshell sandbox rm demo
```

Testcontainers cleanup is exposed through the CLI rather than a broad Docker
prune:

```bash
./whaleshell doctor cleanup --dry-run
./whaleshell -o json doctor cleanup --yes
```

Only dangling anonymous test volumes and stopped containers labeled
`whaleshell=1` are in scope.

### Cursor agent

```bash
./whaleshell sandbox create --name cursor --from cursor --workspace "$PWD" \
  --provider cursor --provider gh --policy ./policies/cursor.yaml
./whaleshell sandbox connect cursor -- agent
```

Agent Dockerfiles live in [`docker/agents/`](./docker/agents/). Policies: [`policies/`](./policies/).
Runnable recipes live in [`examples/`](./examples/).

---

## Package Structure

| Path | Purpose |
|------|---------|
| `cmd/whaleshell` | CLI entrypoint |
| `internal/app` | Commands (sandbox, proxy, gateway, rules, …) |
| `policies/` | Builtin policy YAML |
| `docker/agents/` | cursor / claude / codex images |


---

## Related

| Resource | Link |
|----------|------|
| Roadmap | [ROADMAP.md](./ROADMAP.md) |
| Organization | [https://github.com/whaleshell](https://github.com/whaleshell) |
| Organization overview | [github.com/whaleshell](https://github.com/whaleshell) |
| pkg.go.dev | [`github.com/whaleshell/whaleshell-cli`](https://pkg.go.dev/github.com/whaleshell/whaleshell-cli) |

## License

[Apache-2.0](./LICENSE) © whaleshell
