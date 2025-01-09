use actix_web::{post, get, web, HttpResponse, Responder, HttpRequest};
use crate::schemas::user::{LoginRequest, RegisterRequest, AuthResponse};
use crate::proto_generated::{
    LoginRequest as GrpcLoginRequest,
    TokenRequest as GrpcTokenRequest,
    RegisterRequest as GrpcRegisterRequest};
use crate::GRPC_CLIENT_USERSERVICE;




// POST /user/authenticate
#[post("/user/authenticate")]
async fn authenticate(body: web::Json<LoginRequest>) -> impl Responder {
    let grpc_request: GrpcLoginRequest = body.into_inner().into();
    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
    match grpc_client.login_and_authenticate(grpc_request).await {
        Ok(response) => {
            let http_response: AuthResponse = response.into_inner().into();
            return HttpResponse::Ok().json(http_response)
        }
        Err(err) => {
            eprintln!("gRPC call failed: {}", err);
        }
    }}
    HttpResponse::InternalServerError().body("Failed to authenticate")
}

// GET /user/authenticate
#[get("/user/authenticate")]
async fn authenticate_token(req: HttpRequest) -> impl Responder {
    let (ok, token) = extract_bearer_token(req);
    if !ok { HttpResponse::Unauthorized().body("Unauthorized, could not find bearer token in header");}
    let grpc_request = GrpcTokenRequest{
        token: token
    };
    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
        match grpc_client.authenticate_token(grpc_request).await {
        Ok(response) => {
            let http_response: AuthResponse = response.into_inner().into();
            return HttpResponse::Ok().json(http_response);
        }
        Err(err) => {
            eprintln!("gRPC call failed: {}", err);
            
        }
    }}
    HttpResponse::InternalServerError().body("Failed to authenticate")
}

// POST /user/create_account
#[post("/user/create_account")]
async fn create_account(body: web::Json<RegisterRequest>) -> impl Responder {
    let grpc_request: GrpcRegisterRequest = body.into_inner().into();
    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
    match grpc_client.create_account(grpc_request).await {
        Ok(response) => {
            let http_response: AuthResponse = response.into_inner().into();
            return HttpResponse::Ok().json(http_response);
        }
        Err(err) => {
            eprintln!("gRPC call failed: {}", err);
            
        }
    }}
    HttpResponse::InternalServerError().body("Failed to create account")
}

fn extract_bearer_token(req: HttpRequest) -> (bool, String) {
    if let Some(auth) = req.headers().get("Authorization") {
        if let Ok(auth_str) = auth.to_str() {
            if auth_str.starts_with("Bearer ") {
                let token = &auth_str[7..]; // 'Bearer ' = 7 chars
                return (true, token.to_string());
            }
        }
    }
    return (false, "".to_string());
}

// Configure routes
pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(authenticate)
        .service(authenticate_token)
        .service(create_account);
}
