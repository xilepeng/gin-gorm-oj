

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


``` shell
// gin-gorm-oj
// @Summary 问题列表
// @Tags 公共方法
// @Param page query int false "请输入当前页，默认第一页"
// @Success 200 {string} json: "{"code":"200","msg":"","data":""}"
// @Router /problem-list GetProblemList
```


``` shell
swag init 
 
```




## MySQL 

``` shell
docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql

docker exec -it mysql bash

root@073ed0c2fc09:/# mysql -uroot -p
```

## Redis

``` shell
docker pull redis

docker run -itd --name redis -p 6379:6379 redis

docker exec -it redis /bin/bash

root@03ac20e1c401:/data# redis-cli
127.0.0.1:6379>
```
