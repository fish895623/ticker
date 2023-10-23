#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}

pub fn get_routes() -> Vec<rocket::Route> {
    routes![index]
}