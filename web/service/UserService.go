package service

import (
	"fmt"
	model "golang_test/web/model"
	"github.com/jinzhu/gorm"
)

/*调用函数出错，编译的时候放在了一个文件里面在，我的推测*/
var db *gorm.DB

func init() {

	fmt.Println("init ,开始调用函数")

}
func Add(user *model.User) string {
	var err error
	db, err := gorm.Open("mysql", "root:123456@/liliangbin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&model.User{})
	db.Create(&user)
	fmt.Println("shuju tianjia eng")
	return user.Info
}
