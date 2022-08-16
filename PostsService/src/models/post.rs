
use serde_derive::{Serialize, Deserialize};

use super::comments::Comment;


#[derive(Serialize, Clone, Deserialize)]
pub struct Post {
  pub id: String,
  pub createdAt: String,
  pub text: String,
  pub writtenBy: String,
  pub likedBy: Vec<String>,
  pub comments: Vec<Comment>
}