package middlewares

import (
	"RestAPI/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Authenticate function takes a gin.Context as input and performs token authentication.
// It retrieves the token from the "Authorization" header of the HTTP request.
// If the token is empty, it aborts the request with a Unauthorized response.
// It then calls the VerifyToken function from utils package to validate the token and extract the userId.
// If token verification fails, it aborts the request with a Unauthorized response.
// If token verification succeeds, it sets the "userId" key in the request context with the extracted userId.
// Finally, it calls the Next method of the gin.Context to proceed to the next middleware or handler.
func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
