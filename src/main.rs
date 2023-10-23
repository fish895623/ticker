#[macro_use]
extern crate rocket;

pub mod api;


#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", api::routes::get_routes())
}
