#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "postgresql://${POSTGRES_USERNAME}:${POSTGRES_PASS}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable" -verbose up

echo "start the app"
# exec "$@"
/app/main