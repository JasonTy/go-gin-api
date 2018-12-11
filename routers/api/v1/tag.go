package v1

import (
	"github.com/gin-gonic/gin"
	"go/gin/api/models"
	"github.com/Unknwon/com"
)


func GetTags(c *gin.Context) {
	page, _ := com.StrTo(c.Query("page")).Int()
	size, _ :=com.StrTo(c.Query("size")).Int()

	result := models.GetTags(page, size, "name = ?", "测试")
	c.JSON(200, gin.H{
		"message": result,
	})
}

func AddTag(c *gin.Context) {

}

func EditTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}