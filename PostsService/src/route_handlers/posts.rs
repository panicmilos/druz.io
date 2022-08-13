use rocket::http::Status;
use rocket_contrib::json::Json;
use serde_json::json;

use crate::{requests::posts::CreatePostRequest, errors::{HandableResult, HandleError}};

use super::api_response::ApiResponse;

#[post("/", format = "application/json", data = "<post_request>")]
pub fn create_post(post_request: Json<CreatePostRequest>) -> ApiResponse {

  let post_request = post_request.into_inner();


  let validationResult: HandableResult<()> = post_request.validate();
  match validationResult.to_api_response() {
    Some(apiResponse) => return apiResponse,
    None => ()
  }

  let post = post_request.to_post();

  ApiResponse { json: Json(json!(post)), status: Status::Ok }
    
}
