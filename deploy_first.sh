#!/usr/bin/env bash
noneList=$(docker images | grep -v images | awk '{print $3}')
  if [ ! -z "$noneList" ]; then
    exit 0
  fi
docker network create my-bridge
#docker compose up --remove-orphans -d
echo "creat basic container"

