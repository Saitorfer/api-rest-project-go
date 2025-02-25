package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"practise.com/rest-api-go/models"
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
	//now before enter here in the routes.go, we use the authentificate to be sure the token is valid
	//and we have the user id in the context
	//so we can get the user id from the context and save the event with the user id
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	//get the user id from the context
	userId := context.GetInt64("userId")
	event.UserID = userId

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

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventIdInt64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event by id"})
		return
	}
	//check if the user is the owner of the event
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not allowed to update this event"})
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
	userId := context.GetInt64("userId")

	event, err := models.GetEventById(eventIdInt64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event by id"})
		return
	}
	//check if the user is the owner of the event
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not allowed to delete this event"})
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
