use actix_web::{App, HttpServer};
mod routes;
mod schemas;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenvy::dotenv().expect("kan .env niet vinden");
    let listen_port = std::env::var("LISTEN_PORT").expect("PORT bestaat niet in .env");
    let mut listen_url = "0.0.0.0:".to_string();
    listen_url.push_str(&listen_port);
    print!("running on {}", listen_url);
    HttpServer::new(|| {
        App::new()
            .configure(routes::user::config) // Configure user routes
    }).bind(listen_url)?.run().await
}
