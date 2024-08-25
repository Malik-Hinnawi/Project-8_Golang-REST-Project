package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
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
	err = event.Register(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for an event."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

func deleteRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id", "error": err})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete user for an event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
