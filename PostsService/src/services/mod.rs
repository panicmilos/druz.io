use std::env;

pub mod posts;
pub mod auth;
pub mod likes;
pub mod comments;

pub fn GetUserServiceUrl() -> String {
  env::var("USER_SERVICE_URL").unwrap().to_string()
}

pub const User: &str = "0";
pub const Administrator: &str = "1";