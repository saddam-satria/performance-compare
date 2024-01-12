import { Injectable, HttpStatus } from '@nestjs/common';
import prismaClient from 'src/config/prisma';

@Injectable()
class PostService {
  async getPosts() {
    const posts: any =
      await prismaClient.$queryRaw`SELECT "public"."Post"."postId", "public"."Post"."title", "public"."Post"."body", "public"."Post"."createdAt", "public"."Post"."updatedAt", "public"."Author"."name" AS author_name FROM "public"."Post" INNER JOIN "public"."Author" ON "public"."Author"."authorId" = "public"."Post"."author_id"`;

    return {
      total: posts.length,
      message: 'welcome to nestjs - by Saddam',
      statusCode: HttpStatus.OK,
      data: {
        posts,
      },
    };
  }
}

export default PostService;
