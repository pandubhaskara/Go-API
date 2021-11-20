package articles

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/pandubhaskara/Go-API/models"
)

func Create(c *gin.Context) {
	db := c.MustGet("warehouse").(*mgo.Database)

	products := models.Product{}
	err := c.Bind(&products)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionProduct).Insert(products)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/product")
}

func List(c *gin.Context) {
	db := c.MustGet("warehouse").(*mgo.Database)
	products := []models.Product{}
	err := db.C(models.CollectionProduct).Find(nil).Sort("-updated_on").All(&products)
	if err != nil {
		c.Error(err)
	}
	c.HTML(http.StatusOK, "product/list", gin.H{
		"title":    "Prticles",
		"products": products,
	})
}

func Delete(c *gin.Context) {
	db := c.MustGet("warehouse").(*mgo.Database)
	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionProduct).Remove(query)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/product")
}
