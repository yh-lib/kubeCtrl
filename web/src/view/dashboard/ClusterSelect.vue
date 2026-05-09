<script setup>
  import { Refresh } from '@element-plus/icons-vue'
  const props = defineProps(['dashboardData'])
  const emit = defineEmits([
    'clusterChange',
    'getClusterList',
    'getClusterItem',
    'refresh',
    'clusterChange',
  ])
</script>

<template>
  <section
    style="
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 20px 22px;
      border: 1px solid #e6eaf0;
      border-radius: 24px;
    "
  >
    <!-- 左半区域 -->
    <div>
      <div style="font-size: 12px; letter-spacing: 0.28em; color: #7c8798">CLUSTER DASHBOARD</div>
      <h2 style="margin: 6px">{{ props.dashboardData.clusterItem.clusterAlias }}</h2>
      <p style="margin: 6px">
        集群概览
        <span style="margin: 0 10px">·</span>
        {{ props.dashboardData.clusterItem.clusterVersion }}
        <span style="margin: 0 10px">·</span>
        白色主题
      </p>
    </div>
    <!-- 右半区域 -->
    <div
      style="
        display: flex;
        align-items: center;
        gap: 12px;
        border: 1px solid #e6eaf0;
        padding: 12px;
        border-radius: 16px;
      "
    >
      <el-select
        v-model="props.dashboardData.clusterId"
        filterable
        placeholder="选择集群"
        @visible-change="emit('getClusterList')"
        @change="emit('clusterChange')"
        style="width: 280px"
      >
        <el-option
          v-for="item in props.dashboardData.clusterItems"
          :key="item.clusterId"
          :label="item.clusterAlias"
          :value="item.clusterId"
        />
      </el-select>
      <el-button style="border-radius: 12px" :icon="Refresh" @click="emit('refresh')">
        刷新
      </el-button>
    </div>
  </section>
</template>
