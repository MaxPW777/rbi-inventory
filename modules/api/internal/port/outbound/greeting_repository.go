package outbound

import "rbi/api/internal/core/domain"

// GreetingRepository defines the outbound dependency needed by the core.
type GreetingRepository interface {
	FindGreeting(name string) (domain.Greeting, error)
}
