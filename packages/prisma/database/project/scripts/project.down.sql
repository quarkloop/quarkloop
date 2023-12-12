-- DropForeignKey
ALTER TABLE "project"."TableDocument"
DROP CONSTRAINT "TableDocument_mainTableId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableDocument"
DROP CONSTRAINT "TableDocument_schemaId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TablePayment"
DROP CONSTRAINT "TablePayment_mainTableId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TablePayment"
DROP CONSTRAINT "TablePayment_schemaId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableForm"
DROP CONSTRAINT "TableForm_mainTableId_fkey";

-- DropForeignKey
ALTER TABLE "project"."TableForm"
DROP CONSTRAINT "TableForm_schemaId_fkey";

-- DropForeignKey
ALTER TABLE "project"."AppDiscussion"
DROP CONSTRAINT "AppDiscussion_submissionId_fkey";

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

-- DropSchema
DROP SCHEMA IF EXISTS "project";