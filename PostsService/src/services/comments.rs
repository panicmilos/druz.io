use crate::{models::{comments::Comment, post::Post}, errors::{HandableResult, HandableErrorType, HandableError}};

use super::posts::PostsService;


pub struct CommentsService {
  postsService: PostsService
}

impl CommentsService {

  pub fn New() -> CommentsService {
    CommentsService {
      postsService: PostsService::New()
    }
  }

  pub fn Create(&self, comment: &Comment) -> HandableResult<Comment> {

    let mut post = match self.postsService.GetByIdOrThrow(&comment.postId) {
      Ok(post) => post,
      Err(err) => return Err(err)
    };

    post.comments.push(comment.clone());

    match self.postsService.Update(&post) {
      Ok(_) => Ok(comment.clone()),
      Err(err) => Err(err)
    }

  }

  pub fn Update(&self, comment: &Comment) -> HandableResult<Comment> {

    let mut post = match self.postsService.GetByIdOrThrow(&comment.postId) {
      Ok(post) => post,
      Err(err) => return Err(err)
    };

    let existingComment = match CommentsService::GetByIdOrThrow(&post, &comment.id) {
      Ok(comment) => comment,
      Err(err) => return Err(err)
    };

    post.comments = post.comments.iter().map(|x|
      if x.id != comment.id
      { x.clone() }
      else { Comment { id: x.id.to_string(), createdAt: x.createdAt.to_string(), postId: x.postId.to_string(), userId: x.userId.to_string(), text: comment.text.to_string() }}
    ).collect();

    match self.postsService.Update(&post) {
      Ok(_) => Ok(comment.clone()),
      Err(err) => Err(err)
    }

  }

  pub fn Delete(&self, comment: &Comment) -> HandableResult<Comment> {

    let mut post = match self.postsService.GetByIdOrThrow(&comment.postId) {
      Ok(post) => post,
      Err(err) => return Err(err)
    };

    let existingComment = match CommentsService::GetByIdOrThrow(&post, &comment.id) {
      Ok(comment) => comment,
      Err(err) => return Err(err)
    };

    let index = post.comments.iter().position(|x| x.id == comment.id).unwrap();
    post.comments.remove(index);

    match self.postsService.Update(&post) {
      Ok(_) => Ok(comment.clone()),
      Err(err) => Err(err)
    }

  }

  fn GetByIdOrThrow(post: &Post, commentId: &String) -> HandableResult<Comment> {
    match post.comments.iter().find(|x| x.id == *commentId) {
      Some(comment) => Ok(comment.clone()),
      None => Err(HandableError {
        error: HandableErrorType::MissingEntity,
        message: format!("Comment with id {0} does not exist.", commentId)
      })
    }
  }

  

}