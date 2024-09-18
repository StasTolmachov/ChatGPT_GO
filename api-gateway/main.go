package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http/httputil"
	"net/url"
)

func reverseProxy(target string, c *gin.Context) {
	targetURL, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	log.Printf("Proxying request to: %s", targetURL.String())
	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	router := gin.Default()

	// Прокси для Task Service
	router.Any("/tasks/*proxyPath", func(c *gin.Context) {
		reverseProxy("http://localhost:8081", c)
	})

	// Прокси для User Service
	router.Any("/users/*proxyPath", func(c *gin.Context) {
		reverseProxy("http://localhost:8082", c)
	})

	router.Run(":8080")
}
