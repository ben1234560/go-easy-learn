// Package registry is an interface for service discovery
package registry

import (
	"errors"
)

// The registry provides an interface for service discovery
// and an abstraction over varying implementations
// {consul, etcd, zookeeper, ...}
type Registry interface {
	Register(*Service, ...RegisterOption) error
	Deregister(*Service) error
	GetService(string) ([]*Service, error)
	ListServices() ([]*Service, error)
	Watch() (Watcher, error)
	String() string
}

type Option func(*Options)

type RegisterOption func(*RegisterOptions)

var (
	DefaultRegistry = newConsulRegistry()

	ErrNotFound = errors.New("not found")
)

func NewRegistry(opts ...Option) Registry {
	return newConsulRegistry(opts...)
}

// Register a service node. Additionally supply options such as TTL.
func Register(s *Service, opts ...RegisterOption) error {
	return DefaultRegistry.Register(s, opts...)
}

// Deregister a service node
func Deregister(s *Service) error {
	return DefaultRegistry.Deregister(s)
}

// Retrieve a service. A slice is returned since we separate Name/Version.
func GetService(name string) ([]*Service, error) {
	return DefaultRegistry.GetService(name)
}

// List the services. Only returns service names
func ListServices() ([]*Service, error) {
	return DefaultRegistry.ListServices()
}

// Watch returns a watcher which allows you to track updates to the registry.
func Watch() (Watcher, error) {
	return DefaultRegistry.Watch()
}

func String() string {
	return DefaultRegistry.String()
}
