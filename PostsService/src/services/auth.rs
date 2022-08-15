use reqwest::StatusCode;

use crate::{errors::{HandableError, HandableResult, HandableErrorType}, models::auth::AuthenticatedUser};

use super::GetUserServiceUrl;



pub struct AuthService {
  client: reqwest::blocking::Client
}

impl AuthService {
    
  pub fn New() -> AuthService {
    AuthService {
      client: reqwest::blocking::Client::new()
    }
  }

  pub fn Authorize(&self, token: &String, roles: &Vec<String>) -> HandableResult<AuthenticatedUser> {

    let response = self.client.post(format!("{0}/authorize", GetUserServiceUrl()))
      .header("Authorization", token)
      .send()
      .unwrap();
      
    match response.status() {
      StatusCode::OK => {},
      _ => return Err(HandableError {
        error: HandableErrorType::Unauthorized,
        message: "You are not authenticated.".to_string()
      })
    };

    let authenticatedUser = response.json::<AuthenticatedUser>().unwrap();

    if roles.contains(&authenticatedUser.Role) {
      Ok(authenticatedUser)
    } else {
      Err(HandableError {
        error: HandableErrorType::Forbidden,
        message: "You are not authorized.".to_string()
      })
    }
  }

}