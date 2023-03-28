package models

type Scenario struct {
	Meta                 Meta
	RequestBodyTemplate  string
	ResponseBodyTemplate string
	BeforeHook           BeforeHook
	AfterHook            AfterHook
}
