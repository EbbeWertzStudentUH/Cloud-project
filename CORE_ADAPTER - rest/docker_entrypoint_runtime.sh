#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: CORE_ADAPTER-REST (RUNTIME STAGE)"
echo "service: core-apapter-rest"
echo "poort : 3001"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "/app/GATEWAY-rest"

/app/GATEWAY-rest

# tail -f /dev/null
