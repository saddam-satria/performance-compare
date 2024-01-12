import { PrismaClient } from '@prisma/client';

(async () => {
  const prismaClient = new PrismaClient();

  const author = await prismaClient.author.create({
    data: {
      name: 'test',
    },
  });

  for (let index = 0; index < 1000; index++) {
    await prismaClient.post.create({
      data: {
        body: 'Test',
        title: 'test',
        postOnCategory: {
          create: {
            category: {
              connectOrCreate: {
                create: {
                  name: 'test',
                },
                where: {
                  name: 'test',
                },
              },
            },
          },
        },
        author: {
          connect: {
            authorId: author.authorId,
          },
        },
      },
    });
  }
})();
