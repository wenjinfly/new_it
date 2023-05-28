<template>
    <div>
        <el-header>
            <div>
                <img src="">
                <span>New IT</span>
            </div>
            <div>
                <router-link v-if="isLogin" class="header-link" :to="{ name: 'person' }">
                    {{ userInfo.NickName }}
                    <el-avatar class="header-link" :src="headerImg" />
                </router-link>
                <router-link v-else="isLogin" class="header-link" :to="{ name: 'Login' }"> 登录/注册</router-link>
            </div>

        </el-header>
        <el-row>
            <el-col :span="2" class="left">
                <h3>New IT</h3>
            </el-col>
            <el-col :span="8" class="center">
                <el-input v-model="searchKey" @focus="focus" @blur="blur" @keyup.enter.native="getTableData"
                    placeholder="搜索任务">
                    <template #append>
                        <el-button :icon="Search" id="search" @click="getTableData"></el-button>
                    </template>
                </el-input>

            </el-col>
        </el-row>

        <div>
            <el-row v-for="(item) in tableData" :key="item.TaskId">
                <el-col :span="24">
                    <el-card shadow="hover">
                        <el-descriptions :title="item.TaskName">
                            <el-descriptions-item label="酬金:">
                                <el-text type="danger" tag="b">{{ item.TaskRewardMin }}-{{ item.TaskRewardMax }} 元</el-text>
                            </el-descriptions-item>

                            <el-descriptions-item label="任务需要人数:">
                                <el-tag size="small">{{ item.TaskRequiredPerson }}</el-tag>
                            </el-descriptions-item>

                            <el-descriptions-item label="任务阶段:">
                                <el-tag size="small">{{ item.TaskPhase }}</el-tag>
                            </el-descriptions-item>

                            <el-descriptions-item label="任务地址:">{{ item.TaskAddress }}</el-descriptions-item>
                            <el-descriptions-item label="任务起止时间:">
                                {{ item.TaskBegin }}-{{ item.TaskEnd }}
                            </el-descriptions-item>
                        </el-descriptions>
                    </el-card>
                </el-col>
            </el-row>
        </div>
        <div class="white-list-tool">
            <el-pagination v-model:current-page="page" v-model:page-size="pageSize" :page-sizes="[10, 30, 50]" background
                :total="total" layout="->,prev, pager, next" @size-change="handleSizeChange"
                @current-change="handleCurrentChange" />
        </div>

    </div>
</template>
  
<script setup>
import { Search } from '@element-plus/icons-vue'

import taskApi from '@/api/task.js'
import { ref, onMounted } from 'vue'
//
import { storeToRefs } from 'pinia';
import { useUsersStore } from "../../pinia/modules/user";
const store = useUsersStore();
const { userInfo } = storeToRefs(store)
const { isLogin } = storeToRefs(store);
//
const tableData = ref([])

const isFocus = ref(false)//是否聚焦
const searchKey = ref("")//当前输入框的值

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)

const headerImg = ref("https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png");

const focus = () => {
    isFocus.value = true
}

const blur = () => {
}

onMounted(() => {
    console.log("-----onMounted----")
    if (userInfo.value.HeaderImg !== "") {
        console.log(userInfo)
        console.log(userInfo.value.HeaderImg)
        headerImg.value = userInfo.value.HeaderImg
        console.log(headerImg)
    }

    getTableData()
})

const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
}

const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
}

const getTableData = async () => {
    console.log(searchKey.value)
    console.log(userInfo)

    const table = await taskApi.GetTaskListByName({
        page: page.value,
        pageSize: pageSize.value,
        name: searchKey.value
    })
    if (table.code === 0) {
        tableData.value = table.data.list
        total.value = table.data.total
        page.value = table.data.page
        pageSize.value = table.data.pageSize

    } else {
        console.log("error:" + table.msg)

    }
}


</script>
  
<style  scoped>
.header-link {
    color: aqua;
    font-size: 20px;
    vertical-align: middle;
}

.el-row {
    margin-bottom: 10px;
}

.el-header {
    background-color: #333541;
    display: flex;
    justify-content: space-between;
    padding-left: 0;
    align-items: center;
    color: #0fff;
    font-size: 30px;
}

.left {
    margin-left: 200px;
}

.center {
    margin-top: 15px;
    margin-left: 200px;
}

#search {
    border-radius: 0%;
}

.search-title {
    color: #bdbaba;
    font-size: 15px;
    margin-bottom: 5px;
}

.remove-history {
    color: #bdbaba;
    font-size: 15px;
    float: right;
    margin-top: -22px;
}

#search-box {
    width: 555px;
    height: 300px;
    margin-top: 0px;
    padding-bottom: 20px;
}
</style>