# SPDX-FileCopyrightText: Copyright (c) 2026 the cauteum authors
# SPDX-License-Identifier: Apache-2.0
#
# OpenAI Codex CLI preinstalled.
# Base (cauteum-runtime sandboxes/base) already ships NodeSource Node.js + npm —
# Debian's `npm` package conflicts with that nodejs (Conflicts: npm).

ARG BASE_IMAGE=cauteum-sandbox:local
FROM ${BASE_IMAGE}

ENTRYPOINT []
USER root
ENV DEBIAN_FRONTEND=noninteractive \
    HOME=/root

# Refresh/ensure @openai/codex on the NodeSource toolchain already in base.
RUN npm install -g @openai/codex \
    && CODE_BIN="$(command -v codex || true)" \
    && if [ -z "$CODE_BIN" ]; then CODE_BIN="$(find "$(npm root -g)" -name 'codex' -type f 2>/dev/null | head -1)"; fi \
    && if [ -z "$CODE_BIN" ]; then \
         echo "codex binary not found after npm install -g @openai/codex" >&2; \
         npm root -g; ls -la "$(npm root -g)/@openai" 2>/dev/null || true; \
         exit 1; \
       fi \
    && ln -sfn "$CODE_BIN" /usr/local/bin/codex \
    && /usr/local/bin/codex --version || true \
    && npm cache clean --force

RUN mkdir -p /cauteum/data/home/.local/bin /cauteum/data/bin /etc/profile.d \
    && ln -sfn /usr/local/bin/codex /cauteum/data/home/.local/bin/codex \
    && printf '%s\n' 'export PATH="/usr/local/bin:/cauteum/data/home/.local/bin:/cauteum/data/bin:$PATH"' > /etc/profile.d/cauteum-agent.sh

ENV HOME=/cauteum/data/home \
    PATH=/usr/local/bin:/cauteum/data/home/.local/bin:/cauteum/data/bin:/usr/local/sbin:/usr/sbin:/usr/bin:/sbin:/bin

WORKDIR /workspace
ENTRYPOINT ["/usr/local/bin/cauteum-init"]
CMD ["--", "sleep", "infinity"]
