#!/bin/bash

source ../../.env 

npx prisma migrate dev \
    --skip-seed \
    --skip-generate \
    --schema QuarkloopSystem.prisma

echo '
ALTER TABLE 
    "system"."UserAssignment" 
ADD CONSTRAINT 
    "UserId_UserGroupId_Check_NotNull" CHECK ("userId" IS NOT NULL OR "userGroupId" IS NOT NULL);
' | \
npx prisma db execute --stdin --schema QuarkloopSystem.prisma
