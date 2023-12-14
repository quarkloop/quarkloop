-- DropForeignKey
ALTER TABLE "project"."TableMain"
DROP CONSTRAINT "TableMain_branchId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableDocument"
DROP CONSTRAINT "TableDocument_mainId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableDocument"
DROP CONSTRAINT "TableDocument_branchId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableDocument"
DROP CONSTRAINT "TableDocument_schemaId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TablePayment"
DROP CONSTRAINT "TablePayment_mainId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TablePayment"
DROP CONSTRAINT "TablePayment_branchId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TablePayment"
DROP CONSTRAINT "TablePayment_schemaId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableForm"
DROP CONSTRAINT "TableForm_mainId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableForm"
DROP CONSTRAINT "TableForm_branchId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableForm"
DROP CONSTRAINT "TableForm_schemaId_fkey";

-- DropForeignKey
ALTER TABLE "project"."AppDiscussion"
DROP CONSTRAINT "AppDiscussion_submissionId_fkey";

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

-- DropSchema
DROP SCHEMA IF EXISTS "project";