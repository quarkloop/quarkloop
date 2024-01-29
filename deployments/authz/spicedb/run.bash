#!/bin/bash


docker run --rm \
    -p 50051:50051 \
    -p 8443:8443 \
    -p 9090:9090 \
    --network="host" \
    authzed/spicedb serve \
        --http-enabled \
        --grpc-preshared-key "somerandomkeyhere" \
        --datastore-engine "postgres" \
        --datastore-conn-uri "postgres://postgres:changeme@localhost:5432/spicedb" 