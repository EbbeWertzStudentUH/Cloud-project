use serde::{Deserialize, Serialize};
use crate::proto_generated::{
    
    MilestoneCreateRequest as GrpcMilestoneCreateRequest,
    ProblemAddRequest as GrpcProblemAddRequest,
    ProjectCreateRequest as GrpcProjectCreateRequest,
    TaskCreateRequest as GrpcTaskCreateRequest,
    Problem as GrpcProblem,
    User as GrpcUser,
    ProjectsList as GrpcProjectsList,
    Milestone as GrpcMilestone,
    Task as GrpcTask,
    Project as GrpcProject,
    AddUserToProjectRequest as GRPCAddUserToProjectRequest
};


#[derive(Deserialize)]
pub struct ProjectCreateRequest {
    pub name: String,
    pub deadline: String,
    pub github_repo: String,
}

impl From<ProjectCreateRequest> for GrpcProjectCreateRequest {
    fn from(rest: ProjectCreateRequest) -> Self {
        GrpcProjectCreateRequest {
            user_id: String::new(),
            name: rest.name,
            deadline: rest.deadline,
            github_repo: rest.github_repo,
        }
    }
}

#[derive(Deserialize)]
pub struct AddUserToProjectRequest {
    pub user_id: String
}

impl From<AddUserToProjectRequest> for GRPCAddUserToProjectRequest {
    fn from(rest: AddUserToProjectRequest) -> Self {
        GRPCAddUserToProjectRequest {
            user_id: rest.user_id,
            project_id: String::new()
        }
    }
}

#[derive(Deserialize)]
pub struct MilestoneCreateRequest {
    pub name: String,
    pub deadline: String,
}

impl From<MilestoneCreateRequest> for GrpcMilestoneCreateRequest {
    fn from(rest: MilestoneCreateRequest) -> Self {
        GrpcMilestoneCreateRequest {
            project_id: String::new(),
            name: rest.name,
            deadline: rest.deadline,
        }
    }
}


#[derive(Deserialize)]
pub struct TaskCreateRequest {
    pub name: String,
}

impl From<TaskCreateRequest> for GrpcTaskCreateRequest {
    fn from(rest: TaskCreateRequest) -> Self {
        GrpcTaskCreateRequest {
            project_id: String::new(),
            milestone_id: String::new(),
            name: rest.name,
        }
    }
}

#[derive(Deserialize)]
pub struct ProblemAddRequest {
    pub problem: Problem,
}

impl From<ProblemAddRequest> for GrpcProblemAddRequest {
    fn from(rest: ProblemAddRequest) -> Self {
        GrpcProblemAddRequest {
            project_id: String::new(),
            task_id: String::new(),
            problem: Some(rest.problem.into()),
        }
    }
}



// REST
#[derive(Serialize)]
pub struct User {
    pub first_name: String,
    pub last_name: String,
    pub id: String,
}

impl From<GrpcUser> for User {
    fn from(grpc: GrpcUser) -> Self {
        User {
            first_name: grpc.first_name,
            last_name: grpc.last_name,
            id: grpc.id,
        }
    }
}

#[derive(Serialize)]
pub struct ProjectsList {
    pub projects: Vec<Project>,
}

impl From<GrpcProjectsList> for ProjectsList {
    fn from(grpc: GrpcProjectsList) -> Self {
        ProjectsList {
            projects: grpc.projects.into_iter().map(Project::from).collect(),
        }
    }
}

#[derive(Serialize)]
pub struct Project {
    pub id: String,
    pub name: String,
    pub users: Vec<User>,
    pub deadline: String,
    pub github_repo: String,
    pub milestones: Vec<Milestone>,
}

impl From<GrpcProject> for Project {
    fn from(grpc: GrpcProject) -> Self {
        Project {
            id: grpc.id,
            name: grpc.name,
            users: grpc.users.into_iter().map(User::from).collect(),
            deadline: grpc.deadline,
            github_repo: grpc.github_repo,
            milestones: grpc.milestones.into_iter().map(Milestone::from).collect(),
        }
    }
}

#[derive(Serialize)]
pub struct Milestone {
    pub id: String,
    pub name: String,
    pub deadline: String,
    pub tasks: Vec<Task>,
    pub num_of_problems: i32,
    pub num_of_tasks: i32,
    pub num_of_finished_tasks: i32,
}

impl From<GrpcMilestone> for Milestone {
    fn from(grpc: GrpcMilestone) -> Self {
        Milestone {
            id: grpc.id,
            name: grpc.name,
            deadline: grpc.deadline,
            tasks: grpc.tasks.into_iter().map(Task::from).collect(),
            num_of_problems: grpc.num_of_problems,
            num_of_tasks: grpc.num_of_tasks,
            num_of_finished_tasks: grpc.num_of_finished_tasks,
        }
    }
}

#[derive(Serialize)]
pub struct Task {
    pub id: String,
    pub name: String,
    pub status: String,
    pub user: Option<User>,
    pub active_period_start: Option<String>,
    pub active_period_end: Option<String>,
    pub problems: Vec<Problem>,
    pub num_of_problems: i32,
    pub is_assigned: bool,
}

impl From<GrpcTask> for Task {
    fn from(grpc: GrpcTask) -> Self {
        Task {
            id: grpc.id,
            name: grpc.name,
            status: grpc.status,
            user: grpc.user.map(User::from),
            active_period_start: grpc.active_period_start,
            active_period_end: grpc.active_period_end,
            problems: grpc.problems.into_iter().map(Problem::from).collect(),
            num_of_problems: grpc.num_of_problems,
            is_assigned: grpc.is_assigned,
        }
    }
}

#[derive(Serialize, Deserialize)]
pub struct Problem {
    pub id: Option<String>,
    pub name: String,
    pub posted_at: String,
}

impl From<GrpcProblem> for Problem {
    fn from(grpc: GrpcProblem) -> Self {
        Problem {
            id: grpc.id,
            name: grpc.name,
            posted_at: grpc.posted_at,
        }
    }
}
impl From<Problem> for GrpcProblem {
    fn from(rest: Problem) -> Self {
        GrpcProblem {
            id: rest.id,
            name: rest.name,
            posted_at: rest.posted_at,
        }
    }
}