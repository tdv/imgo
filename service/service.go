package service

// Service interface provides methods for a service.
type Service interface {
	// Start starts service. Calls 'panic' if the problem
	// has been detected while the service was starting.
	Start()
	// Stop stops service. Returns nil on success otherwise error object
	Stop() error
	// Started return true if service is startd otherwise false
	Started() bool
}
