package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not reach events", "error": err})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id", "error": err})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, event)

}

func createEvent(ctx *gin.Context) {

	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	event.ID = 1
	event.UserID = ctx.GetInt64("userId")
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event", "error": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id", "error": err})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event", "error": err})
		return
	}

	userId := ctx.GetInt64("userId")

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to update"})
		return
	}

	var updateEvent models.Event
	err = ctx.ShouldBindJSON(&updateEvent)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	updateEvent.ID = eventId
	updateEvent.UserID = userId
	err = updateEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event Updated Successfully", "event": updateEvent})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id", "error": err})
		return
	}
	userId := ctx.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if event.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to delete"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event", "error": err})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Event has been deleted succesfully"})
}
