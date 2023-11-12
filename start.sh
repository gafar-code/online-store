#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "postgresql://root:secret@3.94.76.231:5432/store_db?sslmode=disable" -verbose up

echo "start the app"
exec "$@"