package service

import (
	"goadmin/entity"
	"goadmin/model"
	"goadmin/lolgo"
)

type ConfigService struct{}

//根据userId 获取用户编号
func (service *ConfigService) FindOne(name string) entity.Config {
	var obj entity.Config
	orm := lolgo.OrmEngin()
	orm.Id(name).Get(&obj)
	return obj
}

//关键字模糊查询
func (service *ConfigService) Query(arg model.ConfigArg) []entity.Config {
	var objs []entity.Config = make([]entity.Config, 0)
	orm := lolgo.OrmEngin()
	t := orm.Where("1=1")
	if 0 < len(arg.Kword) {
		t = t.Where("name like ? or label like ?", "%"+arg.Kword+"%", "%"+arg.Kword+"%")
	}
	t.Limit(arg.GetPageFrom()).Find(&objs)
	return objs
}

//查询所有
func (service *ConfigService) All() []entity.Config {
	var objs []entity.Config = make([]entity.Config, 0)
	orm := lolgo.OrmEngin()
	orm.Where("1=1").Find(&objs)
	return objs
}

//根据ID 更新数据 
func (service *ConfigService) Update(name string, value string) (int64, error) {
	var obj entity.Config
	obj.Name = name
	obj.Value = value
	orm := lolgo.OrmEngin()
	r, e := orm.ID(name).Update(&obj)
	return r, e
}

//插入配置信息
func (service *ConfigService) Add(obj entity.Config) (int64, error) {

	orm := lolgo.OrmEngin()
	r, e := orm.InsertOne(&obj)
	return r, e
}
