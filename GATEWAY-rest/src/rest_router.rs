use axum::{
    Router,
    routing::get,
    extract::{Path, Query},
};
use url::Url;
use std::collections::HashMap;

pub fn create_router() -> Router {
    Router::new().route("/*path", get(handle_dynamic_path))
}

async fn handle_dynamic_path(Path(path): Path<String>, Query(query_params): Query<HashMap<String, String>>) -> String {
    // localhost staat er als dummy. Enkel path is belangrijk
    let url_string: String = format!("http://localhost/{}", path);
    let url_obj: Url = Url::parse(&url_string).expect("Failed to parse URL");

    // path segments
    let path_segments: Vec<String> = url_obj.path_segments()
        .map(|segments| segments.map(|s| s.to_string()).collect::<Vec<String>>())
        .unwrap_or_default();

    format_response(path_segments, query_params)
}
    
fn format_response(path_segments: Vec<String>, query_params: HashMap<String, String>) -> String {
    
    let mut response: String = "".to_string();

    response.push_str("Segments: \n");
    for segment in path_segments {
        response.push_str(format!("   - {}\n", segment).as_str());
    }
    response.push_str("Params: \n");

    for (param_key, param_val) in query_params {
        response.push_str(format!("   - {} = {}\n",param_key, param_val).as_str());
    }
    
    response
}
