use reqwest::StatusCode;

use crate::{models::{user_friends::UserFriend}, errors::{HandableResult, HandableError, HandableErrorType}};

use super::GetUserRelationServiceUrl;


pub struct UserFriendsService {
  client: reqwest::blocking::Client
}


impl UserFriendsService {
    
  pub fn New() -> UserFriendsService {
    UserFriendsService {
      client: reqwest::blocking::Client::new()
    }
  }

  pub fn ReadByIds(&self, token: &String, userFriend: &UserFriend) -> HandableResult<UserFriend> {

    let response = self.client.get(format!("{0}/users/{1}/friends/{2}", GetUserRelationServiceUrl(), &userFriend.UserId, &userFriend.FriendId))
      .header("Authorization", token)
      .send()
      .unwrap();
      
    match response.status() {
      StatusCode::OK => Ok(userFriend.clone()),
      _ => Err(HandableError {
        error: HandableErrorType::MissingEntity,
        message: "You are not friend with given user.".to_string()
      })
    }
  }
}