export type Post = {

  id: string,
  createdAt: string,
  text: string,
  writtenBy: string

  likedBy: string[],
  comments: Comment[]
}

export type Comment = {
  id: string,
  createdAt: string,
  postId: string,
  userId: string,
  text: string
}