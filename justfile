# bootstrap both apps
bootstrap:
  cd modules/web && pnpm install

# start infra (db/redis/traefik)
up:
  docker compose -f docker-compose.dev.yml up -d

# stop infra
down:
  docker compose -f docker-compose.dev.yml down

# run both apps (two terminals recommended)
api:
  cd modules/api && go run cmd/api/main.go

web:
  cd modules/web && pnpm dev

# lint & tests
lint:
  cd modules/web && pnpm lint && pnpm typecheck

test:
  cd modules/web && pnpm test
