use rocket::{http::{Status, RawStr}, request::{FromRequest, self}, Request, Outcome};
use rocket_contrib::json::Json;
use serde_json::json;

use crate::{requests::posts::{CreatePostRequest, UpdatePostRequest}, errors::{HandableResult, HandleError}, route_handlers::api_response::ApiResponse, services::{auth::AuthService, User, user_friends::UserFriendsService}, models::post::Post};

use crate::services::posts::PostsService;



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

#[get("/", format = "application/json")]
pub fn get_posts(token: Token) -> ApiResponse {

  let authService = AuthService::New();
  let authenticatedUser = match authService.Authorize(&token.0.to_string(), &vec![User.to_string()]) {
    Ok(authenticatedUser) => authenticatedUser,
    Err(err) => return err.to_api_response().unwrap()
  };

  let postsService = PostsService::New();
  let posts = match postsService.GetAll() {
    Ok(posts) => posts,
    Err(err) => return err.to_api_response().unwrap()
  };

  // let userFriendsService = UserFriendsService::New();
  // let userFriends = match userFriendsService.ReadById(&token.0.to_string(), &authenticatedUser.Id) {
  //   Ok(userFriends) => userFriends,
  //   Err(err) => return err.to_api_response().unwrap()
  // };

  // let filteredPosts = posts.iter().filter(|p| {
  //   if (p.writtenBy == authenticatedUser.Id) {
  //     return true;
  //   }

  //   if (userFriends.iter().any(|uf| p.writtenBy == uf.FriendId)) {
  //     return true;
  //   }

  //   return false;

  // }).collect::<Vec::<&Post>>();

  // ApiResponse { json: Json(json!(filteredPosts)), status: Status::Ok }

    ApiResponse { json: Json(json!(posts)), status: Status::Ok }

    
}

#[post("/", format = "application/json", data = "<post_request>")]
pub fn create_post(token: Token, post_request: Json<CreatePostRequest>) -> ApiResponse {

  let authService = AuthService::New();
  let authenticatedUser = match authService.Authorize(&token.0.to_string(), &vec![User.to_string()]) {
    Ok(authenticatedUser) => authenticatedUser,
    Err(err) => return err.to_api_response().unwrap()
  };

  let post_request = post_request.into_inner();

  let validationResult: HandableResult<()> = post_request.validate();
  match validationResult {
    Ok(_) => (),
    Err(err) => return err.to_api_response().unwrap()
  }

  let mut post = post_request.to_post();
  post.writtenBy = authenticatedUser.Id.to_string();
  let postsService = PostsService::New();

  match postsService.Create(&post) {
    Ok(post) => ApiResponse { json: Json(json!(post)), status: Status::Ok },
    Err(err) => err.to_api_response().unwrap()
  }
    
}

#[put("/<id>", format = "application/json", data = "<post_request>")]
pub fn update_post(id: &RawStr, token: Token, post_request: Json<UpdatePostRequest>) -> ApiResponse {

  let authService = AuthService::New();
  match authService.Authorize(&token.0.to_string(), &vec![User.to_string()]) {
    Ok(_) => {},
    Err(err) => return err.to_api_response().unwrap()
  };

  let post_request = post_request.into_inner();

  let validationResult: HandableResult<()> = post_request.validate();
  match validationResult {
    Ok(_) => (),
    Err(err) => return err.to_api_response().unwrap()
  }

  let mut post = post_request.to_post();
  post.id = id.to_string();
  let postsService = PostsService::New();

  match postsService.Update(&post) {
    Ok(post) => ApiResponse { json: Json(json!(post)), status: Status::Ok },
    Err(err) => err.to_api_response().unwrap()
  }
    
}

#[delete("/<id>", format = "application/json")]
pub fn delete_post(id: &RawStr, token: Token) -> ApiResponse {

  let authService = AuthService::New();
  match authService.Authorize(&token.0.to_string(), &vec![User.to_string()]) {
    Ok(_) => {},
    Err(err) => return err.to_api_response().unwrap()
  };

  let postsService = PostsService::New();

  match postsService.Delete(&id.to_string()) {
    Ok(post) => ApiResponse { json: Json(json!(post)), status: Status::Ok },
    Err(err) => err.to_api_response().unwrap()
  }
    
}