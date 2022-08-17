use std::env;

pub mod posts;
pub mod auth;
pub mod likes;
pub mod comments;
pub mod user_friends;
pub mod react;

pub fn GetUserServiceUrl() -> String {
  env::var("USER_SERVICE_URL").unwrap().to_string()
}

pub fn GetUserRelationServiceUrl() -> String {
  env::var("USER_RELATION_SERVICE_URL").unwrap().to_string()
}

pub const User: &str = "0";
pub const Administrator: &str = "1";