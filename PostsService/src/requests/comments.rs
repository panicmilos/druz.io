use serde_derive::{Serialize, Deserialize};

use crate::{models::comments::Comment, errors::{HandableResult, HandableError, HandableErrorType}};



#[derive(Serialize, Clone, Deserialize)]
pub struct CreateCommentRequest {
  pub text: String
}

impl CreateCommentRequest {
  pub fn to_comment(&self) -> Comment {
    Comment {
      id: "".to_string(),
      createdAt: "".to_string(),
      postId: "".to_string(),
      userId: "".to_string(),
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

#[derive(Serialize, Clone, Deserialize)]
pub struct UpdateCommentRequest {
  pub text: String
}

impl UpdateCommentRequest {
  pub fn to_comment(&self) -> Comment {
    Comment {
      id: "".to_string(),
      createdAt: "".to_string(),
      postId: "".to_string(),
      userId: "".to_string(),
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