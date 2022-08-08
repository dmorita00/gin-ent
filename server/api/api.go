package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/name/pjName/server/pkg/configutil"
)

func SetHandlers(r *gin.Engine) {
	r.GET("/auth/user", AuthUser)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	sc := configutil.GetServerConfig()
	err := r.Run(sc.Host + ":" + sc.Port)
	if err != nil {
		log.Fatalf("router run error: %+v\n", err)
	}
}
