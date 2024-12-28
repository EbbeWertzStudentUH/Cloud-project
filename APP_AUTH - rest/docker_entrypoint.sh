#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: APP_AUTH-REST"
echo "service: app-auth-rest"
echo "poort : 3003"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "ruby main.rb"

ruby main.rb
# tail -f /dev/null
