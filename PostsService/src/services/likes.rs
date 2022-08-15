use crate::{models::{likes::Like, post::Post}, errors::{HandableResult, HandableError, HandableErrorType}};

use super::posts::PostsService;

pub struct LikesService {
  postsService: PostsService
}


impl LikesService {

  pub fn New() -> LikesService {
    LikesService {
      postsService: PostsService::New()
    }
  }

  pub fn Like(&self, like: &Like) -> HandableResult<Like> {

    let mut post = match self.postsService.GetByIdOrThrow(&like.postId) {
      Ok(post) => post,
      Err(err) => return Err(err)
    };

    if LikesService::hasLikedPost(&post, &like.userId) {
      return Err(HandableError {
        error: HandableErrorType::BadLogic,
        message: "You have already liked the post.".to_string()
      });
    };

    post.likedBy.push(like.userId.to_string());

    match self.postsService.Update(&post) {
      Ok(_) => Ok(like.clone()),
      Err(err) => Err(err)
    }

  }

  pub fn Dislike(&self, like: &Like) -> HandableResult<Like> {

    let mut post = match self.postsService.GetByIdOrThrow(&like.postId) {
      Ok(post) => post,
      Err(err) => return Err(err)
    };

    if !LikesService::hasLikedPost(&post, &like.userId) {
      return Err(HandableError {
        error: HandableErrorType::BadLogic,
        message: "You didn't liked the post in the first place.".to_string()
      });
    };

    let index = post.likedBy.iter().position(|x| *x == like.userId).unwrap();
    post.likedBy.remove(index);

    match self.postsService.Update(&post) {
      Ok(_) => Ok(like.clone()),
      Err(err) => Err(err)
    }

  }

  fn hasLikedPost(post: &Post, userId: &String) -> bool {
    post.likedBy.contains(&userId)
  }

}