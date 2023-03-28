package modedls

type Register struct {
    Responses map[string]*Response
    Parameters   map[string]*Parameter
    Requests map[string]*Request
    Headers      map[string]*Header
}
