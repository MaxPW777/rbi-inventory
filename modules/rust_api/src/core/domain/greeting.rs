use serde::Serialize;

/// Greeting captures the core data returned to callers.
#[derive(Debug, Clone, Serialize)]
pub struct Greeting {
    pub name: String,
    pub message: String,
}

impl Greeting {
    /// Builds a greeting for the provided name.
    pub fn new<N, M>(name: N, message: M) -> Self
    where
        N: Into<String>,
        M: Into<String>,
    {
        Self {
            name: name.into(),
            message: message.into(),
        }
    }
}
