package application

import (
	"errors"

	"rbi/api/internal/core/domain"
	"rbi/api/internal/port/outbound"
)

// GreetingService orchestrates the greeting use case.
type GreetingService struct {
	repository outbound.GreetingRepository
}

func NewGreetingService(repo outbound.GreetingRepository) *GreetingService {
	return &GreetingService{repository: repo}
}

func (s *GreetingService) Greet(name string) (domain.Greeting, error) {
	if name == "" {
		return domain.Greeting{}, errors.New("name is required")
	}

	greeting, err := s.repository.FindGreeting(name)
	if err != nil {
		return domain.Greeting{}, err
	}

	return greeting, nil
}
