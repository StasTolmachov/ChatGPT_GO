package main

import (
	"ChatGPT_GO/user-service/auth"
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
	//todo logger init
	
	router := gin.Default()

	// Маршрут для авторизации (JWT не проверяется)
	router.POST("/users/login", func(c *gin.Context) {
		reverseProxy("http://localhost:8081", c) //http://user-service:8081 docker
	})

	// Маршрут для регистрации
	router.POST("/users/register", func(c *gin.Context) {
		reverseProxy("http://localhost:8081", c) //http://user-service:8081
	})

	// Защищенные маршруты
	protected := router.Group("/tasks")
	protected.Use(auth.JWTAuthMiddleware()) // Применение JWT middleware
	{
		protected.Any("/*proxyPath", func(c *gin.Context) {
			reverseProxy("http://localhost:8082", c) // http://task-service:8082
		})
	}

	router.Run(":8080")
}
