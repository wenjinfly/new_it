<template>
    <h3>task details</h3>
    <span>{{ TaskId }}</span>
    <span>11 - {{ TaskId }}</span>

    <div>{{ taskdetail.TaskId }}</div>

    <div>{{ taskdetail.TaskName }}</div>
    <div>{{ taskdetail.TaskDuty }}</div>
    <div class="line" />
    <div v-for="(val, key, index) in taskdetail">{{ key }}: {{ val }}</div>
</template>



<script setup >
import { onMounted, ref } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import taskApi from '@/api/task.js'

const TaskId = ref(0)

const taskdetail = ref({})

const route = useRoute();


onMounted(() => {
    console.log(route.query)
    console.log(route.query.TaskId)
    TaskId.value = route.query.TaskId
    GetTaskByTaskID(route.query.TaskId)
}
)

const GetTaskByTaskID = async (id) => {
    console.log(id)

    var num = new Number(id);

    console.log(num)

    const detail = await taskApi.GetTaskByTaskID({
        TaskId: num
    })

    if (detail.code === 0) {
        console.log(detail)

        taskdetail.value = detail.data.Task
        console.log(taskdetail)


    } else {
        console.log("error:" + detail.msg)

    }

}

</script>

<style scoped>
/* 横线 */
.line {
    float: right;
    width: 100%;
    height: 1px;
    margin-top: -0.5em;
    background: #d4c4c4;
    position: relative;
    text-align: center;
}
</style>