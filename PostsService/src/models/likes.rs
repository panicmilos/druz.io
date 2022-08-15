use serde_derive::{Serialize, Deserialize};


#[derive(Serialize, Clone, Deserialize)]
pub struct Like {
  pub userId: String,
  pub postId: String
}