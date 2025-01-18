use actix_web::{put, web, HttpResponse, Responder, HttpRequest};
use crate::schemas::notifier::{SimpleResponse, ProjectSubscribeRequest};
use jsonwebtoken::{decode, DecodingKey, Validation, Algorithm};
use serde::Deserialize;
use crate::proto_generated::{UserId as GrpcUserID, ProjectSubscribeRequest as GRPCProjectSubscribeRequest};

use crate::GRPC_CLIENT_NOTIFICATIONSERVICE;
#[derive(Debug, Deserialize)]
struct TokenContent {
    user_id: String,
}

// PUT /notifier/subscribe/friends
#[put("/notifier/subscribe/friends")]
async fn subscribe_friends(req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {return HttpResponse::InternalServerError().body("Failed to subscribe friends list")}
    let grpc_request: GrpcUserID = GrpcUserID { user_id: user_id.clone() };
    
    if let Some(grpc_client) = &mut *GRPC_CLIENT_NOTIFICATIONSERVICE.lock().await {
        match grpc_client.subscribe_friend_list(grpc_request).await {
            Ok(_) => {
                
                return HttpResponse::Ok().json(SimpleResponse { message:"subscribed notifier to friends list".to_string()});
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
            }
        }
    }
    HttpResponse::InternalServerError().body("Failed to subscribe friends list")
}
// PUT /notifier/subscribe/projects
#[put("/notifier/subscribe/projects")]
async fn subscribe_projects_list(req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {return HttpResponse::InternalServerError().body("Failed to subscribe friends list")}
    let grpc_request: GrpcUserID = GrpcUserID { user_id: user_id.clone() };
    
    if let Some(grpc_client) = &mut *GRPC_CLIENT_NOTIFICATIONSERVICE.lock().await {
        match grpc_client.subscribe_projects_list(grpc_request).await {
            Ok(_) => {
                
                return HttpResponse::Ok().json(SimpleResponse { message:"subscribed notifier to projects list".to_string()});
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
            }
        }
    }
    HttpResponse::InternalServerError().body("Failed to subscribe projects list")
}
// PUT /notifier/subscribe/all
#[put("/notifier/subscribe/all")]
async fn subscribe_all(req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {return HttpResponse::InternalServerError().body("Failed to subscribe friends list")}
    let grpc_request: GrpcUserID = GrpcUserID { user_id: user_id.clone() };
    
    if let Some(grpc_client) = &mut *GRPC_CLIENT_NOTIFICATIONSERVICE.lock().await {
        match grpc_client.subscribe_all_initial(grpc_request).await {
            Ok(_) => {
                
                return HttpResponse::Ok().json(SimpleResponse { message:"subscribed notifier to all".to_string()});
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
            }
        }
    }
    HttpResponse::InternalServerError().body("Failed to subscribe all")
}
// PUT /notifier/subscribe/project
#[put("/notifier/subscribe/project")]
async fn subscribe_project(req: HttpRequest, body: web::Json<ProjectSubscribeRequest>) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {return HttpResponse::InternalServerError().body("Failed to subscribe friends list")}
    let mut grpc_request: GRPCProjectSubscribeRequest = body.into_inner().into();
    grpc_request.user_id = user_id;
    
    if let Some(grpc_client) = &mut *GRPC_CLIENT_NOTIFICATIONSERVICE.lock().await {
        match grpc_client.switch_project_subscription(grpc_request).await {
            Ok(_) => {
                
                return HttpResponse::Ok().json(SimpleResponse { message:"subscribed notifier to all".to_string()});
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
            }
        }
    }
    HttpResponse::InternalServerError().body("Failed to subscribe all")
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
    cfg.service(subscribe_friends)
    .service(subscribe_project)
    .service(subscribe_projects_list)
    .service(subscribe_all);
}