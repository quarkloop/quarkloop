-- CreateTable
CREATE TABLE "project"."TableBranch" (
    "id" BIGSERIAL NOT NULL,
    "projectId" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "default" BOOLEAN NOT NULL DEFAULT false,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "createdBy" TEXT NOT NULL,
    "updatedAt" TIMESTAMP(3),
    "updatedBy" TEXT,

    CONSTRAINT "TableBranch_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "project"."TableSchema" (
    "id" BIGSERIAL NOT NULL,
    "projectId" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "createdBy" TEXT NOT NULL,
    "updatedAt" TIMESTAMP(3),
    "updatedBy" TEXT,

    CONSTRAINT "TableSchema_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "project"."TableMain" (
    "id" BIGSERIAL NOT NULL,
    "projectId" BIGINT NOT NULL,
    "branchId" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "createdBy" TEXT NOT NULL,
    "updatedAt" TIMESTAMP(3),
    "updatedBy" TEXT,

    CONSTRAINT "TableMain_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "project"."TableDocument" (
    "id" BIGSERIAL NOT NULL,
    "projectId" BIGINT NOT NULL,
    "branchId" BIGINT NOT NULL,
    "mainId" BIGINT NOT NULL,
    "schemaId" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "createdBy" TEXT NOT NULL,
    "updatedAt" TIMESTAMP(3),
    "updatedBy" TEXT,

    CONSTRAINT "TableDocument_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "project"."TablePayment" (
    "id" BIGSERIAL NOT NULL,
    "projectId" BIGINT NOT NULL,
    "branchId" BIGINT NOT NULL,
    "mainId" BIGINT NOT NULL,
    "schemaId" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "createdBy" TEXT NOT NULL,
    "updatedAt" TIMESTAMP(3),
    "updatedBy" TEXT,

    CONSTRAINT "TablePayment_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "project"."TableForm" (
    "id" BIGSERIAL NOT NULL,
    "projectId" BIGINT NOT NULL,
    "branchId" BIGINT NOT NULL,
    "mainId" BIGINT NOT NULL,
    "schemaId" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "createdBy" TEXT NOT NULL,
    "updatedAt" TIMESTAMP(3),
    "updatedBy" TEXT,

    CONSTRAINT "TableForm_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "project"."App" (
    "id" BIGSERIAL NOT NULL,
    "projectId" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "createdBy" TEXT NOT NULL,
    "updatedAt" TIMESTAMP(3),
    "updatedBy" TEXT,

    CONSTRAINT "App_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "project"."AppSubmission" (
    "id" BIGSERIAL NOT NULL,
    "projectId" BIGINT NOT NULL,
    "title" TEXT NOT NULL,
    "status" INTEGER NOT NULL DEFAULT 1,
    "labels" JSONB,
    "dueDate" TIMESTAMP(3),
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "createdBy" TEXT NOT NULL,
    "updatedAt" TIMESTAMP(3),
    "updatedBy" TEXT,

    CONSTRAINT "AppSubmission_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "project"."AppDiscussion" (
    "id" BIGSERIAL NOT NULL,
    "submissionId" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT,
    "metadata" JSONB,
    "data" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "createdBy" TEXT NOT NULL,
    "updatedAt" TIMESTAMP(3),
    "updatedBy" TEXT,

    CONSTRAINT "AppDiscussion_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "TableBranch_projectId_idx" ON "project"."TableBranch"("projectId");

-- CreateIndex
CREATE UNIQUE INDEX "TableBranch_projectId_name_key" ON "project"."TableBranch"("projectId", "name");

-- CreateIndex
CREATE INDEX "TableSchema_projectId_idx" ON "project"."TableSchema"("projectId");

-- CreateIndex
CREATE INDEX "TableMain_projectId_branchId_idx" ON "project"."TableMain"("projectId", "branchId");

-- CreateIndex
CREATE INDEX "TableDocument_projectId_branchId_idx" ON "project"."TableDocument"("projectId", "branchId");

-- CreateIndex
CREATE INDEX "TableDocument_projectId_branchId_mainId_idx" ON "project"."TableDocument"("projectId", "branchId", "mainId");

-- CreateIndex
CREATE INDEX "TablePayment_projectId_branchId_idx" ON "project"."TablePayment"("projectId", "branchId");

-- CreateIndex
CREATE INDEX "TablePayment_projectId_branchId_mainId_idx" ON "project"."TablePayment"("projectId", "branchId", "mainId");

-- CreateIndex
CREATE INDEX "TableForm_projectId_branchId_idx" ON "project"."TableForm"("projectId", "branchId");

-- CreateIndex
CREATE INDEX "TableForm_projectId_branchId_mainId_idx" ON "project"."TableForm"("projectId", "branchId", "mainId");

-- AddForeignKey
ALTER TABLE "project"."TableMain" ADD CONSTRAINT "TableMain_branchId_fkey" FOREIGN KEY ("branchId") REFERENCES "project"."TableBranch"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project"."TableDocument" ADD CONSTRAINT "TableDocument_mainId_fkey" FOREIGN KEY ("mainId") REFERENCES "project"."TableMain"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project"."TableDocument" ADD CONSTRAINT "TableDocument_branchId_fkey" FOREIGN KEY ("branchId") REFERENCES "project"."TableBranch"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project"."TableDocument" ADD CONSTRAINT "TableDocument_schemaId_fkey" FOREIGN KEY ("schemaId") REFERENCES "project"."TableSchema"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project"."TablePayment" ADD CONSTRAINT "TablePayment_mainId_fkey" FOREIGN KEY ("mainId") REFERENCES "project"."TableMain"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project"."TablePayment" ADD CONSTRAINT "TablePayment_branchId_fkey" FOREIGN KEY ("branchId") REFERENCES "project"."TableBranch"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project"."TablePayment" ADD CONSTRAINT "TablePayment_schemaId_fkey" FOREIGN KEY ("schemaId") REFERENCES "project"."TableSchema"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project"."TableForm" ADD CONSTRAINT "TableForm_mainId_fkey" FOREIGN KEY ("mainId") REFERENCES "project"."TableMain"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project"."TableForm" ADD CONSTRAINT "TableForm_branchId_fkey" FOREIGN KEY ("branchId") REFERENCES "project"."TableBranch"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project"."TableForm" ADD CONSTRAINT "TableForm_schemaId_fkey" FOREIGN KEY ("schemaId") REFERENCES "project"."TableSchema"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "project"."AppDiscussion" ADD CONSTRAINT "AppDiscussion_submissionId_fkey" FOREIGN KEY ("submissionId") REFERENCES "project"."AppSubmission"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
