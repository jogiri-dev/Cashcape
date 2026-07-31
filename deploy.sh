#!/bin/bash
set -euo pipefail

ENV="${1:-staging}"

if [[ "$ENV" != "staging" && "$ENV" != "prod" ]]; then
  echo "Usage: ./deploy.sh [staging|prod]"
  exit 1
fi

cd /srv/docker/cashcape

echo "[$(date)] Deploying $ENV..."
docker compose --env-file ".env.$ENV" -f "docker-compose.$ENV.yml" pull
docker compose --env-file ".env.$ENV" -f "docker-compose.$ENV.yml" up -d
echo "[$(date)] $ENV deploy complete."