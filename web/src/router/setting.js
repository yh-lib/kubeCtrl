// /node/dashboard
export const settingRoutes = {
  path: '/setting',
  component: () => import('../view/layout/Layout.vue'),
  children: [
    {
      path: 'namespaces',
      component: () => import('../view/namespace/Namespace.vue'),
    },
    {
      path: 'rbac',
      component: () => import('../view/setting/Rbac.vue'),
    },
    {
      path: 'user',
      component: () => import('../view/user/List.vue'),
    },
  ],
}
