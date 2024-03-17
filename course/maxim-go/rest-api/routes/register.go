package routes

import (
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	userID := ctx.GetInt64("userID")
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse id."})
		return
	}

	event, err := models.GetEvent(eventID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func cancelRegistration(ctx *gin.Context) {
	userID := ctx.GetInt64("userID")
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse id."})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration of user for event."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}
