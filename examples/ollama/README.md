# Ollama in whaleshell

Run Ollama on the host, reach it from the sandbox via `host.whaleshell.internal:11434`

Policy: `filesystem_policy` / `network_policies` + `inference.providers: [local]` — see `policy.yaml`.

**Запуск CLI:** [профили провайдеров](https://whaleshell.github.io/ru/guides/provider-profiles/).

```bash
export GOWORK=$PWD/go.work
go build -C whaleshell-cli -o ../whaleshell ./cmd/whaleshell && ./whaleshell install
task runtime:image:cli

# host
ollama serve   # listens on 11434

whaleshell sandbox create --name ollama-demo \
  --policy whaleshell-cli/examples/ollama/policy.yaml \
  --workspace .

whaleshell sandbox exec ollama-demo -- curl -s http://host.whaleshell.internal:11434/api/tags
```

No managed URL rewrite — use the native Ollama HTTP API. See [provider profiles](https://whaleshell.github.io/guides/provider-profiles/).
