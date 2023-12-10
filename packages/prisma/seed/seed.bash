#!/bin/bash

seed() {
    npx ts-node --compiler-options {\"module\":\"CommonJS\"} $1
}

echo

seed $PWD/seed/seed.ts

