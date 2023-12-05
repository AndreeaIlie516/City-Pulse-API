package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type event struct {
	ID          string `json:"id"`
	Time        string `json:"time"`
	Band        string `json:"band"`
	Location    string `json:"location"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}

type eventModel struct {
	ID         string `json:"id"`
	Event      event  `json:"event"`
	IsFavorite bool   `json:"is_favorite"`
	IsPrivate  bool   `json:"is_private"`
}

var events = []event{
	{ID: "1", Time: "27 octombrie - 20:30", Band: "Coma + Om la Luna", Location: "/Form Space", ImageUrl: "Url 1", Description: "Description"},
}

var eventModels = []eventModel{
	{ID: events[0].ID, Event: events[0], IsFavorite: false, IsPrivate: false},
}

func getEvents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, events)
}

func getEventModels(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, eventModels)
}

func eventById(id string) (*event, error) {
	for i, b := range events {
		if b.ID == id {
			return &events[i], nil
		}
	}

	return nil, errors.New("event not found")
}

func getEventById(c *gin.Context) {
	id := c.Param("id")
	event, err := eventById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, event)
}

func eventModelById(id string) (*eventModel, error) {
	for i, b := range eventModels {
		if b.ID == id {
			return &eventModels[i], nil
		}
	}

	return nil, errors.New("event model not found")
}

func getEventModelById(c *gin.Context) {
	id := c.Param("id")
	eventModel, err := eventModelById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "event model not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, eventModel)
}

func createPrivateEvent(c *gin.Context) {
	var newEvent event

	if err := c.Bind(&newEvent); err != nil {
		return
	}

	newEvent.ID = strconv.Itoa(len(events) + 1)

	events = append(events, newEvent)
	eventModels = append(eventModels, eventModel{ID: newEvent.ID, Event: newEvent, IsFavorite: false, IsPrivate: true})
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func createPublicEvent(c *gin.Context) {
	var newEvent event

	if err := c.Bind(&newEvent); err != nil {
		return
	}

	newEvent.ID = strconv.Itoa(len(events) + 1)

	events = append(events, newEvent)
	eventModels = append(eventModels, eventModel{ID: newEvent.ID, Event: newEvent, IsFavorite: false, IsPrivate: false})
	c.IndentedJSON(http.StatusCreated, newEvent)
}

func main() {
	router := gin.Default()
	router.GET("/getEvents", getEvents)
	router.GET("/getEventModels", getEventModels)
	router.GET("/getEventById/:id", getEventById)
	router.GET("/getEventModelById/:id", getEventModelById)
	router.POST("/createPrivateEvent", createPrivateEvent)
	router.POST("/createPublicEvent", createPublicEvent)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
