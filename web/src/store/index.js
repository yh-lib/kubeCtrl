// 如何从pinia中提取数据？
// step1: 安装pinia
import { defineStore } from 'pinia'

const defaultWorkLoadItem = () => ({
  clusterId: '',
  nameSpace: '',
  name: '',
  item: {
    kind: 'Deployment',
    apiVersion: 'apps/v1',
    metadata: {
      name: '',
      namespace: '',
      labels: {},
      annotations: {},
    },
    spec: {
      replicas: 3,
      selector: {
        matchLabels: {},
      },
      template: {
        metadata: {
          labels: {},
          annotations: {},
        },
        spec: {
          nodeSelector: {},
          hostNetwork: false,
          containers: [
            {
              name: 'Container-0',
              ports: [],
              env: [],
              envFrom: [],
              volumeMounts: [],
              livenessProbe: {},
              resources: {
                requests: {},
                limits: {},
              },
              readinessProbe: {},
              startupProbe: {},
              lifecycle: {
                postStart: {},
                preStop: {},
              },
            },
          ],
          volumes: [],
          restartPolicy: '',
          terminationGracePeriodSeconds: 30,
          dnsPolicy: 'Default',
          securityContext: {},
          schedulerName: 'default-scheduler',
          tolerations: [],
          imagePullSecrets: [{ name: '' }],
        },
      },
      strategy: {
        type: 'RollingUpdate',
        rollingUpdate: {
          maxUnavailable: '25%',
          maxSurge: '25%',
        },
      },
      revisionHistoryLimit: 10,
      progressDeadlineSeconds: 600,
    },
  },
})

// step2: 创建一个容器
export const useDemoStore = defineStore('demo', {
  // 容器的内容
  // state: 用来存储全局状态/数据，可以理解为数据配置的位置
  // data
  state: () => {
    return {
      msg: 'value1',
      msg2: 'value2',
    }
  },
  getters: {
    // getters: 用来存储计算属性，可以理解为对数据进行加工处理的方法 相当于Vue中的computed计算属性
  },
  actions: {
    // actions: 用来存储修改全局状态/数据的方法，可以理解为修改数据的方法 相当于Vue中的methods方法
    // step11: 定义修改msg数据的方法
    changeMsg(v1, v2) {
      this.msg = v1
      this.msg2 = v2
    },
  },
})

// 左侧菜单栏状态
export const useIsCollapse = defineStore('isCollapse', {
  // 容器的内容
  // state: 用来存储全局状态/数据，可以理解为数据配置的位置
  // data
  state: () => {
    return {
      isCollapse: false,
    }
  },
  getters: {},
  actions: {
    changeisCollapse() {
      this.isCollapse = !this.isCollapse
    },
  },
})

// 工作负载数据模型
export const useWorkLoadData = defineStore('workLoadData', {
  state: () => {
    return {
      workLoadItem: defaultWorkLoadItem(),
    }
  },
  getters: {},
  actions: {
    resetWorkLoadItem() {
      this.workLoadItem = defaultWorkLoadItem()
    },
  },
})
