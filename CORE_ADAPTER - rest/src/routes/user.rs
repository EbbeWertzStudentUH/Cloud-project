use actix_web::{post, get, delete, web, HttpResponse, Responder, HttpRequest};
use crate::schemas::user::{LoginRequest, RegisterRequest, AuthResponse, FriendEditRequest, FriendsResponse};
use crate::proto_generated::{
    LoginRequest as GrpcLoginRequest,
    TokenRequest as GrpcTokenRequest,
    RegisterRequest as GrpcRegisterRequest,
    UserId as GrpcUserID,
    FriendEditRequest as GrpcFriendEditRequest,
};
use serde::Deserialize;

use crate::GRPC_CLIENT_USERSERVICE;

#[derive(Deserialize)]
struct UserIDQueryParams {
    user_id: String,
}


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

// GET /user/friends
#[get("/user/friends")]
async fn get_friends(query: web::Query<UserIDQueryParams>) -> impl Responder {
    let grpc_request: GrpcUserID = GrpcUserID { user_id: query.user_id.clone() };
    
    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
        match grpc_client.get_friends(grpc_request).await {
            Ok(response) => {
                let http_response: FriendsResponse = response.into_inner().into();
                return HttpResponse::Ok().json(http_response);
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
            }
        }
    }
    HttpResponse::InternalServerError().body("Failed to fetch friends")
}

// GET /user/friend-requests
#[get("/user/friend-requests")]
async fn get_friend_requests(query: web::Query<UserIDQueryParams>) -> impl Responder {
    let grpc_request: GrpcUserID = GrpcUserID { user_id: query.user_id.clone() };

    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
        match grpc_client.get_friend_requests(grpc_request).await {
            Ok(response) => {
                let http_response: FriendsResponse = response.into_inner().into();
                return HttpResponse::Ok().json(http_response);
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
            }
        }
    }
    HttpResponse::InternalServerError().body("Failed to fetch friend requests")
}

// POST /user/friend-request
#[post("/user/friend-request")]
async fn add_friend_request(body: web::Json<FriendEditRequest>) -> impl Responder {
    let grpc_request: GrpcFriendEditRequest = body.into_inner().into();
    
    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
        match grpc_client.add_friend_request(grpc_request).await {
            Ok(response) => {
                let http_response: FriendsResponse = response.into_inner().into();
                return HttpResponse::Ok().json(http_response);
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
            }
        }
    }
    HttpResponse::InternalServerError().body("Failed to add friend request")
}

// POST /user/friend
#[post("/user/friend")]
async fn add_friend(body: web::Json<FriendEditRequest>) -> impl Responder {
    let grpc_request: GrpcFriendEditRequest = body.into_inner().into();

    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
        match grpc_client.add_friend(grpc_request).await {
            Ok(response) => {
                let http_response: FriendsResponse = response.into_inner().into();
                return HttpResponse::Ok().json(http_response);
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
            }
        }
    }
    HttpResponse::InternalServerError().body("Failed to add friend")
}

// DELETE /user/remove-friend-request
#[delete("/user/remove-friend-request")]
async fn remove_friend_request(body: web::Json<FriendEditRequest>) -> impl Responder {
    let grpc_request: GrpcFriendEditRequest = body.into_inner().into();
    
    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
        match grpc_client.remove_friend_request(grpc_request).await {
            Ok(response) => {
                let http_response: FriendsResponse = response.into_inner().into();
                return HttpResponse::Ok().json(http_response);
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
            }
        }
    }
    HttpResponse::InternalServerError().body("Failed to remove friend request")
}

// DELETE /user/remove-friend
#[delete("/user/remove-friend")]
async fn remove_friend(body: web::Json<FriendEditRequest>) -> impl Responder {
    let grpc_request: GrpcFriendEditRequest = body.into_inner().into();

    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
        match grpc_client.remove_friend(grpc_request).await {
            Ok(response) => {
                let http_response: FriendsResponse = response.into_inner().into();
                return HttpResponse::Ok().json(http_response);
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
            }
        }
    }
    HttpResponse::InternalServerError().body("Failed to remove friend")
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
