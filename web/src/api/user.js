import http from '../utils/http';

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
const login = (data) => {

    return http.post(
        '/login',
        data
    )
}

const register = (data) => {
    return http.post(
        '/user/register',
        data
    )
}

export default {
    login,
    register
};
