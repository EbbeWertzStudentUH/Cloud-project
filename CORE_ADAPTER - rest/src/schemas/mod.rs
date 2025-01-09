use serde::{Deserialize, Serialize};

// OpenAPI Schema for LoginRequest
#[derive(Deserialize)]
pub struct LoginRequest {
    pub email: String,
    pub password: String,
}

// OpenAPI Schema for RegisterRequest
#[derive(Deserialize)]
pub struct RegisterRequest {
    pub email: String,
    pub password: String,
    pub first_name: String,
    pub last_name: String,
}

// OpenAPI Schema for AuthResponse
#[derive(Serialize)]
pub struct AuthResponse {
    pub valid: bool,
    pub token: String,
    pub first_name: String,
    pub last_name: String,
}
