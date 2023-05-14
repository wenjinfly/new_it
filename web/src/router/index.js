import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
    {
        path: '/',
        redirect: '/home'
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/login/index.vue')
    },
    {
        path: '/home',
        name: 'home',
        component: () => import('@/views/task/index.vue')
    },
    {
        path: "/404",
        name: "NotFound",
        component: () => import('@/views/base/404.vue'),
        meta: { title: 'Page not found', isLogin: false },
    }

]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

// 路由拦截
router.beforeEach((to, from, next) => {
    if (to.matched.length === 0) {
        // 不存在的路由地址
        next('/404')
    } else {
        next()
    }
})

export default router