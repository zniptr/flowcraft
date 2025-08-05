package executableregistry

import (
	"github.com/zniptr/flowcraft/internal/helpers"
	"github.com/zniptr/flowcraft/pkg/executable"
)

type ExecutableRegistry interface {
	Get(name string) executable.ExecutableFactory
	Register(name string, executable executable.ExecutableFactory)
	Unregister(name string)
}

type executableRegistryImpl struct {
	mutex    helpers.MutexHelper
	registry map[string]func() executable.Executable
}

var (
	instance *executableRegistryImpl

	getMutexFunc = helpers.NewMutexHelper
)

func GetInstance() ExecutableRegistry {
	if instance == nil {
		mutex := getMutexFunc()
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			instance = &executableRegistryImpl{
				registry: make(map[string]func() executable.Executable),
				mutex:    mutex,
			}
		}
	}

	return instance
}

func (registry *executableRegistryImpl) Get(name string) executable.ExecutableFactory {
	registry.mutex.Lock()
	defer registry.mutex.Unlock()
	return registry.registry[name]
}

func (registry *executableRegistryImpl) Register(name string, factory executable.ExecutableFactory) {
	registry.mutex.Lock()
	defer registry.mutex.Unlock()
	registry.registry[name] = factory
}

func (registry *executableRegistryImpl) Unregister(name string) {
	registry.mutex.Lock()
	defer registry.mutex.Unlock()
	delete(registry.registry, name)
}
