#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: CORE_NOTIFIER-WS-REST"
echo "service: core-notifier-ws-rest"
echo "poort : 3004"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "pyhton main.py"

pyhton main.py
# tail -f /dev/null
