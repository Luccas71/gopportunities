package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//inicializa o router com a config padrão do gin
	router := gin.Default()
	// handler
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// roda o servidor
	router.Run() // listen and serve on 0.0.0.0:8080
}
