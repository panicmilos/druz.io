use crate::{errors::HandableResult, models::{self, user_friends::UserFriend}};

use super::{posts::PostsService, user_friends::UserFriendsService};

pub struct ReactService {
  postsService: PostsService,
  userFriendsService: UserFriendsService
}


impl ReactService {
    
  pub fn New() -> ReactService {
    ReactService {
      postsService: PostsService::New(),
      userFriendsService: UserFriendsService::New()
    }
  }

  pub fn CanReact(&self, postId: &String, userId: &String, token: &String) -> HandableResult<()> {

    let post = match self.postsService.GetByIdOrThrow(postId) {
      Ok(post) => post,
      Err(err) => return Err(err)
    };

    if post.writtenBy == userId.to_string() {
      return Ok(());
    }

    match self.userFriendsService.ReadByIds(token, &UserFriend {
      UserId: userId.to_string(),
      FriendId: post.writtenBy.to_string()
    }) {
      Ok(_) => Ok(()),
      Err(err) => Err(err)
    }
  }

}