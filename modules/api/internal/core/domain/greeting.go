package domain

// Greeting captures the core data returned to callers.
type Greeting struct {
	Name    string
	Message string
}

// NewGreeting builds a greeting for the provided name.
func NewGreeting(name, message string) Greeting {
	return Greeting{Name: name, Message: message}
}
