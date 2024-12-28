#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: CORE_GATEWAY-GRPC"
echo "service: core-gateway-grpc"
echo "poort : 3004"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "ruby main.rb"

ruby main.rb
# tail -f /dev/null
