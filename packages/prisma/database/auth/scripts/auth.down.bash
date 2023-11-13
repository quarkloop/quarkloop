#!/bin/bash

source ../../.env 

npx prisma migrate diff \
    --from-schema-datamodel QuarkloopAuth.prisma \
    --to-migrations migrations \
    --shadow-database-url $PG_QUARKLOOP_AUTH_URL \
    --script > scripts/auth.down_genenerated.sql

npx prisma db execute \
    --file scripts/auth.down.sql \
    --schema QuarkloopAuth.prisma

