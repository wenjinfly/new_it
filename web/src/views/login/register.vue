<template>
    <div class="container">
        <el-form ref="ruleFormRef" class="register" :model="userInfo" :rules="userInfoRules" status-icon>

            <h3 class="title">用户注册</h3>
            <el-form-item prop="username">
                <el-input v-model="userInfo.username" size="large" placeholder="请输入账号" clearable />
            </el-form-item>
            <el-form-item prop="password">
                <el-input v-model="userInfo.password" size="large" placeholder="请输入密码" show-password />
            </el-form-item>
            <el-form-item prop="cheackPassword">
                <el-input v-model="userInfo.cheackPassword" size="large" placeholder="请确认密码" show-password />
            </el-form-item>
            <el-form-item>
                <el-button style="width: 100%" size="large" type="primary" @click="onrRegister"
                    :icon="UploadFilled">注册</el-button>
            </el-form-item>
            <el-row>
                <router-link :to="{ name: 'Login' }"> 已有密码？登录</router-link>

            </el-row>
        </el-form>


    </div>
</template >


<script setup>
import { UploadFilled } from '@element-plus/icons-vue'
import userApi from '@/api/user.js';
import router from '@/router/index.js'
import { ElMessage } from 'element-plus'
import { reactive, ref } from 'vue'
import md5 from 'js-md5'

const ruleFormRef = ref(null)

const userInfo = ref({
    username: "",
    password: "",
    cheackPassword: "",
    md5Password: ""
})


const onrRegister = async () => {

    if (!ruleFormRef.value) return
    ruleFormRef.value.validate(async (valid, fields) => {
        if (valid) {
            console.log('submit!')
            console.log(userInfo)
            let salt = 'new_it'
            //将密码加盐做md5
            userInfo.value.md5Password = md5(userInfo.value.password + salt);

            console.log(userInfo)

            const table = await userApi.register({
                username: userInfo.value.username,
                password: userInfo.value.md5Password,
                nickname: "user001",
                authorityId: "888"
            })

            if (table.code === 0) {
                console.log(table)
                ElMessage.success("注册成功，请登录！")
                router.push({ name: "Login" })

            } else {
                console.log("error:" + table.msg)

            }



        } else {
            console.log('error submit!', fields)
        }
    })

}

const checkUsername = (rule, value, callback) => {
    if (value.length < 4) {
        return callback(new Error('请输入正确的账号'))
    } else {
        callback()
    }
}


const checkPassword = (rule, value, callback) => {
    console.log(value)

    if (value === '' || value.length < 6) {
        callback(new Error('Please input the password'))
    } else {
        if (userInfo.value.cheackPassword !== '') {
            if (!ruleFormRef.value) return
            ruleFormRef.value.validateField('cheackPassword', () => null)
        }
        callback()
    }
}
const validatePass2 = (rule, value, callback) => {
    if (value === '' || value.length < 6) {
        callback(new Error('Please input the password again'))
    } else if (value !== userInfo.value.password) {
        callback(new Error("Two inputs don't match!"))
    } else {
        callback()
    }
}

const userInfoRules = reactive({
    username: [{ validator: checkUsername, trigger: 'blur', required: true, message: '请输入正确的账号' }],
    password: [{ validator: checkPassword, trigger: 'blur', required: true, }],
    cheackPassword: [{ validator: validatePass2, trigger: 'blur', required: true }]
})

</script>

<style scoped >
.container {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100vw;
    height: 100vh;
}

.register {
    width: 350px;
    padding: 36px 36px 36px 36px;
    border-radius: 6px;
    box-shadow: 0 0 18px #cac6c6;
}

.title {
    font-size: 1.5rem;
    text-align: center;
    padding-bottom: 30px;
}
</style>
