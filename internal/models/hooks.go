package models

import "github.com/qatoolist/RouTest/internal/interfaces"

// BeforeHookFunc is a function type that wraps around BeforeHook and enforces its signature.
type BeforeHookFunc func(*interfaces.Route) (*interfaces.Route, error)

// Execute method executes the BeforeHook function with the given Route as argument.
// It returns the modified Route object and any error occurred during the execution of BeforeHook.
func (bh BeforeHookFunc) Execute(route *interfaces.Route) (*interfaces.Route, error) {
	return bh(route)
}

// AfterHookFunc is a function type that wraps around AfterHook and enforces its signature.
type AfterHookFunc func(*interfaces.Response) (*interfaces.Response, error)

// Execute executes the AfterHook with the given Response as argument.
func (ah AfterHookFunc) Execute(resp *interfaces.Response) (*interfaces.Response, error) {
	return ah(resp)
}

/* Example usage -

package mypackage

import (
    "path/to/types"
    "path/to/interfaces"
)

func main() {
    bhFunc := types.BeforeHookFunc(myBeforeHook)
    route := &interfaces.Route{}
    modifiedRoute, err := bhFunc.Execute(route)
    if err != nil {
        // Handle error
    }
    // Use modifiedRoute
}

func myBeforeHook(route *interfaces.Route) (*interfaces.Route, error) {
    // Implement BeforeHook logic
    return route, nil
}

*/
