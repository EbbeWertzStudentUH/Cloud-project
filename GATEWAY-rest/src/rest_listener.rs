use tokio::net::TcpListener;

pub async fn create_listener() -> TcpListener {
    let listen_url = std::env::var("LISTEN_URL").expect("PORT bestaat niet in .env");
    let listener = TcpListener::bind(&listen_url)
        .await
        .expect("kon TCP listener niet maken");

    println!("Listening on {}", listener.local_addr().unwrap());

    listener
}
