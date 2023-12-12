-- DropForeignKey
ALTER TABLE "project"."TableDocument" DROP CONSTRAINT "TableDocument_mainTableId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableDocument" DROP CONSTRAINT "TableDocument_schemaId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TablePayment" DROP CONSTRAINT "TablePayment_mainTableId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TablePayment" DROP CONSTRAINT "TablePayment_schemaId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableForm" DROP CONSTRAINT "TableForm_mainTableId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableForm" DROP CONSTRAINT "TableForm_schemaId_fkey";

-- DropForeignKey
ALTER TABLE "project"."AppDiscussion" DROP CONSTRAINT "AppDiscussion_submissionId_fkey";

-- DropTable
DROP TABLE "project"."Table";

-- DropTable
DROP TABLE "project"."TableDocumentSchema";

-- DropTable
DROP TABLE "project"."TableDocument";

-- DropTable
DROP TABLE "project"."TablePaymentSchema";

-- DropTable
DROP TABLE "project"."TablePayment";

-- DropTable
DROP TABLE "project"."TableFormSchema";

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
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
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
    "metadata" JSONB,
    "data" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
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
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "AppSubmission_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Table" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "type" INTEGER NOT NULL,
    "description" TEXT NOT NULL,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "Table_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TableDocument" (
    "id" SERIAL NOT NULL,
    "mainTableId" INTEGER NOT NULL,
    "schemaId" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "TableDocument_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TableDocumentSchema" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "TableDocumentSchema_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TableForm" (
    "id" SERIAL NOT NULL,
    "mainTableId" INTEGER NOT NULL,
    "schemaId" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "TableForm_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TableFormSchema" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "TableFormSchema_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TablePayment" (
    "id" SERIAL NOT NULL,
    "mainTableId" INTEGER NOT NULL,
    "schemaId" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "TablePayment_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "TablePaymentSchema" (
    "id" SERIAL NOT NULL,
    "projectId" TEXT NOT NULL,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "createdBy" TEXT NOT NULL,
    "updatedBy" TEXT,

    CONSTRAINT "TablePaymentSchema_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "App_createdBy_updatedBy_idx" ON "App"("createdBy" ASC, "updatedBy" ASC);

-- CreateIndex
CREATE INDEX "AppDiscussion_createdBy_updatedBy_idx" ON "AppDiscussion"("createdBy" ASC, "updatedBy" ASC);

-- CreateIndex
CREATE INDEX "AppSubmission_createdBy_updatedBy_idx" ON "AppSubmission"("createdBy" ASC, "updatedBy" ASC);

-- CreateIndex
CREATE INDEX "Table_createdBy_updatedBy_idx" ON "Table"("createdBy" ASC, "updatedBy" ASC);

-- CreateIndex
CREATE INDEX "Table_projectId_idx" ON "Table"("projectId" ASC);

-- CreateIndex
CREATE INDEX "TableDocument_createdBy_updatedBy_idx" ON "TableDocument"("createdBy" ASC, "updatedBy" ASC);

-- CreateIndex
CREATE INDEX "TableDocumentSchema_createdBy_updatedBy_idx" ON "TableDocumentSchema"("createdBy" ASC, "updatedBy" ASC);

-- CreateIndex
CREATE INDEX "TableDocumentSchema_projectId_idx" ON "TableDocumentSchema"("projectId" ASC);

-- CreateIndex
CREATE INDEX "TableForm_createdBy_updatedBy_idx" ON "TableForm"("createdBy" ASC, "updatedBy" ASC);

-- CreateIndex
CREATE INDEX "TableFormSchema_createdBy_updatedBy_idx" ON "TableFormSchema"("createdBy" ASC, "updatedBy" ASC);

-- CreateIndex
CREATE INDEX "TableFormSchema_projectId_idx" ON "TableFormSchema"("projectId" ASC);

-- CreateIndex
CREATE INDEX "TablePayment_createdBy_updatedBy_idx" ON "TablePayment"("createdBy" ASC, "updatedBy" ASC);

-- CreateIndex
CREATE INDEX "TablePaymentSchema_createdBy_updatedBy_idx" ON "TablePaymentSchema"("createdBy" ASC, "updatedBy" ASC);

-- CreateIndex
CREATE INDEX "TablePaymentSchema_projectId_idx" ON "TablePaymentSchema"("projectId" ASC);

-- AddForeignKey
ALTER TABLE "AppDiscussion" ADD CONSTRAINT "AppDiscussion_submissionId_fkey" FOREIGN KEY ("submissionId") REFERENCES "AppSubmission"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TableDocument" ADD CONSTRAINT "TableDocument_mainTableId_fkey" FOREIGN KEY ("mainTableId") REFERENCES "Table"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TableDocument" ADD CONSTRAINT "TableDocument_schemaId_fkey" FOREIGN KEY ("schemaId") REFERENCES "TableDocumentSchema"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TableForm" ADD CONSTRAINT "TableForm_mainTableId_fkey" FOREIGN KEY ("mainTableId") REFERENCES "Table"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TableForm" ADD CONSTRAINT "TableForm_schemaId_fkey" FOREIGN KEY ("schemaId") REFERENCES "TableFormSchema"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TablePayment" ADD CONSTRAINT "TablePayment_mainTableId_fkey" FOREIGN KEY ("mainTableId") REFERENCES "Table"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "TablePayment" ADD CONSTRAINT "TablePayment_schemaId_fkey" FOREIGN KEY ("schemaId") REFERENCES "TablePaymentSchema"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

