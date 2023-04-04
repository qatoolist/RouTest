package models

import (
	"sync"

	"github.com/qatoolist/RouTest/internal/interfaces"
)

func NewHooksRegistry() interfaces.HooksRegistry {
	return &HooksRegistryImpl{
		beforeHooks: make([]interfaces.BeforeHook, 0),
		afterHooks:  make([]interfaces.AfterHook, 0),
		mutex:       &sync.RWMutex{},
	}
}

// HooksRegistry holds a list of hooks that can be executed in a specific order.
type HooksRegistryImpl struct {
	beforeHooks []interfaces.BeforeHook
	afterHooks  []interfaces.AfterHook
	mutex       *sync.RWMutex
}

// RegisterBeforeHook registers a new BeforeHook function to the hooks registry.
func (hr *HooksRegistryImpl) RegisterBeforeHook(hook interfaces.BeforeHook) {
	hr.mutex.Lock()
	defer hr.mutex.Unlock()
	hr.beforeHooks = append(hr.beforeHooks, hook)
}

// RegisterAfterHook registers a new AfterHook function to the hooks registry.
func (hr *HooksRegistryImpl) RegisterAfterHook(hook interfaces.AfterHook) {
	hr.mutex.Lock()
	defer hr.mutex.Unlock()
	hr.afterHooks = append(hr.afterHooks, hook)
}

// RunBeforeHooks executes all the registered BeforeHook functions in the order they were added.
func (hr *HooksRegistryImpl) RunBeforeHooks(route interfaces.Route) (interfaces.Route, error) {
	hr.mutex.RLock()
	defer hr.mutex.RUnlock()
	for _, hook := range hr.beforeHooks {
		var err error
		route, err = hook(route)
		if err != nil {
			return nil, err
		}
	}
	return route, nil
}

// RunAfterHooks executes all the registered AfterHook functions in the order they were added.
func (hr *HooksRegistryImpl) RunAfterHooks(resp interfaces.Response) (interfaces.Response, error) {
	hr.mutex.RLock()
	defer hr.mutex.RUnlock()
	for _, hook := range hr.afterHooks {
		var err error
		resp, err = hook(resp)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}
