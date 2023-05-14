import http from '../utils/http';

// @Summary 用户登录 获取动态路由
// @Produce  application/json
// @Param 可以什么都不填 调一下即可
// @Router /menu/getViewMenu [post]
export const getViewMenu = () => {
    return http.post(
        '/menu/getViewMenu'
    )
}



export default {
    getViewMenu,
};
