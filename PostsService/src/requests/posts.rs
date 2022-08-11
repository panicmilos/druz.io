use serde_derive::{Serialize, Deserialize};

use crate::models::post::Post;



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
}