package routes

import (
	models "RestAPI/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// registerForEvents is a handler function that registers a user for a specific event.
// It expects a `userId` parameter to be set in the request context and an `id` parameter
// in the URL path which represents the event id. It fetches the event by the provided id,
// registers the user for the event, and returns a success message or an error message if any
// error occurs during the process.
func registerForEvents(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Registered"})
}

// cancelRegistration cancels the registration of a user for an event.
// It retrieves the userId from the context and the eventId from the URL parameter.
// If the eventId cannot be parsed, it returns an error response.
// It creates a new Event struct with the retrieved eventId.
// It then calls the CancelRegistration method on the event, passing the userId, to cancel the registration.
// If the cancellation fails, it returns an error response.
// Finally, it returns a success response indicating the event registration was cancelled successfully.
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Registration Cancelled"})
}
