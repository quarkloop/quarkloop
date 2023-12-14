-- DropForeignKey
ALTER TABLE "project"."TableMain" DROP CONSTRAINT "TableMain_branchId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableDocument" DROP CONSTRAINT "TableDocument_mainId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableDocument" DROP CONSTRAINT "TableDocument_branchId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableDocument" DROP CONSTRAINT "TableDocument_schemaId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TablePayment" DROP CONSTRAINT "TablePayment_mainId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TablePayment" DROP CONSTRAINT "TablePayment_branchId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TablePayment" DROP CONSTRAINT "TablePayment_schemaId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableForm" DROP CONSTRAINT "TableForm_mainId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableForm" DROP CONSTRAINT "TableForm_branchId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableForm" DROP CONSTRAINT "TableForm_schemaId_fkey";

-- DropForeignKey
ALTER TABLE "project"."AppDiscussion" DROP CONSTRAINT "AppDiscussion_submissionId_fkey";

-- DropTable
DROP TABLE "project"."TableBranch";

-- DropTable
DROP TABLE "project"."TableSchema";

-- DropTable
DROP TABLE "project"."TableMain";

-- DropTable
DROP TABLE "project"."TableDocument";

-- DropTable
DROP TABLE "project"."TablePayment";

-- DropTable
DROP TABLE "project"."TableForm";

-- DropTable
DROP TABLE "project"."App";

-- DropTable
DROP TABLE "project"."AppSubmission";

-- DropTable
DROP TABLE "project"."AppDiscussion";

-- CreateTable
CREATE TABLE "App" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "App_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "AppDiscussion" (
    "id" SERIAL NOT NULL,
    "submissionId" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "AppDiscussion_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "AppSubmission" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "status" INTEGER NOT NULL DEFAULT 1,
    "labels" JSONB,
    "dueDate" TIMESTAMP(3),
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "AppSubmission_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TableBranch" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "default" BOOLEAN NOT NULL DEFAULT false,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "TableBranch_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TableDocument" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "branchId" INTEGER NOT NULL,
    "mainId" INTEGER NOT NULL,
    "schemaId" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "TableDocument_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TableForm" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "branchId" INTEGER NOT NULL,
    "mainId" INTEGER NOT NULL,
    "schemaId" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "TableForm_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TableMain" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "branchId" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "TableMain_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TablePayment" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "branchId" INTEGER NOT NULL,
    "mainId" INTEGER NOT NULL,
    "schemaId" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "TablePayment_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TableSchema" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "TableSchema_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "TableBranch_projectId_idx" ON "TableBranch"("projectId" ASC);

-- CreateIndex
CREATE UNIQUE INDEX "TableBranch_projectId_name_key" ON "TableBranch"("projectId" ASC, "name" ASC);

-- CreateIndex
CREATE INDEX "TableDocument_projectId_branchId_idx" ON "TableDocument"("projectId" ASC, "branchId" ASC);

-- CreateIndex
CREATE INDEX "TableDocument_projectId_branchId_mainId_idx" ON "TableDocument"("projectId" ASC, "branchId" ASC, "mainId" ASC);

-- CreateIndex
CREATE INDEX "TableForm_projectId_branchId_idx" ON "TableForm"("projectId" ASC, "branchId" ASC);

-- CreateIndex
CREATE INDEX "TableForm_projectId_branchId_mainId_idx" ON "TableForm"("projectId" ASC, "branchId" ASC, "mainId" ASC);

-- CreateIndex
CREATE INDEX "TableMain_projectId_branchId_idx" ON "TableMain"("projectId" ASC, "branchId" ASC);

-- CreateIndex
CREATE INDEX "TablePayment_projectId_branchId_idx" ON "TablePayment"("projectId" ASC, "branchId" ASC);

-- CreateIndex
CREATE INDEX "TablePayment_projectId_branchId_mainId_idx" ON "TablePayment"("projectId" ASC, "branchId" ASC, "mainId" ASC);

-- CreateIndex
CREATE INDEX "TableSchema_projectId_idx" ON "TableSchema"("projectId" ASC);

-- AddForeignKey
ALTER TABLE "AppDiscussion" ADD CONSTRAINT "AppDiscussion_submissionId_fkey" FOREIGN KEY ("submissionId") REFERENCES "AppSubmission"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TableDocument" ADD CONSTRAINT "TableDocument_branchId_fkey" FOREIGN KEY ("branchId") REFERENCES "TableBranch"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TableDocument" ADD CONSTRAINT "TableDocument_mainId_fkey" FOREIGN KEY ("mainId") REFERENCES "TableMain"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TableDocument" ADD CONSTRAINT "TableDocument_schemaId_fkey" FOREIGN KEY ("schemaId") REFERENCES "TableSchema"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TableForm" ADD CONSTRAINT "TableForm_branchId_fkey" FOREIGN KEY ("branchId") REFERENCES "TableBranch"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TableForm" ADD CONSTRAINT "TableForm_mainId_fkey" FOREIGN KEY ("mainId") REFERENCES "TableMain"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TableForm" ADD CONSTRAINT "TableForm_schemaId_fkey" FOREIGN KEY ("schemaId") REFERENCES "TableSchema"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TableMain" ADD CONSTRAINT "TableMain_branchId_fkey" FOREIGN KEY ("branchId") REFERENCES "TableBranch"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TablePayment" ADD CONSTRAINT "TablePayment_branchId_fkey" FOREIGN KEY ("branchId") REFERENCES "TableBranch"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TablePayment" ADD CONSTRAINT "TablePayment_mainId_fkey" FOREIGN KEY ("mainId") REFERENCES "TableMain"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TablePayment" ADD CONSTRAINT "TablePayment_schemaId_fkey" FOREIGN KEY ("schemaId") REFERENCES "TableSchema"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

