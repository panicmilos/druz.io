use rocket::http::Status;
use rocket_contrib::json::Json;
use serde_json::json;

use crate::requests::posts::CreatePostRequest;

use super::api_response::ApiResponse;

#[post("/", format = "application/json", data = "<post_request>")]
pub fn create_post(post_request: Json<CreatePostRequest>) -> ApiResponse {

  let post = post_request.into_inner().to_post();

  ApiResponse { json: Json(json!(post)), status: Status::Ok }
    
}
