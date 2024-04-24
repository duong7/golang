package controllers

import (
	"go_htmx/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	notes := []models.Note{
		{
			Name:    "Traveling",
			Content: "Will go to 3 domestic tourist destinations: Da Nang, Da Lat, Hue.",
		},
		{
			Name:    "Weight-gain",
			Content: "Gain 5 kg in the next 3 months.",
		},
		{
			Name:    "Apply for job",
			Content: "There are internships this year.",
		},
	}

	c.HTML(
		http.StatusOK,
		"notes/index.html",
		gin.H{
			"notes": notes,
		},
	)
}

type FormData struct {
	Name    string `form:"name"`
	Content string `form:"content"`
}

func NotesCreate(c *gin.Context) {
	var data FormData
	c.Bind(&data)
	// Simulate a delay
	time.Sleep(2 * time.Second)
	c.HTML(http.StatusOK,
		"notes/note.html",
		gin.H{
			"Name":    data.Name,
			"Content": data.Content,
		})
}
