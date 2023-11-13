-- CreateTable
CREATE TABLE "system"."OperatingSystem" (
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
CREATE TABLE "system"."Workspace" (
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

-- CreateTable
CREATE TABLE "system"."App" (
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
CREATE TABLE "system"."AppComponent" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "config" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "appId" TEXT,

    CONSTRAINT "AppComponent_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "system"."AppIssue" (
    "id" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "appId" TEXT,

    CONSTRAINT "AppIssue_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "system"."AppSubmission" (
    "id" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "appId" TEXT,

    CONSTRAINT "AppSubmission_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "system"."AppDeployment" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "appId" TEXT,

    CONSTRAINT "AppDeployment_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "OperatingSystem_path_key" ON "system"."OperatingSystem"("path");

-- CreateIndex
CREATE UNIQUE INDEX "Workspace_path_key" ON "system"."Workspace"("path");

-- CreateIndex
CREATE UNIQUE INDEX "Workspace_osId_id_key" ON "system"."Workspace"("osId", "id");

-- CreateIndex
CREATE UNIQUE INDEX "App_path_key" ON "system"."App"("path");

-- AddForeignKey
ALTER TABLE "system"."Workspace" ADD CONSTRAINT "Workspace_osId_fkey" FOREIGN KEY ("osId") REFERENCES "system"."OperatingSystem"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "system"."App" ADD CONSTRAINT "App_osId_fkey" FOREIGN KEY ("osId") REFERENCES "system"."OperatingSystem"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "system"."App" ADD CONSTRAINT "App_workspaceId_fkey" FOREIGN KEY ("workspaceId") REFERENCES "system"."Workspace"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "system"."AppComponent" ADD CONSTRAINT "AppComponent_appId_fkey" FOREIGN KEY ("appId") REFERENCES "system"."App"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "system"."AppIssue" ADD CONSTRAINT "AppIssue_appId_fkey" FOREIGN KEY ("appId") REFERENCES "system"."App"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "system"."AppSubmission" ADD CONSTRAINT "AppSubmission_appId_fkey" FOREIGN KEY ("appId") REFERENCES "system"."App"("id") ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "system"."AppDeployment" ADD CONSTRAINT "AppDeployment_appId_fkey" FOREIGN KEY ("appId") REFERENCES "system"."App"("id") ON DELETE SET NULL ON UPDATE CASCADE;
