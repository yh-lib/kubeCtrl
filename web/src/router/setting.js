// /node/dashboard
export const settingRoutes = {
  path: '/setting',
  component: () => import('../view/layout/Layout.vue'),
  children: [    
    {
      path: 'rbac',
      component: () => import('../view/setting/Rbac.vue'),
    },
    {
      path: 'user',
      component: () => import('../view/user/List.vue'),
    },
    {
      path: 'events',
      component: () => import('../view/event/Events.vue'),
    },
  ],
}
