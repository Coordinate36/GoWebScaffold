package api

import (
	"log"
	"net/http"
	. "web/config"
	. "web/models"

	"github.com/gin-gonic/gin"
)

func (r *Router) addPerson(c *gin.Context) {
	var person Person
	if err := c.BindJSON(&person); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "err": err})
		return
	}

	if err := r.db.Create(&person).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "err": Constant.DbInsertErrMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": ""})
}

func (r *Router) removePerson(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := r.db.Delete(&Person{}, "id = ?", id).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "err": Constant.DbDeleteErrMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": ""})
}

func (r *Router) getPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person

	if err := r.db.First(&person, "id = ?", id).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "err": Constant.DbQueryErrMsg})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": person})
}

func (r *Router) batchGetPerson(c *gin.Context) {
	var persons []Person

	if err := r.db.Find(&persons).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, Constant.DbDeleteErrMsg)
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": persons})
}
