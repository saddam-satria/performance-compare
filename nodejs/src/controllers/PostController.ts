import { Controller, Get, HttpCode, HttpStatus } from '@nestjs/common';
import PostService from 'src/services/PostService';

@Controller('/api/v1')
class PostController {
  constructor(private postService: PostService) {}
  @Get('posts')
  @HttpCode(HttpStatus.OK)
  async getPost() {
    return this.postService.getPosts();
  }
}

export default PostController;
