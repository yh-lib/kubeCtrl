<script setup>
  const props = defineProps(['tableData'])

  const clampPercent = (value, maxRange) => {
    // 取出所有非数字字符
    if (typeof value === 'string') {
      const result = value.replace(/[^0-9]/g, '')
      const numValue = Number(result) || 0
      // 计算百分比
      let percentage = (numValue / maxRange) * 100
      // 将百分比限制在 0 到 100 之间
      percentage = Math.max(0, Math.min(100, percentage))
      // 精确到小数点后一位
      return parseFloat(percentage.toFixed(1))
    }
  }
</script>

<template>
  <article style="border: 1px solid #e6eaf0; border-radius: 22px; padding: 18px">
    <!-- 标题 -->
    <h2 style="margin: 0 0 14px; font-size: 18px">节点状态</h2>
    <!-- 表格 -->
    <el-table :data="props.tableData">
      <el-table-column prop="metadata.name" label="节点名称" min-width="180" />
      <!-- <el-table-column prop="usage.cpu" label="节点名称" min-width="180" /> -->
      <el-table-column prop="cpuUsage" label="CPU 使用率" min-width="200">
        <template #default="scope">
          <div class="table-progress">
            <el-progress :percentage="clampPercent(scope.row.usage.cpu, 1000)" />
          </div>
        </template>
      </el-table-column>
      <el-table-column prop="memoryUsage" label="内存使用率" min-width="220">
        <template #default="scope">
          <div class="table-progress">
            <el-progress
              :percentage="clampPercent(scope.row.usage.memory, 8137576)"
              color="#d97706"
            />
          </div>
        </template>
      </el-table-column>
    </el-table>
  </article>
</template>

<style scoped></style>
