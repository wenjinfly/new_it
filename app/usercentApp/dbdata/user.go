package dbdata

import (
	userModel "new_it/app/usercentApp/model"
	"new_it/global"

	"errors"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

var User = new(DBDATA_USER)

type DBDATA_USER struct{}

func (u *DBDATA_USER) TableName() string {
	return "sys_users"
}

func (u *DBDATA_USER) Initialize() error {

	uid, _ := uuid.NewV4()
	uid2, _ := uuid.NewV4()
	uid3, _ := uuid.NewV4()
	uid4, _ := uuid.NewV4()

	entities := []userModel.SysUsers{
		{Uuid: uid2.String(), UserName: "root", Password: "2361bd3bf151f0ec52a3ee762bca3644", NickName: "研发者", HeaderImg: "https://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "666"},
		{Uuid: uid.String(), UserName: "admin", Password: "2361bd3bf151f0ec52a3ee762bca3644", NickName: "超级管理员", HeaderImg: "https://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "777"},
		{Uuid: uid3.String(), UserName: "user", Password: "2361bd3bf151f0ec52a3ee762bca3644", NickName: "普通用户", HeaderImg: "https://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "888"},
		{Uuid: uid4.String(), UserName: "test", Password: "2361bd3bf151f0ec52a3ee762bca3644", NickName: "测试人员", HeaderImg: "https://qmplusimg.henrongyi.top/gva_header.jpg", AuthorityId: "999"},
	}
	if err := global.GLB_DB.Create(&entities).Error; err != nil {
		return errors.New(u.TableName() + "表数据初始化失败!")
	}
	return nil
}

func (u *DBDATA_USER) CheckDataExist() bool {
	if errors.Is(global.GLB_DB.Where("user_name = ?", "admin").First(&userModel.SysUsers{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
