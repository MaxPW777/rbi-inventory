use std::sync::Arc;

use axum::{
    extract::{Path, State},
    http::StatusCode,
    response::IntoResponse,
    routing::get,
    Json, Router,
};
use serde::Serialize;

use crate::port::inbound::GreetingUseCase;

/// GreetingHandler wires the inbound port to an HTTP transport.
#[derive(Clone)]
pub struct GreetingHandler {
    service: Arc<dyn GreetingUseCase>,
}

impl GreetingHandler {
    pub fn new(service: Arc<dyn GreetingUseCase>) -> Self {
        Self { service }
    }

    pub fn register_routes(&self, router: Router<()>) -> Router<Arc<dyn GreetingUseCase>> {
        router.route("/greet/:name", get(Self::handle_greet))
    }

    async fn handle_greet(
        State(service): State<Arc<dyn GreetingUseCase>>,
        Path(name): Path<String>,
    ) -> impl IntoResponse {
        match service.greet(&name) {
            Ok(greeting) => (StatusCode::OK, Json(greeting)).into_response(),
            Err(err) => {
                (StatusCode::BAD_REQUEST, Json(ErrorResponse { error: err })).into_response()
            }
        }
    }
}

#[derive(Serialize)]
struct ErrorResponse {
    error: String,
}
