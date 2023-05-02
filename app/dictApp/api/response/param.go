package response

import "new_it/app/dictApp/model"

type ParamDicttypeResponse struct {
	Dicttype model.DictType `json:"DictType"`
}

type ParamDictInfoResponse struct {
	Dictinfo model.DictInfo `json:"DictInfo"`
}
