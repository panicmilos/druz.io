import axios from "axios";
import { useEffect, useState } from "react";
import { POSTS_SERVICE_URL } from "../imports";
import { Like } from "../models/Post";

export const useLikesService = (postId: string) => {

  const [likesService, setLikesService] = useState(new LikesService(postId));

  useEffect(() => {
    setLikesService(new LikesService(postId));
  }, [postId]);

  return likesService;
}


export class LikesService {

  public ID: string;
  private baseUrl: string;

  constructor(postId: string) {
    this.ID = 'CommentsService';
    this.baseUrl = `${POSTS_SERVICE_URL}/posts/${postId}`;
  }

  public async like(): Promise<Like> {
    return (await axios.post(`${this.baseUrl}/like`)).data;
  }

  public async dislike(): Promise<Like> {
    return (await axios.post(`${this.baseUrl}/dislike`)).data;
  }

}