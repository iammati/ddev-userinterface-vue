package main

import (
	"net/http"

	"github.com/zserge/lorca"
	"github.com/gin-gonic/gin"
    "github.com/drud/ddev/pkg/ddevapp"
)

func server() {
	router := gin.Default()
	router.StaticFS("/dist", http.Dir("./dist"))
	router.LoadHTMLGlob("./dist/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Retrieving ddevapp to do some fancy stuff with it
	// in case its necessary. For the moment not at all (yet)
    // app := &ddevapp.DdevApp{}

    // Simple router status retriever via ddev's GetRouterStatus func
	router.GET("/api/router_status", func(c *gin.Context) {
		status, warning := ddevapp.GetRouterStatus()

		c.JSON(200, gin.H{
			"status": status,
			"warning": warning,
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
