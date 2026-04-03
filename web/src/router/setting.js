// /node/dashboard
export const settingRoutes = {
    path: "/setting",
    component: () => import('../view/layout/Layout.vue'),
    children: [
        {
            path: "namespaces",
            component: () => import('../view/setting/Namespace.vue'),
        }
    ]
}