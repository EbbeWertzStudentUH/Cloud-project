use serde::Serialize;

// REST
#[derive(Serialize)]
pub struct SimpleResponse {
    pub message: String,
}