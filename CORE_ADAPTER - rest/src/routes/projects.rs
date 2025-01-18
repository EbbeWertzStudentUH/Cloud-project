use actix_web::{post, get, put, delete, web, HttpResponse, Responder, HttpRequest};
use crate::schemas::projects::*;
use jsonwebtoken::{decode, DecodingKey, Validation, Algorithm};
use serde::Deserialize;
use crate::proto_generated::{
    ProjectCreateRequest as GRPCProjectCreateRequest,
    ProjectId as GRPCProjectId,
    AddUserToProjectRequest as GRPCAddUserToProjectRequest,
    MilestoneAddRequest as GRPCMilestoneAddRequest,
    TaskAddRequest as GRPCTaskAddRequest,
    ProblemAddRequest as GRPCProblemAddRequest,
    ResolveProblemRequest as GRPCResolveProblemRequest,
    TaskAssignRequest as GRPCTaskAssignRequest,
    TaskCompleteRequest as GRPCTaskCompleteRequest,
    UserId as GRPCUserId
};
use crate::GRPC_CLIENT_PROJECTSERVICE;
#[derive(Debug, Deserialize)]
struct TokenContent {
    user_id: String,
}
// POST /project/create
#[post("/project/create")]
async fn create_project(body: web::Json<ProjectCreateRequest>, req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {
        return HttpResponse::InternalServerError().body("Failed to create project");
    }
    let mut grpc_request: GRPCProjectCreateRequest = body.into_inner().into();
    grpc_request.user_id = user_id;
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.create_project(grpc_request).await {
            Ok(_) => HttpResponse::Ok().finish(),
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to create project")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
}

// GET /project/full/{id}
#[get("/project/full/{id}")]
async fn get_full_project_by_id(project_id: web::Path<String>) -> impl Responder {
    let grpc_request = GRPCProjectId { project_id: project_id.to_string() };
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.get_full_project_by_id(grpc_request).await {
            Ok(response) => HttpResponse::Ok().json(Project::from(response.into_inner())),
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to fetch project")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
}

// GET /projects/user
#[get("/projects/user")]
async fn get_projects_from_user(req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {
        return HttpResponse::InternalServerError().body("Failed to fetch projects");
    }
    let grpc_request = GRPCUserId { user_id };
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.get_projects_from_user(grpc_request).await {
            Ok(response) => HttpResponse::Ok().json(ProjectsList::from(response.into_inner())),
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to fetch projects")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
}

// POST /project/user
#[post("/project/user")]
async fn add_user_to_project(body: web::Json<AddUserToProjectRequest>, req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {
        return HttpResponse::InternalServerError().body("Failed to add user to project");
    }
    let mut grpc_request: GRPCAddUserToProjectRequest = body.into_inner().into();
    grpc_request.user_id = user_id;
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.add_user_to_project(grpc_request).await {
            Ok(_) => HttpResponse::Ok().finish(),
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to add user to project")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
}

// POST /project/milestone
#[post("/project/milestone")]
async fn create_milestone_in_project(body: web::Json<MilestoneAddRequest>) -> impl Responder {
    let grpc_request: GRPCMilestoneAddRequest = body.into_inner().into();
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.create_milestone_in_project(grpc_request).await {
            Ok(_) => HttpResponse::Ok().finish(),
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to create milestone")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
}

// POST /project/task
#[post("/project/task")]
async fn create_task_in_milestone(body: web::Json<TaskAddRequest>) -> impl Responder {
    let grpc_request: GRPCTaskAddRequest = body.into_inner().into();
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.create_task_in_milestone(grpc_request).await {
            Ok(_) => HttpResponse::Ok().finish(),
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to create task")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
}

// POST /project/task/problem
#[post("/project/task/problem")]
async fn add_problem_to_task(body: web::Json<ProblemAddRequest>) -> impl Responder {
    let grpc_request: GRPCProblemAddRequest = body.into_inner().into();
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.add_problem_to_task(grpc_request).await {
            Ok(_) => HttpResponse::Ok().finish(),
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to add problem")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
}

// POST /project/task/problem/resolve
#[delete("/project/task/problem/resolve")]
async fn resolve_problem(body: web::Json<ResolveProblemRequest>) -> impl Responder {
    let grpc_request: GRPCResolveProblemRequest = body.into_inner().into();
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.resolve_problem(grpc_request).await {
            Ok(_) => HttpResponse::Ok().finish(),
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to resolve problem")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
}

// POST /project/task/assign
#[put("/project/task/assign")]
async fn assign_task(body: web::Json<TaskAssignRequest>) -> impl Responder {
    let grpc_request: GRPCTaskAssignRequest = body.into_inner().into();
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.assign_task(grpc_request).await {
            Ok(_) => HttpResponse::Ok().finish(),
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to assign task")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
}

// POST /project/task/complete
#[put("/project/task/complete")]
async fn complete_task(body: web::Json<TaskCompleteRequest>) -> impl Responder {
    let grpc_request: GRPCTaskCompleteRequest = body.into_inner().into();
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.complete_task(grpc_request).await {
            Ok(_) => HttpResponse::Ok().finish(),
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to complete task")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
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
    cfg.service(create_project)
        .service(get_full_project_by_id)
        .service(get_projects_from_user)
        .service(add_user_to_project)
        .service(create_milestone_in_project)
        .service(create_task_in_milestone)
        .service(add_problem_to_task)
        .service(resolve_problem)
        .service(complete_task);
}