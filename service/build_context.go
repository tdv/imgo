package service

// BuildContext interface provides methods for getting configuration
// and requery other entities.
// The interface is send into entity creator function, which
// has been registered by RegisterEntity.
// It happens when the entity is created by the application builder.
type BuildContext interface {
	// GetConfig returns configuration branch for the entity.
	GetConfig() Config
	// GetEntity returns the dependent entity by id.
	GetEntity(id string) interface{}
}
