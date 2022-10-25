package service

import (
	"gin-gorm-oj/define"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// GetSubmitList
// @Tags 公共方法
// @Summary 提交列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param problem_identity query string false "problem_identity"
// @Param user_identity query string false "user_identity"
// @Param status query int false "status"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /submit-list [get]
func GetSubmitList(c *gin.Context) {
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

	problemIdentity := c.Query("problem_identity")
	userIdentity := c.Query("user_identity")
	status := c.Query("status")
}
