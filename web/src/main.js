//import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
// 引入封装的router
import router from '@/router/index'
//
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import 'element-plus/dist/index.css'
import { createPinia } from "pinia";

const app = createApp(App)
const pinia = createPinia();

app
    .use(router)
    .use(pinia)
    .use(ElementPlus, { locale: zhCn })
    .mount('#app')

