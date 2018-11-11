package api

import (
	. "web/config"
	"web/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Router struct {
	db     *gorm.DB
	router *gin.Engine
}

func (r *Router) Setup() {
	r.db = models.Db
	gin.DisableConsoleColor()
	r.router = gin.Default()

	r.router.GET("/people/:id", r.getPerson)
	r.router.GET("/people", r.batchGetPerson)
	r.router.POST("/people/add", r.addPerson)
	r.router.POST("/people/remove/:id", r.removePerson)
}

func (r *Router) Run() {
	r.router.Run(Config.ServerAddr)
}

func (r *Router) Close() {

}
