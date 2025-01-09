use tonic::transport::Channel;
use crate::proto_generated::user_service_client::UserServiceClient;
use tokio::time::{sleep, Duration};
use tokio::sync::Mutex;
use std::sync::Arc;
use actix_web::{web, HttpResponse, Responder, HttpRequest};


pub async fn try_to_connect(gateway_url: &str) -> UserServiceClient<Channel> {
    loop {
        match UserServiceClient::connect(gateway_url.to_string()).await {
            Ok(client) => {
                return client;
            }
            Err(_) => {
                println!("could not connect to gRPC server. Trying again in 3s");
                sleep(Duration::from_secs(3)).await;
            }
        }
    }
}