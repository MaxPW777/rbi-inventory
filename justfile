# bootstrap both apps
bootstrap:
  cd modules/api && mix deps.get
  cd modules/web && pnpm install

# start infra (db/redis/traefik)
up:
  docker compose -f docker-compose.dev.yml up -d

# stop infra
down:
  docker compose -f docker-compose.dev.yml down

# DB tasks
db-reset:
  cd modules/api && MIX_ENV=dev mix ecto.drop && mix ecto.create && mix ecto.migrate && mix run priv/repo/seeds.exs

# run both apps (two terminals recommended)
api:
  cd modules/api && mix phx.server

web:
  cd modules/web && pnpm dev

# lint & tests
lint:
  cd modules/api && mix format --check-formatted && mix credo --strict
  cd modules/web && pnpm lint && pnpm typecheck

test:
  cd modules/api && mix test
  cd modules/web && pnpm test
