// /cluster/dashboard
export const clusterRoutes = {
  path: '/cluster',
  component: () => import('../view/layout/Layout.vue'),
  children: [
    {
      path: 'dashboard',
      component: () => import('../view/dashboard/Dashboard.vue'),
    },
    {
      path: 'list',
      component: () => import('../view/cluster/Cluster.vue'),
    },
    {
      path: 'namespace',
      component: () => import('../view/namespace/Namespace.vue'),
    },
    {
      path: 'node',
      component: () => import('../view/node/Node.vue'),
    },
    {
      path: 'event',
      component: () => import('../view/event/Event.vue'),
    },
  ],
}
