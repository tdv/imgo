// Service is the images service package (ImGo)
package service

// Config interface provides methods for getting values from
// configuration branch or get other configuration branch.
// In the all methods the path parameter might be a complex,
// and his parts must be separated by dot.
type Config interface {
	// GetStrVal returns a string value.
	GetStrVal(path string) string
	// GetIntVal returns an int value.
	GetIntVal(path string) int
	// GetBranch returns an other branch.
	GetBranch(path string) Config
}
