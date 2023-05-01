package service

import (
	"errors"
	"new_it/app/usercentApp/model"
	"new_it/common"
	"new_it/global"

	"gorm.io/gorm"
)

type AuthorityService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth model.SysAuthority
//@return: err error, authority model.SysAuthority

func (authorityService *AuthorityService) CreateAuthority(auth model.SysAuthorities) (err error, authority model.SysAuthorities) {
	var authorityBox model.SysAuthorities
	if !errors.Is(global.GLB_DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), auth
	}
	err = global.GLB_DB.Create(&auth).Error
	return err, auth
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CopyAuthority
//@description: 复制一个角色
//@param: copyInfo response.SysAuthorityCopyResponse
//@return: err error, authority model.SysAuthority
/*
func (authorityService *AuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (err error, authority model.SysAuthorities) {
	var authorityBox model.SysAuthorities
	if !errors.Is(global.GLB_DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), authority
	}
	copyInfo.Authority.Children = []model.SysAuthority{}
	err, menus := MenuServiceApp.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	if err != nil {
		return
	}
	var baseMenu []system.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = global.GLB_DB.Create(&copyInfo.Authority).Error
	if err != nil {
		return
	}
	paths := CasbinServiceApp.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = CasbinServiceApp.UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Authority)
	}
	return err, copyInfo.Authority
}
*/
//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth model.SysAuthority
//@return: err error, authority model.SysAuthority

func (authorityService *AuthorityService) UpdateAuthority(auth model.SysAuthorities) (authority model.SysAuthorities, err error) {

	if auth.ParentId != "" {
		if errors.Is(global.GLB_DB.Where("authority_id = ?", auth.ParentId).First(&model.SysAuthorities{}).Error, gorm.ErrRecordNotFound) {
			// 判断父角色是否存在
			return model.SysAuthorities{}, errors.New("需要设置的父角色不存在")
		}

	}
	err = global.GLB_DB.Where("authority_id = ?", auth.AuthorityId).First(&model.SysAuthorities{}).Updates(&auth).Error
	return auth, err
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: DeleteAuthority
// @description: 删除角色
// @param: auth *model.SysAuthority
// @return: err error
func (authorityService *AuthorityService) DeleteAuthority(auth *model.SysAuthorities) (err error) {
	if errors.Is(global.GLB_DB.Debug().Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GLB_DB.Where("authority_id = ?", auth.AuthorityId).First(&model.SysUsers{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GLB_DB.Where("parent_id = ?", auth.AuthorityId).First(&model.SysAuthorities{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := global.GLB_DB.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)

	if len(auth.SysBaseMenus) > 0 {
		//先删除关联数据
		err = global.GLB_DB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		if err != nil {
			return
		}
		err = db.Unscoped().Delete(auth).Error

	} else {
		err = db.Error
		if err != nil {
			return
		}
	}
	//err = global.GLB_DB.Delete(&[]model.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (a *AuthorityService) GetAuthorityInfoList(info common.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GLB_DB.Model(&model.SysAuthorities{})
	db.Where("parent_id = ?", "0").Count(&total)

	var authority []model.SysAuthorities
	err = db.Limit(limit).Offset(offset).Where("parent_id = ?", "0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = a.findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth model.SysAuthority
//@return: err error, sa model.SysAuthority

func (authorityService *AuthorityService) GetAuthorityInfo(auth model.SysAuthorities) (err error, sa model.SysAuthorities) {
	err = global.GLB_DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

func (authorityService *AuthorityService) GetAuthorityInfoByID(authid string) (err error, sa model.SysAuthorities) {
	err = global.GLB_DB.Where("authority_id = ?", authid).First(&sa).Error
	return err, sa
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.SysAuthority
//@return: error

func (authorityService *AuthorityService) SetMenuAuthority(auth *model.SysAuthorities) error {
	var s model.SysAuthorities
	global.GLB_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GLB_DB.Model(&s).Association("SysBaseMenus").Append(&auth.SysBaseMenus)

	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *model.SysAuthority
//@return: err error

func (a *AuthorityService) findChildrenAuthority(authority *model.SysAuthorities) (err error) {
	err = global.GLB_DB.Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = a.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}
