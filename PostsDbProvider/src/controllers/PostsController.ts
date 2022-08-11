import { Controller, Get, Post } from "nimbly-api";
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

  @Post('')
  async createPost(post: any): Promise<any> {
    return this.postsService.add(post);
  }

}