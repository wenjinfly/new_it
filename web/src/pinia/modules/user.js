import { defineStore } from 'pinia'
import { ref, computed, watch } from 'vue'

// 第一个参数是应用程序中 store 的唯一 id
export const useUsersStore = defineStore('users', {
    // 其它配置项
    state: () => {
        //用户信息
        const userInfo = ref({
            userid: '',
            uuid: '',
            nickName: '',
            headerImg: '',
            authority: {},
            sideMode: 'dark',
            activeColor: '#4D70FF',
            baseColor: '#fff'
        })
        const setUserInfo = (val) => {
            userInfo.value = val
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
            console.log("111111")

            console.log(val)
            console.log(token)
            console.log("22222")

            console.log(token.value)
        }




        return {
            userInfo,
            token,
            //
            setUserInfo,
            setToken,
            //测试使用
            name: "小猪课堂",
            age: 25,
            sex: "男",
        }
    },

})