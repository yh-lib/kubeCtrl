<script setup>
  import { reactive } from 'vue'
  import TabOfBasic from './TabOfBasic.vue'
  import TabOfPort from './TabOfPort.vue'
  const data = reactive({
    activeName: 'Basic',
  })
  const props = defineProps(['containerItem'])
</script>

<template>
  <el-tabs tab-position="left" style="height: 560px" v-model="data.activeName">
    <el-tab-pane label="基础配置" name="Basic">
      <TabOfBasic style="padding: 20px 40px" :container-item="props.containerItem" />
    </el-tab-pane>
    <el-tab-pane label="端口配置" name="Port">
      <TabOfPort style="padding: 20px 40px" :container-item="props.containerItem" />
    </el-tab-pane>
    <el-tab-pane label="健康检查" name="Health">Health</el-tab-pane>
    <el-tab-pane label="环境变量" name="Env">Env</el-tab-pane>
    <el-tab-pane label="存储配置" name="Volume">Volume</el-tab-pane>
    <el-tab-pane label="生命周期" name="LifeCycle">LifeCycle</el-tab-pane>
  </el-tabs>
</template>
