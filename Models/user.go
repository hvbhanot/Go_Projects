package models

import (
	"RestAPI/db"
	"RestAPI/utils"
	"errors"
)

// User represents a user with an ID, email, and password.
// - ID: The unique identifier of the user.
// - Email: The email address of the user. (required)
// - Password: The password of the user. (required)
type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Save saves the user to the database by inserting a new row into the "users" table.
// It hashes the user's password using the HashPassword function from the utils package.
// The user's email and hashed password are then used as parameters for the database query.
// If the query execution is successful, the last inserted row ID is retrieved and assigned to the user's ID.
// If any error occurs during the preparation, execution, or retrieval of the query, it is returned.
// Otherwise, nil is returned.
func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	HashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, HashedPassword)
	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()
	u.ID = userID

	return err
}

// ValidateCredentials validates the credentials of a User by checking if the email exists in the database
// and if the password matches the hashed password stored in the database.
// It performs a SELECT query to retrieve the user's ID and hashed password from the database based on the email.
// If the query fails or the password is invalid, it returns an error indicating that the credentials are invalid.
// Otherwise, it returns nil to indicate that the credentials are valid.
//
// The ValidateCredentials method should be called on an instance of the User struct.
// It relies on the global variable db.DB, which is an instance of *sql.DB that is used for database operations.
// It also uses the utils.CheckPasswordHash function to check if the provided password matches the hashed password.
//
// Example usage:
//
//	var user models.User
//	err := context.ShouldBindJSON(&user)
//	if err != nil {
//	    context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
//	    return
//	}
//	err = user.ValidateCredentials()
//	if err != nil {
//	    context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
//	    return
//	}
//	token, err := utils.GenerateToken(user.Email, user.ID)
//	if err != nil {
//	    context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
//	    return
//	}
//	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
