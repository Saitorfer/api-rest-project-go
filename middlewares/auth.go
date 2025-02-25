package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"practise.com/rest-api-go/utils"
)

func Authenticate(context *gin.Context) {
	//We ecpect the token to be part of the header request
	token := context.Request.Header.Get("Authorization")

	//fiest lets se if there is a token, if there is no token, we send an error
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Autorized"})
		return
	}
	//now lets se if the token is valid
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Autorized"})
		return
	}

	//if the token is valid, we will set the user id in the context
	context.Set("userId", userId)
	context.Next()
}
