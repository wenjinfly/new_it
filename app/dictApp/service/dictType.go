package service

import (
	"errors"
	"new_it/app/dictApp/api/request"
	"new_it/app/dictApp/model"
	"new_it/global"

	"gorm.io/gorm"
)

type DICT_TYPE_SERVICE struct{}

// AddDictType
//
//	@receiver us
//	@param dt
//	@return error
func (us *DICT_TYPE_SERVICE) AddDictType(dt model.DictType) error {
	var dttype model.DictType
	if !errors.Is(global.GLB_DB.Where("type_code = ?", dt.TypeCode).First(&dttype).Error, gorm.ErrRecordNotFound) {
		//
		return errors.New("当前字典类型已经存在")
	}

	var err = global.GLB_DB.Create(&dt).Error

	return err
}

// DeleteDictType
//
//	@receiver us
//	@param dt
//	@return error
func (us *DICT_TYPE_SERVICE) DeleteDictType(dt model.DictType) error {

	if errors.Is(global.GLB_DB.Where("type_code = ?", dt.TypeCode).First(&model.DictType{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("该类型不存在")
	}

	//model.DictInfo
	if !errors.Is(global.GLB_DB.Where("type_code = ?", dt.TypeCode).First(&model.DictInfo{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此类型正在DictInfo中使用禁止删除")
	}

	err := global.GLB_DB.Where("type_code = ?", dt.TypeCode).Delete(&dt).Error

	return err
}

// UpdateDictType
//
//	@receiver us
//	@param dt
//	@return error
func (us *DICT_TYPE_SERVICE) UpdateDictType(dt model.DictType) error {
	if errors.Is(global.GLB_DB.Where("type_code = ?", dt.TypeCode).First(&model.DictType{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("该类型不存在")
	}

	var err = global.GLB_DB.Updates(&dt).Error
	return err
}

// GetDictTypeByCode
//
//	@receiver us
//	@param typecode
//	@return dt
//	@return err
func (us *DICT_TYPE_SERVICE) GetDictTypeByCode(typecode string) (dt model.DictType, err error) {

	var reqUser model.DictType
	err = global.GLB_DB.First(&reqUser, "type_code = ?", typecode).Error
	return reqUser, err
}

// GetDictTypeList
//
//	@receiver us
//	@param info
//	@return list
//	@return total
//	@return err
func (us *DICT_TYPE_SERVICE) GetDictTypeList(info request.ParamDictTypeList) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GLB_DB.Model(&model.DictType{})

	if info.TypeCode != "" {
		db = db.Where("type_code = ?", info.TypeCode)
	}

	if info.TypeName != "" {
		db = db.Where("type_name = ?", info.TypeName)
	}

	if info.TypeCNName != "" {
		db = db.Where("type_cn_name = ?", info.TypeCNName)
	}

	var userList []model.DictType
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&userList).Error

	return userList, total, err
}
