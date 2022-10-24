

``` s


go get -u github.com/gin-gonic/gin

go get -u gorm.io/gorm

go get -u gorm.io/driver/mysql

go get -u github.com/swaggo/swag/cmd/swag

go install github.com/swaggo/swag/cmd/swag

swag init

go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files


```


swagger 接口访问地址：
``` s

http://localhost:8080/swagger/index.html



```


``` s
// gin-gorm-oj
// @Summary 问题列表
// @Tags 公共方法
// @Param page query int false "请输入当前页，默认第一页"
// @Success 200 {string} json: "{"code":"200","msg":"","data":""}"
// @Router /problem-list GetProblemList
```