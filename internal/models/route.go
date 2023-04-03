package models

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/qatoolist/RouTest/internal/interfaces"
)

// Route represents a custom HTTP request with additional fields.
type Route struct {
	// Info contains the request information.
	Info Info

	// ParentRoute provides the reference to the parent route.
	ParentApplication interfaces.Application

	// Meta contains the metadata about the route.
	Meta interfaces.Meta

	// Scenario contains the scenario that should be executed for the route.
	ScenarioRegistry interfaces.ScenarioRegistry

	// RouteParametersRegistry are the route level parameters and
	// These Parameters are available through all scenarios registered under this route
	// But can be overriden by providing the scenario level parameters having same keys.
	RouteParametersRegistry interfaces.ParametersRegistry

	// RouteHooksRegistry is a registry of Before and After Hooks defined at route level
	// The Before Hooks are triggered for every scenario of every route defined under this application
	// The Order of Hooks Execution is -  BeforeApplicationHooks, BeforeRouteHooks, BeforeScenarioHooks
	RouteHooksRegistry interfaces.HooksRegistry

	// Body is the request body.
	Body []byte

	// Response is the channel for the HTTP response.
	Response <-chan *http.Response
}

// SetReqBodySchema sets the request body schema for the route.
func (r *Route) SetReqBodySchema(schema string) error {
	reqBodySchema, err := NewRequestBodySchema(schema)
	if err != nil {
		return err
	}
	r.Info.SetRequestBodySchema(reqBodySchema)
	return nil
}

// SetResBodySchema sets the response body schema for the route.
func (r *Route) SetResBodySchema(schema string) error {
	resBodySchema, err := NewResponseBodySchema(schema)
	if err != nil {
		return err
	}
	r.Info.SetResponseBodySchema(resBodySchema)
	return nil
}

// ValidateReqBody validates the request body against the request body schema.
func (r *Route) ValidateReqBody(body interface{}) error {
	if r.Info.GetRequestBodySchema() == interfaces.RequestBodySchema(&RequestBodySchema{}) {
		return nil
	}
	return r.Info.GetRequestBodySchema().Validate(body)
}

// ValidateResBody validates the response body against the response body schema.
func (r *Route) ValidateResBody(body interface{}) error {
	if r.Info.GetResponseBodySchema() == interfaces.ResponseBodySchema(&ResponseBodySchema{}) {
		return nil
	}
	return r.Info.GetResponseBodySchema().Validate(body)
}

// Send sends the HTTP request and returns the HTTP response.
func (r *Route) Send() (*http.Response, error) {
	client := http.DefaultClient
	req, err := http.NewRequest(r.Info.GetMethod().String(), r.Info.GetPath(), bytes.NewReader(r.Body))
	if err != nil {
		return nil, err
	}

	if r.Info.GetRequestBodySchema() != interfaces.RequestBodySchema(&RequestBodySchema{}) {
		err := r.ValidateReqBody(req.Body)
		if err != nil {
			return nil, err
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if r.Info.GetResponseBodySchema() != interfaces.ResponseBodySchema(&ResponseBodySchema{}) {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = r.ValidateResBody(body)
		if err != nil {
			return nil, err
		}
		resp.Body = ioutil.NopCloser(bytes.NewReader(body))
	}
	return resp, nil
}

func (r *Route) GetScenarioRegistry() interfaces.ScenarioRegistry {
	return r.ScenarioRegistry
}

func (r *Route) GetRouteParametersRegistry() interfaces.ParametersRegistry {
	return r.RouteParametersRegistry
}

func (r *Route) GetRouteHooksRegistry() interfaces.HooksRegistry {
	return r.RouteHooksRegistry
}

func (r *Route) GetParentApplication() interfaces.Application {
	return r.ParentApplication
}

func (r *Route) GetName() string {
	return r.Info.GetName()
}
