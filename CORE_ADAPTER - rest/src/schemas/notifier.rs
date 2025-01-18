use serde::{Deserialize, Serialize};
use crate::proto_generated::{
    ProjectSubscribeRequest as GRPCProjectSubscribeRequest
};

// REST
#[derive(Serialize)]
pub struct SimpleResponse {
    pub message: String,
}


// REST
#[derive(Deserialize)]
pub struct ProjectSubscribeRequest {
    pub subscribe_project: String,
    pub unsubscribe_project: Option<String>,
    pub user_id: String,
}
// REST -> gRPC
impl From<ProjectSubscribeRequest> for GRPCProjectSubscribeRequest {
    fn from(rest: ProjectSubscribeRequest) -> Self {
        GRPCProjectSubscribeRequest {
            subscribe_project: rest.subscribe_project,
            unsubscribe_project: rest.unsubscribe_project,
            user_id: rest.user_id,
        }
    }
}