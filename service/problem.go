package service

import (
	"gin-gorm-oj/define"
	"gin-gorm-oj/helper"
	"gin-gorm-oj/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
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
	categoryIdentity := c.Query("category_identity")
	list := make([]*models.ProblemBasic, 0)
	tx := models.GetProblemList(keyword, categoryIdentity)
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

// GetProblemDetail
// @Tags 公共方法
// @Summary 问题详情
// @Param identity query string false "identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-detail [get]
func GetProblemDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题唯一标识不能为空",
		})
		return
	}
	// 查找表
	data := new(models.ProblemBasic)
	err := models.DB.Where("identity = ?", identity).
		Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "问题不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get ProblemDetail Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

// ProblemCreate
// @Tags 管理员私有方法
// @Summary 问题创建
// @Param token header string true "token"
// @Param title formData string true "title"
// @Param content formData string true "content"
// @Param max_runtime formData string true "max_runtime"
// @Param max_mem formData string true "max_mem"
// @Param category_ids formData array false "category_ids"
// @Param test_case formData array true "test_case"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-create [post]
func ProblemCreate(c *gin.Context) {
	in := new(define.ProblemBasic)
	err := c.ShouldBindJSON(in)
	if err != nil {
		log.Println("[JsonBind Error] : ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}

	if in.Title == "" || in.Content == "" || len(in.ProblemCategories) == 0 || len(in.TestCases) == 0 || in.MaxRuntime == 0 || in.MaxMem == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	identity := helper.GetUUID()
	data := &models.ProblemBasic{
		Identity:   identity,
		Title:      in.Title,
		Content:    in.Content,
		MaxRuntime: in.MaxRuntime,
		MaxMem:     in.MaxMem,
		CreatedAt:  models.MyTime(time.Now()),
		UpdatedAt:  models.MyTime(time.Now()),
	}
	// 处理分类
	categoryBasics := make([]*models.ProblemCategory, 0)
	for _, id := range in.ProblemCategories {
		categoryBasics = append(categoryBasics, &models.ProblemCategory{
			ProblemId:  data.ID,
			CategoryId: uint(id),
			CreatedAt:  models.MyTime(time.Now()),
			UpdatedAt:  models.MyTime(time.Now()),
		})
	}
	data.ProblemCategories = categoryBasics
	// 处理测试用例
	testCaseBasics := make([]*models.TestCase, 0)
	for _, v := range in.TestCases {
		// 举个例子 {"input":"1 2\n","output":"3\n"}
		testCaseBasic := &models.TestCase{
			Identity:        helper.GetUUID(),
			ProblemIdentity: identity,
			Input:           v.Input,
			Output:          v.Output,
			CreatedAt:       models.MyTime(time.Now()),
			UpdatedAt:       models.MyTime(time.Now()),
		}
		testCaseBasics = append(testCaseBasics, testCaseBasic)
	}
	data.TestCases = testCaseBasics

	// 创建问题
	err = models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Problem Create Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"identity": data.Identity,
		},
	})
}
