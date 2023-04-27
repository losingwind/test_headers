package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	localAPIRsp := os.Getenv("LOCAL_API_RSP")
	remoteServer := os.Getenv("REMOTE_SERVER")
	keepHeader := os.Getenv("KEEP_HEADER")
	router.GET("/local", func(c *gin.Context) {
		for k, v := range c.Request.Header {
			fmt.Println(k, v)
		}
		c.String(200, localAPIRsp)
	})

	router.GET("/remote", func(c *gin.Context) {
		for k, v := range c.Request.Header {
			fmt.Println(k, v)
		}

		client := &http.Client{}
		req, _ := http.NewRequest("GET", "http://"+remoteServer+"/local", nil)

		if keepHeader == "true" {
			req.Header = c.Request.Header
		} else {
			req.Header["X-Request-Id"] = c.Request.Header["X-Request-Id"]
			req.Header["X-B3-Traceid"] = c.Request.Header["X-B3-Traceid"]
		}

		rsp, err := client.Do(req)

		//rsp, err := http.Get("http://" + remoteServer + "/local")

		if err != nil {
			c.String(500, err.Error())
		} else {
			body, _ := ioutil.ReadAll(rsp.Body)
			c.String(200, string(body))
		}
	})
	router.Run(":8080")
}
