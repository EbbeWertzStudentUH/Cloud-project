#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: APP_GHSTATS-REST"
echo "service: app-ghstats-rest"
echo "poort : 3010"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "node dist/app.js"

node dist/app.js

# tail -f /dev/null
