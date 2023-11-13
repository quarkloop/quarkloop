#!/bin/bash

source ../../.env 

npx prisma migrate diff \
    --from-schema-datamodel QuarkloopSheets.prisma \
    --to-migrations migrations \
    --shadow-database-url $PG_QUARKLOOP_SHEETS_URL \
    --script > scripts/app_sheets.down_genenerated.sql

npx prisma db execute \
    --file scripts/app_sheets.down.sql \
    --schema QuarkloopSheets.prisma

# echo 'DROP DATABASE IF EXISTS "QuarkloopSheets";' | \
# npx prisma db execute \
#     --stdin \
#     --url="$PG_QUARKLOOP_SHEETS_URL"    

