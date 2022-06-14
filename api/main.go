package main

import(
	"fmt"
	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID uint64 `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Age uint64 `json:"age"`
	Email string `json:"email"`
	Tel string `json:"tel"` 
}

var data = []User{
	{
		ID: 1,
		Firstname: "Test",
		Lastname: "User",
		Age: 12,
		Email: "test@user.com",
		Tel: "0819991199",
	},
}

func main() {
	fmt.Println("Test Server.")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000/"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type", "Origin"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
	router.GET("/users", getUser)
	router.Run(":8000")
}

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, data)
}