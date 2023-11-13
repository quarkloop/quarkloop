-- DropForeignKey
ALTER TABLE IF EXISTS "auth"."Account"
DROP CONSTRAINT "Account_userId_fkey";

-- DropForeignKey
ALTER TABLE IF EXISTS "auth"."Session"
DROP CONSTRAINT "Session_userId_fkey";

-- DropTable
DROP TABLE IF EXISTS "auth"."User";

-- DropTable
DROP TABLE IF EXISTS "auth"."Account";

-- DropTable
DROP TABLE IF EXISTS "auth"."Session";

-- DropTable
DROP TABLE IF EXISTS "auth"."VerificationToken";

-- DropSchema
DROP SCHEMA IF EXISTS "auth";