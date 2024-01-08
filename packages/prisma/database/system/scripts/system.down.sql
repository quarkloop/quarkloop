-- DropForeignKey
ALTER TABLE "system"."Workspace"
DROP CONSTRAINT "Workspace_orgId_fkey";

-- DropConstraint
ALTER TABLE "system"."UserAssignment"
DROP CONSTRAINT "UserId_UserGroupId_Check_NotNull";

-- Permission
DROP TABLE "system"."Permission";

-- UserAssignment
DROP TABLE "system"."UserAssignment";

-- UserRole
DROP TABLE "system"."UserRole";

-- UserAssignment
DROP TABLE "system"."UserGroup";

-- DropForeignKey
ALTER TABLE "system"."Project"
DROP CONSTRAINT "Project_orgId_fkey";

-- DropForeignKey
ALTER TABLE "system"."Project"
DROP CONSTRAINT "Project_workspaceId_fkey";

-- DropTable
DROP TABLE "system"."Project";

-- DropTable
DROP TABLE "system"."Workspace";

-- DropTable
DROP TABLE "system"."Organization";

-- _prisma_migrations
DROP TABLE "system"."_prisma_migrations";

-- DropSchema
DROP SCHEMA IF EXISTS "system";