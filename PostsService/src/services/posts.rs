
use crate::{repository, errors::{HandableResult, HandableErrorType, HandableError}, models::post::Post};


pub struct PostsService {
  repository: repository::posts::PostsRepository
}


impl PostsService {

  pub fn New() -> PostsService {
    PostsService {
      repository: repository::posts::PostsRepository::New()
    }
  }

  pub fn GetAll(&self) -> HandableResult<Vec<Post>> {

    match self.repository.GetAll() {
      Some(posts) => Ok(posts),
      None => Err(HandableError { error: HandableErrorType::FatalError, message: "Some fatal error has occured.".to_string() })
    }

  }

  pub fn GetById(&self, id: &String) -> HandableResult<Post> {

    match self.repository.GetById(&id) {
      Some(post) => Ok(post),
      None => Err(HandableError { error: HandableErrorType::FatalError, message: "Some fatal error has occured.".to_string() })
    }

  }

  pub fn GetByIdOrThrow(&self, id: &String) -> HandableResult<Post> {

    match self.repository.GetById(&id) {
      Some(post) => Ok(post),
      None => Err(HandableError {
        message: format!("Post with id {0} does not exist.", id),
        error: HandableErrorType::MissingEntity
      })
    }

  }

  pub fn Create(&self, post: &Post) -> HandableResult<Post> {

    match self.repository.Create(&post) {
      Some(post) => Ok(post),
      None => Err(HandableError { error: HandableErrorType::FatalError, message: "Some fatal error has occured.".to_string() })
    }
  }

  pub fn Update(&self, post: &Post) -> HandableResult<Post> {

    let mut existingPost = match self.GetByIdOrThrow(&post.id) {
      Ok(post) => post,
      Err(err) => return Err(err)
    };

    existingPost.text = post.text.to_string();
    existingPost.likedBy = post.likedBy.clone();
    existingPost.comments = post.comments.clone();

    match self.repository.Update(&existingPost) {
      Some(post) => Ok(post),
      None => Err(HandableError { error: HandableErrorType::FatalError, message: "Some fatal error has occured.".to_string() })
    }
  }

  pub fn Delete(&self, id: &String) -> HandableResult<Post> {

    match self.GetByIdOrThrow(&id) {
      Ok(post) => post,
      Err(err) => return Err(err)
    };

    match self.repository.Delete(&id) {
      Some(post) => Ok(post),
      None => Err(HandableError { error: HandableErrorType::FatalError, message: "Some fatal error has occured.".to_string() })
    }
  }


}