/*
  Warnings:

  - Added the required column `accessType` to the `App` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "system"."App" ADD COLUMN     "accessType" INTEGER NOT NULL;
