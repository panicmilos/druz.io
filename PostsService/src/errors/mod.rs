use rocket::{http::Status, response::status::Unauthorized};
use rocket_contrib::json::Json;
use serde_json::json;

use crate::route_handlers::api_response::ApiResponse;


pub struct HandableError {
  pub message: String,
  pub error: HandableErrorType

}

pub enum HandableErrorType {
  BadLogic,
  Unauthorized,
  Forbidden,
  MissingEntity
}

pub type HandableResult<T> = Result<T, HandableError>;

pub trait HandleError {
  fn to_api_response(&self) -> Option<ApiResponse>;
}

impl<T> HandleError for HandableResult<T> {
  fn to_api_response(&self) -> Option<ApiResponse> {

    match &self
    {
      Ok(_) => None,
      Err(handableError) => match handableError.error {
        HandableErrorType::BadLogic => Some(ApiResponse { json: Json(json!({ "message": handableError.message })), status: Status::BadRequest }),
        HandableErrorType::Unauthorized => Some(ApiResponse { json: Json(json!({ "message": handableError.message })), status: Status::Unauthorized }),
        HandableErrorType::Forbidden => Some(ApiResponse { json: Json(json!({ "message": handableError.message })), status: Status::Forbidden }),
        HandableErrorType::MissingEntity => Some(ApiResponse { json: Json(json!({ "message": handableError.message })), status: Status::NotFound })
      }
    }
  }
}