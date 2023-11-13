#!/bin/bash

source ../../.env 

npx prisma migrate diff \
    --from-schema-datamodel QuarkloopSystem.prisma \
    --to-migrations migrations \
    --shadow-database-url $PG_QUARKLOOP_SYSTEM_URL \
    --script > scripts/system.down_genenerated.sql

npx prisma db execute \
    --file scripts/system.down.sql \
    --schema QuarkloopSystem.prisma

