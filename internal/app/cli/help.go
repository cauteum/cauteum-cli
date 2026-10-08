package cli

import (
	"fmt"
	"strings"
)

// wantsHelp reports whether argv asks for help (-h/--help/help).
func wantsHelp(args []string) bool {
	for _, a := range args {
		switch a {
		case "-h", "--help", "help":
			return true
		}
	}
	return false
}

// helpPath strips help flags and returns the command path for nested help.
func helpPath(args []string) []string {
	out := make([]string, 0, len(args))
	for _, a := range args {
		switch a {
		case "-h", "--help", "help":
			continue
		default:
			out = append(out, a)
		}
	}
	return out
}

func printHelp(args []string) error {
	path := helpPath(args)
	fmt.Print(helpText(path))
	return nil
}

func helpText(path []string) string {
	if len(path) == 0 {
		return rootHelp
	}
	key := strings.Join(path, " ")
	if t, ok := helpTree[key]; ok {
		return t
	}
	for i := len(path) - 1; i >= 1; i-- {
		parent := strings.Join(path[:i], " ")
		if t, ok := helpTree[parent]; ok {
			return t
		}
	}
	if t, ok := helpTree[path[0]]; ok {
		return t
	}
	return rootHelp
}

const rootHelp = `cauteum — OpenShell-compatible agent sandbox CLI

Usage:
  cauteum [global flags] <command> [flags]

Global flags:
  -g, --gateway NAME     Select gateway (or OPENSHELL_GATEWAY / CAUTEUM_GATEWAY)
  --workspace NAME       Default workspace (or OPENSHELL_WORKSPACE)
  -o, --output FORMAT    text|json|yaml

Commands:
  sandbox (sb)       Create and manage sandboxes
  exec               Execute a command in a sandbox
  provider           Provider instances and profiles
  policy (pol)       Network policy get/set/update
  gateway (gw)       Add/select/login gateways
  workspace (ws)     Workspaces and members
  service (svc)      Expose sandbox ports via *.openshell.localhost
  forward (fwd)      TCP port forwards
  inference          Route inference.local
  settings           Gateway/local settings
  logs (lg)          Sandbox logs
  status             Gateway connectivity
  health             Docker/Podman and gateway health probe
  init               Write an agent starter policy
  doctor (dr)        Environment checks
  whoami             Identity (supports -o json)
  term               Interactive TUI
  install            Install CLI + ensure local gateway
  completions        Shell completions
  version            Print version

Run 'cauteum <command> --help' for details.
`

var helpTree = map[string]string{
	"sandbox": `cauteum sandbox — manage sandboxes

Usage:
  cauteum sandbox create|list|get|stop|start|delete|exec|connect|upload|download|ssh-config|provider|template …

Aliases: sb

Examples:
  cauteum sandbox create --name app --from ollama
  cauteum sandbox list
  cauteum sandbox exec app -- ls /
`,
	"sandbox create": `cauteum sandbox create — create a sandbox

Usage:
  cauteum sandbox create --name NAME [flags]

Flags (OpenShell-aligned):
  --name NAME
  --from base|ollama|cursor|claude|…
  --image IMAGE
  --policy PATH
  --cpu N --memory SIZE   (or set defaults.memory in config / CAUTEUM_DEFAULT_MEMORY)
  --pids-limit N          (-1 unlimited; default 2048 via driver)
  --provider NAME (repeatable)
  --forward PORT (repeatable)
  --workspace PATH
  --label KEY=VALUE
  --driver-config-json JSON
`,
	"sandbox template": `cauteum sandbox template — workload templates

Usage:
  cauteum sandbox template create|list|get|delete …
`,
	"sandbox provider": `cauteum sandbox provider — attach providers to a sandbox

Usage:
  cauteum sandbox provider list|attach|detach …
`,
	"provider": `cauteum provider — provider instances

Usage:
  cauteum provider create|list|get|update|delete|profile|refresh|effective …

Examples:
  cauteum provider create --name gh --type github --from-existing
  cauteum provider refresh configure NAME --credential-key K --strategy oauth2-refresh-token
`,
	"provider profile": `cauteum provider profile — custom provider YAML profiles

Usage:
  cauteum profile list|show|import|update|export|delete|lint …
  cauteum profile import --url https://example.org/profile.yaml
  cauteum profile import --from ./provider-profiles
`,
	"profile": `cauteum profile — reusable provider definitions

Usage:
  cauteum profile list|show|import|update|export|delete|lint …
`,
	"provider refresh": `cauteum provider refresh — credential refresh strategies

Usage:
  cauteum provider refresh status|configure|rotate|delete …

Strategies: env | oauth2-refresh-token | oauth2-client-credentials | aws-sts-assume-role
`,
	"policy": `cauteum policy — network policy

Usage:
  cauteum policy get|set|update|list|delete|check …
`,
	"gateway": `cauteum gateway — manage gateways

Usage:
  cauteum gateway ensure|add|remove|select|info|list|login|logout

  ensure   start/select local gateway on 127.0.0.1:7443 if needed
`,
	"workspace": `cauteum workspace — workspaces (gateway-backed)

Usage:
  cauteum workspace create --name NAME
  cauteum workspace list|get|delete NAME
  cauteum workspace member add|remove|list …
`,
	"workspace member": `cauteum workspace member — manage members

Usage:
  cauteum workspace member add --workspace NAME --subject SUBJECT [--role user|admin]
  cauteum workspace member remove --workspace NAME --subject SUBJECT
  cauteum workspace member list --workspace NAME
`,
	"service": `cauteum service — expose HTTP services

Usage:
  cauteum service expose <sandbox> <port> [name]
  cauteum service list|get|delete …

Edge URL (gateway Host router):
  http://<name>.openshell.localhost:<gateway-port>/
`,
	"forward": `cauteum forward — TCP forwards into a sandbox

Usage:
  cauteum forward start <host-port> <sandbox> [-d]
  cauteum forward stop <id>
  cauteum forward list
`,
	"inference": `cauteum inference — inference.local routing

Usage:
  cauteum inference get|set|update|list|show|local
`,
	"settings": `cauteum settings — key/value settings

Usage:
  cauteum settings get|set|delete …
`,
	"logs": `cauteum logs — sandbox logs

Usage:
  cauteum logs <name> [--tail] [-n N] [--since 5m]
`,
	"doctor": `cauteum doctor — environment checks

Usage:
  cauteum doctor check
  cauteum doctor cleanup [--dry-run|--yes]

cleanup is a scoped dry-run by default. With --yes it removes only dangling
anonymous Testcontainers volumes and stopped containers labeled cauteum=1.
`,
	"install": `cauteum install — install CLI binary and ensure local gateway

Usage:
  cauteum install [--force]

Copies cauteum to ~/.local/share/cauteum/bin, symlinks ~/.local/bin/cauteum,
then starts/selects a local cauteum-gateway if needed.
`,
	"whoami": `cauteum whoami — print identity

Usage:
  cauteum whoami
  cauteum -o json whoami
`,
	"completions": `cauteum completions — shell completions

Usage:
  cauteum completions <bash|zsh|fish|powershell>
`,
	"status": `cauteum status — gateway connectivity

Usage:
  cauteum status
  cauteum -o json status
`,
	"health": `cauteum health — engine and gateway health probe

Usage:
  cauteum health
  cauteum -o json health
`,
	"init": `cauteum init — write an agent starter policy

Usage:
  cauteum init --agent cursor [--dir DIR] [--force]
`,
	"exec": `cauteum exec — execute a command in a sandbox

Usage:
  cauteum exec [--name] NAME -- COMMAND [ARG ...]
`,
	"term": `cauteum term — interactive TUI

Usage:
  cauteum term
`,
	"rule": `cauteum rule — approval rules (MVP)

Usage:
  cauteum rule get|approve|approve-all|reject|history|clear …
`,
}
