import { useState } from "react";
import { CrudService, POSTS_SERVICE_URL } from "../imports";
import { Post } from "../models/Post";


export const usePostsService = () => {

  const [postsService] = useState(new PostsService());
  
  return postsService;
}

export class PostsService extends CrudService<Post> {


  constructor() {
    super('PostsService', `${POSTS_SERVICE_URL}/posts`);
  }

}