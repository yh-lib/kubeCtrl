<script setup>
  const props = defineProps(['dashboardData'])
</script>
<template>
  <section class="stats-grid">
    <article class="info-card">
      <div class="label">节点数</div>
      <div class="value">{{ props.dashboardData.clusterItem.clusterNodeNum }}</div>
    </article>

    <article class="info-card">
      <div class="label">Pod 数</div>
      <div class="value">{{ props.dashboardData.clusterItem.clusterPodNum }}</div>
    </article>

    <article class="info-card">
      <div class="label">命名空间数</div>
      <div class="value">{{ props.dashboardData.clusterItem.clusterNsNum }}</div>
    </article>

    <article class="info-card">
      <div class="label">版本</div>
      <div class="value">{{ props.dashboardData.clusterItem.clusterVersion }}</div>
    </article>
  </section>
</template>

<style scoped>
  .info-card,
  .status-card {
    padding: 16px;
  }

  .info-card,
  .status-card,
  .section-card {
    border: 1px solid #e6eaf0;
    border-radius: 22px;
    background: #ffffff;
    padding: 16px;
    text-align: center;
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
</style>
