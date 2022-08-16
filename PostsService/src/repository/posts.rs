use std::collections::HashMap;

use reqwest::StatusCode;

use crate::models::post::Post;

use super::GetDbProviderUrl;
use serde_json::{Value, json};



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

    let mut map = HashMap::new();
    map.insert("text", Value::String(post.text.to_string()));
    map.insert("writtenBy", Value::String(post.writtenBy.to_string()));
    map.insert("likedBy", Value::Array(vec![]));
    map.insert("comments", Value::Array(vec![]));

    Some(
      self.client.post(format!("{0}/posts", GetDbProviderUrl()))
      .json(&map)
      .send()
      .unwrap()
      .json::<Post>()
      .unwrap()
    )

  }

  pub fn Update(&self, post: &Post) -> Option<Post> {

    let mut map = HashMap::new();
    map.insert("text", Value::String(post.text.to_string()));
    map.insert("writtenBy", Value::String(post.writtenBy.to_string()));
    map.insert("likedBy", Value::Array(post.likedBy.iter().map(|x| Value::String(x.to_string())).collect()));
    map.insert("comments", Value::Array(post.comments.iter().map(|x| json!(x)).collect()));

    Some(
      self.client.put(format!("{0}/posts/{1}", GetDbProviderUrl(), post.id))
      .json(&map)
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