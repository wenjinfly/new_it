import http from '../utils/http';

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => {

    return http.post(
        '/login',
        data
    )
}


export default {
    login,
};
