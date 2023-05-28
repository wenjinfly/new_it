import { defineStore } from 'pinia'
import { ref, computed, watch } from 'vue'

// 第一个参数是应用程序中 store 的唯一 id
export const useUsersStore = defineStore('users', {
    // 其它配置项
    state: () => {
        //用户信息
        const userInfo = ref({
            UserId: '',
            UserName: '',
            Uuid: '',
            NickName: '',
            HeaderImg: '',
            authority: {},
            sideMode: 'dark',
            activeColor: '#4D70FF',
            baseColor: '#fff'
        })
        const setUserInfo = (val) => {
            userInfo.value = val

            console.log(userInfo)
        }
        const ResetUserInfo = (value = {}) => {
            userInfo.value = {
                ...userInfo.value,
                ...value
            }
        }

        //token
        const token = ref(window.localStorage.getItem('token') || '')
        const setToken = (val) => {
            token.value = val
        }

        const isLogin = ref(false)
        const setIsLogin = (val) => {
            isLogin.value = val
        }




        return {
            userInfo,
            token,
            isLogin,
            //
            setUserInfo,
            setToken,
            setIsLogin,
            //测试使用
            name: "小猪课堂",
            age: 25,
            sex: "男",
        }
    },

})