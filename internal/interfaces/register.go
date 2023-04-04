package interfaces

// ResponseMap maps response names to their respective Response structs.
type ResponseMap map[string]*Response

// ParameterMap maps parameter names to their respective Parameter structs.
type ParameterMap map[string]*Parameter

type Register interface {
	AddResponse(name string, resp *Response)
	GetResponse(name string) (*Response, error)
	AddParameter(name string, param *Parameter)
	GetParameter(name string) (*Parameter, error)
}
