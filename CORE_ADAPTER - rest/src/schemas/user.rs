use serde::{Deserialize, Serialize};
use crate::proto_generated::{
    AuthResponse as GrpcAuthResponse,
    LoginRequest as GrpcLoginRequest,
    RegisterRequest as GrpcRegisterRequest};

// REST
#[derive(Deserialize)]
pub struct LoginRequest {
    pub email: String,
    pub password: String,
}
// REST -> gRPC
impl From<LoginRequest> for GrpcLoginRequest {
    fn from(rest: LoginRequest) -> Self {
        GrpcLoginRequest {
            email: rest.email,
            password: rest.password,
        }
    }
}

// REST
#[derive(Deserialize)]
pub struct RegisterRequest {
    pub email: String,
    pub password: String,
    pub first_name: String,
    pub last_name: String,
}
// REST -> gRPC
impl From<RegisterRequest> for GrpcRegisterRequest {
    fn from(rest: RegisterRequest) -> Self {
        GrpcRegisterRequest {
            email: rest.email,
            password: rest.password,
            first_name: rest.first_name,
            last_name: rest.last_name
        }
    }
}

// REST
#[derive(Serialize)]
pub struct AuthResponse {
    pub valid: bool,
    pub token: String,
    pub first_name: String,
    pub last_name: String,
}
// GRPC -> REST
impl From<GrpcAuthResponse> for AuthResponse {
    fn from(grpc: GrpcAuthResponse) -> Self {
        AuthResponse {
            valid: grpc.valid,
            token: grpc.token,
            first_name: grpc.first_name,
            last_name: grpc.last_name,
        }
    }
}