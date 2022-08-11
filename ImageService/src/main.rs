#![feature(proc_macro_hygiene, decl_macro)]

use std::env;
use dotenv::dotenv;
use rocket::{Config, config::Environment};

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

    rocket::custom(cfg)
        .mount("/images",  routes![
            route_handlers::get_image, route_handlers::upload_image
        ])
        .launch();
}
