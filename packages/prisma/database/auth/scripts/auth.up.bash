#!/bin/bash

set -a # automatically export all variables

mkdir -p ./migrations/test ./migrations/dev

create_test_db() {
    source ../../.env.test
    rm -rf ./migrations/test/*
    cp QuarkloopAuth.prisma ./migrations/test/QuarkloopAuth.prisma

    # db down
    npx prisma db execute \
        --file scripts/auth.down.sql \
        --schema ./migrations/test/QuarkloopAuth.prisma

    # db up
    npx prisma migrate dev \
        --skip-seed \
        --skip-generate \
        --schema ./migrations/test/QuarkloopAuth.prisma
}

create_dev_db() {
    source ../../.env 
    cp QuarkloopAuth.prisma ./migrations/dev/QuarkloopAuth.prisma

    npx prisma migrate dev \
        --skip-seed \
        --skip-generate \
        --schema ./migrations/dev/QuarkloopAuth.prisma
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