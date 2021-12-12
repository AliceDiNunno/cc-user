package rest

import (
	"github.com/AliceDiNunno/cc-user/src/adapters/rest/request"
	e "github.com/AliceDiNunno/go-nested-traced-error"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (rH RoutesHandler) verifyAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			rH.handleError(c, e.Wrap(ErrAuthorizationHeaderMissing))
			return
		}

		payload, err := rH.usecases.CheckJwtToken(authorizationHeader)

		if err != nil {
			rH.handleError(c, err.Append(ErrInvalidAuthorizationHeader))
			return
		}

		c.Set("userID", payload.UserID)
	}
}

func (rH RoutesHandler) createAuthTokenHandler(c *gin.Context) {
	var tokenRequest request.AccessTokenRequest

	if stderr := c.ShouldBindJSON(&tokenRequest); stderr != nil {
		rH.handleError(c, e.Wrap(stderr).Append(ErrFormValidation))
		return
	}

	token, err := rH.usecases.CreateAuthToken(tokenRequest.ToDomain())

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, token)
}

func (rH RoutesHandler) createJwtTokenHandler(c *gin.Context) {
	var jwtRequest request.JwtTokenRequest

	if stderr := c.ShouldBindJSON(&jwtRequest); stderr != nil {
		rH.handleError(c, e.Wrap(stderr).Append(ErrFormValidation))
		return
	}

	token, err := rH.usecases.CreateJwtToken(jwtRequest.ToDomain())

	if err != nil {
		rH.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, token)
}

func (rH RoutesHandler) deleteJwtTokenHandler(c *gin.Context) {

}
