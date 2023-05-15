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
import { ElMessage } from 'element-plus'

///
import { useRouterStore } from '@/pinia/modules/router.js'

///
const handleLogin = async () => {

    loginLoading.value = true;
    let salt = 'new_it'
    //将密码加盐做md5
    userInfo.password = md5(userInfo.password + salt);

    console.log(import.meta.env.VITE_APP_BASE_URL)

    const store = useUsersStore();
    console.log(userInfo)
    try {
        //http的 login从服务请求登录
        const res = await userApi.login(userInfo);

        if (res.data.code === 0) {
            // 将用户信息和token放入store
            store.setUserInfo(res.data.data.user)
            store.setToken(res.data.data.token)

            // 再去请求动态菜单信息
            const menus = await menuApi.getViewMenu()
            if (menus.data.code === 0) {
                const routerStore = useRouterStore()
                //存储菜单信息
                await routerStore.SetAsyncRouter(menus.data)

                const asyncRouters = routerStore.asyncRouters
                //将菜单信息加入动态路由
                asyncRouters.forEach(asyncRouter => {
                    router.addRoute(asyncRouter)
                })

                console.log(res.data.data.user.authority.DefaultRouter)
                //跳转到当前角色的默认页面
                router.push({ name: res.data.data.user.authority.DefaultRouter })
                return true

            } else {
                console.log("----get--menus---error --")

            }

        }//end of if (res.data.code === 0)
        else {
            ElMessage({
                type: 'error',
                message: '登录失败，请正确填写登录信息',
                showClose: true,
            })
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