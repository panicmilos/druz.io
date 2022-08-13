
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
      None => Err(HandableError {
        message: "Error while fetching posts".to_string(),
        error: HandableErrorType::BadLogic
      })
    }

  }

  pub fn Create(&self, post: &Post) -> HandableResult<Post> {

    match self.repository.Create(&post) {
      Some(posts) => Ok(posts),
      None => Err(HandableError {
        message: "Error while creating a new post".to_string(),
        error: HandableErrorType::BadLogic
      })
    }
  }
}