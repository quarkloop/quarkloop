-- DropForeignKey
ALTER TABLE "app"."SheetInstance"
DROP CONSTRAINT "SheetInstance_dataId_fkey";

-- DropForeignKey
ALTER TABLE "app"."SheetComponent"
DROP CONSTRAINT "SheetComponent_instanceId_fkey";

-- DropTable
DROP TABLE "app"."SheetInstance";

-- DropTable
DROP TABLE "app"."SheetData";

-- DropTable
DROP TABLE "app"."SheetComponent";

-- DropSchema
DROP SCHEMA IF EXISTS "app";