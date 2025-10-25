#!/bin/sh
set -e

echo "🚀 Running database migrations..."
/app/migrate -path /app/migration -database "postgresql://root:mysecret@postgres:5432/simple_bank?sslmode=disable" -verbose up


echo "start the app"
exec "$@"
