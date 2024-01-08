-- DropForeignKey
ALTER TABLE "auth"."Account"
DROP CONSTRAINT "Account_userId_fkey";

-- DropForeignKey
ALTER TABLE "auth"."Session"
DROP CONSTRAINT "Session_userId_fkey";

-- DropTable
DROP TABLE "auth"."VerificationToken";

-- DropTable
DROP TABLE "auth"."Session";

-- DropTable
DROP TABLE "auth"."Account";

-- DropTable
DROP TABLE "auth"."User";

-- _prisma_migrations
DROP TABLE "auth"."_prisma_migrations";

-- DropSchema
DROP SCHEMA IF EXISTS "auth";