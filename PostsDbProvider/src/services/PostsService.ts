import { inject, injectable } from "inversify";
import DocumentStore from "ravendb";
import { MissingEntityError } from "../errors/MissingEntityError";
import { Post } from "../models/Post";
import { v4 as uuidv4 } from 'uuid';

const COLLECTION_NAME = "Posts";

@injectable()
export class PostsService {

  constructor(@inject("DocumentStore") private store: DocumentStore) {
    
  }

  async get(): Promise<Post[]> {

    const session = this.store.openSession();
    const posts = await session
      .query({ collection: COLLECTION_NAME })
      .all() as Post[];

    return posts;

  }

  async getById(id: string): Promise<Post> {

    const session = this.store.openSession();
    const post = await session.load(id) as Post;

    return post;
  }

  async getByIdOrThrow(id: string): Promise<Post> {

    const session = this.store.openSession();
    const post = await session.load(id) as Post;
    if (!post) throw new MissingEntityError(`Post with id ${id} does not exist.`);

    return post;
  }

  async add(post: Post): Promise<Post> {
    delete post?.id;
    !post.createdAt && (post.createdAt = new Date().toISOString());

    const session = this.store.openSession();
    await session.store(post);
    const metadata = session.advanced.getMetadataFor(post);
    metadata["@collection"] = COLLECTION_NAME;
    await session.saveChanges();

    return post;
  }

  async update(id: string, post: Post): Promise<Post> {

    const session = this.store.openSession();
    const existingPost = await session.load(id) as Post;
    existingPost.text = post.text;
    existingPost.likedBy = post.likedBy;
    existingPost.comments = post.comments;
    existingPost.comments.forEach(comment => {
      !comment.id && (comment.id = uuidv4().toString());
      !comment.createdAt && (comment.createdAt = new Date().toISOString());
    });
    
    await session.saveChanges();

    return existingPost;
  }

  async delete(id: string): Promise<Post> {

    const post = await this.getById(id);
    const session = this.store.openSession();
    await session.delete(id);
    await session.saveChanges();

    return post;
  }

}