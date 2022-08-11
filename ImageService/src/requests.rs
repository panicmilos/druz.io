use serde_derive::{Serialize, Deserialize};
use uuid::Uuid;

use super::models::Image;

#[derive(Serialize, Clone, Deserialize)]
pub struct UploadImageRequest {
  pub data: String,
  pub ext: String
}

impl UploadImageRequest {
  pub fn to_image(&self) -> Image {
    Image {
      name: Uuid::new_v4().to_string(),
      extension: self.ext.clone(),
      data: self.data.clone()
    }
  }
}