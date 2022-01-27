package main

import (
	"fmt"
	"net/http"

	"github.com/zserge/lorca"
	"github.com/gin-gonic/gin"
    "github.com/drud/ddev/pkg/ddevapp"
	"github.com/drud/ddev/pkg/exec"
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

	// List of ddev projects
	router.GET("/api/list_projects", func(c *gin.Context) {
		// apps := ddevapp.GetActiveProjects()
		DdevBin := "/usr/local/bin/ddev"
		jsonOut, err := exec.RunHostCommand(DdevBin, "list", "-j")

		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{
			"list": jsonOut,
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

func Var_dump(expression ...interface{} ) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
