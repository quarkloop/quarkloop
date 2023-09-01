import { PrismaClient, Prisma } from "@prisma/client";

const prisma = new PrismaClient();

const dataToSeed: Prisma.AppCreateArgs[] = [];

async function main() {
  console.log(`[App] Start seeding ...`);

  for (const data of dataToSeed) {
    const record = await prisma.app.create(data);
    console.log(`[App] Created with id: ${record.id}`);
  }

  console.log(`[App] Seeding finished.`);
}

main()
  .then(async () => {
    await prisma.$disconnect();
  })
  .catch(async (e) => {
    console.error(e);
    await prisma.$disconnect();
    process.exit(1);
  });
