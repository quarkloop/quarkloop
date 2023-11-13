/*
  Warnings:

  - You are about to drop the column `dataId` on the `SheetInstance` table. All the data in the column will be lost.

*/
-- DropIndex
DROP INDEX "app"."SheetInstance_dataId_key";

-- AlterTable
ALTER TABLE "app"."SheetInstance" DROP COLUMN "dataId";
