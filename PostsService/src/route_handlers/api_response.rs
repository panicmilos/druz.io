use rocket::{http::{Status, ContentType}, Request, response, Response};
use rocket::response::Responder;
use rocket_contrib::json::Json;
use serde_json::Value;

#[derive(Debug)]
pub struct ApiResponse {
  pub json: Json<Value>,
  pub status: Status,
}

impl<'r> Responder<'r> for ApiResponse {
  fn respond_to(self, req: &Request) -> response::Result<'r> {
    Response::build_from(self.json.respond_to(&req).unwrap())
      .status(self.status)
      .header(ContentType::JSON)
      .ok()
  }
}