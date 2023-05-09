//import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
// 引入封装的router
import router from '@/router/index'
//
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import 'element-plus/dist/index.css'

const app = createApp(App)

app
    .use(router)
    .use(ElementPlus, { locale: zhCn })
    .mount('#app')

