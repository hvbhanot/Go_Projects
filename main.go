package main

import (
	"RestAPI/Modals"
	"RestAPI/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents) //GET , POST, PUT , PATCH, DELETE
	server.POST("/events", createEvents)
	err := server.Run(":8080") // localhost:8080
	if err != nil {
		fmt.Println(err)
		return
	}
}

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
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Couldn't parse the data"})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.SaveNewEvent()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event, Try again later !!"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"Message": "Event created", "Event": event})
}
