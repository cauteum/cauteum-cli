# Bring Your Own Container

Run a sandbox with a **custom** image (OpenShell-style BYOC). whaleshell does not require
its first-party agent layers — any standard Linux image works if it meets the
contract in the [image reference](https://whaleshell.github.io/reference/images/).

## Quick start

```bash
docker build -t whaleshell-byoc:latest whaleshell-cli/examples/bring-your-own-container

whaleshell sandbox create --name byoc \
  --image whaleshell-byoc:latest \
  --workspace "$PWD" \
  --policy whaleshell-cli/policies/default.yaml \
  --no-proxy \
  -- python /sandbox/app.py

whaleshell sandbox exec byoc -- curl -sf http://127.0.0.1:8080/hello
whaleshell sandbox rm byoc
```

Or register an alias:

```yaml
# ~/.config/whaleshell/config.yaml
images:
  byoc: whaleshell-byoc:latest
```

```bash
whaleshell sandbox create --name byoc --from byoc --workspace . --policy whaleshell-cli/policies/default.yaml --no-proxy -- python /sandbox/app.py
```

## Requirements

| Rule | Why |
|------|-----|
| Non-root `USER` (or policy `run_as_*`) | Matches Docker/Podman sandbox identity |
| Writable `/sandbox` or `/workspace` | Agent/workdir |
| `iproute2` recommended | Netns / routing |
| No distroless / `FROM scratch` | Need a real userland |
| Pass command after `--` | Image CMD is replaced by `whaleshell-init` |

First-party images (`--from cursor`, …): build locally or `task images:pull` / GHCR — [image reference](https://whaleshell.github.io/reference/images/).
