# Claude / Codex-style agents

Use builtin Anthropic + OpenAI presets **or** a `claude-code` provider instance.
API keys stay on the host / gateway; the sandbox sees `whaleshell:resolve:env:…` placeholders.

**Запуск + providers:** [профили провайдеров](https://whaleshell.github.io/ru/guides/provider-profiles/).

## A. Inference presets in policy (no gateway provider)

```bash
export GOWORK=$PWD/go.work
# keys on host for rewrite when using local proxy path — prefer gateway provider (B)

whaleshell sandbox create --name agent-demo \
  --policy whaleshell-cli/examples/claude-codex/policy.yaml \
  --workspace .
```

## B. Gateway provider (рекомендуется)

```bash
whaleshell-gateway --listen 127.0.0.1:7443 &
whaleshell gateway add http://127.0.0.1:7443 --local --name local && whaleshell gateway select local

ANTHROPIC_API_KEY=… whaleshell provider create --name claude --type claude-code --credential ANTHROPIC_API_KEY

whaleshell sandbox create --name agent-demo \
  --policy whaleshell-cli/examples/claude-codex/policy.yaml \
  --workspace . \
  --gateway http://127.0.0.1:7443 \
  --provider claude

whaleshell sandbox exec agent-demo -- env | grep -E 'ANTHROPIC|OPENAI'
# expect: ANTHROPIC_API_KEY=whaleshell:resolve:env:ANTHROPIC_API_KEY
```

Custom host model (vLLM): `inference.profiles` — see [provider profiles](https://whaleshell.github.io/guides/provider-profiles/).
