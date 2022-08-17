import { useEffect, useState } from "react";
import { CrudService, POSTS_SERVICE_URL } from "../imports";
import { Comment } from "../models/Post";


export const useCommentsService = (postId: string) => {

  const [commentsService, setCommentsService] = useState(new CommentsService(postId));

  useEffect(() => {
    setCommentsService(new CommentsService(postId));
  }, [postId]);

  return commentsService;
}

export class CommentsService extends CrudService<Comment> {


  constructor(postId: string) {
    super('CommentsService', `${POSTS_SERVICE_URL}/posts/${postId}/comments`);
  }

}