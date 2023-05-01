package dbdata

import (
	"errors"
	"fmt"
	"new_it/app/usercentApp/model"
	"new_it/global"
	"strings"
)

var ViewAuthorityMenuMysql = new(viewAuthorityMenuMysql)

type viewAuthorityMenuMysql struct{}

func (v *viewAuthorityMenuMysql) TableName() string {
	var entity model.ViewAuthorityMenu
	return entity.TableName()
}

func (v *viewAuthorityMenuMysql) Initialize() error {
	var entity model.SysAuthorityMenus

	sql := `
	CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW @table_name AS
	select @menus.menu_id           AS menu_id,
		   @menus.path              AS path,
		   @menus.icon              AS icon,
		   @menus.name              AS name,
		   @menus.sort              AS sort,
		   @menus.title             AS title,
		   @menus.hidden            AS hidden,
		   @menus.component         AS component,
		   @menus.parent_id         AS parent_id,
		   @menus.created_at        AS created_at,
		   @menus.updated_at        AS updated_at,
		   @menus.deleted_at        AS deleted_at,
		   @menus.keep_alive        AS keep_alive,
		   @menus.menu_level        AS menu_level,
		   @menus.default_menu      AS default_menu,
		   @menus.close_tab      	AS close_tab,
		   @sys_authority_menus.authority_id AS authority_id
	from (@sys_authority_menus
			 join @menus on ((@sys_authority_menus.menu_id = @menus.menu_id)));
	`
	sql = strings.ReplaceAll(sql, "@table_name", v.TableName())
	sql = strings.ReplaceAll(sql, "@menus", "sys_base_menus")
	sql = strings.ReplaceAll(sql, "@sys_authority_menus", entity.TableName())
	if err := global.GLB_DB.Exec(sql).Error; err != nil {
		return errors.New(err.Error() + v.TableName() + "视图创建失败!")
	}
	return nil
}

func (v *viewAuthorityMenuMysql) CheckDataExist() bool {
	err1 := global.GLB_DB.Find(&[]model.ViewAuthorityMenu{}).Error
	err2 := fmt.Errorf("Error 1146: Table '%v.%v' doesn't exist", global.GLB_CFG_INFO.Mysql.Dbname, v.TableName())
	if errors.As(err1, &err2) {
		return false
	}
	return true
}
