package lolgo

import (
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
)

//session的操作

//设置session信息
func SetSession(ctx *gin.Context, k string, o interface{}) {
	session := sessions.Get(ctx) //实例
	session.Set(k, o) //设置
	session.Save()  //保存
}

//
func GetSession(ctx *gin.Context, k string) interface{} {
	session := sessions.Get(ctx)  //实例
	return session.Get(k)  //获取session信息
}

//保存用户的信息  继承是user 数据源 接口 
func SaveUser(ctx *gin.Context, user interface{}) {
	SetSession(ctx, "user", user)
}

//加载用户sessioon 信息
func LoadUser(ctx *gin.Context) interface{} {
	return GetSession(ctx, "user")
}

//保存角色ID
func SaveRoleId(ctx *gin.Context, roleId interface{}) {
	session := sessions.Get(ctx)

	session.Set("roleid",roleId)
	session.Save()
}

//加载角色ID
func LoadRoleId(ctx *gin.Context) interface{} {
	session := sessions.Get(ctx)
	o := session.Get("roleid")
	return o

}

//清除所有的session信息
func ClearAllSession(ctx *gin.Context) {
	session := sessions.Get(ctx)
	session.Clear()
	session.Save()
	return
}
