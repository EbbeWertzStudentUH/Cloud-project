use actix_web::{post, get, delete, web, HttpResponse, Responder, HttpRequest};
use crate::schemas::notifier::SimpleResponse;
use crate::schemas::user::{LoginRequest, RegisterRequest, AuthResponse, FriendEditRequest, FriendsResponse, User};
use jsonwebtoken::{decode, DecodingKey, Validation, Algorithm};
use serde::Deserialize;
use crate::proto_generated::{
    FriendEditRequest as GrpcFriendEditRequest, LoginRequest as GrpcLoginRequest, RegisterRequest as GrpcRegisterRequest, TokenRequest as GrpcTokenRequest, UserId as GrpcUserID
};

use crate::GRPC_CLIENT_USERSERVICE;
#[derive(Debug, Deserialize)]
struct TokenContent {
    user_id: String,
}

#[derive(Deserialize)]
struct UserIDQueryParam {
    id: String,
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

// GET /user?id
#[get("/user")]
async fn get_user_name(query: web::Query<UserIDQueryParam>) -> impl Responder {
    let grpc_request = GrpcUserID{
        user_id: query.id.to_string()
    };
    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
        match grpc_client.get_user_name(grpc_request).await {
        Ok(response) => {
            let http_response: User = response.into_inner().into();
            return HttpResponse::Ok().json(http_response);
        }
        Err(err) => {
            eprintln!("gRPC call failed: {}", err);
            
        }
    }}
    HttpResponse::InternalServerError().body("Failed to get user info")
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
async fn get_friends(req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {return HttpResponse::InternalServerError().body("Failed to fetch friends")}
    let grpc_request: GrpcUserID = GrpcUserID { user_id: user_id.clone() };
    
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
async fn get_friend_requests(req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {return HttpResponse::InternalServerError().body("Failed to fetch friend requests")}
    let grpc_request: GrpcUserID = GrpcUserID { user_id: user_id.clone() };

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



// POST /user/friend-requests/send
#[post("/user/friend-requests/send")]
async fn send_friend_request(body: web::Json<FriendEditRequest>, req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {return HttpResponse::InternalServerError().body("Failed to send friend request")}
    let mut grpc_request: GrpcFriendEditRequest = body.into_inner().into();
    grpc_request.user_id = user_id;
    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
        match grpc_client.send_friend_request(grpc_request).await {
            Ok(_) => {
                return HttpResponse::Ok().json(SimpleResponse { message:"friend request is sent".to_string()});
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
            }
        }
    }
    HttpResponse::InternalServerError().body("Failed to send friend request")
}

// POST /user/friend-requests/accept
#[post("/user/friend-requests/accept")]
async fn accept_friend_request(body: web::Json<FriendEditRequest>, req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {return HttpResponse::InternalServerError().body("Failed to accept friend request")}
    let mut grpc_request: GrpcFriendEditRequest = body.into_inner().into();
    grpc_request.user_id = user_id;
    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
        match grpc_client.accept_friend_request(grpc_request).await {
            Ok(_) => {
                return HttpResponse::Ok().json(SimpleResponse { message:"you are now friends".to_string()});
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
            }
        }
    }
    HttpResponse::InternalServerError().body("Failed to accept friend request")
}

// DELETE /user/friend-requests/reject
#[delete("/user/friend-requests/reject")]
async fn reject_friend_request(body: web::Json<FriendEditRequest>, req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {return HttpResponse::InternalServerError().body("Failed to remove friend request")}
    let mut grpc_request: GrpcFriendEditRequest = body.into_inner().into();
    grpc_request.user_id = user_id;
    if let Some(grpc_client) = &mut *GRPC_CLIENT_USERSERVICE.lock().await {
        match grpc_client.reject_friend_request(grpc_request).await {
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

// DELETE /user/friend
#[delete("/user/friends")]
async fn remove_friend(body: web::Json<FriendEditRequest>, req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {return HttpResponse::InternalServerError().body("Failed to remove friend")}
    let mut grpc_request: GrpcFriendEditRequest = body.into_inner().into();
    grpc_request.user_id = user_id;
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

fn get_and_decode_token(req: HttpRequest) -> (bool, String) {
    let (ok, token) = extract_bearer_token(req);
    if !ok {
        return (false, "".to_string());
    }
    let secret = std::env::var("JWT_SECRET").expect("JWT_SECRET bestaat niet in .env");
    let decoded = decode::<TokenContent>(
        token.as_str(),
        &DecodingKey::from_secret(secret.as_ref()),
        &Validation::new(Algorithm::HS256),
    );
    match decoded {
        Ok(data) => (true, data.claims.user_id),
        Err(err) => {
            eprintln!("Failed to decode token: {}", err);
            (false, "".to_string())
        }
    }
}

// Configure routes
pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(authenticate)
        .service(authenticate_token)
        .service(create_account)
        .service(get_friends)
        .service(get_friend_requests)
        .service(send_friend_request)
        .service(accept_friend_request)
        .service(remove_friend)
        .service(reject_friend_request)
        .service(get_user_name);
}
