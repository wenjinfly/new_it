<template>
    <h3>dictonary info</h3>
    <div>
        <div>
            <el-form :inline="true">

            </el-form>
        </div>

        <dic>
            <el-table></el-table>
        </dic>

        <dic>
            <el-pagination>

            </el-pagination>
        </dic>




    </div>
</template>

<script setup>
import ref from 'vue'

const searchInfo = ref({})
</script>