
use serde_derive::{Serialize, Deserialize};


#[derive(Serialize, Clone, Deserialize)]
pub struct Post {
  pub id: String,
  pub createdAt: String,
  pub text: String,
  pub writtenBy: String,
  pub likedBy: Vec<String>
}