import { Module } from '@nestjs/common';
import PostController from 'src/controllers/PostController';
import PostService from 'src/services/PostService';

@Module({
  controllers: [PostController],
  providers: [PostService],
  exports: [PostModule],
})
class PostModule {}

export default PostModule;
