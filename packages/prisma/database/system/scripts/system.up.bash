#!/bin/bash

set -a # automatically export all variables

mkdir -p ./migrations/test ./migrations/dev


create_test_db() {
    source ../../.env.test
    rm -rf ./migrations/test/*
    cp QuarkloopSystem.prisma ./migrations/test/QuarkloopSystem.prisma

    # db down
    npx prisma db execute \
        --file scripts/system.down.sql \
        --schema ./migrations/test/QuarkloopSystem.prisma

    # db up
    npx prisma migrate dev \
        --skip-seed \
        --skip-generate \
        --schema ./migrations/test/QuarkloopSystem.prisma

    echo '
    ALTER TABLE 
        "system"."UserAssignment" 
    ADD CONSTRAINT 
        "UserId_UserGroupId_Check_NotNull" CHECK ("userId" IS NOT NULL OR "userGroupId" IS NOT NULL);
    ' | \
    npx prisma db execute --stdin --schema ./migrations/test/QuarkloopSystem.prisma       
}

create_dev_db() {
    source ../../.env 
    cp QuarkloopSystem.prisma ./migrations/dev/QuarkloopSystem.prisma

    # db up
    npx prisma migrate dev \
        --skip-seed \
        --skip-generate \
        --schema ./migrations/dev/QuarkloopSystem.prisma

    echo '
    ALTER TABLE 
        "system"."UserAssignment" 
    ADD CONSTRAINT 
        "UserId_UserGroupId_Check_NotNull" CHECK ("userId" IS NOT NULL OR "userGroupId" IS NOT NULL);
    ' | \
    npx prisma db execute --stdin --schema ./migrations/dev/QuarkloopSystem.prisma
}

case "$1" in
  "test")
    create_test_db
    ;;
  "dev")
    create_dev_db
    ;;
  *)
    echo "You have failed to specify what to do correctly."
    exit 1
    ;;
esac

set +a # automatically unexport all variables