#!/bin/bash

source ../../.env 

npx prisma migrate dev \
    --skip-seed \
    --skip-generate \
    --schema QuarkloopProject.prisma