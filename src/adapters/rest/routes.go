package rest

func SetRoutes(server GinServer, routesHandler RoutesHandler) {
	r := server.Router

	r.NoRoute(routesHandler.endpointNotFound)

	authenticationEndpoint := r.Group("/auth")

	//This is not oauth, but we need to be able to authenticate with a token
	//this is a private project, so we don't need to worry about security too much

	//fetching a token with a username and password
	authenticationEndpoint.POST("/token", routesHandler.createAuthTokenHandler)
	authenticationEndpoint.POST("/jwt", routesHandler.createJwtTokenHandler)

	authenticatedEndpoint := r.Group("/", routesHandler.verifyAuthenticationMiddleware())
	authenticationEndpoint.DELETE("/jwt", routesHandler.deleteJwtTokenHandler)

	profileEndpoint := authenticatedEndpoint.Group("/me")
	profileEndpoint.GET("", routesHandler.getProfileHandler)
	profileEndpoint.GET("/roles", routesHandler.getRolesHandler)
	profileEndpoint.GET("/permissions", routesHandler.getPermissionsHandler)

}
