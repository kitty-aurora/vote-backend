package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"vote-backend/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 跨域
	r.Use(cors.Default())

	// 路由
	r.GET("/health", func(c *gin.Context) { c.String(200, "ok") })
	r.GET("/candidates", controllers.GetCandidates)
	r.POST("/vote/:name", controllers.VoteCandidate)
	r.POST("/reset", controllers.ResetVotes)

	return r
}
