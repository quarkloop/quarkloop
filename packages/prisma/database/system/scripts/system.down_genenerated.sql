-- DropForeignKey
ALTER TABLE "system"."Workspace" DROP CONSTRAINT "Workspace_orgId_fkey";

-- DropForeignKey
ALTER TABLE "system"."Project" DROP CONSTRAINT "Project_orgId_fkey";

-- DropForeignKey
ALTER TABLE "system"."Project" DROP CONSTRAINT "Project_workspaceId_fkey";

-- DropForeignKey
ALTER TABLE "system"."ProjectService" DROP CONSTRAINT "ProjectService_projectId_fkey";

-- DropForeignKey
ALTER TABLE "system"."ProjectSubmission" DROP CONSTRAINT "ProjectSubmission_projectId_fkey";

-- DropForeignKey
ALTER TABLE "system"."ProjectDiscussion" DROP CONSTRAINT "ProjectDiscussion_projectId_fkey";

-- DropForeignKey
ALTER TABLE "system"."ProjectForm" DROP CONSTRAINT "ProjectForm_projectId_fkey";

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

-- CreateTable
CREATE TABLE "Organization" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "path" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "accessType" INTEGER NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "Organization_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Project" (
    "id" TEXT NOT NULL,
    "orgId" TEXT NOT NULL,
    "workspaceId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "path" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "accessType" INTEGER NOT NULL,
    "metadata" JSONB,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "Project_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "ProjectDiscussion" (
    "id" SERIAL NOT NULL,
    "projectId" INTEGER NOT NULL,
    "userId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "metadata" JSONB,
    "data" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "ProjectDiscussion_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "ProjectForm" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "projectId" TEXT NOT NULL,

    CONSTRAINT "ProjectForm_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "ProjectService" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "type" INTEGER NOT NULL,
    "description" TEXT NOT NULL,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "isBuiltin" BOOLEAN NOT NULL DEFAULT false,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "projectId" TEXT NOT NULL,

    CONSTRAINT "ProjectService_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "ProjectSubmission" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "userId" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "ProjectSubmission_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Workspace" (
    "id" TEXT NOT NULL,
    "orgId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "path" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "accessType" INTEGER NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "Workspace_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Organization_path_key" ON "Organization"("path" ASC);

-- CreateIndex
CREATE UNIQUE INDEX "Project_path_key" ON "Project"("path" ASC);

-- CreateIndex
CREATE UNIQUE INDEX "Workspace_orgId_id_key" ON "Workspace"("orgId" ASC, "id" ASC);

-- CreateIndex
CREATE UNIQUE INDEX "Workspace_path_key" ON "Workspace"("path" ASC);

-- AddForeignKey
ALTER TABLE "Project" ADD CONSTRAINT "Project_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES "Organization"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Project" ADD CONSTRAINT "Project_workspaceId_fkey" FOREIGN KEY ("workspaceId") REFERENCES "Workspace"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "ProjectDiscussion" ADD CONSTRAINT "ProjectDiscussion_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES "ProjectSubmission"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "ProjectForm" ADD CONSTRAINT "ProjectForm_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES "Project"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "ProjectService" ADD CONSTRAINT "ProjectService_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES "Project"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "ProjectSubmission" ADD CONSTRAINT "ProjectSubmission_projectId_fkey" FOREIGN KEY ("projectId") REFERENCES "Project"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "Workspace" ADD CONSTRAINT "Workspace_orgId_fkey" FOREIGN KEY ("orgId") REFERENCES "Organization"("id") ON DELETE CASCADE ON UPDATE CASCADE;

