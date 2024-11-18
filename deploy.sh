#!/bin/bash

COMPOSE_FILE="docker/docker-compose.yml"
CONTAINER_NAME="app_container"
SCRIPT_PATH="/migrations/migrate.sh"

echo "Starting docker-compose..."
docker-compose -f $COMPOSE_FILE up -d --build

if [ $? -ne 0 ]; then
  echo "Failed to start docker-compose."
  exit 1
fi

echo "Waiting for container to start..."
sleep 10

echo "Running migration script inside the container..."
docker exec -it $CONTAINER_NAME sh -c "$SCRIPT_PATH"

if [ $? -ne 0 ]; then
  echo "Failed to execute migration script."
  exit 1
fi

echo "Migration script executed successfully."
