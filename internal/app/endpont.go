package app

type Endpoint[Services any] struct {
	Method  string
	Route   string
	Handler Handler[Services]
}
