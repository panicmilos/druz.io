use serde_derive::{Serialize, Deserialize};


#[derive(Serialize, Clone, Deserialize)]
pub struct Comment {
  pub id: String,
  pub createdAt: String,
  pub postId: String,
  pub userId: String,
  pub text: String
}