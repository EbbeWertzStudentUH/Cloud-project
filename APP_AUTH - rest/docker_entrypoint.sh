#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: MICRO-auth-rest"
echo "service: auth-rest-svc"
echo "interne poort : 4567"
echo "externe poort: 4567"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "ruby main.rb"

ruby main.rb
# tail -f /dev/null
