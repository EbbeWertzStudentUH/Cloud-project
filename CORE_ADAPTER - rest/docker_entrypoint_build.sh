#!/bin/sh

echo "=== Docker Container Info ==="
echo "container: CORE_ADAPTER-REST (BUILDER STAGE)"
echo "service: core-apapter-rest"
echo "==============================="
echo
echo "=== content van /app ==="
ls /app
echo "=== content van /app/target ==="
ls /app/target
echo "=== content van /app/target/release ==="
ls /app/target/release

# tail -f /dev/null
