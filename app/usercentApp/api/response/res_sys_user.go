package response

import "new_it/app/usercentApp/model"

type LoginResponse struct {
	User      model.SysUsers `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
