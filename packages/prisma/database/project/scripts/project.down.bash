#!/bin/bash

source ../../.env 

npx prisma migrate diff \
    --from-schema-datamodel QuarkloopProject.prisma \
    --to-migrations migrations \
    --shadow-database-url $PG_QUARKLOOP_PROJECT_URL \
    --script > scripts/project.down_genenerated.sql

npx prisma db execute \
    --file scripts/project.down.sql \
    --schema QuarkloopProject.prisma

# echo 'DROP DATABASE IF EXISTS "QuarkloopProject";' | \
# npx prisma db execute \
#     --stdin \
#     --url="$PG_QUARKLOOP_SHEETS_URL"    

