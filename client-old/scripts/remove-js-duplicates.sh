#!/bin/bash

# Set the root directory (change if needed)
ROOT_DIR="./src"

# Find all .js files
find "$ROOT_DIR" -type f -name "*.jsx" | while read jsfile; do
  base="${jsfile%.jsx}"
  if [ -f "$base.tsx" ]; then
    echo "🗑️ Deleting: $jsfile"
    rm "$jsfile"
  fi
done