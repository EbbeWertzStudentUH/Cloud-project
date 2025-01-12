#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: APP_DEVSTATS-SOAP"
echo "service: app-devstats-soap"
echo "poort : 3007"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "gradle run --no-daemon "

gradle run --no-daemon 
# tail -f /dev/null
