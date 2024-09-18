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
	router := gin.Default()

	// Маршрут для авторизации (JWT не проверяется)
	router.POST("/users/login", func(c *gin.Context) {
		reverseProxy("http://localhost:8081", c)
	})

	// Маршрут для регистрации
	router.POST("/users/register", func(c *gin.Context) {
		reverseProxy("http://localhost:8081", c)
	})

	// Защищенные маршруты
	protected := router.Group("/tasks")
	protected.Use(auth.JWTAuthMiddleware()) // Применение JWT middleware
	{
		protected.Any("/*proxyPath", func(c *gin.Context) {
			reverseProxy("http://localhost:8082", c) // Прокси на Task Service
		})
	}
	/*
		// Прокси для Task Service
		router.Any("/tasks/*proxyPath", func(c *gin.Context) {
			reverseProxy("http://localhost:8081", c) // Прокси на User Service
		})

		// Прокси для User Service
		router.Any("/users/*proxyPath", func(c *gin.Context) {
			reverseProxy("http://localhost:8082", c)
		})
	*/
	router.Run(":8080")
}
