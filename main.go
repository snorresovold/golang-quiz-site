package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type answer struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Correct bool   `json:"correct"`
}

type question struct {
	ID      int      `json:"id"`
	Content string   `json:"content"`
	Answers []answer `json:"answer"`
}

type quiz struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	Questions []question `json:"question"`
}

var answers = []answer{
	{ID: 0, Content: "Yes", Correct: true},
	{ID: 1, Content: "No", Correct: false},
}

var questions = []question{
	{ID: 0, Content: "Is Python a good language", Answers: answers},
}

var quizzes = []quiz{
	{ID: 0, Title: "python quiz", Author: "SnorreSovold", Questions: questions},
}

func getQuizzes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, quizzes)
}

func main() {
	router := gin.Default()
	router.GET("/quizzes", getQuizzes)
	router.Run("localhost:8080")
}
