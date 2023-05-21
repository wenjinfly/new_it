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
                <el-button :icon="Search" type="primary" @click="onSubmit">查询</el-button>
                <el-button :icon="Refresh" @click="onReset">重置</el-button>
            </el-form-item>

        </el-form>
    </div>

    <div>
        <div>
            <el-button size="small" type="primary" :icon="Plus" @click="openDialog">新增</el-button>
        </div>
        <el-table :data="tableData" style="width:100%" tooltip-effect="dark" row-key="TypeCode">
            <el-table-column type="selection" width="55"></el-table-column>
            <el-table-column align="left" label="字典编码" prop="TypeCode"></el-table-column>
            <el-table-column align="left" label="字典名（中）" prop="TypeCNName"></el-table-column>
            <el-table-column align="left" label="字典名（英）" prop="TypeName"></el-table-column>
            <el-table-column align="left" label="日期" :formatter="formatmyData"> </el-table-column>
            <el-table-column align="left" label="按钮组">
                <template #default="scope">
                    <el-button size="small" :icon="Document" type="primary" link @click="toDetile(scope.row)">详情</el-button>
                    <el-button size="small" :icon="Edit" type="primary" link
                        @click="updateSysDictionaryFunc(scope.row)">变更</el-button>
                    <el-popover v-model:visible="scope.row.visible" placement="top" width="160">
                        <p>确定要删除吗？</p>
                        <div style="text-align: right; margin-top: 8px">
                            <el-button size="small" link @click="scope.row.visible = false">取消</el-button>
                            <el-button type="primary" size="small">确定</el-button>
                        </div>
                        <template #reference>
                            <el-button link icon="delete" size="small" style="margin-left: 10px" type="primary"
                                @click="scope.row.visible = true">删除</el-button>
                        </template>
                    </el-popover>
                </template>

            </el-table-column>

        </el-table>
        <div class="white-list-tool">
            <el-pagination v-model:current-page="page" v-model:page-size="pageSize" :page-sizes="[10, 30, 50]" background
                :total="total" layout="->,total, sizes, prev, pager, next, jumper" small @size-change="handleSizeChange"
                @current-change="handleCurrentChange" />
        </div>

        <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
            <el-form ref="dialogForm" :model="addFormData" :rules="rules" size="default" label-width="110px">
                <el-form-item align="left" label="字典编码" prop="TypeCode">
                    <el-input v-model="addFormData.TypeCode" placeholder="请输入字典编码" clearable />
                </el-form-item>
                <el-form-item align="left" label="字典名（中）" prop="TypeCNName">
                    <el-input v-model="addFormData.TypeCNName" placeholder="请输入字典名称" clearable />
                </el-form-item>
                <el-form-item align="left" label="字典名（英）" prop="TypeName">
                    <el-input v-model="addFormData.TypeName" placeholder="请输入字典名称" clearable />
                </el-form-item>
            </el-form>
            <el-button size="small" @click="closeDialog">取消</el-button>
            <el-button size="small" type="primary" @click="enterDialog">确定</el-button>
        </el-dialog>
    </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { Document, Delete, Edit, Search, Share, Upload, Refresh, Plus } from '@element-plus/icons-vue'
import { formatDate } from '@/utils/format'

import dictTypeApi from '@/api/dictType.js'
const searchInfo = reactive({
    TypeCode: "",
    TypeName: "",
    TypeCNName: ""
})

const formatmyData = (row) => {
    return formatDate(row.CreatedAt)
}

const onReset = () => {
    searchInfo.TypeCode = ""
    searchInfo.TypeCNName = ""
    searchInfo.TypeName = ""
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

const addFormData = ref({
    TypeCode: "",
    TypeCNName: "",
    TypeName: "",
})

const rules = ref({
    TypeCode: [
        { required: true, message: "请输入编码", trigger: 'blur' }
    ],
    TypeCNName: [
        { required: true, message: "请输入字典名（中）", trigger: 'blur' }
    ],
    TypeName: [
        { required: true, message: "请输入字典名（英）", trigger: 'blur' }
    ]
})


const typeOpt = ref('')
const updateSysDictionaryFunc = async (row) => {

    const res = await dictTypeApi.GetDictTypeByCode({ TypeCode: row.TypeCode })
    typeOpt.value = 'update'
    if (res.code === 0) {
        console.log(res)
        addFormData.value = res.data.DictType
        dialogFormVisible.value = true
    }
}

import { useRouter, useRoute } from 'vue-router'

const toDetile = (row) => {


    router.push({
        name: 'dictInfo',
        params: {
            id: row.TypeCode,
        },
    })
}

const dialogForm = ref(null)
const enterDialog = async () => {
    dialogForm.value.validate(async (valid) => {
        console.log(valid)
        if (!valid) return
        let res
        switch (typeOpt.value) {
            case 'create':
                res = await dictTypeApi.AddDictType(addFormData.value)
                break
            case 'update':
                console.log(addFormData.value)
                res = await dictTypeApi.UpdateDictType(addFormData.value)
                break
            default:
                res = await dictTypeApi.AddDictType(addFormData.value)
                break
        }
        if (res.code === 0) {
            closeDialog()
            getTableData()
        }
    })
}
//
const dialogFormVisible = ref(false)
const openDialog = () => {
    typeOpt.value = 'create'
    dialogFormVisible.value = true
}
const closeDialog = () => {
    dialogFormVisible.value = false
    addFormData.value = {
        TypeCode: "",
        TypeCNName: "",
        TypeName: "",
    }
}

const onSubmit = () => {
    page.value = 1
    pageSize.value = 10
    getTableData()
}

// 分页
const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
}

const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
}

const getTableData = async () => {
    console.log(page.value)
    const table = await dictTypeApi.getDictTypeList({
        page: page.value,
        pageSize: pageSize.value,
        ...searchInfo
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

<style></style>
