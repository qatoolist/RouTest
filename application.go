package routest

import (
	"github.com/qatoolist/RouTest/internal/interfaces"
	"github.com/qatoolist/RouTest/internal/loaders"
	"github.com/qatoolist/RouTest/internal/models"
)

type application struct {
	app *models.Application
}

func NewApplication(meta string) interfaces.Application {

	env, err := loaders.LoadRoutestsEnv()

	if err != nil {
		panic(err)
	}
	configDir := "./config"

	loader, err := loaders.ConfigLoaderFactory(env, configDir)
	if err != nil {
		panic(err)
	}

	new_config, err := loader.LoadConfig(env, configDir)
	if err != nil {
		panic(err)
	}

	config := models.NewConfig()
	config.CopyFromTemp(new_config)

	requirements_path := "./config/requirements.yaml"

	requirements, err := models.NewRequirementsFromPath(requirements_path)
	if err != nil {
		panic(err)
	}

	new_meta, err := models.NewMetaFromString(meta)
	if err != nil {
		panic(err)
	}

	host, err := config.GetHost()
	if err != nil {
		panic(err)
	}

	app, err := models.NewApplication(env, config, requirements, new_meta, host)
	if err != nil {
		panic(err)
	}

	return &application{app: app}
}

func (a *application) NewRoute(info interfaces.Info, meta string) interfaces.Route {
	route := a.app.NewRoute(info, meta)
	return route
}

func (a *application) GetRouteByName(name string) (interfaces.Route, bool) {
	return a.app.GetRouteByName(name)
}

func (a *application) AddRoute(name string, route *interfaces.Route) interfaces.Route {
	return a.app.AddRoute(name, route)
}

func (a *application) LoadRequirements(configPath string) error {
	return a.app.LoadRequirements(configPath)
}

func (a *application) RegisterResponse(name string, resp *interfaces.Response) {
	a.app.RegisterResponse(name, resp)
}

func (a *application) GetResponse(name string) (*interfaces.Response, error) {
	return a.app.GetResponse(name)
}

func (a *application) RegisterParameter(name string, param *interfaces.Parameter) {
	a.app.RegisterParameter(name, param)
}

func (a *application) GetParameter(name string) (*interfaces.Parameter, error) {
	return a.app.GetParameter(name)
}

func (a *application) GetApplicationParametersRegistry() interfaces.ParametersRegistry {
	return a.app.GetApplicationParametersRegistry()
}

func (a *application) GetApplicationHooksRegistry() interfaces.HooksRegistry {
	return a.app.GetApplicationHooksRegistry()
}

func (a *application) GetMeta() interfaces.Meta {
	return a.app.GetMeta()
}
