<script setup>
  import { onBeforeMount } from 'vue'
  import { utcTime2BjTime } from '../../utils/utctimeConv/utctimeConv'

  const props = defineProps(['dashboardData'])
  onBeforeMount(() => {
    console.log('获取事件', props.dashboardData.eventItems)
  })
</script>

<template>
  <el-card shadow="never" style="border-radius: 22px; height: 580px">
    <h2 style="margin: 0px 0px 14px 60px; font-size: 18px">当前集群事件</h2>
    <el-timeline>
      <el-timeline-item
        v-for="event in dashboardData.eventItems"
        :key="event.metadata.name"
        :type="event.type"
        :timestamp="utcTime2BjTime(event.metadata.creationTimestamp)"
        placement="top"
      >
        <el-card shadow="never" style="border-radius: 16px">
          <div style="display: flex; justify-content: space-between">
            <span style="font-size: 14px; font-weight: 700">{{ event.reason }}</span>
            <el-tag :type="event.type === 'Warning' ? 'warning' : 'success'">
              {{ event.type }}
            </el-tag>
          </div>
          <div style="margin-top: 8px; font-size: 13px; color: #6b7280">{{ event.message }}</div>
        </el-card>
      </el-timeline-item>
    </el-timeline>
  </el-card>
</template>
