package service

// Storage provides methods for load and store some buffer in the storage.
type Storage interface {
	// Put puts a buffer into storage.
	// Returns nil on success otherwise error object.
	Put(id string, buf []byte) error
	// Get returns some buffer and nil as error object.
	// If the data not found, will returned nil data and error object.
	Get(id string) ([]byte, error)
}
