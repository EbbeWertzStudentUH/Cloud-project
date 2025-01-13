#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: APP_PROJECT_FACADE-RPC"
echo "service: app-project-facade-rpc"
echo "poort : 3009"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "go run ."
go run .
# tail -f /dev/null
