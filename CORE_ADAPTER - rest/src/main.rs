use actix_web::{web, App, HttpServer};
mod routes;
mod schemas;
mod grpc_client;
mod proto_generated {
    tonic::include_proto!("facade_service");
}


#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenvy::dotenv().expect("kan .env niet vinden");
    let listen_port = std::env::var("LISTEN_PORT").expect("PORT bestaat niet in .env");
    let gateway_url = std::env::var("GATEWAY_URL").expect("GATEWAY_URL bestaat niet in .env");
    let mut listen_url = "0.0.0.0:".to_string();
    listen_url.push_str(&listen_port);
    println!("running on {listen_url}");

    let grpc_client = grpc_client::try_to_connect(gateway_url.as_str()).await;
    println!("connected to grpc");

    let grpc_client_data = web::Data::new(grpc_client);

    HttpServer::new(move || {
        App::new()
            .app_data(grpc_client_data.clone())
            .configure(routes::user::config)
    }).bind(listen_url)?.run().await
}
