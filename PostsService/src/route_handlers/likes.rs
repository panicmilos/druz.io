use rocket::{request::{FromRequest, self}, Request, Outcome, http::{Status, RawStr}};
use rocket_contrib::json::Json;
use serde_json::json;

use crate::{models::likes::Like, services::{likes::LikesService, auth::AuthService, User, react::ReactService}, errors::HandleError};

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

#[post("/<post_id>/like")]
pub fn like_post(post_id: &RawStr, token: Token) -> ApiResponse {

  let authService = AuthService::New();
  let authenticatedUser = match authService.Authorize(&token.0.to_string(), &vec![User.to_string()]) {
    Ok(authenticatedUser) => authenticatedUser,
    Err(err) => return err.to_api_response().unwrap()
  };

  let reactService = ReactService::New();
  match reactService.CanReact(&post_id.to_string(), &authenticatedUser.Id, &token.0.to_string()) {
    Ok(_) => {},
    Err(err) => return err.to_api_response().unwrap()
  };

  let like = Like {
    postId: post_id.to_string(),
    userId: authenticatedUser.Id.to_string()
  };

  let likesService = LikesService::New();

  match likesService.Like(&like) {
    Ok(like) => ApiResponse { json: Json(json!(like)), status: Status::Ok },
    Err(err) => err.to_api_response().unwrap()
  }
    
}

#[post("/<post_id>/dislike")]
pub fn dislike_post(post_id: &RawStr, token: Token) -> ApiResponse {

  let authService = AuthService::New();
  let authenticatedUser = match authService.Authorize(&token.0.to_string(), &vec![User.to_string()]) {
    Ok(authenticatedUser) => authenticatedUser,
    Err(err) => return err.to_api_response().unwrap()
  };

  let reactService = ReactService::New();
  match reactService.CanReact(&post_id.to_string(), &authenticatedUser.Id, &token.0.to_string()) {
    Ok(_) => {},
    Err(err) => return err.to_api_response().unwrap()
  };

  let like = Like {
    postId: post_id.to_string(),
    userId: authenticatedUser.Id.to_string()
  };

  let likesService = LikesService::New();

  match likesService.Dislike(&like) {
    Ok(like) => ApiResponse { json: Json(json!(like)), status: Status::Ok },
    Err(err) => err.to_api_response().unwrap()
  }
    
}