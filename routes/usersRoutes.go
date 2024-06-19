package routes

import (
	"RestAPI/Models"
	"RestAPI/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// signup handles the signup functionality by parsing the request JSON body into a User struct,
// saving the user to the database, and returning a JSON response.
// If parsing the request data fails, a bad request response is returned with a corresponding message.
// If saving the user to the database fails, an internal server error response is returned with a corresponding message.
// If all operations are successful, a created response is returned with a corresponding message.
func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// login handles the login functionality by parsing the JSON request body into a User struct.
// It then calls the ValidateCredentials method on the user to check if the credentials are valid.
// If the credentials are valid, a token is generated using the GenerateToken function.
// The generated token is returned in the response along with a success message.
// If any error occurs during parsing, credential validation, or token generation, an appropriate error message is returned in the response.
func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
