<script setup>
  import { utcTime2BjTime } from '../../utils/utctimeConv/utctimeConv'

  const props = defineProps(['tableData'])

  const getAgeText = (row) => {
    const creationTimestamp = row?.metadata?.creationTimestamp
    if (!creationTimestamp) {
      return '-'
    }

    const createdAt = new Date(creationTimestamp).getTime()
    if (Number.isNaN(createdAt)) {
      return '-'
    }

    const diffSeconds = Math.max(0, Math.floor((Date.now() - createdAt) / 1000))
    if (diffSeconds < 60) {
      return `${diffSeconds}秒钟前`
    }

    const diffMinutes = Math.floor(diffSeconds / 60)
    if (diffMinutes < 60) {
      return `${diffMinutes}分钟前`
    }

    const diffHours = Math.floor(diffMinutes / 60)
    if (diffHours < 24) {
      return `${diffHours}小时前`
    }

    const diffDays = Math.floor(diffHours / 24)
    return `${diffDays}天前`
  }
</script>

<template>
  <el-table :data="props.tableData" style="width: 100%" height="1010px">
    <el-table-column label="类型" prop="type" width="100px" />
    <el-table-column label="消息" prop="message" />
    <el-table-column label="命名空间" prop="metadata.namespace" width="200px" />
    <el-table-column label="涉及对象" prop="involvedObject" width="300px">
      <template #default="scope">
        {{ scope.row.involvedObject.kind }}:
        {{ scope.row.involvedObject.name }}
      </template>
    </el-table-column>
    <el-table-column label="来源" prop="source" width="200px">
      <template #default="scope">
        {{ scope.row.source.component }}
        {{ scope.row.source.host }}
      </template>
    </el-table-column>
    <el-table-column label="发生的次数" prop="count" width="100px" />
    <el-table-column label="距今" prop="age" width="100px">
      <template #default="scope">
        {{ getAgeText(scope.row) }}
      </template>
    </el-table-column>
    <el-table-column label="最后发生时间" prop="lastTimestamp" width="200px">
      <template #default="scope">
        {{ utcTime2BjTime(scope.row.lastTimestamp) }}
      </template>
    </el-table-column>

    <!-- 未开放 -->
    <!-- <el-table-column label="名称" prop="metadata.name" />
    <el-table-column label="标识符" prop="metadata.uid" />
    <el-table-column label="事件创建的时间" prop="metadata.creationTimestamp" />
    <el-table-column label="事件的原因" prop="reason" />
    <el-table-column label="首次发生时间" prop="firstTimestamp" /> -->
  </el-table>
</template>
