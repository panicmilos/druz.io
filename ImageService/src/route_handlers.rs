use std::path::{PathBuf, Path};

use rocket::Request;
use rocket::Response;
use rocket::http::ContentType;
use rocket::http::Status;
use rocket::response;
use rocket::response::NamedFile;
use rocket_contrib::json::Json;
use serde_json::{Value, json};
use rocket::response::Responder;

use super::utils;

use super::requests;

#[derive(Debug)]
pub struct ApiResponse {
    json: Json<Value>,
    status: Status,
}

impl<'r> Responder<'r> for ApiResponse {
    fn respond_to(self, req: &Request) -> response::Result<'r> {
        Response::build_from(self.json.respond_to(&req).unwrap())
            .status(self.status)
            .header(ContentType::JSON)
            .ok()
    }
}


#[get("/<name..>", rank = 5)]
pub fn get_image(name: PathBuf) -> Option<NamedFile> {
    NamedFile::open(Path::new("public/").join(name)).ok()
}

#[post("/", format = "application/json", data = "<new_image>", rank = 10)]
pub fn upload_image(new_image: Json<requests::UploadImageRequest>) -> ApiResponse {

    let image = new_image.into_inner().to_image();

    let saveImageResult = utils::save_image(&image);
    match saveImageResult {
        Ok(_) => ApiResponse { json: Json(json!(image)), status: Status::Ok },
        Err(errorMessage) => ApiResponse { json: Json(json!({ "message": errorMessage })), status: Status::BadRequest }
    }
}
