package lolgo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回数据的方法  Result(obj,code,data,msg)
func Result(ctx * gin.Context,code int,data interface{},msg string){
	ctx.JSON(http.StatusOK, gin.H{"code": code, "data": data, "msg":msg})
}

//返回成功的信息
func ResultOk(ctx * gin.Context,data interface{}){
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": data, "msg": ""})
}

//返回数据列表 以及数量 rows行数 total 总共的数据 猜测可能是做分页用的
func ResultList(ctx * gin.Context,data interface{},total int64){
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "rows": data, "msg": "","total":total})
}

//返回成功的信息
func ResultOkMsg(ctx * gin.Context,data interface{},msg string){
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": data, "msg": msg})
}

//返回获取失败的信息
func ResultFail(ctx * gin.Context,err interface{}){
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg":err})
}

//操作访问失败的返回信息
func ResultFailData(ctx * gin.Context,data interface{},err interface{}){
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": data, "msg":err})
}