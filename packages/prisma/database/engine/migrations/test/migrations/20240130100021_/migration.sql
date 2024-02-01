-- CreateTable
CREATE TABLE "engine"."TableColumn" (
    "id" BIGSERIAL NOT NULL,
    "tableId" BIGINT NOT NULL,
    "columns" JSONB NOT NULL,

    CONSTRAINT "TableColumn_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "engine"."TableRecord" (
    "id" BIGSERIAL NOT NULL,
    "tableId" BIGINT,
    "data" JSONB NOT NULL,

    CONSTRAINT "TableRecord_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "engine"."Table" (
    "id" BIGSERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "createdBy" TEXT NOT NULL,
    "updatedAt" TIMESTAMP(3),
    "updatedBy" TEXT,

    CONSTRAINT "Table_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "TableColumn_tableId_key" ON "engine"."TableColumn"("tableId");

-- AddForeignKey
ALTER TABLE "engine"."TableColumn" ADD CONSTRAINT "TableColumn_tableId_fkey" FOREIGN KEY ("tableId") REFERENCES "engine"."Table"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "engine"."TableRecord" ADD CONSTRAINT "TableRecord_tableId_fkey" FOREIGN KEY ("tableId") REFERENCES "engine"."Table"("id") ON DELETE SET NULL ON UPDATE CASCADE;
