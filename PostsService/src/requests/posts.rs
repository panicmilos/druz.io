use serde_derive::{Serialize, Deserialize};

use crate::{models::{post::{Post}, comments::Comment}, errors::{HandableResult, HandableError, HandableErrorType}};



#[derive(Serialize, Clone, Deserialize)]
pub struct CreatePostRequest {
  pub text: String
}

impl CreatePostRequest {
  pub fn to_post(&self) -> Post {
    Post {
      id: "".to_string(),
      createdAt: "".to_string(),
      text: self.text.clone(),
      writtenBy: "".to_string(),
      likedBy: Vec::<String>::with_capacity(0),
      comments: Vec::<Comment>::with_capacity(0)
    }
  }

  pub fn validate(&self) -> HandableResult<()> {

    if self.text.len() < 5 {
      return Err(HandableError {
        message: "Text should have at least 5 charactes.".to_string(),
        error: HandableErrorType::BadLogic
      })
    }

    return Ok(());
  }
}

#[derive(Serialize, Clone, Deserialize)]
pub struct UpdatePostRequest {
  pub text: String
}

impl UpdatePostRequest {
  pub fn to_post(&self) -> Post {
    Post {
      id: "".to_string(),
      createdAt: "".to_string(),
      text: self.text.clone(),
      writtenBy: "".to_string(),
      likedBy: Vec::<String>::with_capacity(0),
      comments: Vec::<Comment>::with_capacity(0)
    }
  }

  pub fn validate(&self) -> HandableResult<()> {

    if self.text.len() < 5 {
      return Err(HandableError {
        message: "Text should have at least 5 charactes.".to_string(),
        error: HandableErrorType::BadLogic
      })
    }

    return Ok(());
  }
}