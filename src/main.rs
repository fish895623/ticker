#[macro_use]
extern crate rocket;
use rocket_dyn_templates::{context, Template};

#[get("/")]
fn index() -> Template {
    Template::render("index", context!(title: "This is Title"))
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .attach(Template::fairing())
        .mount("/", routes![index])
}
