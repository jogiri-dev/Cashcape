#!/bin/bash

# Set target directory from argument or default to 'src'
TARGET_DIR="${1:-src}"

# Find all .js files in the target directory
find "$TARGET_DIR" -name "*.js" | while read file; do
  # Check if the file contains JSX-like syntax
  if grep -q "<[A-Za-z]" "$file" || grep -q "<>" "$file"; then
    # Rename the file to .jsx
    mv "$file" "${file%.js}.jsx"
    echo "Renamed: $file → ${file%.js}.jsx"
  fi
done