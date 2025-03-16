#!/bin/bash

# Display starting message
echo "🚀 Starting Software Slayer development environment..."

# Navigate to server directory
cd "$(dirname "$0")/server"
echo "📦 Building server Docker container..."

# Build the Docker image for the server
docker build -t software-slayer-server ./app

# Start docker-compose in detached mode
echo "🐳 Starting Docker containers..."
docker-compose up -d

# Check if Docker containers started successfully
if [ $? -ne 0 ]; then
  echo "❌ Failed to start Docker containers. Exiting..."
  exit 1
fi

echo "✅ Server is running!"

# Navigate to client directory and start Expo
cd ../client
echo "📱 Starting Expo client..."
npm start

# Handle script exit
cleanup() {
  echo "🛑 Shutting down development environment..."
  cd ../server
  docker-compose down
  echo "👋 Done!"
}

# Set up trap to catch exit and run cleanup
trap cleanup EXIT INT TERM