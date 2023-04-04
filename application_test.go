package routest

import (
	"testing"

	"github.com/qatoolist/RouTest/internal/models"
)

func TestApplication(t *testing.T) {

	appMetaYaml := `
	assignee: "Jane Doe"
	automation_status: "manual-only"
	component: "shopping cart"
	importance: "medium"
	requirements: "User must be able to add items to the cart and checkout"
	requirements_override: "None"
	setup: "Navigate to the shopping cart page"
	test_steps: "Add items to the cart, go to the checkout page, enter shipping and payment information, and place order"
	expected_results: "Order should be placed successfully and confirmation page should be displayed"
	negative: false
	type: "smoke"
	tags: "cart, checkout, user info"
	`
	mockApp := NewApplication(appMetaYaml)

	infoYaml := `
	name: "Sample API endpoint"
	description: "This is a sample API endpoint for demonstration purposes."
	path: "/sample"
	method: "GET"
	requestBodySchema: |
	  {
	    "type": "object",
	    "properties": {
	      "param1": {
	        "type": "string",
	        "description": "Description for param1"
	      },
	      "param2": {
	        "type": "integer",
	        "description": "Description for param2"
	      }
	    }
	  }
	responseBodySchema: |
	  {
	    "type": "object",
	    "properties": {
	      "result": {
	        "type": "string",
	        "description": "Description for result"
	      }
	    }
	  }
	`

	info := models.NewInfo(infoYaml)

	routeMetaYaml := `
	assignee: "Anand Chavan"
	automation_status: "manual-only"
	component: "shopping cart"
	importance: "medium"
	requirements: "User must be able to add items to the cart and checkout"
	requirements_override: "None"
	setup: "Navigate to the shopping cart page"
	test_steps: "Add items to the cart, go to the checkout page, enter shipping and payment information, and place order"
	expected_results: "Order should be placed successfully and confirmation page should be displayed"
	negative: false
	type: "smoke"
	tags: "cart, checkout, user info"
	`

	route1 := mockApp.NewRoute(info, routeMetaYaml)
	route1_name := route1.GetName()

	mockApp.AddRoute(route1_name, &route1)
	if _, ok := mockApp.GetRouteByName(route1_name); !ok {
		t.Errorf("AddRoute() did not add the route as expected")
	}

	// Test LoadRequirements
	err := mockApp.LoadRequirements("path/to/requirements.yaml")
	if err != nil {
		t.Errorf("LoadRequirements() returned unexpected error: %v", err)
	}

	// Test RegisterResponse
	respName := "myResponse"
	resp, err := models.NewResponse()
	if err != nil {
		panic(err)
	}

	mockApp.RegisterResponse(respName, &resp)
	if _, err := mockApp.GetResponse(respName); err != nil {
		t.Errorf("RegisterResponse() did not register the response as expected")
	}
	/*
		// Test GetResponse
		if r, err := mockApp.GetResponse("nonexistentResponse"); !errors.Is(err, interfaces.ErrResponseNotFound) || r != nil {
			t.Errorf("GetResponse() did not return the expected error")
		}

		// Test RegisterParameter
		paramName := "myParam"
		param := &interfaces.Parameter{}
		mockApp.RegisterParameter(paramName, param)
		if p, err := mockApp.GetParameter(paramName); err != nil || p != param {
			t.Errorf("RegisterParameter() did not register the parameter as expected")
		}

		// Test GetParameter
		if p, err := mockApp.GetParameter("nonexistentParam"); !errors.Is(err, interfaces.ErrParameterNotFound) || p != nil {
			t.Errorf("GetParameter() did not return the expected error")
		}

		// Test GetApplicationParametersRegistry
		if registry := mockApp.GetApplicationParametersRegistry(); registry == nil {
			t.Errorf("GetApplicationParametersRegistry() returned nil")
		}

		// Test GetApplicationHooksRegistry
		if registry := mockApp.GetApplicationHooksRegistry(); registry == nil {
			t.Errorf("GetApplicationHooksRegistry() returned nil")
		}

		// Test GetMeta
		if meta := mockApp.GetMeta(); meta == nil {
			t.Errorf("GetMeta() returned nil")
		}
	*/
}
