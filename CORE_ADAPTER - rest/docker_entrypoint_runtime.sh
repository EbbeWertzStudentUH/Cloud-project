#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: GATEWAY-rest (RUNTIME STAGE)"
echo "service: rest-gateway-svc"
echo "interne poort : 3001"
echo "externe poort: (intern)"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "/app/GATEWAY-rest"

/app/GATEWAY-rest

# tail -f /dev/null
