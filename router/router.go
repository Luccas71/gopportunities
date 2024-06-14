package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//inicializa o router com a config padrão do gin
	router := gin.Default()
	// inicializa as rotas
	initializeRoutes(router)
	// roda o servidor
	router.Run() // listen and serve on 0.0.0.0:8080
}
