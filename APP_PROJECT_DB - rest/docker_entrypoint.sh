#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: APP_PROJECT_DB-REST"
echo "service: app-project-db-rest"
echo "poort : 3009"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "python main.py"

python main.py
# tail -f /dev/null
