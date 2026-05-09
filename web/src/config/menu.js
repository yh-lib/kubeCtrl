export const MENU_CONFIG = [
  // 1. 集群概览
  {
    title: '集群概览',
    index: '/cluster/overview',
    icon: 'Monitor', // 显示器代表监控/概览
    items: [
      {
        title: '仪表盘',
        index: '/cluster/dashboard',
        icon: 'DataAnalysis',
      },
      {
        title: '集群列表',
        index: '/cluster/list',
        icon: 'Connection',
      },
      {
        title: '节点列表',
        index: '/node/list',
        icon: 'Grid',
      },
      {
        title: '事件中心',
        index: '/cluster/events',
        icon: 'Bell',
      },
    ],
  },
  // 2. 工作负载
  {
    title: '工作负载',
    index: '/workloads',
    icon: 'Box', // 盒子代表容器/负载
    items: [
      {
        title: '容器集 (Pod)',
        index: '/workloads/pod',
        icon: 'Aim',
      },
      {
        title: '无状态集 (Deployment)',
        index: '/workloads/deployment',
        icon: 'Suitcase', // 火箭代表发布/部署
      },
      {
        title: '有状态集 (StatefulSet)',
        index: '/workloads/statefulset',
        icon: 'Coin', // 数据库代表有状态
      },
      {
        title: '守护进程集 (DaemonSet)',
        index: '/workloads/daemonset',
        icon: 'Lock', // 盾牌代表守护/安全
      },
      {
        title: '定时任务 (cronJob)',
        index: '/workloads/cronJob',
        icon: 'Timer', // 定时器代表任务
      },
    ],
  },
  // // 3. 服务与网络
  // {
  //     title: '服务与网络',
  //     index: '/network',
  //     icon: 'Connection', // 连接代表网络
  //     subMenu: [
  //         {
  //             title: '服务 (Services)',
  //             index: '/network/services',
  //             icon: 'Link', // 链接代表服务发现
  //             items: [
  //                 {
  //                     title: '创建服务',
  //                     index: '/network/services/create'
  //                 },
  //                 {
  //                     title: '服务列表',
  //                     index: '/network/services/list'
  //                 }
  //             ]
  //         },
  //         {
  //             title: '入口 (Ingresses)',
  //             index: '/network/ingresses',
  //             icon: 'Position', // 大门代表流量入口
  //             items: [
  //                 {
  //                     title: '创建入口',
  //                     index: '/network/ingresses/create'
  //                 },
  //                 {
  //                     title: '入口列表',
  //                     index: '/network/ingresses/list'
  //                 }
  //             ]
  //         },
  //         {
  //             title: '网络策略',
  //             index: '/network/policies',
  //             icon: 'Guide', // 指南针代表规则导向
  //             items: [
  //                 {
  //                     title: '策略列表',
  //                     index: '/network/policies/list'
  //                 }
  //             ]
  //         }
  //     ]
  // },
  // // 4. 存储管理
  // {
  //     title: '存储管理',
  //     index: '/storage',
  //     icon: 'Coin', // 硬盘代表存储
  //     subMenu: [
  //         {
  //             title: '持久化声明 (PVC)',
  //             index: '/storage/pvc',
  //             icon: 'Ticket', // 票据代表申请/Claim
  //             items: [
  //                 {
  //                     title: '创建 PVC',
  //                     index: '/storage/pvc/create'
  //                 },
  //                 {
  //                     title: 'PVC 列表',
  //                     index: '/storage/pvc/list'
  //                 }
  //             ]
  //         },
  //         {
  //             title: '存储类 (StorageClass)',
  //             index: '/storage/class',
  //             icon: 'List', // 列表代表分类定义
  //             items: [
  //                 {
  //                     title: '存储类列表',
  //                     index: '/storage/class/list'
  //                 }
  //             ]
  //         }
  //     ]
  // },
  // // 5. 配置与密钥
  // {
  //     title: '配置中心',
  //     index: '/config',
  //     icon: 'Edit', // 齿轮代表配置
  //     subMenu: [
  //         {
  //             title: '配置映射 (ConfigMap)',
  //             index: '/config/configmaps',
  //             icon: 'Document', // 文档代表配置文件
  //             items: [
  //                 {
  //                     title: '创建配置',
  //                     index: '/config/configmaps/create'
  //                 },
  //                 {
  //                     title: '配置列表',
  //                     index: '/config/configmaps/list'
  //                 }
  //             ]
  //         },
  //         {
  //             title: '密钥管理 (Secrets)',
  //             index: '/config/secrets',
  //             icon: 'Key', // 钥匙代表密钥
  //             items: [
  //                 {
  //                     title: '创建密钥',
  //                     index: '/config/secrets/create'
  //                 },
  //                 {
  //                     title: '密钥列表',
  //                     index: '/config/secrets/list'
  //                 }
  //             ]
  //         }
  //     ]
  // },
  // 6. 系统设置
  {
    title: '系统设置',
    index: '/setting',
    icon: 'Setting', // 工具代表系统级设置
    items: [
      {
        title: '命名空间',
        index: '/setting/namespaces',
        icon: 'Box',
      },
      {
        title: '用户管理',
        index: '/setting/user',
        icon: 'User',
      },
      // {
      //   title: '角色与权限 (RBAC)',
      //   index: '/setting/rbac',
      //   icon: 'UserFilled',
      // },
    ],
  },
]
