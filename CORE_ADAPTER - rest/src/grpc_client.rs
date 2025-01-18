use tonic::transport::Channel;
use crate::proto_generated::{notification_service_client::NotificationServiceClient, project_service_client::ProjectServiceClient, user_service_client::UserServiceClient};
use tokio::time::{sleep, Duration};

pub async fn try_to_connect_userservice(gateway_url: &str) -> UserServiceClient<Channel> {
    loop {
        match UserServiceClient::connect(gateway_url.to_string()).await {
            Ok(client) => {
                println!("connected to gRPC: UserService.");
                return client;
            }
            Err(_) => {
                println!("could not connect to gRPC server: UserService. Trying again in 3s");
                sleep(Duration::from_secs(3)).await;
            }
        }
    }
}

pub async fn try_to_connect_notifierservice(gateway_url: &str) -> NotificationServiceClient<Channel> {
    loop {
        match NotificationServiceClient::connect(gateway_url.to_string()).await {
            Ok(client) => {
                println!("connected to gRPC: NotfierService.");
                return client;
            }
            Err(_) => {
                println!("could not connect to gRPC server: NotfierService. Trying again in 3s");
                sleep(Duration::from_secs(3)).await;
            }
        }
    }
}

pub async fn try_to_connect_projectservice(gateway_url: &str) -> ProjectServiceClient<Channel> {
    loop {
        match ProjectServiceClient::connect(gateway_url.to_string()).await {
            Ok(client) => {
                println!("connected to gRPC: ProjectService.");
                return client;
            }
            Err(_) => {
                println!("could not connect to gRPC server:ProjectService. Trying again in 3s");
                sleep(Duration::from_secs(3)).await;
            }
        }
    }
}