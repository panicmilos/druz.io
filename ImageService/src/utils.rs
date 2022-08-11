
use std::fs::File;
use std::io::Write;

use super::models::Image;

pub fn save_image(image: &Image) -> Result<(), String> {
  
  let decodedData = base64::decode(&image.data).unwrap();
  let imageName = format!("{0}.{1}", image.name, image.extension);
  
  let createFileResult = File::create(format!("public/{imageName}"));
  let mut f = match createFileResult {
      Ok(file) => file,
      Err(_) => return Err("Unable to create file.".to_string())
  };

  let writeBytesResult = f.write_all(&decodedData);
  match writeBytesResult {
    Ok(_) => Ok(()),
    Err(_) => Err("Unable to write data to the file.".to_string())
  }
}