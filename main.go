package main

import (
	
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/shamalgithub/go_jwt.git/initializer"
)



func init(){
	initializer.LoadEnvVariables()
	DB := initializer.ConnectToDb()
	initializer.SyncDatabase(DB)
}


type todo struct{
	ID string  `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
}

var todos = []todo{
    {ID: "1", Item: "clean dfdjldffff f99y nfow roomyy", Completed: false},
    {ID: "2", Item: "read book", Completed: false},
    {ID: "3", Item: "take a np", Completed: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK , todos)
}

func main() {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/todo" , getTodos)
	router.Run("localhost:8080")
}