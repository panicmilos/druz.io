import { Controller, Delete, Get, Path, Post, Put } from "nimbly-api";
import { AppContainer } from "../AppContainer";
import { IController } from "../contracts/IController";
import { PostsService } from "../services/PostsService";


@Controller('/posts')
export class PostsController implements IController {

  private postsService: PostsService;

  constructor() {
    this.postsService = AppContainer.get("PostsService") as PostsService;
  }

  @Get('')
  async getPosts(): Promise<any> {
    return this.postsService.get();
  }

  @Get('/:id')
  async getPost(@Path('id') id: string): Promise<any> {
    return this.postsService.getByIdOrThrow(id);
  }

  @Post('')
  async createPost(post: any): Promise<any> {
    return this.postsService.add(post as Post);
  }

  @Put('/:id')
  async updatePost(@Path('id') id: string, post: any): Promise<any> {
    return this.postsService.update(id, post);
  }

  @Delete('/:id')
  async deletePost(@Path('id') id: string): Promise<any> {
    return this.postsService.delete(id);
  }
}