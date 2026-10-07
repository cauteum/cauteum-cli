# whaleshell-cli/docker/

Agent sandbox images for the `whaleshell` CLI (`--from cursor|claude|codex`).  
Runtime base (`whaleshell-sandbox:local` / GHCR `:cli`) is built from the [runtime image sources](https://github.com/whaleshell/whaleshell-runtime/tree/main/images/sandbox) via `task runtime:image:cli` in the multi-repo workspace.

| Path | Tag | Contents |
|------|-----|----------|
| [`agents/cursor`](./agents/cursor/) | `whaleshell-sandbox:cursor` | + Cursor Agent CLI under `/opt/cursor-agent` |
| [`agents/claude`](./agents/claude/) | `whaleshell-sandbox:claude` | + Claude Code CLI |
| [`agents/codex`](./agents/codex/) | `whaleshell-sandbox:codex` | + OpenAI Codex CLI |

Published images and BYOC: [image reference](https://whaleshell.github.io/reference/images/).

## Build

```bash
export GOWORK=$PWD/go.work
task runtime:image:cli          # base → whaleshell-sandbox:local
task docker:agent:cursor        # → whaleshell-sandbox:cursor
task docker:agent:claude
task docker:agent:codex
task docker:agent:all
```

## Use

```bash
whaleshell sandbox create --name cursor \
  --from cursor \
  --workspace . \
  --policy whaleshell-cli/policies/cursor.yaml \
  -- agent
```
