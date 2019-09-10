// Service is the images service package (ImGo)
package service

// Converter provides method for converting input images into common format.
type Converter interface {
	// Convert converting images into common format.
	// Returns new image, image id and error if the error happened by image converting,
	// otherwise nil, empty string as id and error object.
	Convert(buf []byte, format string, width *int, height *int) ([]byte, string, error)
}
