package main

import (
	"errors"
	"fmt"
	"net"
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

		v1.POST("/start_project", func(c *gin.Context) {
			Var_dump(c.PostFormMap("bla"))
			Var_dump(c.PostForm("bla"))

			c.JSON(200, gin.H{
				"state": "starting",
			})
		})
    }

	// Since the UI only runs local, this shouldn't
	// be that much of a security-issue, true?
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	go router.Run(":8080")
}

func main() {
	ip, err := hostIP()
	if err != nil {
		fmt.Println(err)
	}

	server()

	var ui lorca.UI
	ui, _ = lorca.New("", "", 1200, 1024)
	defer ui.Close()

	ui.Load("http://" + ip + ":8080")
	<-ui.Done()
}

func Var_dump(expression ...interface{} ) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}

// @credits https://stackoverflow.com/a/23558495/12100192
func hostIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
