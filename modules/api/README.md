# RBI inventory API module

This module bootstraps a tiny HTTP service organised using Hexagonal Architecture. It gives you concrete files in the usual folders so you can expand from here.

## Layout

```
.
├── api.go                  # Application entrypoint that wires the adapters and core
├── internal
│   ├── adapter
│   │   ├── http            # Inbound transport adapters
│   │   └── memory          # Outbound infrastructure adapters
│   ├── core
│   │   ├── app             # Use case coordination
│   │   └── domain          # Enterprise business rules
│   └── port
│       ├── inbound         # Interfaces consumed by adapters
│       └── outbound        # Interfaces implemented by adapters
└── README.md
```

## Flow

1. `main` creates the outbound adapter (`memory.GreetingRepository`).
2. The outbound adapter is injected into the core `application.GreetingService` (use case).
3. The HTTP adapter receives the inbound request and calls the `GreetingService` through the inbound port.
4. The service returns a domain `Greeting`, which the HTTP adapter serialises back to the client.

Each package has a representative file so you can use it as a template for additional features.

## Running the example

```bash
cd modules/api
go run api.go
```

You can then hit `GET /greet/{name}` (for example `curl http://localhost:8080/greet/World`) to see the wiring in action.
