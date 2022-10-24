package service

import (
	"gin-gorm-oj/define"
	"gin-gorm-oj/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProblemList
// @Summary 问题列表
// @Tags 公共方法
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
	size, err := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	if err != nil {
		log.Println("GetProblemList size strconv error:", err)
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList page strconv error:", err)
		return
	}
	page = (page - 1) * size
	var count int64

	keyword := c.Query("keyword")

	list := make([]*models.Problem, 0)
	tx := models.GetProblemList(keyword)
	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("GetProblemList error:", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}
