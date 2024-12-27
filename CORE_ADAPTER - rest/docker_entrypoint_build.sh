#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: GATEWAY-rest (BUILDER STAGE)"
echo "service: rest-gateway-svc"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo "=== content van /app/target ==="
ls /app/target
echo "=== content van /app/target/release ==="
ls /app/target/release

# tail -f /dev/null
