use serde::{Deserialize, Serialize};
use crate::proto_generated::{
    AuthResponse as GrpcAuthResponse,
    LoginRequest as GrpcLoginRequest,
    RegisterRequest as GrpcRegisterRequest,
    User as GrpcUser,
    FriendsResponse as GrpcFriendsResponse,
    FriendEditRequest as GrpcFriendEditRequest
};

// REST
#[derive(Deserialize)]
pub struct LoginRequest {
    pub email: String,
    pub password: String,
}
// REST -> gRPC
impl From<LoginRequest> for GrpcLoginRequest {
    fn from(rest: LoginRequest) -> Self {
        GrpcLoginRequest {
            email: rest.email,
            password: rest.password,
        }
    }
}

// REST
#[derive(Deserialize)]
pub struct RegisterRequest {
    pub email: String,
    pub password: String,
    pub first_name: String,
    pub last_name: String,
}
// REST -> gRPC
impl From<RegisterRequest> for GrpcRegisterRequest {
    fn from(rest: RegisterRequest) -> Self {
        GrpcRegisterRequest {
            email: rest.email,
            password: rest.password,
            first_name: rest.first_name,
            last_name: rest.last_name
        }
    }
}

// REST
#[derive(Serialize)]
pub struct User {
    pub first_name: String,
    pub last_name: String,
    pub id: String
}
// GRPC -> REST
impl From<GrpcUser> for User {
    fn from(grpc: GrpcUser) -> Self {
        User {
            first_name: grpc.first_name,
            last_name: grpc.last_name,
            id: grpc.id
        }
    }
}

// REST
#[derive(Serialize)]
pub struct AuthResponse {
    pub valid: bool,
    pub token: String,
    pub user: User
}
// GRPC -> REST
impl From<GrpcAuthResponse> for AuthResponse {
    fn from(grpc: GrpcAuthResponse) -> Self {
        AuthResponse {
            valid: grpc.valid,
            token: grpc.token,
            user: grpc.user
            .map(|u| u.into())
            .unwrap_or_else(|| User {
                first_name: String::new(),
                last_name: String::new(),
                id: String::new(),
            }),
        }
    }
}


// REST

#[derive(Serialize)]
pub struct FriendsResponse {
    pub users: Vec<User>,
}

// GRPC -> REST
impl From<GrpcFriendsResponse> for FriendsResponse {
    fn from(grpc: GrpcFriendsResponse) -> Self {
        let users = grpc.users.into_iter().map(|user| user.into()).collect();

        FriendsResponse {
            users,
        }
    }
}

// REST

#[derive(Deserialize)]
pub struct FriendEditRequest {
    pub friend_id: String,
}

// REST -> gRPC
impl From<FriendEditRequest> for GrpcFriendEditRequest {
    fn from(rest: FriendEditRequest) -> Self {
        GrpcFriendEditRequest {
            user_id: String::new(),
            friend_id: rest.friend_id,
        }
    }
}