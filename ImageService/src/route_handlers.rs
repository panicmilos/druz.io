use std::path::{PathBuf, Path};

use rocket::response::NamedFile;
use rocket_contrib::json::Json;
use serde_json::{Value, json};

use super::utils;

use super::requests;


#[get("/<name..>", rank = 5)]
pub fn get_image(name: PathBuf) -> Option<NamedFile> {
    NamedFile::open(Path::new("public/").join(name)).ok()
}

#[post("/", format = "application/json", data = "<new_image>", rank = 10)]
pub fn upload_image(new_image: Json<requests::UploadImageRequest>) -> Json<Value> {

    let image = new_image.into_inner().to_image();

    let saveImageResult = utils::save_image(&image);
    match saveImageResult {
        Ok(_) => Json(json!({ "status": 200, "result": json!(image) })),
        Err(errorMessage) => Json(json!({ "status": 400, "result": json!({ "message": errorMessage }) }))
    }
}
