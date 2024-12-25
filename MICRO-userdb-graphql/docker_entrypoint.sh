#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: MICRO-userdb-graphql"
echo "service: userdb-graphql-svc"
echo "interne poort : 3002"
echo "externe poort: /"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo
echo "=== Server start command: ==="
echo "node index.js"

node index.js

# tail -f /dev/null
