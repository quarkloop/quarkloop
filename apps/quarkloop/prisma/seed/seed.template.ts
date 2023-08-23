import { PrismaClient, Prisma } from "@prisma/client";

const prisma = new PrismaClient();

const dataToSeed: Prisma.UserCreateArgs[] = [];

async function main() {
  console.log(`[User] Start seeding ...`);

  for (const data of dataToSeed) {
    const record = await prisma.user.create(data);
    console.log(`[User] Created with id: ${record.id}`);
  }

  console.log(`[User] Seeding finished.`);
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
