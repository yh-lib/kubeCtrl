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
          <div class="label">创建中</div>
          <div class="value">{{ itemForm.podStatus.pending }}</div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never" class="status-card">
          <div class="label">已成功</div>
          <div class="value">{{ itemForm.podStatus.succeeded }}</div>
        </el-card>
      </el-col>
    </el-row>
  </article>
</template>

<style scoped>
  .section-card h2 {
    margin: 0 0 14px;
    font-size: 18px;
    color: #111827;
  }
  .status-card {
    padding: 16px;
  }

  .status-card,
  .section-card {
    border: 1px solid #e6eaf0;
    border-radius: 22px;
    background: #ffffff;
  }

  .label {
    font-size: 16px; /* 字体大小 */
    color: #6b7280; /* 灰色文字 */
    margin-bottom: 8px; /* 底部间距 */
  }
  .value {
    font-size: 32px; /* 字体大小 */
    font-weight: 700; /* 加粗 */
    color: #111827; /* 深灰色文字 */
  }
</style>
