package main

import (
	"fmt"
	"net/http"

	"github.com/zserge/lorca"
	"github.com/gin-gonic/gin"
    "github.com/drud/ddev/pkg/ddevapp"
)

func server() {
	router := gin.Default()

	router.LoadHTMLGlob("./public/index.html")
	router.StaticFS("/public/dist", http.Dir("./public/dist"))
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Retrieving ddevapp to do some fancy stuff with it
	// in case its necessary, not yet
    // app := &ddevapp.DdevApp{}

    // Simple router status retriever via ddev's GetRouterStatus func
	router.GET("/api/router_status", func(c *gin.Context) {
		status, warning := ddevapp.GetRouterStatus()

        // fmt.Println(status, warning)

		c.JSON(200, gin.H{
			"data": status,
		})
	})

	go router.Run(":8080")
}

func main() {
	server()

	var ui lorca.UI
	ui, _ = lorca.New("", "", 1200, 1024)
	defer ui.Close()

	ui.Load("http://localhost:8080")
	<-ui.Done()
}
