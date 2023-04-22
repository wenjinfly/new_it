package response

import "new_it/app/usercentApp/model"

type SysAuthorityResponse struct {
	Authority model.SysAuthorities `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      model.SysAuthorities `json:"authority"`
	OldAuthorityId string               `json:"oldAuthorityId"` // 旧角色ID
}
