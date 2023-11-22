-- DropForeignKey
ALTER TABLE "system"."Workspace"
DROP CONSTRAINT "Workspace_orgId_fkey";

-- DropForeignKey
ALTER TABLE "system"."Project"
DROP CONSTRAINT "Project_orgId_fkey";

-- DropForeignKey
ALTER TABLE "system"."Project"
DROP CONSTRAINT "Project_workspaceId_fkey";

-- DropForeignKey
ALTER TABLE "system"."ProjectService"
DROP CONSTRAINT "ProjectService_projectId_fkey";

-- DropForeignKey
ALTER TABLE "system"."ProjectSubmission"
DROP CONSTRAINT "ProjectSubmission_projectId_fkey";

-- DropForeignKey
ALTER TABLE "system"."ProjectDiscussion"
DROP CONSTRAINT "ProjectDiscussion_projectId_fkey";

-- DropForeignKey
ALTER TABLE "system"."ProjectForm"
DROP CONSTRAINT "ProjectForm_projectId_fkey";

-- DropTable
DROP TABLE "system"."Organization";

-- DropTable
DROP TABLE "system"."Workspace";

-- DropTable
DROP TABLE "system"."Project";

-- DropTable
DROP TABLE "system"."ProjectService";

-- DropTable
DROP TABLE "system"."ProjectSubmission";

-- DropTable
DROP TABLE "system"."ProjectDiscussion";

-- DropTable
DROP TABLE "system"."ProjectForm";

-- DropSchema
DROP SCHEMA IF EXISTS "system";