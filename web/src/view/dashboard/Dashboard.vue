<script setup>
  import { computed, ref, onBeforeMount, reactive } from 'vue'
  import { Refresh, CircleCheckFilled, WarningFilled } from '@element-plus/icons-vue'
  import {
    getHandler,
    getListHandler,
    getMetricsHandler,
    getStatusHandler,
  } from '../../api/generic'
  import ClusterSelect from './ClusterSelect.vue'
  import ClusterInfo from './ClusterInfo.vue'
  import NodeInfo from './NodeInfo.vue'
  import PodInfo from './PodInfo.vue'
  import Event from './Event.vue'

  const refresh = async () => {
    getClusterItem() // 集群状态
    getNodesMetrics() // 节点状态
    getPodStatus() // Pod 状态
    getEventList() // 事件状态
  }

  onBeforeMount(async () => {
    await getClusterList()
    dashboardData.clusterId = dashboardData.clusterItems[0].clusterId
    dashboardData.clusterItem = dashboardData.clusterItems[0]
    getNodesMetrics()
    getPodStatus()
    getEventList()
  })

  const dashboardData = reactive({
    // annotaions 中获取集群状态
    clusterItem: {},
    clusterItems: [],
    clusterId: '',
    // Node 状态
    nodeStatus: [],
    // Pod 状态
    podStatus: [],
    // 事件 list
    eventItems: [],
  })

  const getEventList = () => {
    dashboardData.eventItems = []
    getListHandler(dashboardData.clusterId, '', 'event').then((res) => {
      if (res.data.status === 200) {
        dashboardData.eventItems = res.data.data.items
      }
    })
  }

  const getPodStatus = () => {
    getStatusHandler(dashboardData.clusterId, '', 'pod').then((res) => {
      if (res.data.status === 200) {
        dashboardData.podStatus = res.data.data.items
      }
    })
  }
  const getClusterList = () => {
    return new Promise((resolve, reject) => {
      getListHandler('', '', 'cluster').then((res) => {
        if (res.data.status === 200) {
          dashboardData.clusterItems = res.data.data.items
          resolve() // 成功时调用 resolve
        } else {
          reject(new Error('请求失败')) // 失败时调用 reject
        }
      })
    })
  }
  const getClusterItem = () => {
    getHandler(dashboardData.clusterId, '', 'cluster', dashboardData.clusterId).then((res) => {
      if (res.data.status === 200) {
        dashboardData.clusterItem = res.data.data.item
      }
    })
  }
  const getNodesMetrics = () => {
    return new Promise((resolve, reject) => {
      getMetricsHandler(dashboardData.clusterId, '', 'node').then((res) => {
        if (res.data.status === 200) {
          dashboardData.nodeStatus = res.data.data.items
          resolve() // 成功时调用 resolve
          console.log('父组件：：：', dashboardData.clusterId, dashboardData.nodeItems)
        } else {
          reject(new Error('请求失败')) // 失败时调用 reject
        }
      })
    })
  }
</script>

<template>
  <div class="dashboard-page">
    <!-- dashboard main 第1行 -->
    <cluster-select
      :dashboard-data="dashboardData"
      @get-cluster-list="getClusterList"
      @cluster-change="refresh"
      @refresh="refresh"
      style="height: 130px"
    />
    <!-- dashboard main 第2行 -->
    <cluster-info :dashboard-data="dashboardData" style="height: 100px" />
    <!-- dashboard main 第3行 -->
    <section
      style="
        display: grid;
        grid-template-columns: 1.4fr 1fr;
        gap: 16px;
        margin-top: 26px;
        height: 260px;
      "
    >
      <node-info :dashboard-data="dashboardData" />
      <pod-info :dashboard-data="dashboardData" />
    </section>
    <!-- dashboard main 第4行 -->
    <section style="margin-top: 14px; width: 100%">
      <event :dashboard-data="dashboardData"></event>
    </section>
  </div>
</template>

<style scoped>
  .page-header h1 {
    margin: 8px 0 8px;
    font-size: 28px;
    line-height: 1.1;
    color: #111827;
  }

  .page-header p {
    margin: 0;
    color: #6b7280;
    font-size: 14px;
  }

  .stats-grid {
    display: grid;
    grid-template-columns: repeat(4, minmax(0, 1fr));
    gap: 16px;
    margin-top: 14px;
  }

  .info-card,
  .status-card,
  .section-card {
    border: 1px solid #e6eaf0;
    border-radius: 22px;
    background: #ffffff;
    box-shadow: 0 12px 24px rgba(15, 23, 42, 0.035);
  }

  .info-card,
  .status-card {
    padding: 16px;
  }

  .section-card {
    padding: 18px;
  }

  .table-progress {
    padding-right: 8px;
  }
</style>
