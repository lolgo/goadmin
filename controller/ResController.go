package controller

import (
	"goadmin/lolgo"

	"github.com/gin-gonic/gin"

	"goadmin/entity"
	"goadmin/service"

	"github.com/gin-gonic/gin/binding"
)

type ResController struct {
	lolgo.Controller
}

var resourceService service.ResourceService

//初始化权限资源
func (ctrl *ResController) init() {
	auth := resourceService.All()
	tmp := make(map[string]int64)
	for _, a := range auth {
		tmp[a.Patern] = a.ID
	}
	lolgo.AllAuth(tmp)
}
func (ctrl *ResController) Router(router *gin.Engine) {
	ctrl.init()
	r := router.Group("resource")
	r.POST("addmod", ctrl.addmod)
	r.POST("addres", ctrl.addres)
	r.POST("search", ctrl.search)
}

//基于全部的搜索
func (ctrl *ResController) addmod(ctx *gin.Context) {
	var res entity.Resource
	ctx.ShouldBindWith(&res, binding.FormPost)
	res.Pid = 0
	res.ResType = "mod"
	res.Stat = 1

	ret, err := resourceService.Add(res)
	if err != nil {
		lolgo.ResultFail(ctx, err.Error())
	} else {
		ctrl.init()
		lolgo.ResultOkMsg(ctx, ret, "模块添加成功")
	}
}

//基于全部的搜索
func (ctrl *ResController) addres(ctx *gin.Context) {
	var res entity.Resource
	ctx.ShouldBindWith(&res, binding.FormPost)
	ret, err := resourceService.Add(res)
	if err != nil {
		lolgo.ResultFail(ctx, err.Error())
	} else {
		ctrl.init()
		lolgo.ResultOkMsg(ctx, ret, "资源添加成功")
	}
}

//基于全部的搜索
func (ctrl *ResController) search(ctx *gin.Context) {
	ret := resourceService.All()
	//最后响应数据列表到前端
	lolgo.ResultList(ctx, ret, int64(len(ret)))
}
