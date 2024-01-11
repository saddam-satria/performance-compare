import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { PORT } from './config/constant';
import prismaClient from './config/prisma';
import { NextFunction, Request, Response } from 'express';
import * as pidusage from 'pidusage';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.use(async (_req: Request, _res: Response, next: NextFunction) => {
    const stats = await pidusage(process.pid);
    console.log(
      `CPU Usage ${stats.cpu.toFixed(2)}% - Memory Usage ${Math.round(
        stats.memory / (1024 * 1024),
      )} MB`,
    );
    next();
  });

  await prismaClient.$connect();
  await app.listen(PORT);
}
bootstrap();
