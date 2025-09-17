package memory

import (
	"fmt"

	"rbi/api/internal/core/domain"
)

// GreetingRepository is a light-weight outbound adapter backed by an in-memory map.
type GreetingRepository struct {
	messages       map[string]string
	defaultMessage string
}

func NewGreetingRepository() *GreetingRepository {
	return &GreetingRepository{
		messages: map[string]string{
			"Max":   "Welcome back, Max!",
			"World": "Hello, World!",
		},
		defaultMessage: "Hello, %s!",
	}
}

func (r *GreetingRepository) FindGreeting(name string) (domain.Greeting, error) {
	if message, ok := r.messages[name]; ok {
		return domain.NewGreeting(name, message), nil
	}

	return domain.NewGreeting(name, fmt.Sprintf(r.defaultMessage, name)), nil
}
