use std::env;


pub mod posts;

pub fn GetDbProviderUrl() -> String {
  env::var("DB_PROVIDER_URL").unwrap().to_string()
}