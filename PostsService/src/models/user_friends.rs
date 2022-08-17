use serde_derive::{Serialize, Deserialize};


#[derive(Serialize, Clone, Deserialize)]
pub struct UserFriend {
  pub UserId: String,
  pub FriendId: String
}