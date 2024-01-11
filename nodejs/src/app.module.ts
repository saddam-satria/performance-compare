import { Module } from '@nestjs/common';
import PostModule from './modules/PostModule';

@Module({
  imports: [PostModule],
  controllers: [],
  providers: [],
})
export class AppModule {}
