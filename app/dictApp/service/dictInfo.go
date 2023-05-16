package service

import (
	"errors"
	"new_it/app/dictApp/api/request"
	"new_it/app/dictApp/model"
	"new_it/global"

	"gorm.io/gorm"
)

type DICT_INFO_SERVICE struct{}

func (us *DICT_INFO_SERVICE) AddDictInfo(dt model.DictInfo) error {
	var dttype model.DictInfo
	if !errors.Is(global.GLB_DB.Where("code = ?", dt.Code).First(&dttype).Error, gorm.ErrRecordNotFound) {
		//
		return errors.New("当前字典数据已经存在")
	}

	if errors.Is(global.GLB_DB.Where("type_code = ?", dt.TypeCode).First(&model.DictType{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此TypeCode在DictType中不存在")
	}

	var err = global.GLB_DB.Create(&dt).Error

	return err
}

func (us *DICT_INFO_SERVICE) DeleteDictInfo(dt model.DictInfo) error {

	if errors.Is(global.GLB_DB.Where("code = ?", dt.Code).First(&dt).Error, gorm.ErrRecordNotFound) {
		return errors.New("该类型不存在")
	}

	err := global.GLB_DB.Where("code = ?", dt.Code).Delete(&dt).Error

	return err
}

func (us *DICT_INFO_SERVICE) UpdateDictInfo(dt model.DictInfo) error {
	if errors.Is(global.GLB_DB.Where("code = ?", dt.Code).First(&dt).Error, gorm.ErrRecordNotFound) {
		return errors.New("该类型不存在")
	}
	var err = global.GLB_DB.Updates(&dt).Error
	return err
}

func (us *DICT_INFO_SERVICE) GetDictInfoByCode(code string) (dt model.DictInfo, err error) {

	var reqUser model.DictInfo
	err = global.GLB_DB.First(&reqUser, "code = ?", code).Error
	return reqUser, err
}

func (us *DICT_INFO_SERVICE) GetDictInfoListByTypeCode(info request.ParamDictInfoList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GLB_DB.Model(&model.DictInfo{})

	var userList []model.DictInfo
	err = db.Where("type_code = ?", info.TypeCode).Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&userList).Error

	return userList, total, err
}

func (us *DICT_INFO_SERVICE) CheckDictCodeExist(code string) bool {
	if errors.Is(global.GLB_DB.Where("code = ?", code).First(&model.DictInfo{}).Error, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
