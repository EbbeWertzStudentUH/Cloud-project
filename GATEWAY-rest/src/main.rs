use axum::Router;
use tokio::net::TcpListener;
use crate::rest_listener::create_listener;
use crate::rest_router::create_router;
mod rest_listener;
mod rest_router;

#[tokio::main]
async fn main() {
    // Expose .env variables naar echte environment
    dotenvy::dotenv().expect("kan .env niet vinden");

    let listener: TcpListener = create_listener().await;
    let app: Router = create_router();

    axum::serve(listener, app)
        .await
        .expect("Kon app niet serven");
}
