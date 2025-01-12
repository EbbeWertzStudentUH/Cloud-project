#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: CLIENT_DEVSTATS-SOAP"
echo "service: client-devstats"
echo "poort : 8082"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "none (tail -f /dev/null)"

tail -f /dev/null
