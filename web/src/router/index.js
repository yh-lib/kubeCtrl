// 基础导入
import { createRouter, createWebHashHistory } from "vue-router";
import { CONFIG } from "../config";
import { userRoutes } from "./user";
import { clusterRoutes } from "./cluster.js";
import { nodeRoutes } from "./node.js";
import { settingRoutes } from "./setting.js";

// 路由配置
const routes = [
    { path: "/", redirect: '/cluster/dashboard' },
    { path: "/login", component: () => import('../view/Login.vue') },
    userRoutes,
    clusterRoutes,
    nodeRoutes,
    settingRoutes
]

// 实例化路由
const router = createRouter({
    history: createWebHashHistory(),
    routes
})

// 全局守卫
router.beforeEach((to, from) => {
    // 1. 访问login页面
    if (to.path == '/login') {
        // 1.1 本地localStorage存有token，跳转至首页
        window.localStorage.getItem(CONFIG.TOKEN_NAME) && router.replace('/cluster/dashboard')
    } else {
        // 2. 访问非login页面
        // 2.1 本地localStorage没有token
        window.localStorage.getItem(CONFIG.TOKEN_NAME) || router.replace('/login')
    }
})


export default router