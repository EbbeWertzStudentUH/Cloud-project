#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: APP_USERDB-GRAPHQL"
echo "service: app-userdb-graphql"
echo "poort : 3002"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "node index.js"

node index.js

# tail -f /dev/null
