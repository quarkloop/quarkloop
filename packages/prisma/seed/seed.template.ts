import { PrismaClient, Prisma } from "@prisma/client";

const prisma = new PrismaClient();

const dataToSeed: Prisma.AppCreateArgs[] = [];

async function main() {
    console.log(`[Project] Start seeding ...`);

    for (const data of dataToSeed) {
        const record = await prisma.project.create(data);
        console.log(`[Project] Created with id: ${record.id}`);
    }

    console.log(`[Project] Seeding finished.`);
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
