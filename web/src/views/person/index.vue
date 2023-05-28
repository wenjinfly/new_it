<template>
    <h3> persion info</h3>
    <p>姓名：{{ name }}</p>
    <p>年龄：{{ age }}</p>
    <p>性别：{{ sex }}</p>
    <p>token:{{ token }}</p>
    <p>isLogin:{{ isLogin }}</p>
    <P>UserId:{{ userInfo.UserId }}</P>
    <P>uuid:{{ userInfo.Uuid }}</P>
    <div>
        <el-header>
            <div>
                <img src="">
                <span>New IT</span>
            </div>
            <div>
                {{ userInfo.NickName }}
            </div>
        </el-header>
    </div>
</template>

<script setup lang="ts">
//import { ref } from "vue";
import { storeToRefs } from 'pinia';
import { useUsersStore } from "../../pinia/modules/user";
const store = useUsersStore();
/*
const name = ref<string>(store.name);
const age = ref<number>(store.age);
const sex = ref<string>(store.sex);
*/
//state中的数据变为了响应式的
const { name, age, sex } = storeToRefs(store);
const { token, isLogin } = storeToRefs(store);
const { userInfo } = storeToRefs(store)
</script>


<style  scoped>
.el-header {
    background-color: #333541;
    display: flex;
    justify-content: space-between;
    padding-left: 0;
    align-items: center;
    color: #0fff;
    font-size: 30px;
}
</style>
