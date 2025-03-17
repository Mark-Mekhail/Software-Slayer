#!/bin/bash

echo "🚀 Starting Docker build and containers..."

# Navigate to server directory
cd "$(dirname "$0")/../server"
echo "📦 Building server Docker container..."
docker build -t software-slayer-server ./app

echo "🐳 Starting Docker containers..."
docker-compose up -d

if [ $? -ne 0 ]; then
  echo "❌ Failed to start Docker containers. Exiting..."
  exit 1
fi

echo "✅ Server is running!"

cleanup() {
  echo "🛑 Shutting down..."
  cd ../server
  docker-compose down
  echo "👋 Done!"
}

trap cleanup EXIT INT TERM

# Keep script running for logs (optional)
while true; do
  sleep 1
done
