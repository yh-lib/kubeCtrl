// 如何从pinia中提取数据？
// step1: 安装pinia
import { defineStore } from 'pinia'
import TcpSocket from '../view/components/workLoads/tabOfHealth/TcpSocket.vue'

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
              name: 'container-0',
              ports: [],
              env: [],
              envFrom: [],
              volumeMounts: [],
              resources: {
                requests: {},
                limits: {},
              },
              startupProbe: {
                // 通用参数
                initialDelaySeconds: 0,
                periodSeconds: 10,
                timeoutSeconds: 1,
                successThreshold: 1,
                failureThreshold: 30,
                terminationGracePeriodSeconds: 30,
                // 探测方法
                exec: { command: '' },
                tcpSocket: {},
                httpGet: {
                  httpHeaders: [],
                },
                grpc: {},
              },
              readinessProbe: {
                // 通用参数
                initialDelaySeconds: 0,
                periodSeconds: 10,
                timeoutSeconds: 1,
                successThreshold: 1,
                failureThreshold: 30,
                terminationGracePeriodSeconds: 30,
                // 探测方法
                exec: { command: [] },
                tcpSocket: {},
                httpGet: {
                  httpHeaders: [],
                },
                grpc: {},
              },
              livenessProbe: {
                // 通用参数
                initialDelaySeconds: 0,
                periodSeconds: 10,
                timeoutSeconds: 1,
                successThreshold: 1,
                failureThreshold: 30,
                terminationGracePeriodSeconds: 30,
                // 探测方法
                exec: { command: [] },
                tcpSocket: {},
                httpGet: {
                  httpHeaders: [],
                },
                grpc: {},
              },
              lifecycle: {
                postStart: {},
                preStop: {},
              },
            },
          ],
          volumes: [],
          restartPolicy: '',
          terminationGracePeriodSeconds: null,
          dnsPolicy: '',
          securityContext: {},
          schedulerName: '',
          tolerations: [],
          imagePullSecrets: [{ name: '' }],
        },
      },
      strategy: {
        type: '',
        rollingUpdate: {},
      },
      revisionHistoryLimit: null,
      progressDeadlineSeconds: null,
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
