use tokio::net::TcpListener;

pub async fn create_listener() -> TcpListener {
    let listen_port = std::env::var("LISTEN_PORT").expect("PORT bestaat niet in .env");
    let mut listen_url = "0.0.0.0:".to_string();
    listen_url.push_str(&listen_port);
    let listener = TcpListener::bind(&listen_url)
        .await
        .expect("kon TCP listener niet maken");

    println!("Listening on {}", listener.local_addr().unwrap());

    listener
}
