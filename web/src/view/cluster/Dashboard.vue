<template>
    <div class="k8s-dashboard">
        <!-- 集群信息 -->
        <div class="section">
          <h2>集群信息</h2>
          <el-row :gutter="20">
            <el-col :span="6">
              <div class="info-card">
                <div class="label">节点数</div>
                <div class="value">{{ clusterInfo.nodeCount }}</div>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="info-card">
                <div class="label">Pod 数</div>
                <div class="value">{{ clusterInfo.podCount }}</div>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="info-card">
                <div class="label">命名空间数</div>
                <div class="value">{{ clusterInfo.namespaceCount }}</div>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="info-card">
                <div class="label">版本</div>
                <div class="value">{{ clusterInfo.version }}</div>
              </div>
            </el-col>
          </el-row>
        </div>

        <!-- 节点状态 -->
        <div class="section">
          <h2>节点状态</h2>
          <el-table :data="nodeStatus" style="width: 100%">
            <el-table-column prop="name" label="节点名称" width="180"></el-table-column>
            <el-table-column prop="cpuUsage" label="CPU 使用率" width="180">
              <template #default="scope">
                <el-progress :percentage="scope.row.cpuUsage" color="#409EFF"></el-progress>
              </template>
            </el-table-column>
            <el-table-column prop="memoryUsage" label="内存使用率" width="180">
              <template #default="scope">
                <el-progress :percentage="scope.row.memoryUsage" color="#67C23A"></el-progress>
              </template>
            </el-table-column>
          </el-table>
        </div>

        <!-- Pod 状态 -->
        <div class="section">
          <h2>Pod 状态</h2>
          <el-row :gutter="20">
            <el-col :span="6">
              <div class="status-card">
                <div class="label">运行中</div>
                <div class="value">{{ podStatus.running }}</div>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="status-card">
                <div class="label">失败</div>
                <div class="value">{{ podStatus.failed }}</div>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="status-card">
                <div class="label">Pending</div>
                <div class="value">{{ podStatus.pending }}</div>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="status-card">
                <div class="label">Succeeded</div>
                <div class="value">{{ podStatus.succeeded }}</div>
              </div>
            </el-col>
          </el-row>
        </div>
      </div> 
</template>

<script setup>
import { ref } from 'vue';

// 假数据：集群信息
const clusterInfo = ref({
  nodeCount: 5,
  podCount: 20,
  namespaceCount: 10,
  version: 'v1.24.0',
});

// 假数据：节点状态
const nodeStatus = ref([
  { name: 'node-1', cpuUsage: 30, memoryUsage: 45 },
  { name: 'node-2', cpuUsage: 60, memoryUsage: 70 },
  { name: 'node-3', cpuUsage: 10, memoryUsage: 20 },
]);

// 假数据：Pod 状态
const podStatus = ref({
  running: 15,
  failed: 2,
  pending: 3,
  succeeded: 0,
});
</script>

<style scoped>
.k8s-dashboard {
  margin-top: 50px;
  padding: 20px;
}

.section {
  margin-bottom: 30px;
}

.info-card,
.status-card {
  border: 1px solid #ebebeb;
  border-radius: 8px;
  padding: 15px;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.label {
  font-size: 14px;
  color: #888;
}

.value {
  font-size: 20px;
  font-weight: bold;
  color: #333;
}

.el-table {
  margin-top: 10px;
}
</style>