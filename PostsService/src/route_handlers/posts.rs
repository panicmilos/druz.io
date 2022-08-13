use rocket::http::Status;
use rocket_contrib::json::Json;
use serde_json::json;

use crate::{requests::posts::CreatePostRequest, errors::{HandableResult, HandleError}};

use super::api_response::ApiResponse;
use crate::services::posts::PostsService;

#[get("/", format = "application/json")]
pub fn get_posts() -> ApiResponse {

  let postsService = PostsService::New();

  match postsService.GetAll() {
    Ok(posts) => ApiResponse { json: Json(json!(posts)), status: Status::Ok },
    Err(err) => err.to_api_response().unwrap()
  }
    
}

#[post("/", format = "application/json", data = "<post_request>")]
pub fn create_post(post_request: Json<CreatePostRequest>) -> ApiResponse {

  let post_request = post_request.into_inner();

  let validationResult: HandableResult<()> = post_request.validate();
  match validationResult {
    Ok(_) => (),
    Err(err) => return err.to_api_response().unwrap()
  }

  let post = post_request.to_post();
  let postsService = PostsService::New();

  match postsService.Create(&post) {
    Ok(post) => ApiResponse { json: Json(json!(post)), status: Status::Ok },
    Err(err) => err.to_api_response().unwrap()
  }
    
}
