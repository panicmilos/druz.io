use serde_derive::{Serialize, Deserialize};


#[derive(Serialize, Clone, Deserialize)]
pub struct AuthenticatedUser {
  pub Id: String,
  pub Role: String
}