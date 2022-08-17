#![feature(proc_macro_hygiene, decl_macro)]

use std::env;
use dotenv::dotenv;
use rocket::{Config, config::Environment};
use rocket::http::Method;
use rocket_cors::{AllowedOrigins, CorsOptions};

#[macro_use] extern crate rocket;
extern crate base64;
extern crate dotenv;
extern crate uuid;

mod requests;
mod route_handlers;
mod utils;
mod models;

fn main() {
    dotenv().ok();

    let cfg = Config::build(Environment::Development)
        .address(env::var("HOST").unwrap())
        .port(env::var("PORT").unwrap().parse().unwrap())   
        .unwrap();

    let cors = CorsOptions::default()
        .allowed_origins(AllowedOrigins::all())
        .allowed_methods(
            vec![Method::Get, Method::Post, Method::Put, Method::Delete, Method::Patch]
                .into_iter()
                .map(From::from)
                .collect(),
        )
        .allow_credentials(true);

    rocket::custom(cfg)
        .attach(cors.to_cors().unwrap())
        .mount("/images",  routes![
            route_handlers::get_image, route_handlers::upload_image
        ])
        .launch();
}
