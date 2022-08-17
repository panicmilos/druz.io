use serde_derive::{Serialize, Deserialize};


#[derive(Serialize, Clone, Deserialize)]
pub struct UserFriend {
  pub UserId: u32,
  pub FriendId: u32
}