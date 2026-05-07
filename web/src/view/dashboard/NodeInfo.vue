<script setup>
  const props = defineProps(['tableData'])
</script>

<template>
  <article style="border: 1px solid #e6eaf0; border-radius: 22px; padding: 18px">
    <!-- 标题 -->
    <h2 style="margin: 0 0 14px; font-size: 18px">节点状态</h2>
    <!-- 表格 -->
    <el-table :data="props.tableData">
      <el-table-column prop="metadata.name" label="节点名称" min-width="180" />
      <!-- <el-table-column prop="cpuUsage" label="CPU 使用率" min-width="200">
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
      </el-table-column> -->
    </el-table>
  </article>
</template>

<style scoped></style>
