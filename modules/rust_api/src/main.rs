mod adapter;
mod core;
mod port;

use std::sync::Arc;

use adapter::{http::GreetingHandler, memory::GreetingRepository};
use axum::Router;
use core::app::GreetingService;
use port::inbound::GreetingUseCase;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let repository = GreetingRepository::new();
    let service: Arc<dyn GreetingUseCase> = Arc::new(GreetingService::new(repository));
    let handler = GreetingHandler::new(service);

    let app = handler.register_routes(Router::new());

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await?;
    println!("listening on http://0.0.0.0:3000");
    axum::serve(listener, app.into_make_service()).await?;

    Ok(())
}
