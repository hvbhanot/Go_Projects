package routes

import (
	"RestAPI/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// getEvents retrieves all events from the database and returns them as a JSON response.
// It calls the GetAllEvents function from the models package to fetch the events from the database.
// If an error occurs during the database query, it returns a JSON response with a 500 Internal Server Error status.
// Otherwise, it returns a JSON response with a 200 OK status and the events data.
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

// getEvent retrieves an event from the database based on the provided event ID.
// It parses the event ID from the request URL and calls models.GetEventByID
// to fetch the event from the database. If the event is found, it is returned
// as a JSON response with status code OK (200). If the event ID cannot be parsed
// or an error occurs during the fetching process, an appropriate error message
// is returned as a JSON response with the corresponding status code (BadRequest
// for parsing error, InternalServerError for fetching error).
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

// createEvent creates a new event based on the request data provided in the JSON body.
// It first binds the JSON data to the event struct using ShouldBindJSON(). If there is an error parsing the data,
// it returns a JSON response with a 400 Bad Request status code and an error message.
// If the data is successfully parsed, it retrieves the user ID from the request context.
// It sets the retrieved user ID as the UserID of the event struct.
// Then, it calls the Save() method of the event, which saves the event to the database.
// If there is an error saving the event, it returns a JSON response with a 500 Internal Server Error status code and an error message.
// If the event is successfully saved, it returns a JSON response with a 201 Created status code,
// a success message, and the event details in the response body.
// This function is typically used to handle the creation of events in the application.
func createEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	userId := context.GetInt64("userId")

	event.UserID = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

// updateEvent updates an event's details in the database based on the provided event ID.
// It first parses the event ID from the request parameter. If parsing fails, it returns an error message.
// It retrieves the user ID from the request context and compares it with the event's user ID.
// If the user IDs do not match, it returns an unauthorized error message.
// It fetches the event from the database using the event ID. If fetching fails, it returns an error message.
// It binds the JSON data from the request body to the updatedEvent struct. If binding fails, it returns an error message.
// It assigns the event ID to the updatedEvent struct and updates the event in the database.
// If updating fails, it returns an error message.
// Finally, it returns a success message if the event is updated successfully.
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to update event"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

// deleteEvent deletes an event from the database based on the event ID provided in the URL parameter.
// It first parses the event ID from the URL parameter and checks for any parsing errors.
// If there is an error, it sends an HTTP response with a 400 status code and an error message.
// Then, it retrieves the user ID from the context and fetches the event details from the database
// using the models.GetEventByID function.
// If the event's user ID doesn't match the authenticated user ID, it sends an HTTP response with
// a 401 status code and an error message indicating that the user is not authorized to delete the event.
// If there is an error fetching the event details, it sends an HTTP response with a 500 status code
// and an error message indicating the failure to fetch the event.
// If all checks pass, it calls the Delete method on the event to delete it from the database.
// If there is an error while deleting the event, it sends an HTTP response with a 500 status code
// and an error message indicating the failure to delete the event.
// Finally, if the event is successfully deleted, it sends an HTTP response with a 200 status code
// and a success message indicating the successful deletion of the event.
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to delete event"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event."})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could Not Delete Event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})
}
