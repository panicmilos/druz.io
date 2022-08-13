use serde_derive::{Serialize, Deserialize};

use crate::{models::post::Post, errors::{HandableResult, HandableError, HandableErrorType}};



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