use rocket::{request::{FromRequest, self}, Request, Outcome, http::{Status, RawStr}};
use rocket_contrib::json::{Json};
use serde_json::json;

use crate::{services::{auth::AuthService, comments::CommentsService, User}, requests::comments::{CreateCommentRequest, UpdateCommentRequest}, errors::{HandleError, HandableResult}, models::comments::Comment};

use super::api_response::ApiResponse;

#[derive(Debug)]
pub struct Token(String);

#[derive(Debug)]
pub enum ApiTokenError {
  Missing,
}

impl<'a, 'r> FromRequest<'a, 'r> for Token {
  type Error = ApiTokenError;

  fn from_request(request: &'a Request<'r>) -> request::Outcome<Self, Self::Error> {
    let token = request.headers().get_one("Authorization");
    match token {
      Some(token) => { Outcome::Success(Token(token.to_string())) },
      None => Outcome::Failure((Status::Unauthorized, ApiTokenError::Missing)),
    }
  }
}


#[post("/<post_id>/comments", format = "application/json", data = "<comment_request>")]
pub fn create_comment(post_id: &RawStr, token: Token, comment_request: Json<CreateCommentRequest>) -> ApiResponse {

  let comment_request = comment_request.into_inner();

  let validationResult: HandableResult<()> = comment_request.validate();
  match validationResult {
    Ok(_) => (),
    Err(err) => return err.to_api_response().unwrap()
  };

  let authService = AuthService::New();
  let authenticatedUser = match authService.Authorize(&token.0.to_string(), &vec![User.to_string()]) {
    Ok(authenticatedUser) => authenticatedUser,
    Err(err) => return err.to_api_response().unwrap()
  };

  let mut comment = comment_request.to_comment();
  comment.postId = post_id.to_string();
  comment.userId = authenticatedUser.Id.to_string();

  let commentsService = CommentsService::New();

  match commentsService.Create(&comment) {
    Ok(comment) => ApiResponse { json: Json(json!(comment)), status: Status::Ok },
    Err(err) => err.to_api_response().unwrap()
  }
}
    
#[put("/<post_id>/comments/<comment_id>", format = "application/json", data = "<comment_request>")]
pub fn update_comment(post_id: &RawStr, comment_id: &RawStr, token: Token, comment_request: Json<UpdateCommentRequest>) -> ApiResponse {

  let comment_request = comment_request.into_inner();

  let validationResult: HandableResult<()> = comment_request.validate();
  match validationResult {
    Ok(_) => (),
    Err(err) => return err.to_api_response().unwrap()
  };

  let authService = AuthService::New();
  let authenticatedUser = match authService.Authorize(&token.0.to_string(), &vec![User.to_string()]) {
    Ok(authenticatedUser) => authenticatedUser,
    Err(err) => return err.to_api_response().unwrap()
  };

  let mut comment = comment_request.to_comment();
  comment.id = comment_id.to_string();
  comment.postId = post_id.to_string();

  let commentsService = CommentsService::New();

  match commentsService.Update(&comment) {
    Ok(comment) => ApiResponse { json: Json(json!(comment)), status: Status::Ok },
    Err(err) => err.to_api_response().unwrap()
  }
    
}

#[delete("/<post_id>/comments/<comment_id>", format = "application/json")]
pub fn delete_comment(post_id: &RawStr, comment_id: &RawStr, token: Token) -> ApiResponse {

  let authService = AuthService::New();
  let authenticatedUser = match authService.Authorize(&token.0.to_string(), &vec![User.to_string()]) {
    Ok(authenticatedUser) => authenticatedUser,
    Err(err) => return err.to_api_response().unwrap()
  };

  let comment = Comment {
    id: comment_id.to_string(),
    createdAt: "".to_string(),
    postId: post_id.to_string(),
    userId: "".to_string(),
    text: "".to_string()
  };

  let commentsService = CommentsService::New();

  match commentsService.Delete(&comment) {
    Ok(comment) => ApiResponse { json: Json(json!(comment)), status: Status::Ok },
    Err(err) => err.to_api_response().unwrap()
  }
    
}
