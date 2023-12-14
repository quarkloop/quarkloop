-- DropForeignKey
ALTER TABLE "system"."Workspace"
DROP CONSTRAINT "Workspace_orgId_fkey";

-- DropForeignKey
ALTER TABLE "system"."Project"
DROP CONSTRAINT "Project_orgId_fkey";

-- DropForeignKey
ALTER TABLE "system"."Project"
DROP CONSTRAINT "Project_workspaceId_fkey";

-- DropTable
DROP TABLE "system"."Organization";

-- DropTable
DROP TABLE "system"."Workspace";

-- DropTable
DROP TABLE "system"."Project";

-- DropSchema
DROP SCHEMA IF EXISTS "system";