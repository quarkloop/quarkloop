-- DropTable
DROP TABLE "app"."SheetInstance";

-- DropTable
DROP TABLE "app"."SheetComponent";

-- CreateTable
CREATE TABLE "SheetComponent" (
    "id" SERIAL NOT NULL,
    "appId" TEXT NOT NULL,
    "type" INTEGER NOT NULL,
    "settings" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "SheetComponent_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "SheetInstance" (
    "id" SERIAL NOT NULL,
    "appId" TEXT NOT NULL,
    "dataId" INTEGER NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),
    "name" TEXT NOT NULL,
    "rowCount" INTEGER NOT NULL,
    "rows" JSONB[],

    CONSTRAINT "SheetInstance_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE INDEX "SheetComponent_appId_idx" ON "SheetComponent"("appId" ASC);

-- CreateIndex
CREATE INDEX "SheetInstance_appId_idx" ON "SheetInstance"("appId" ASC);

-- CreateIndex
CREATE UNIQUE INDEX "SheetInstance_dataId_key" ON "SheetInstance"("dataId" ASC);

