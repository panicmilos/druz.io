use crate::models::post::Post;

use super::GetDbProviderUrl;



pub struct PostsRepository {
  client: reqwest::blocking::Client
}

impl PostsRepository {

  pub fn New() -> PostsRepository {
    PostsRepository {
      client: reqwest::blocking::Client::new()
    }
  }

  pub fn GetAll(&self) -> Option<Vec<Post>> {

    Some(
      self.client.get(format!("{0}/posts", GetDbProviderUrl()))
      .send()
      .unwrap()
      .json::<Vec<Post>>()
      .unwrap()
    )

  }

  pub fn Create(&self, post: &Post) -> Option<Post> {

    Some(
      self.client.post(format!("{0}/posts", GetDbProviderUrl()))
      .form(&post)
      .send()
      .unwrap()
      .json::<Post>()
      .unwrap()
    )

  }

}