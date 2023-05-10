import { defineStore } from 'pinia'

// 第一个参数是应用程序中 store 的唯一 id
export const useUsersStore = defineStore('users', {
    // 其它配置项
    state: () => {

        return {
            name: "小猪课堂",
            age: 25,
            sex: "男",
        }
    },

})