#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: CORE_GATEWAY-GRPC"
echo "service: core-gateway-grpc"
echo "poort : 3006"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "go run ."

go run .
# tail -f /dev/null
