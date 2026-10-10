# Providers (local copies)

Canonical builtin profiles live in **`cauteum-providers/profiles/`**.

This directory keeps copies for older docs/paths; prefer importing from:

```bash
cauteum provider profile import ../cauteum-providers/profiles/github.yaml
```

`provider.FindBuiltinDir()` resolves `cauteum-providers/profiles` first.
