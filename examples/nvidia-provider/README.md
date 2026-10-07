# NVIDIA provider profile example

Builtin profile: `whaleshell-providers/profiles/nvidia.yaml` (also via `whaleshell provider profile show nvidia`).

**Запуск + create:** [профили провайдеров](https://whaleshell.github.io/ru/guides/provider-profiles/).

```bash
export GOWORK=$PWD/go.work
go build -C whaleshell-cli -o whaleshell ./cmd/whaleshell && ./whaleshell install
go build -C whaleshell-gateway -o whaleshell-gateway ./cmd/whaleshell-gateway
./whaleshell-gateway --listen 127.0.0.1:7443 &
whaleshell gateway add http://127.0.0.1:7443 --local --name local && whaleshell gateway select local

NVIDIA_API_KEY=… whaleshell provider create --name nv --type nvidia --credential NVIDIA_API_KEY

whaleshell sandbox create --name nim \
  --policy whaleshell-cli/policies/default.yaml \
  --workspace . \
  --gateway http://127.0.0.1:7443 \
  --provider nv

whaleshell provider effective nim
whaleshell sandbox exec nim -- env | grep NVIDIA || true
# expect: NVIDIA_API_KEY=whaleshell:resolve:env:NVIDIA_API_KEY
```

См. [профили провайдеров](https://whaleshell.github.io/ru/guides/provider-profiles/).
