<script lang="ts" setup>
  import { reactive, onBeforeMount, computed } from 'vue'
  import ElCard from '../components/ElCard.vue'
  import { getListHandler } from '../../api/generic.js'
  import Table from './Table.vue'

  // 获取Elcard上传的数据
  const getSelectValue = (selectValue) => {
    Object.assign(data, selectValue)
    getEventItems()
  }
  const getEventItems = () => {
    data.items = []
    getListHandler(data.clusterId, data.nameSpace, 'event').then((res) => {
      if (res.data.status == 200) {
        data.items = res.data.data.items
      }
    })
  }

  // 需要的数据变量
  const data = reactive({
    items: [], // 后端返回的 list
    item: {}, // 后端返回的 obj
    search: '', // 接收 header 搜索框中数据
    curHostName: '',
    // 集群选择器
    clusterOptions: [], // 集群选择 options
    clusterId: '', // 当前选择的集群id
    nameSpace: '',
  })

  // // // 子组件加载前自动获取数据
  onBeforeMount(async () => {
    await getclusterOptions() // 获取集群列表
    data.clusterId = data.clusterOptions[0].clusterId
    getEventItems()
  })

  // 事件过滤
  const filterTableData = computed(() =>
    data.items.filter(
      (item) =>
        !data.search ||
        item.type.toLowerCase().includes(data.search.toLowerCase()) ||
        item.message.toLowerCase().includes(data.search.toLowerCase()) ||
        item.metadata.namespace.toLowerCase().includes(data.search.toLowerCase()) ||
        item.involvedObject.kind.toLowerCase().includes(data.search.toLowerCase()) ||
        item.involvedObject.name.toLowerCase().includes(data.search.toLowerCase()) ||
        item.source.component.toLowerCase().includes(data.search.toLowerCase()) ||
        (item.source.host && item.source.host.toLowerCase().includes(data.search.toLowerCase()))
    )
  )

  // 获取集群列表
  const getclusterOptions = async () => {
    await getListHandler('', '', 'cluster').then((res) => {
      if (res.data.status === 200) {
        data.clusterOptions = res.data.data.items
      }
    })
  }
</script>

<template>
  <ElCard
    title="事件列表"
    :op-cluster="true"
    :op-ns="true"
    :op-search="true"
    :op-refresh="true"
    @change="getSelectValue"
    @refresh="getEventItems"
  >
    <template #mainData>
      <Table :table-data="filterTableData" v-if="data.items != null"></Table>
    </template>
  </ElCard>
</template>

<style lang="less" scoped>
  :deep(.no-border-input .el-input__wrapper) {
    box-shadow: none;
    border: none;
  }
</style>
