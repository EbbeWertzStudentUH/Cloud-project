use actix_web::{post, get, web, HttpResponse, Responder, HttpRequest};
use crate::schemas::{LoginRequest, RegisterRequest, AuthResponse};



// POST /user/authenticate
#[post("/user/authenticate")]
async fn authenticate(_body: web::Json<LoginRequest>) -> impl Responder {
    // Add your authentication logic here
    let response = AuthResponse {
        valid: true,
        token: "example_token".to_string(),
        first_name: "John".to_string(),
        last_name: "Doe".to_string(),
    };
    HttpResponse::Ok().json(response)
}

// GET /user/authenticate
#[get("/user/authenticate")]
async fn authenticate_token(req: HttpRequest) -> impl Responder {
    let (ok, token) = extract_bearer_token(req);
    if !ok { HttpResponse::Unauthorized().body("Unauthorized, could not find bearer token in header");}

    let response = AuthResponse {
        valid: true,
        token: token.to_string(),
        first_name: "John".to_string(),
        last_name: "Doe".to_string(),
    };
    return HttpResponse::Ok().json(response);
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


// POST /user/create_account
#[post("/user/create_account")]
async fn create_account(body: web::Json<RegisterRequest>) -> impl Responder {
    // Add your account creation logic here
    let response = AuthResponse {
        valid: true,
        token: "new_user_token".to_string(),
        first_name: body.first_name.clone(),
        last_name: body.last_name.clone(),
    };
    HttpResponse::Created().json(response)
}

// Configure routes
pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(authenticate)
        .service(authenticate_token)
        .service(create_account);
}
