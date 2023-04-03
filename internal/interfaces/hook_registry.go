package interfaces

type HooksRegistry interface {
	RegisterBeforeHook(hook BeforeHook)
	RegisterAfterHook(hook AfterHook)
	RunBeforeHooks(route Route) (Route, error)
	RunAfterHooks(resp Response) (Response, error)
}
