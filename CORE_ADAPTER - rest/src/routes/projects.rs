use jsonwebtoken::{decode, Algorithm, DecodingKey, Validation};
use serde::Deserialize;
use actix_web::{get, post, put, web, HttpRequest, HttpResponse, Responder};
use crate::proto_generated::{
    ProjectCreateRequest as GRPCProjectCreateRequest,
    ProjectId as GRPCProjectID,
    UserId as GRPCUserID,
    AddUserToProjectRequest as GRPCAddUserToProjectRequest,
    MilestoneCreateRequest as GRPCMilestoneCreateRequest,
    TaskCreateRequest as GRPCTaskCreateRequest,
    ProblemAddRequest as GRPCProblemAddRequest,
    ResolveProblemRequest as GRPCResolveProblemRequest,
    TaskAssignRequest as GRPCTaskAssignRequest,
    TaskCompleteRequest as GRPCTaskCompleteRequest,
};
use crate::schemas::projects::{
    AddUserToProjectRequest, MilestoneCreateRequest, ProblemAddRequest, Project, ProjectCreateRequest, ProjectsList, TaskCreateRequest
};
use crate::GRPC_CLIENT_PROJECTSERVICE;

// POST /project
#[post("/project")]
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

// GET /project/{project_id}
#[get("/project/{project_id}")]
async fn get_full_project_by_id(project_id: web::Path<String>) -> impl Responder {
    let grpc_request = GRPCProjectID { project_id: project_id.to_string() };
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.get_full_project_by_id(grpc_request).await {
            Ok(response) => {
                let http_response: Project = response.into_inner().into();
                return HttpResponse::Ok().json(http_response)
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to fetch project")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
}

// GET /projects
#[get("/projects")]
async fn get_projects_from_user(req: HttpRequest) -> impl Responder {
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {
        return HttpResponse::InternalServerError().body("Failed to fetch projects");
    }
    let grpc_request = GRPCUserID { user_id };
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.get_projects_from_user(grpc_request).await {
            Ok(response) => {
                let http_response: ProjectsList = response.into_inner().into();
                return HttpResponse::Ok().json(http_response)
            }
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to fetch projects")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
}

// POST /project/{project_id}/user
#[post("/project/{project_id}/user")]
async fn add_user_to_project(body: web::Json<AddUserToProjectRequest>, project_id: web::Path<String>) -> impl Responder {
    let mut grpc_request: GRPCAddUserToProjectRequest = body.into_inner().into();
    grpc_request.project_id = project_id.to_string();
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

// POST /project/{project_id}/milestone
#[post("/project/{project_id}/milestone")]
async fn create_milestone_in_project(body: web::Json<MilestoneCreateRequest>, project_id: web::Path<String>) -> impl Responder {
    let mut grpc_request: GRPCMilestoneCreateRequest = body.into_inner().into();
    grpc_request.project_id = project_id.to_string();
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

// POST /project/{project_id}/milestone/{milestone_id}/task
#[post("/project/{project_id}/milestone/{milestone_id}/task")]
async fn create_task_in_milestone(body: web::Json<TaskCreateRequest>, ids: web::Path<(String, String)>) -> impl Responder {
    let (project_id, milestone_id) = ids.into_inner();
    let mut grpc_request: GRPCTaskCreateRequest = body.into_inner().into();
    grpc_request.project_id = project_id.to_string();
    grpc_request.milestone_id = milestone_id.to_string();
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

// POST /project/{project_id}/task/{task_id}/problem
#[post("/project/{project_id}/task/{task_id}/problem")]
async fn add_problem_to_task(body: web::Json<ProblemAddRequest>, ids: web::Path<(String, String)>) -> impl Responder {
    let (project_id, task_id) = ids.into_inner();
    let mut grpc_request: GRPCProblemAddRequest = body.into_inner().into();
    grpc_request.project_id = project_id.to_string();
    grpc_request.task_id = task_id.to_string();
    if let Some(grpc_client) = &mut *GRPC_CLIENT_PROJECTSERVICE.lock().await {
        match grpc_client.add_problem_to_task(grpc_request).await {
            Ok(_) => HttpResponse::Ok().finish(),
            Err(err) => {
                eprintln!("gRPC call failed: {}", err);
                HttpResponse::InternalServerError().body("Failed to add problem to task")
            }
        }
    } else {
        HttpResponse::InternalServerError().body("Failed to connect to gRPC client")
    }
}

// PUT /project/{project_id}/task/{task_id}/problem/{problem_id}/resolve
#[put("/project/{project_id}/task/{task_id}/problem/{problem_id}/resolve")]
async fn resolve_problem(ids: web::Path<(String, String, String)>) -> impl Responder {
    let (project_id, task_id, problem_id) = ids.into_inner();
    let grpc_request = GRPCResolveProblemRequest{
        project_id: project_id.to_string(),
        problem_id: problem_id.to_string(),
        task_id: task_id.to_string()
    };
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

// PUT /project/{project_id}/task/{task_id}/assign
#[put("/project/{project_id}/task/{task_id}/assign")]
async fn assign_task(req: HttpRequest, ids: web::Path<(String, String)>) -> impl Responder {
    let (project_id, task_id) = ids.into_inner();
    let (ok, user_id) = get_and_decode_token(req);
    if !ok {
        return HttpResponse::InternalServerError().body("Failed to assign task");
    }
    let grpc_request = GRPCTaskAssignRequest{
        user_id: user_id,
        project_id: project_id.to_string(),
        task_id: task_id.to_string()
    };
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

// PUT /project/{project_id}/task/{task_id}/complete
#[post("/project/{project_id}/task/{task_id}/complete")]
async fn complete_task(ids: web::Path<(String, String)>) -> impl Responder {
    let (project_id, task_id) = ids.into_inner();
    let grpc_request = GRPCTaskCompleteRequest{
        project_id: project_id.to_string(),
        task_id: task_id.to_string()
    };
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

#[derive(Debug, Deserialize)]
struct TokenContent {
    user_id: String,
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
        .service(complete_task)
        .service(assign_task);
}