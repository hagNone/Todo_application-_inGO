package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks = []Task{}
var currentID = 1

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"tasks": tasks,
		})
	})

	router.POST("/add", func(c *gin.Context) {
		title := c.PostForm("title")
		if title != "" {
			tasks = append(tasks, Task{ID: currentID, Title: title})
			currentID++
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.POST("/complete/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Completed = true
				break
			}
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.POST("/delete/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		for i := range tasks {
			if tasks[i].ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":8080")
}
