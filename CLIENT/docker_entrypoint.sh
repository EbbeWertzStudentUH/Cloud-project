#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: MAIN-CLIENT-svelte"
echo "service: svelte-app-svc"
echo "interne poort : 3000"
echo "externe poort: 80"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "node build/index.js"

node build/index.js

# tail -f /dev/null
