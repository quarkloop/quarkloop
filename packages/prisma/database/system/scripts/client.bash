#!/usr/bin/env bash

set -a
source ../../.env

PGDATABASE=${PG_QUARKLOOP_SYSTEM_DB}

docker-compose \
    -f ./scripts/docker-compose.yml \
    run --rm \
    pg_client -c 'SELECT version();' 

set +a # automatically unexport all variables    