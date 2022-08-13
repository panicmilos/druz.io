#![feature(proc_macro_hygiene, decl_macro)]

use std::env;

use dotenv::dotenv;
use rocket::{Config, config::Environment};

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

    rocket::custom(cfg)
        .mount("/posts",  routes![
            route_handlers::posts::get_posts,
            route_handlers::posts::create_post
        ])
        .launch();


}
