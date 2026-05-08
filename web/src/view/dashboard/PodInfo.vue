<script setup>
  const props = defineProps(['itemForm'])
</script>

<template>
  <article class="section-card">
    <h2>Pod 状态</h2>
    <el-row :gutter="16">
      <el-col :span="6">
        <el-card shadow="never" class="status-card">
          <div class="label">运行中</div>
          <div class="value">{{ itemForm.podStatus.running }}</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never" class="status-card">
          <div class="label">失败</div>
          <div class="value">{{ itemForm.podStatus.failed }}</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never" class="status-card">
          <div class="label">Pending</div>
          <div class="value">{{ itemForm.podStatus.pending }}</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never" class="status-card">
          <div class="label">Succeeded</div>
          <div class="value">{{ itemForm.podStatus.succeeded }}</div>
        </el-card>
      </el-col>
    </el-row>
  </article>
</template>

<style scoped></style>
