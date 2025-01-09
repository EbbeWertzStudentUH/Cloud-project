use actix_web::{web, App, HttpServer};
use tokio::sync::Mutex;
use std::sync::Arc;
use tonic::{client, transport::Channel};
use once_cell::sync::Lazy;
use crate::proto_generated::user_service_client::UserServiceClient;

mod routes;
mod schemas;
mod grpc_client;
mod proto_generated {
    tonic::include_proto!("facade_service");
}

pub static GRPC_CLIENT_USERSERVICE: Lazy<Arc<Mutex<Option<UserServiceClient<Channel>>>>> = Lazy::new(|| {
    Arc::new(Mutex::new(None)) // Initially None
});


#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenvy::dotenv().ok(); // ok negeert errors want .env is enkel voor local. In docker word .env in environment geload en is er geen .env file (docker kan .env files niet eens COPY-en)
    
    let gateway_url = std::env::var("GATEWAY_URL").expect("GATEWAY_URL bestaat niet in .env");

    let grpc_client_userservice = grpc_client::try_to_connect_userservice(gateway_url.as_str()).await;
    {
        let mut grpc_client_userservice_lock = GRPC_CLIENT_USERSERVICE.lock().await;
        *grpc_client_userservice_lock = Some(grpc_client_userservice);
    }
    println!("connected to grpc");
    
    let listen_port = std::env::var("LISTEN_PORT").expect("LISTEN_PORT bestaat niet in .env");
    let mut listen_url = "0.0.0.0:".to_string();
    listen_url.push_str(&listen_port);
    println!("running on {listen_url}");


    HttpServer::new(move || {
        App::new()
            .configure(routes::user::config)
    }).bind(listen_url)?.run().await
}
