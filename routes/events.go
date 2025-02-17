package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"practise.com/rest-api-go/models"
	"practise.com/rest-api-go/utils"
)

func GetEvents(context *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, try again later"})
		return
	}
	//we want to send back a JSON response
	//first we set the code (200=sucess) and the content
	//gin.H give us back a map, god for encode
	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {
	//get the id
	eventIdInt64, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventIdInt64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event, try again later"})
		return
	}

	context.JSON(http.StatusOK, event)

}

func CreateEvent(context *gin.Context) {
	//We ecpect the token to be part of the header request
	token := context.Request.Header.Get("Authorization")

	//fiest lets se if there is a token, if there is no token, we send an error
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Autorized"})
		return
	}
	//now lets se if the token is valid
	err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Autorized"})
		return
	}

	var event models.Event
	//if the structure and JSON we receive from FE are the same, this method do all the work to parse it
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	//TODO GENERATE ID
	event.ID = 1
	event.UserID = 1

	//save event
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event, try again later"})
		return
	}
	//Send back a created message
	context.JSON(http.StatusCreated, gin.H{"message": "event created"})
}

func UpdateEvent(context *gin.Context) {
	//get the id
	eventIdInt64, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	_, err = models.GetEventById(eventIdInt64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event by id"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	updatedEvent.ID = eventIdInt64

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event by id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func DeleteEvent(context *gin.Context) {
	//get the id
	eventIdInt64, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	//Be sure event exist and get it

	event, err := models.GetEventById(eventIdInt64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event by id"})
		return
	}

	//delete the event

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event by id"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}
