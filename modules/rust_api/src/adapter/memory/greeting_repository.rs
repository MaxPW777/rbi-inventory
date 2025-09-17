use std::collections::HashMap;

use crate::core::domain::Greeting;
use crate::port::outbound::GreetingRepository as GreetingRepositoryPort;

/// GreetingRepository is a light-weight outbound adapter backed by an in-memory map.
pub struct GreetingRepository {
    messages: HashMap<String, String>,
    default_message: String,
}

impl GreetingRepository {
    pub fn new() -> Self {
        let messages = HashMap::from([
            ("Max".to_string(), "Welcome back, Max!".to_string()),
            ("World".to_string(), "Hello, World!".to_string()),
        ]);

        Self {
            messages,
            default_message: "Hello, {}!".to_string(),
        }
    }
}

impl GreetingRepositoryPort for GreetingRepository {
    fn find_greeting(&self, name: &str) -> Result<Greeting, String> {
        if let Some(message) = self.messages.get(name) {
            return Ok(Greeting::new(name, message));
        }

        Ok(Greeting::new(
            name,
            self.default_message.replace("{}", name),
        ))
    }
}
