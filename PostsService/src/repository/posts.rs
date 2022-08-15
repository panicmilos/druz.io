use reqwest::StatusCode;

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

  pub fn GetById(&self, id: &String) -> Option<Post> {

    let response = self.client.get(format!("{0}/posts/{1}", GetDbProviderUrl(), id))
      .send()
      .unwrap();

    match response.status() {
      StatusCode::OK => Some(response.json::<Post>().unwrap()),
      _ => None
    }

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

  pub fn Update(&self, post: &Post) -> Option<Post> {

    Some(
      self.client.put(format!("{0}/posts/{1}", GetDbProviderUrl(), post.id))
      .form(&post)
      .send()
      .unwrap()
      .json::<Post>()
      .unwrap()
    )

  }

  pub fn Delete(&self, id: &String) -> Option<Post> {

    Some(
      self.client.delete(format!("{0}/posts/{1}", GetDbProviderUrl(), id))
      .send()
      .unwrap()
      .json::<Post>()
      .unwrap()
    )

  }

}