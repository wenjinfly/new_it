package api

type UsercentApiGroup struct {
	UsercentApi
	AuthorityInfoApi
}

var (
	UserApi      = UsercentApi{}
	AuthorityApi = AuthorityInfoApi{}
)
