#!/bin/bash

seed() {
    npx ts-node --compiler-options {\"module\":\"CommonJS\"} $1
}

echo


# seed $PWD/prisma/seed/seed.App.ts

