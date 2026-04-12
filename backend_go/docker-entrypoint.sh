#!/bin/sh
set -eu

MIGRATIONS_PATH="${MIGRATIONS_PATH:-/app/migrations}"

echo "Applying database migrations from ${MIGRATIONS_PATH}"
set +e
MIGRATE_OUTPUT="$(migrate -path "${MIGRATIONS_PATH}" -database "${DATABASE_URL}" up 2>&1)"
MIGRATE_STATUS=$?
set -e

if [ "${MIGRATE_STATUS}" -ne 0 ]; then
  echo "${MIGRATE_OUTPUT}"
  case "${MIGRATE_OUTPUT}" in
    *"no change"*)
      ;;
    *)
      exit "${MIGRATE_STATUS}"
      ;;
  esac
fi

exec /app/taskflow-api
