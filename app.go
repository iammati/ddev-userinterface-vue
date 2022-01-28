package main

import (
	"fmt"
	"net/http"

	"github.com/drud/ddev/pkg/ddevapp"
    "github.com/drud/ddev/pkg/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/zserge/lorca"
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

	router.Use(cors.Default())

	// Endpoints for REST api v1
    v1 := router.Group("/api")
    {
	    // Simple router status retriever via ddev's GetRouterStatus func
        v1.GET("/router_status", func(c *gin.Context) {
			status, warning := ddevapp.GetRouterStatus()

			c.JSON(200, gin.H{
				"status": status,
				"warning": warning,
			})
        })

		// List of ddev projects
		v1.GET("/list_projects", func(c *gin.Context) {
			DdevBin := "/usr/local/bin/ddev"
			jsonOutput, err := exec.RunHostCommand(DdevBin, "list", "-j")

			if err != nil {
				fmt.Println(err)
			}

			c.JSON(200, gin.H{
				"list": jsonOutput,
			})
		})
    }

	// Note that this should only be used in development
	// and MUST BE changed as soon as it goes into production
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

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
