package rest

import "github.com/AliceDiNunno/cc-user/src/core/usecases"

type RoutesHandler struct {
	usecases usecases.Usecases
}

func NewRouter(ucHandler usecases.Usecases) RoutesHandler {
	return RoutesHandler{usecases: ucHandler}
}
