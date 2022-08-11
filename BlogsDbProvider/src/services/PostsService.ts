import { inject, injectable } from "inversify";
import DocumentStore from "ravendb";
import { Post } from "../models/Post";

const COLLECTION_NAME = "Posts";

@injectable()
export class PostsService {

  constructor(@inject("DocumentStore") private store: DocumentStore) {
    
  }

  async get(): Promise<Post[]> {
    const session = this.store.openSession();

    return session
      .query({ collection: COLLECTION_NAME })
      .all() as Promise<Post[]>; 
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

}