package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type toDo struct {
	ID string `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
}

var toDoList = []toDo{
	{ID: "1", Item: "Learn go",Completed: true},
	{ID: "2", Item: "Sent Message",Completed: true},
	{ID: "3", Item: "Start to job",Completed: false},
}

func geToDoList(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, toDoList)
	}
func addToDo(context *gin.Context) {
	var newToDo toDo

	if err := context.BindJSON(&newToDo); err != nil {
		return 
	}

	toDoList = append(toDoList, newToDo)

	context.IndentedJSON(http.StatusCreated,newToDo)
}

func toggleStatus(context *gin.Context) {
	id := context.Param("id")
	todo,err := getToDoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK,todo)
}
func getToDoById(id string) (*toDo,error) {
	for i,t := range toDoList {
		if t.ID == id {
			return &toDoList[i], nil
		}
	}
	return	nil, fmt.Errorf("no such to-do: %v", id)
}

func getToDo(context *gin.Context) {
	id := context.Param("id")
	todo,err := getToDoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
	}

	context.IndentedJSON(http.StatusOK,todo)
}
func main() {
	router := gin.Default()
	router.GET("/toDoList",geToDoList)
	router.GET("/toDoLis/:id",getToDo)
	router.PATCH("/toDoLis/:id",toggleStatus)
	router.POST("/toDoList",addToDo)
	router.Run("localhost:9090")
}