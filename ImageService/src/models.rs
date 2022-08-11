use serde_derive::{Serialize, Deserialize};


#[derive(Serialize, Clone, Deserialize)]
pub struct Image {
  pub name: String,
  pub extension: String,
  pub data: String,
}