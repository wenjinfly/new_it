package model

type ViewAuthorityMenu struct {
	SysBaseMenus
	AuthorityId  string              `json:"-" gorm:"comment:角色ID"`
	ChildrenView []ViewAuthorityMenu `json:"children" gorm:"-"`
}

func (s ViewAuthorityMenu) TableName() string {
	return "view_authority_menu"
}
