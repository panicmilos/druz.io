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

mod errors;
mod requests;
mod models;
mod route_handlers;
mod repository;
mod services;

fn main() {
    println!("Hello, world!");

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
        .mount("/posts",  routes![
            route_handlers::posts::get_posts,
            route_handlers::posts::create_post,
            route_handlers::posts::update_post,
            route_handlers::posts::delete_post,

            route_handlers::likes::like_post,
            route_handlers::likes::dislike_post,

            route_handlers::comments::create_comment,
            route_handlers::comments::update_comment,
            route_handlers::comments::delete_comment
        ])
        .launch();


}
