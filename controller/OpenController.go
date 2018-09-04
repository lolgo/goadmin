package controller

import (
	"goadmin/lolgo"
	"github.com/gin-gonic/gin"

)

type OpenController struct {
	lolgo.Controller
}


func (ctrl *OpenController)Router(router *gin.Engine){

	r := router.Group("open")
	r.GET("verify",ctrl.verify)

}

func (ctrl *OpenController)verify(ctx *gin.Context){
	lolgo.LoadVerify(ctx)
}
