<template>
    <h3> persion info</h3>
    <p>姓名：{{ name }}</p>
    <p>年龄：{{ age }}</p>
    <p>性别：{{ sex }}</p>
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
</script>

