<template>
    <h3>缺少一个查询接口，根据这三个查询</h3>
    <div>
        <el-form :inline="true" :model="searchInfo">
            <el-form-item label="字典类型编码">
                <el-input v-model="searchInfo.TypeCode" placeholder="搜索编码" clearable />
            </el-form-item>
            <el-form-item label="字典名（中）">
                <el-input v-model="searchInfo.TypeCNName" placeholder="搜索名称" clearable />
            </el-form-item>
            <el-form-item label="字典名（英）">
                <el-input v-model="searchInfo.TypeName" placeholder="搜索名称" clearable />
            </el-form-item>

            <el-form-item>
                <el-button :icon="Search">查询</el-button>
                <el-button :icon="Refresh" @click="onReset">重置</el-button>

            </el-form-item>

        </el-form>
    </div>

    <div>
        <div>
            <el-button>新增</el-button>
        </div>
        <el-table>

        </el-table>
        <div>
            <el-pagination></el-pagination>
        </div>
    </div>
</template>


<script setup>
import { reactive, ref } from 'vue'
import { Delete, Edit, Search, Share, Upload, Refresh } from '@element-plus/icons-vue'

const searchInfo = reactive({
    TypeCode: "",
    TypeName: "",
    TypeCNName: ""
})


const onReset = () => {
    searchInfo.TypeCode = ""
    searchInfo.TypeCNName = ""
    searchInfo.TypeName = ""
}

</script>

<style></style>
