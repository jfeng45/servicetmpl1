// package container use dependency injection to create concrete type and wire the whole application together
package container

type Container interface {
	// BuildUseCase creates concrete types for use case and it is included types.
	// For each call, it will create a new instance, which means it is not a singleton
	// Only exceptions are data store handlers, which are singletons. They are cached in container.
	BuildUseCase(code string) (interface{}, error)

	// This should only be used by container and it's sub-package
	// Get instance by code from container. Only data store handler can be retrieved from container
	Get(code string) (interface{}, bool)

	// This should only be used by container and it's sub-package
	// Put value into container with code as the key. Only data store handler is saved in container
	Put(code string, value interface{})
}
