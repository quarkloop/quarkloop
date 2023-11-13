-- CreateTable
CREATE TABLE "app"."SheetInstance" (
    "id" SERIAL NOT NULL,
    "appId" TEXT NOT NULL,
    "dataId" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    "rowCount" INTEGER NOT NULL,
    "rows" JSONB[],
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "SheetInstance_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "app"."SheetComponent" (
    "id" SERIAL NOT NULL,
    "appId" TEXT NOT NULL,
    "type" INTEGER NOT NULL,
    "settings" JSONB NOT NULL,
    "createdAt" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3),

    CONSTRAINT "SheetComponent_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "SheetInstance_dataId_key" ON "app"."SheetInstance"("dataId");

-- CreateIndex
CREATE INDEX "SheetInstance_appId_idx" ON "app"."SheetInstance"("appId");

-- CreateIndex
CREATE INDEX "SheetComponent_appId_idx" ON "app"."SheetComponent"("appId");
