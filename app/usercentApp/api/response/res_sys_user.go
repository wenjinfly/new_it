package response

import "new_it/app/configApp/model"

type LoginResponse struct {
	User      model.SysUsers `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
