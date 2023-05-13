import http from '../utils/http';

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => {
    console.log("#####1###")
    console.log(data)
    console.log("#####2###")

    return http({
        url: '/login',
        method: 'post',
        data: data
    })

    // return http.post(
    //     '/login',
    //     data
    // )
}


export default {
    login,
};
