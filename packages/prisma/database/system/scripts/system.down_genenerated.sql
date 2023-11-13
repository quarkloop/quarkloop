-- DropForeignKey
ALTER TABLE "system"."Workspace" DROP CONSTRAINT "Workspace_osId_fkey";

-- DropForeignKey
ALTER TABLE "system"."App" DROP CONSTRAINT "App_osId_fkey";

-- DropForeignKey
ALTER TABLE "system"."App" DROP CONSTRAINT "App_workspaceId_fkey";

-- DropForeignKey
ALTER TABLE "system"."AppComponent" DROP CONSTRAINT "AppComponent_appId_fkey";

-- DropForeignKey
ALTER TABLE "system"."AppIssue" DROP CONSTRAINT "AppIssue_appId_fkey";

-- DropForeignKey
ALTER TABLE "system"."AppSubmission" DROP CONSTRAINT "AppSubmission_appId_fkey";

-- DropForeignKey
ALTER TABLE "system"."AppDeployment" DROP CONSTRAINT "AppDeployment_appId_fkey";

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

-- CreateTable
CREATE TABLE "App" (
    "id" TEXT NOT NULL,
    "osId" TEXT NOT NULL,
    "workspaceId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "path" TEXT NOT NULL,
    "description" TEXT,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "App_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "AppComponent" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "config" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "appId" TEXT,

    CONSTRAINT "AppComponent_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "AppDeployment" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "appId" TEXT,

    CONSTRAINT "AppDeployment_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "AppIssue" (
    "id" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "appId" TEXT,

    CONSTRAINT "AppIssue_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "AppSubmission" (
    "id" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "appId" TEXT,

    CONSTRAINT "AppSubmission_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "OperatingSystem" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "path" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "overview" TEXT,
    "imageUrl" TEXT,
    "websiteUrl" TEXT,
    "isVerified" BOOLEAN NOT NULL DEFAULT false,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "OperatingSystem_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Workspace" (
    "id" TEXT NOT NULL,
    "osId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "path" TEXT NOT NULL,
    "description" TEXT,
    "accessType" INTEGER,
    "imageUrl" TEXT,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "Workspace_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "App_path_key" ON "App"("path" ASC);

-- CreateIndex
CREATE UNIQUE INDEX "OperatingSystem_path_key" ON "OperatingSystem"("path" ASC);

-- CreateIndex
CREATE UNIQUE INDEX "Workspace_osId_id_key" ON "Workspace"("osId" ASC, "id" ASC);

-- CreateIndex
CREATE UNIQUE INDEX "Workspace_path_key" ON "Workspace"("path" ASC);

-- AddForeignKey
ALTER TABLE "App" ADD CONSTRAINT "App_osId_fkey" FOREIGN KEY ("osId") REFERENCES "OperatingSystem"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "App" ADD CONSTRAINT "App_workspaceId_fkey" FOREIGN KEY ("workspaceId") REFERENCES "Workspace"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "AppComponent" ADD CONSTRAINT "AppComponent_appId_fkey" FOREIGN KEY ("appId") REFERENCES "App"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "AppDeployment" ADD CONSTRAINT "AppDeployment_appId_fkey" FOREIGN KEY ("appId") REFERENCES "App"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "AppIssue" ADD CONSTRAINT "AppIssue_appId_fkey" FOREIGN KEY ("appId") REFERENCES "App"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "AppSubmission" ADD CONSTRAINT "AppSubmission_appId_fkey" FOREIGN KEY ("appId") REFERENCES "App"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Workspace" ADD CONSTRAINT "Workspace_osId_fkey" FOREIGN KEY ("osId") REFERENCES "OperatingSystem"("id") ON DELETE CASCADE ON UPDATE CASCADE;

