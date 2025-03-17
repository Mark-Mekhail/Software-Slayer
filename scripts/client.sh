#!/bin/bash

if [ $# -ne 1 ] || [[ ! "$1" =~ ^(ios|android)$ ]]; then
  echo "Usage: $0 <platform>"
  echo "Supported platforms: ios, android"
  exit 1
fi

PLATFORM=$1

echo "📱 Starting Expo client for $PLATFORM..."
cd "$(dirname "$0")/../client"
npm run $PLATFORM
