<script setup>
  import { ref, watch } from 'vue'
  import { getHandler } from '../../api/generic'
  const props = defineProps(['itemForm'])

  const tableData = ref([])
  // 监听 itemForm.nodeItems 的变化
  watch(
    () => props.itemForm.nodeItems,
    async (newNodeItems) => {
      if (!newNodeItems || newNodeItems.length === 0) {
        tableData.value = []
        return
      }
      const result = []
      for (const item of newNodeItems) {
        // 获取主机名
        const hostName = item.metadata.name
        // 获取 CPU、Mem 使用率
        let useageCpu
        if (item.usage.cpu.slice(-1) === 'm') {
          useageCpu = item.usage.cpu.replace('m', '') // cpu 使用量
        } else if (item.usage.cpu.slice(-1) === 'n') {
          useageCpu = item.usage.cpu.replace('n', '') / 1000 / 1000 // cpu 使用量
        } else {
          useageCpu = item.usage.cpu.replace('n', '') / 1000
        }
        const useageMem = item.usage.memory.replace('Ki', '') / 1024 // mem 使用量
        // 异步获取节点容量信息
        const res = await getHandler(props.itemForm.clusterId, '', 'node', hostName)
        if (res.data.status === 200) {
          const totalCpu = res.data.data.items.status.capacity.cpu * 1000
          const totalMem = res.data.data.items.status.capacity.memory.replace('Ki', '') / 1024
          const cpuPercent = Math.round(parseFloat((useageCpu / totalCpu) * 100))
          const memPercent = Math.round(parseFloat((useageMem / totalMem) * 100))
          // 将结果推入数组
          result.push({
            hostName,
            cpuPercent,
            memPercent,
          })
        }
      }
      // 更新 tableData
      tableData.value = result
      console.log('result', result)
    },
    { immediate: true }
  )
</script>
<template>
  <article style="border: 1px solid #e6eaf0; border-radius: 22px; padding: 18px">
    <!-- 标题 -->
    <h2 style="margin: 0 0 14px; font-size: 18px">节点状态</h2>
    <!-- 表格 -->
    <el-table :data="tableData">
      <el-table-column prop="hostName" label="节点名称" min-width="180" />
      <el-table-column prop="cpuPercent" label="CPU 使用率" min-width="200">
        <template #default="scope">
          <el-progress :percentage="scope.row.cpuPercent" />
        </template>
      </el-table-column>
      <el-table-column prop="memPercent" label="内存使用率" min-width="220">
        <template #default="scope">
          <el-progress :percentage="scope.row.memPercent" color="#d97706" />
        </template>
      </el-table-column>
    </el-table>
  </article>
</template>
