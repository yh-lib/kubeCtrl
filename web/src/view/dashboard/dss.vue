<script setup>
  import { computed, ref, onBeforeMount, reactive } from 'vue'
  import { Refresh, CircleCheckFilled, WarningFilled } from '@element-plus/icons-vue'
  import { getListHandler } from '../../api/generic'

  onBeforeMount(async () => {
    await getClusterList()
    console.log('测试数据：：：', data.clusterItems)
  })

  const data = reactive({
    clusterItems: [],
    curClusterItem: {},
    curClusterId: '',
  })

  // 获取集群列表
  const getClusterList = async () => {
    await getListHandler('', '', 'cluster').then((res) => {
      if (res.data.status === 200) {
        data.clusterItems = res.data.data.items // 获取集群列表
        data.curClusterId = data.clusterItems[0].clusterId // 默认集群ID
      }
    })
  }

  // 获取集群详情
  const handleClusterChange = async (clusterId) => {
    data.curClusterItem = data.clusterItems.find((item) => item.clusterId === clusterId) || {}
  }

  // 获取节点列表
  const getNodeList = () => {}

  // ++++++++++++++++++++++++++++++++++++++
  const clusterOptions = [
    { clusterId: 'cluster-01', clusterAlias: '生产集群 A' },
    { clusterId: 'cluster-02', clusterAlias: '测试集群 B' },
    { clusterId: 'cluster-03', clusterAlias: '开发集群 C' },
  ]

  const dashboardMap = {
    'cluster-01': {
      clusterInfo: {
        nodeCount: 6,
        podCount: 84,
        namespaceCount: 12,
        version: 'v1.28.3',
      },
      nodeStatus: [
        { name: 'prod-node-01', cpuUsage: 48, memoryUsage: 62 },
        { name: 'prod-node-02', cpuUsage: 39, memoryUsage: 53 },
        { name: 'prod-node-03', cpuUsage: 45, memoryUsage: 58 },
        { name: 'prod-node-04', cpuUsage: 29, memoryUsage: 42 },
      ],
      podStatus: {
        running: 78,
        failed: 1,
        pending: 4,
        succeeded: 1,
      },
      events: [
        { time: '10:32', type: 'Normal', reason: 'NodeReady', message: 'prod-node-02 已恢复就绪' },
        {
          time: '10:18',
          type: 'Warning',
          reason: 'Unhealthy',
          message: 'prod-node-04 存在短暂网络抖动',
        },
        {
          time: '09:50',
          type: 'Normal',
          reason: 'ScalingReplicaSet',
          message: '前端工作负载完成扩容',
        },
      ],
    },
    'cluster-02': {
      clusterInfo: {
        nodeCount: 4,
        podCount: 41,
        namespaceCount: 8,
        version: 'v1.27.7',
      },
      nodeStatus: [
        { name: 'stage-node-01', cpuUsage: 31, memoryUsage: 39 },
        { name: 'stage-node-02', cpuUsage: 29, memoryUsage: 35 },
        { name: 'stage-node-03', cpuUsage: 38, memoryUsage: 49 },
        { name: 'stage-node-04', cpuUsage: 19, memoryUsage: 28 },
      ],
      podStatus: {
        running: 38,
        failed: 1,
        pending: 2,
        succeeded: 0,
      },
      events: [
        {
          time: '11:05',
          type: 'Normal',
          reason: 'NodeReady',
          message: 'stage-node-03 完成维护后恢复',
        },
        {
          time: '10:41',
          type: 'Warning',
          reason: 'BackOff',
          message: 'qa 命名空间中的某个 Pod 拉起失败',
        },
        { time: '10:11', type: 'Normal', reason: 'Pulled', message: '镜像拉取完成，任务开始运行' },
      ],
    },
    'cluster-03': {
      clusterInfo: {
        nodeCount: 3,
        podCount: 26,
        namespaceCount: 6,
        version: 'v1.29.1',
      },
      nodeStatus: [
        { name: 'dev-node-01', cpuUsage: 24, memoryUsage: 31 },
        { name: 'dev-node-02', cpuUsage: 19, memoryUsage: 27 },
        { name: 'dev-node-03', cpuUsage: 22, memoryUsage: 34 },
      ],
      podStatus: {
        running: 24,
        failed: 1,
        pending: 1,
        succeeded: 0,
      },
      events: [
        {
          time: '09:58',
          type: 'Normal',
          reason: 'NodeReady',
          message: '开发节点完成重启并恢复正常',
        },
        { time: '09:34', type: 'Normal', reason: 'Created', message: '新版本服务已部署到开发集群' },
        {
          time: '08:59',
          type: 'Warning',
          reason: 'Evicted',
          message: '一个临时 Pod 因资源不足被驱逐',
        },
      ],
    },
  }

  const selectedClusterId = ref(clusterOptions[0].clusterId)
  const dashboard = ref(dashboardMap[selectedClusterId.value])

  const dashboardEvents = computed(() => dashboard.value.events || [])

  const clampPercent = (value) => Math.max(0, Math.min(100, Number(value || 0)))
  const clusterLabel = (item) => item?.clusterAlias || item?.clusterId || '未命名集群'

  const loadDashboard = async (clusterId) => {
    window.setTimeout(() => {
      dashboard.value = dashboardMap[clusterId] || dashboardMap[clusterOptions[0].clusterId]
    }, 180)
  }

  const refreshDashboard = async () => {
    await loadDashboard(selectedClusterId.value)
  }
</script>

<template>
  <div class="dashboard-page">
    <!-- 获取集群信息 -->
    <section class="page-header">
      <div>
        <div class="page-kicker">Cluster Dashboard</div>
        <h1>{{ data.curClusterItem.clusterAlias }}</h1>
        <p>
          集群概览
          <span class="hero-dot">·</span>
          {{ data.curClusterItem.clusterVersion }}
          <span class="hero-dot">·</span>
          白色主题
        </p>
      </div>

      <div class="hero-actions">
        <el-select
          v-model="data.curClusterId"
          class="cluster-select"
          filterable
          placeholder="选择集群"
          @change="handleClusterChange"
        >
          <el-option
            v-for="item in data.clusterItems"
            :key="item.clusterId"
            :label="item.clusterAlias"
            :value="item.clusterId"
          />
        </el-select>
        <el-button style="border-radius: 12px" :icon="Refresh" @click="getClusterList">
          刷新
        </el-button>
      </div>
    </section>

    <section class="stats-grid">
      <article class="info-card">
        <div class="label">节点数</div>
        <div class="value">{{ dashboard.clusterInfo.nodeCount }}</div>
      </article>

      <article class="info-card">
        <div class="label">Pod 数</div>
        <div class="value">{{ dashboard.clusterInfo.podCount }}</div>
      </article>

      <article class="info-card">
        <div class="label">命名空间数</div>
        <div class="value">{{ dashboard.clusterInfo.namespaceCount }}</div>
      </article>

      <article class="info-card">
        <div class="label">版本</div>
        <div class="value">{{ dashboard.clusterInfo.version }}</div>
      </article>
    </section>

    <section class="content-grid">
      <article class="section-card">
        <h2>节点状态</h2>
        <el-table :data="dashboard.nodeStatus" style="width: 100%">
          <el-table-column prop="name" label="节点名称" min-width="180" />
          <el-table-column prop="cpuUsage" label="CPU 使用率" min-width="200">
            <template #default="scope">
              <div class="table-progress">
                <el-progress :percentage="clampPercent(scope.row.cpuUsage)" />
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="memoryUsage" label="内存使用率" min-width="220">
            <template #default="scope">
              <div class="table-progress">
                <el-progress :percentage="clampPercent(scope.row.memoryUsage)" color="#d97706" />
              </div>
            </template>
          </el-table-column>
        </el-table>
      </article>

      <article class="section-card">
        <h2>Pod 状态</h2>
        <el-row :gutter="16">
          <el-col :span="6">
            <el-card shadow="never" class="status-card">
              <div class="label">运行中</div>
              <div class="value">{{ dashboard.podStatus.running }}</div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="never" class="status-card">
              <div class="label">失败</div>
              <div class="value">{{ dashboard.podStatus.failed }}</div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="never" class="status-card">
              <div class="label">Pending</div>
              <div class="value">{{ dashboard.podStatus.pending }}</div>
            </el-card>
          </el-col>
          <el-col :span="6">
            <el-card shadow="never" class="status-card">
              <div class="label">Succeeded</div>
              <div class="value">{{ dashboard.podStatus.succeeded }}</div>
            </el-card>
          </el-col>
        </el-row>
      </article>
    </section>

    <section class="events-section">
      <el-card shadow="never" class="section-card events-card">
        <div class="events-title">当前集群事件</div>
        <el-timeline>
          <el-timeline-item
            v-for="event in dashboardEvents"
            :key="`${event.time}-${event.reason}`"
            :type="event.type === 'Warning' ? 'warning' : 'primary'"
            :timestamp="event.time"
            placement="top"
          >
            <el-card shadow="never" class="event-card">
              <div class="event-head">
                <span class="event-reason">{{ event.reason }}</span>
                <el-tag :type="event.type === 'Warning' ? 'warning' : 'success'" effect="light">
                  {{ event.type }}
                </el-tag>
              </div>
              <div class="event-message">{{ event.message }}</div>
            </el-card>
          </el-timeline-item>
        </el-timeline>
      </el-card>
    </section>
  </div>
</template>

<style scoped>
  .dashboard-page {
    height: 100%;
    padding: 18px 0 20px;
    box-sizing: border-box;
    overflow: hidden;
  }

  .page-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 24px;
    padding: 20px 22px;
    border-radius: 24px;
    background: #ffffff;
    border: 1px solid #e6eaf0;
    box-shadow: 0 12px 28px rgba(15, 23, 42, 0.04);
  }

  .page-kicker {
    font-size: 12px;
    letter-spacing: 0.28em;
    text-transform: uppercase;
    color: #7c8798;
  }

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

  .hero-dot {
    margin: 0 10px;
  }

  .hero-actions {
    display: flex;
    align-items: center;
    gap: 12px;
    flex-shrink: 0;
    padding: 12px;
    border-radius: 16px;
    border: 1px solid #e6eaf0;
  }

  .cluster-select {
    width: 260px;
  }

  .cluster-select :deep(.el-input__wrapper) {
    background: #ffffff;
    box-shadow: inset 0 0 0 1px #d7dde6;
    border-radius: 12px;
  }

  .cluster-select :deep(.el-input__inner) {
    color: #111827;
  }

  .cluster-select :deep(.el-select__placeholder) {
    color: #64748b;
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

  .label {
    font-size: 14px;
    color: #6b7280;
  }

  .value {
    margin-top: 10px;
    font-size: 28px;
    font-weight: 700;
    color: #111827;
  }

  .content-grid {
    display: grid;
    grid-template-columns: 1.4fr 1fr;
    gap: 16px;
    margin-top: 14px;
    width: 100%;
  }

  .events-section {
    margin-top: 14px;
    width: 100%;
  }

  .events-card {
    width: 100%;
    height: 620px;
    min-width: 0;
    box-sizing: border-box;
  }

  .section-card h2 {
    margin: 0 0 14px;
    font-size: 18px;
    color: #111827;
  }

  .events-title {
    margin-bottom: 12px;
    font-size: 14px;
    font-weight: 700;
    color: #111827;
  }

  .event-card {
    border-radius: 16px;
    border: 1px solid #e6eaf0;
    background: #ffffff;
  }

  .event-head {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 12px;
  }

  .event-reason {
    font-size: 14px;
    font-weight: 700;
    color: #111827;
  }

  .event-message {
    margin-top: 8px;
    font-size: 13px;
    color: #6b7280;
  }

  :deep(.el-progress__text) {
    min-width: 44px;
  }

  @media (max-width: 1200px) {
    .content-grid {
      grid-template-columns: 1fr;
    }
  }

  @media (max-width: 960px) {
    .dashboard-page {
      height: auto;
      overflow: auto;
    }
    event-reason .page-header {
      flex-direction: column;
      align-items: stretch;
    }

    .hero-actions {
      width: 100%;
      flex-direction: column;
      align-items: stretch;
    }

    .cluster-select {
      width: 100%;
    }

    .stats-grid {
      grid-template-columns: 1fr 1fr;
    }

    .content-grid {
      grid-template-columns: 1fr;
    }
  }
</style>
