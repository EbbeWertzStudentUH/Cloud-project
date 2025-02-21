use actix_web::{App, HttpServer};
use actix_cors::Cors;

use proto_generated::{notification_service_client::NotificationServiceClient, project_service_client::ProjectServiceClient};
use tokio::sync::Mutex;
use std::sync::Arc;
use tonic::transport::Channel;
use once_cell::sync::Lazy;
use crate::proto_generated::user_service_client::UserServiceClient;

mod routes;
mod schemas;
mod grpc_client;
mod proto_generated {
    tonic::include_proto!("gateway_service");
}

pub static GRPC_CLIENT_USERSERVICE: Lazy<Arc<Mutex<Option<UserServiceClient<Channel>>>>> = Lazy::new(|| {
    Arc::new(Mutex::new(None)) // Initially None
});
pub static GRPC_CLIENT_NOTIFICATIONSERVICE: Lazy<Arc<Mutex<Option<NotificationServiceClient<Channel>>>>> = Lazy::new(|| {
    Arc::new(Mutex::new(None)) // Initially None
});
pub static GRPC_CLIENT_PROJECTSERVICE: Lazy<Arc<Mutex<Option<ProjectServiceClient<Channel>>>>> = Lazy::new(|| {
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
    let grpc_client_notificationservice = grpc_client::try_to_connect_notifierservice(gateway_url.as_str()).await;
    {
        let mut grpc_client_notificationservice_lock = GRPC_CLIENT_NOTIFICATIONSERVICE.lock().await;
        *grpc_client_notificationservice_lock = Some(grpc_client_notificationservice);
    }
    let grpc_client_projectservice = grpc_client::try_to_connect_projectservice(gateway_url.as_str()).await;
    {
        let mut grpc_client_projectservice_lock = GRPC_CLIENT_PROJECTSERVICE.lock().await;
        *grpc_client_projectservice_lock = Some(grpc_client_projectservice);
    }
    println!("connected to grpc");
    
    let listen_port = std::env::var("LISTEN_PORT").expect("LISTEN_PORT bestaat niet in .env");
    let mut listen_url = "0.0.0.0:".to_string();
    listen_url.push_str(&listen_port);
    println!("running on {listen_url}");


    HttpServer::new(move || {
        App::new()
            .wrap(
            Cors::default()
                .allow_any_origin()
                .allow_any_method()
                .allow_any_header()
                .max_age(3600),
            )
            .configure(routes::user::config)
            .configure(routes::notifiier::config)
            .configure(routes::projects::config)
    }).bind(listen_url)?.run().await
}
