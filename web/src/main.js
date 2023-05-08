import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
// 引入封装的router
import router from '@/router/index'

const app = createApp(App)

app
    .use(router)
    .mount('#app')

