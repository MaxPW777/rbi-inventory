package inbound

import "rbi/api/internal/core/domain"

// GreetingUseCase exposes the inbound port consumed by the delivery adapters.
type GreetingUseCase interface {
	Greet(name string) (domain.Greeting, error)
}
