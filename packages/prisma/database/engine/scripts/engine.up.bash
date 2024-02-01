#!/bin/bash

set -a # automatically export all variables

mkdir -p ./migrations/test ./migrations/dev

create_test_db() {
    source ../../.env.test
    rm -rf ./migrations/test/*
    cp QuarkloopEngine.prisma ./migrations/test/QuarkloopEngine.prisma

    # db down
    npx prisma db execute \
        --file scripts/engine.down.sql \
        --schema ./migrations/test/QuarkloopEngine.prisma

    # db up
    npx prisma migrate dev \
        --skip-seed \
        --skip-generate \
        --schema ./migrations/test/QuarkloopEngine.prisma    
}

create_dev_db() {
    source ../../.env 
    cp QuarkloopEngine.prisma ./migrations/dev/QuarkloopEngine.prisma

    # db up
    npx prisma migrate dev \
        --skip-seed \
        --skip-generate \
        --schema ./migrations/dev/QuarkloopEngine.prisma
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