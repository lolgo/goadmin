package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/tommy351/gin-sessions"
	"goadmin/lolgo"
	"net/http"
	"strconv"

	//系统包和第三方包
	"github.com/gin-gonic/gin"

	//自定义包
	"goadmin/controller"
)

func registerRouter(router *gin.Engine) {
	new(controller.PageController).Router(router)


}

//在main方法中初始化数据库  配置文件  路由  公共函数库  外部库文件等
func main() {

	cfg := new(lolgo.Config)  //Config 类 包含读取各种类型的配置方法
	cfg.Parse("config/app.properties")  //解析配置文件
	fmt.Println("[ok] load config ")	//输出提示信息
	lolgo.SetCfg(cfg)  //设置配置文件

	lolgo.Configuration(cfg.Logger["filepath"]) //配置文件路径

	gin.SetMode(cfg.App["mode"])  //gin package set development mode is configure's appliation mode

	//初始化数据源
	for k, ds := range cfg.Datasource {
		//xorm 的NewEngine 方法
		e, err := xorm.NewEngine(ds["driveName"], ds["dataSourceName"])
		if err != nil {
			fmt.Println("data source init error", err.Error())
			return
		}

		//显示SQL语句
		e.ShowSQL(ds["showSql"] == "true")

		//字符串转换成功整数
		n, _ := strconv.Atoi(ds["maxIdle"])

		//设置最大的连接数
		e.SetMaxIdleConns(n)

		//字符串转换成功整数
		n, _ = strconv.Atoi(ds["maxOpen"])

		//设置最大的开启数
		e.SetMaxOpenConns(n)


		//这个没有弄懂
		err = e.Sync2(new(entity.User), new(entity.Config), new(entity.RefRoleRes), new(entity.Resource), new(entity.Role))
		if err != nil {
			fmt.Println("data source init error", err.Error())
			return
		}

		//具体查看restgo中的setEngin 类
		lolgo.SetEngin(k, e)
	}

	//初始化数据源成功
	fmt.Println("[ok] init datasource")

	//初始化路由
	router := gin.Default()

	//这个没有弄明白
	for k, v := range cfg.Static {
		router.Static(k, v)
	}

	//静态文件
	for k, v := range cfg.StaticFile {
		router.StaticFile(k, v)
	}

	//设置函数映射(待解释)
	router.SetFuncMap(lolgo.GetFuncMap())
	//路由找不到 显示的界面 和信息
	router.NoRoute(lolgo.NoRoute)

	//方法找不到显示的界面和信息
	router.NoMethod(lolgo.NoMethod)

	//设置视图的路径
	router.LoadHTMLGlob(cfg.View["path"] + "/**/*")

	//使用配置中的分割符  配置在视图中
	router.Delims(cfg.View["deliml"], cfg.View["delimr"])

	//cookie 存储 使用byte数组存储 字符串的信息
	store := sessions.NewCookieStore([]byte(cfg.Session["name"]))
	//这个暂时没有弄明白
	router.Use(sessions.Middleware(cfg.Session["name"], store))
	//验证信息
	router.Use(lolgo.Auth())

	//注册路由
	registerRouter(router)

	//应用启动
	fmt.Println("[ok] app run", cfg.App["addr"]+":"+cfg.App["port"])

	//监听端口
	http.ListenAndServe(cfg.App["addr"]+":"+cfg.App["port"], router)
}
