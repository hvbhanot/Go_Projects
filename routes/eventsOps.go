package routes

import (
	"RestAPI/Modals"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(ctx *gin.Context) {
	events, err := Modals.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events try again later"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvents(ctx *gin.Context) {
	var event Modals.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the data"})
		return
	}

	err = event.SaveNewEvent()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event, try again later!"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func getEvent(ctx *gin.Context) {
	EventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": ", Could not parse event ID , try again later!"})
	}
	event, err := Modals.GetEventByID(EventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch ID!"})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func updateEvent(ctx *gin.Context) {
	EventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": ", Could not parse event ID , try again later!"})
		return
	}
	_, err = Modals.GetEventByID(EventID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": ", Could not fetch the event"})
		return
	}

	var updatedEvent Modals.Event
	err = ctx.ShouldBindJSON(&updatedEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse the data"})
		return
	}

	updatedEvent.ID = EventID
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}
