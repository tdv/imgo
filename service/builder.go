// Service is the images service package (ImGo)
package service

// Builder interface provides method for getting the service instance.
type Builder interface {
	// Build - build the application and returns service interface
	Build() (interface{}, error)
}
