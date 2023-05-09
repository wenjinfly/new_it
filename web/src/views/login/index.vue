<template>
    <section class="container">

        <el-form ref="loginFormRef" class="login" :model="userInfo" :rules="userInfoRules">
            <h3 class="title">New IT 登录</h3>
            <el-form-item prop="account">
                <el-input v-model="userInfo.account" size="large" placeholder="账号" clearable />
            </el-form-item>
            <el-form-item prop="password">
                <el-input v-model="userInfo.password" size="large" placeholder="密码" show-password />
            </el-form-item>
            <el-button style="width: 100%" size="large" type="primary" @click="handleLogin">
                登录
            </el-button>
        </el-form>
    </section>
</template>



<script setup>
import { reactive, ref } from 'vue'
import { ElForm, ElFormItem, ElInput, ElButton } from 'element-plus';
const loginFormRef = ref(null)
const userInfo = reactive(
    {
        account: '',
        password: ''
    })

const checkAccount = (rule, value, callback) => {
    console.log(value)

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
    account: [{ validator: checkAccount, trigger: 'blur', required: true, message: '请输入正确的账号' }],
    password: [{ validator: checkPassword, trigger: 'blur', required: true, message: '请输入正确的密码' }]
})


const handleLogin = async () => {
    //console.log('submit!')
    // console.log(loginFormRef)
    // console.log(userInfo)
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