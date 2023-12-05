package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func main() {
	router := gin.Default()
	router.GET("/getEvents", getEvents)
	router.GET("/getEventModels", getEventModels)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
