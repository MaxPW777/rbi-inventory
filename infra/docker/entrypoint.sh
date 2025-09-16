#!/usr/bin/env bash
set -euo pipefail

# run Ecto migrations before boot
if [ -x "/app/bin/inventory" ]; then
  /app/bin/inventory eval "Inventory.Release.migrate"
fi

exec /app/bin/inventory start
