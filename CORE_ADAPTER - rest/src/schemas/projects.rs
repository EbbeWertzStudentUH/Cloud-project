use serde::{Deserialize, Serialize};
use crate::proto_generated::{
    ProjectCreateRequest as GRPCProjectCreateRequest,
    ProjectId as GRPCProjectId,
    Project as GRPCProject,
    ProjectsList as GRPCProjectsList,
    AddUserToProjectRequest as GRPCAddUserToProjectRequest,
    MilestoneAddRequest as GRPCMilestoneAddRequest,
    TaskAddRequest as GRPCTaskAddRequest,
    ProblemAddRequest as GRPCProblemAddRequest,
    ResolveProblemRequest as GRPCResolveProblemRequest,
    TaskAssignRequest as GRPCTaskAssignRequest,
    TaskCompleteRequest as GRPCTaskCompleteRequest,
    Milestone as GRPCMilestone,
    Task as GRPCTask,
    User as GRPCUser,
    Problem as GRPCProblem,
};

// REST
#[derive(Deserialize)]
pub struct ProjectCreateRequest {
    pub name: String,
    pub deadline: String,
    pub github_repo: String,
}

// REST -> gRPC
impl From<ProjectCreateRequest> for GRPCProjectCreateRequest {
    fn from(rest: ProjectCreateRequest) -> Self {
        GRPCProjectCreateRequest {
            user_id: String::new(),
            name: rest.name,
            deadline: rest.deadline,
            github_repo: rest.github_repo,
        }
    }
}

// REST
#[derive(Serialize)]
pub struct Project {
    pub id: String,
    pub name: String,
    pub users: Vec<User>,
    pub deadline: String,
    pub github_repo: String,
    pub milestones: Vec<Milestone>,
}

// gRPC -> REST
impl From<GRPCProject> for Project {
    fn from(grpc: GRPCProject) -> Self {
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

// REST
#[derive(Serialize, Deserialize)]
pub struct Milestone {
    pub id: Option<String>,
    pub name: String,
    pub deadline: String,
    pub tasks: Vec<Task>,
}

// gRPC -> REST
impl From<GRPCMilestone> for Milestone {
    fn from(grpc: GRPCMilestone) -> Self {
        Milestone {
            id: grpc.id,
            name: grpc.name,
            deadline: grpc.deadline,
            tasks: grpc.tasks.into_iter().map(Task::from).collect(),
        }
    }
}

// REST -> gRPC
impl From<Milestone> for GRPCMilestone {
    fn from(rest: Milestone) -> Self {
        GRPCMilestone {
            id: None,
            name: rest.name,
            deadline: rest.deadline,
            tasks: rest.tasks.into_iter().map(GRPCTask::from).collect(),
        }
    }
}

// REST
#[derive(Serialize, Deserialize)]
pub struct Task {
    pub id: Option<String>,
    pub name: String,
    pub status: Option<String>,
    pub user: Option<User>,
    pub active_period_start: Option<String>,
    pub active_period_end: Option<String>,
    pub problems: Vec<Problem>,
}

// gRPC -> REST
impl From<GRPCTask> for Task {
    fn from(grpc: GRPCTask) -> Self {
        Task {
            id: grpc.id,
            name: grpc.name,
            status: grpc.status,
            user: grpc.user.map(User::from),
            active_period_start: grpc.active_period_start,
            active_period_end: grpc.active_period_end,
            problems: grpc.problems.into_iter().map(Problem::from).collect(),
        }
    }
}

// REST -> gRPC
impl From<Task> for GRPCTask {
    fn from(rest: Task) -> Self {
        GRPCTask {
            id: None,
            name: rest.name,
            status: None,
            user: rest.user.map(GRPCUser::from),
            active_period_start: None,
            active_period_end: None,
            problems: rest.problems.into_iter().map(GRPCProblem::from).collect(),
        }
    }
}

// REST
#[derive(Serialize, Deserialize)]
pub struct User {
    pub first_name: String,
    pub last_name: String,
    pub id: String,
}

// gRPC -> REST
impl From<GRPCUser> for User {
    fn from(grpc: GRPCUser) -> Self {
        User {
            first_name: grpc.first_name,
            last_name: grpc.last_name,
            id: grpc.id,
        }
    }
}

// REST -> gRPC
impl From<User> for GRPCUser {
    fn from(rest: User) -> Self {
        GRPCUser {
            first_name: rest.first_name,
            last_name: rest.last_name,
            id: rest.id,
        }
    }
}

// REST
#[derive(Serialize, Deserialize)]
pub struct Problem {
    pub id: Option<String>,
    pub name: String,
    pub posted_at: String,
}

// gRPC -> REST
impl From<GRPCProblem> for Problem {
    fn from(grpc: GRPCProblem) -> Self {
        Problem {
            id: grpc.id,
            name: grpc.name,
            posted_at: grpc.posted_at,
        }
    }
}

// REST -> gRPC
impl From<Problem> for GRPCProblem {
    fn from(rest: Problem) -> Self {
        GRPCProblem {
            id: None,
            name: rest.name,
            posted_at: rest.posted_at,
        }
    }
}

// REST
#[derive(Deserialize)]
pub struct AddUserToProjectRequest {
    pub project_id: String,
}

// REST -> gRPC
impl From<AddUserToProjectRequest> for GRPCAddUserToProjectRequest {
    fn from(rest: AddUserToProjectRequest) -> Self {
        GRPCAddUserToProjectRequest {
            project_id: rest.project_id,
            user_id: String::new(),
        }
    }
}

// REST
#[derive(Deserialize)]
pub struct TaskAddRequest {
    pub project_id: String,
    pub milestone_id: String,
    pub task: Task,
}

// REST -> gRPC
impl From<TaskAddRequest> for GRPCTaskAddRequest {
    fn from(rest: TaskAddRequest) -> Self {
        GRPCTaskAddRequest {
            project_id: rest.project_id,
            milestone_id: rest.milestone_id,
            task: Some(GRPCTask::from(rest.task)),
        }
    }
}

// REST
#[derive(Deserialize)]
pub struct ProblemAddRequest {
    pub project_id: String,
    pub task_id: String,
    pub problem: Problem,
}

// REST -> gRPC
impl From<ProblemAddRequest> for GRPCProblemAddRequest {
    fn from(rest: ProblemAddRequest) -> Self {
        GRPCProblemAddRequest {
            project_id: rest.project_id,
            task_id: rest.task_id,
            problem: Some(GRPCProblem::from(rest.problem)),
        }
    }
}

// REST
#[derive(Deserialize)]
pub struct ResolveProblemRequest {
    pub project_id: String,
    pub task_id: String,
    pub problem_id: String,
}

// REST -> gRPC
impl From<ResolveProblemRequest> for GRPCResolveProblemRequest {
    fn from(rest: ResolveProblemRequest) -> Self {
        GRPCResolveProblemRequest {
            project_id: rest.project_id,
            task_id: rest.task_id,
            problem_id: rest.problem_id,
        }
    }
}

// REST
#[derive(Deserialize)]
pub struct TaskAssignRequest {
    pub project_id: String,
    pub task_id: String,
}

// REST -> gRPC
impl From<TaskAssignRequest> for GRPCTaskAssignRequest {
    fn from(rest: TaskAssignRequest) -> Self {
        GRPCTaskAssignRequest {
            project_id: rest.project_id,
            task_id: rest.task_id,
            user_id: String::new(),
        }
    }
}

// REST
#[derive(Deserialize)]
pub struct TaskCompleteRequest {
    pub project_id: String,
    pub task_id: String,
}

// REST -> gRPC
impl From<TaskCompleteRequest> for GRPCTaskCompleteRequest {
    fn from(rest: TaskCompleteRequest) -> Self {
        GRPCTaskCompleteRequest {
            project_id: rest.project_id,
            task_id: rest.task_id,
        }
    }
}

// REST
#[derive(Deserialize)]
pub struct MilestoneAddRequest {
    pub project_id: String,
    pub milestone: Milestone,
}

// REST -> gRPC
impl From<MilestoneAddRequest> for GRPCMilestoneAddRequest {
    fn from(rest: MilestoneAddRequest) -> Self {
        GRPCMilestoneAddRequest {
            project_id: rest.project_id,
            milestone: Some(GRPCMilestone::from(rest.milestone)),
        }
    }
}

// REST
#[derive(Serialize)]
pub struct ProjectsList {
    pub projects: Vec<Project>,
}

// gRPC -> REST
impl From<GRPCProjectsList> for ProjectsList {
    fn from(grpc: GRPCProjectsList) -> Self {
        ProjectsList {
            projects: grpc.projects.into_iter().map(Project::from).collect(),
        }
    }
}

// REST
#[derive(Deserialize)]
pub struct ProjectId {
    pub project_id: String,
}

// REST -> gRPC
impl From<ProjectId> for GRPCProjectId {
    fn from(rest: ProjectId) -> Self {
        GRPCProjectId {
            project_id: rest.project_id,
        }
    }
}
