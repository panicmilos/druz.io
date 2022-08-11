import { Controller, Get, Post } from "nimbly-api";
import { IController } from "../contracts/IController";


@Controller('/posts')
export class PostsController implements IController {

  constructor() {
  }

  @Get('')
  async getPosts(): Promise<any> {
    return [{ "posts": "CAO"}];
  }

  @Post('')
  async createPost(post: any): Promise<any> {
    console.log(post)
    return post;
  }

}