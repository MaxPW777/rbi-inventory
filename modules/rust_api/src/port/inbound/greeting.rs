use crate::core::domain::Greeting;

/// GreetingUseCase exposes the inbound port consumed by delivery adapters.
pub trait GreetingUseCase: Send + Sync {
    fn greet(&self, name: &str) -> Result<Greeting, String>;
}
