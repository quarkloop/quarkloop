#!/bin/bash

source ../../.env 

npx prisma migrate diff \
    --from-schema-datamodel QuarkloopEngine.prisma \
    --to-migrations migrations \
    --shadow-database-url $PG_QUARKLOOP_ENGINE_URL \
    --script > scripts/engine.down_genenerated.sql

npx prisma db execute \
    --file scripts/engine.down.sql \
    --schema QuarkloopEngine.prisma

