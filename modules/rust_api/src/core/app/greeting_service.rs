use crate::core::domain::Greeting;
use crate::port::{inbound::GreetingUseCase, outbound::GreetingRepository};

/// GreetingService orchestrates the greeting use case.
pub struct GreetingService<R> {
    repository: R,
}

impl<R> GreetingService<R>
where
    R: GreetingRepository,
{
    pub fn new(repository: R) -> Self {
        Self { repository }
    }
}

impl<R> GreetingUseCase for GreetingService<R>
where
    R: GreetingRepository,
{
    fn greet(&self, name: &str) -> Result<Greeting, String> {
        if name.is_empty() {
            return Err("name is required".into());
        }

        self.repository.find_greeting(name)
    }
}
