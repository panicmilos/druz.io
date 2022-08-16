export type Post = {
  id: string;
  createdAt: string;
  text: string;
  writtenBy: string;
  likedBy: string[];
  comments: any;
}