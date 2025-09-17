use crate::core::domain::Greeting;

/// GreetingRepository defines the outbound dependency needed by the core.
pub trait GreetingRepository: Send + Sync {
    fn find_greeting(&self, name: &str) -> Result<Greeting, String>;
}
