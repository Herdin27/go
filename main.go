package main

import (
	"api/learning/config"
	"api/learning/controllers"
	"fmt"
	"rsc.io/quote"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBinit()
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()
	fmt.Println("hello word")
	fmt.Println(quote.Go())
	router.GET("/person", inDB.GetPerson)
	router.Run(":3000")
}
