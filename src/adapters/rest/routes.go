package rest

func SetRoutes(server GinServer, routesHandler RoutesHandler) {
	r := server.Router

	r.NoRoute(routesHandler.endpointNotFound)
}
