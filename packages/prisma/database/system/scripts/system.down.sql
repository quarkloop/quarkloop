-- DropForeignKey
ALTER TABLE "system"."Workspace"
DROP CONSTRAINT "Workspace_osId_fkey";

-- DropForeignKey
ALTER TABLE "system"."App"
DROP CONSTRAINT "App_osId_fkey";

-- DropForeignKey
ALTER TABLE "system"."App"
DROP CONSTRAINT "App_workspaceId_fkey";

-- DropForeignKey
ALTER TABLE "system"."AppComponent"
DROP CONSTRAINT "AppComponent_appId_fkey";

-- DropForeignKey
ALTER TABLE "system"."AppIssue"
DROP CONSTRAINT "AppIssue_appId_fkey";

-- DropForeignKey
ALTER TABLE "system"."AppSubmission"
DROP CONSTRAINT "AppSubmission_appId_fkey";

-- DropForeignKey
ALTER TABLE "system"."AppDeployment"
DROP CONSTRAINT "AppDeployment_appId_fkey";

-- DropTable
DROP TABLE "system"."OperatingSystem";

-- DropTable
DROP TABLE "system"."Workspace";

-- DropTable
DROP TABLE "system"."App";

-- DropTable
DROP TABLE "system"."AppComponent";

-- DropTable
DROP TABLE "system"."AppIssue";

-- DropTable
DROP TABLE "system"."AppSubmission";

-- DropTable
DROP TABLE "system"."AppDeployment";

-- DropSchema
DROP SCHEMA IF EXISTS "system";