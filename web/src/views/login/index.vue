<template>
    <section class="container">

        <el-form ref="loginFormRef" class="login" :model="userInfo" :rules="userInfoRules">
            <h3 class="title">New IT 登录</h3>
            <el-form-item prop="username">
                <el-input v-model="userInfo.username" size="large" placeholder="账号" clearable />
            </el-form-item>
            <el-form-item prop="password">
                <el-input v-model="userInfo.password" size="large" placeholder="密码" show-password />
            </el-form-item>
            <el-button style="width: 100%" size="large" type="primary" :loading="loginLoading" @click="handleLogin">
                登录
            </el-button>
        </el-form>
    </section>
</template>



<script setup>
import { reactive, ref } from 'vue'
import { ElForm, ElFormItem, ElInput, ElButton } from 'element-plus';
import md5 from 'js-md5'

const loginLoading = ref(false);

const loginFormRef = ref(null)
const userInfo = reactive(
    {
        username: '',
        password: ''
    })

const checkUsername = (rule, value, callback) => {

    if (value.length < 4) {
        return callback(new Error('请输入正确的账号'))
    } else {
        callback()
    }
}

const checkPassword = (rule, value, callback) => {

    if (value === "" || value.length < 6) {
        return callback(new Error('请输入正确的密码'))
    } else {
        callback()
    }
}

const userInfoRules = reactive({
    username: [{ validator: checkUsername, trigger: 'blur', required: true, message: '请输入正确的账号' }],
    password: [{ validator: checkPassword, trigger: 'blur', required: true, message: '请输入正确的密码' }]
})

import userApi from '@/api/user.js';
import menuApi from '@/api/menu.js';
import { useUsersStore } from '@/pinia/modules/user.js'

import router from '@/router/index.js'

///
import { useRouterStore } from '@/pinia/modules/router.js'

///
const handleLogin = async () => {

    loginLoading.value = true;
    let salt = 'new_it'
    userInfo.password = md5(userInfo.password + salt);

    console.log(import.meta.env.VITE_APP_BASE_URL)

    const store = useUsersStore();
    console.log(userInfo)
    console.log("------2-----")
    try {
        const res = await userApi.login(userInfo);

        if (res.data.code === 0) {

            store.setUserInfo(res.data.data.user)
            store.setToken(res.data.data.token)

            const menus = await menuApi.getViewMenu()
            if (menus.data.code === 0) {
                console.log("------menus-----")
                console.log(menus)
            } else {
                console.log("----get--menus---error --")

            }

            const routerStore = useRouterStore()
            console.log("======222222===")

            await routerStore.SetAsyncRouter(menus.data)
            console.log("======2222223333322===")

            const asyncRouters = routerStore.asyncRouters
            console.log(asyncRouters)

            const uuu = router.getRoutes()
            console.log(uuu)

            asyncRouters.forEach(asyncRouter => {
                router.addRoute(asyncRouter)
            })
            console.log("=====22====")
            const uuu2 = router.getRoutes()
            console.log(uuu2)
            console.log("=====33====")

            console.log(router)
            console.log("=====45====")
            console.log(asyncRouters)

            console.log(res.data.data.user.authority.DefaultRouter)

            router.push({ name: res.data.data.user.authority.DefaultRouter })
            return true
        }

    } catch (e) {
        console.warn(e);
    } finally {
        loginLoading.value = false;
    }


    if (!loginFormRef.value) return


}


</script>

<style scoped>
.container {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100vw;
    height: 100vh;
}

.login {
    width: 350px;
    padding: 36px 36px 36px;
    border-radius: 6px;
    box-shadow: 0 0 18px #cac6c6;
}

.title {
    font-size: 1.5rem;
    text-align: center;
    padding-bottom: 30px;
}
</style>