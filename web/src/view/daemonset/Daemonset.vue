<script setup>
  import ElCard from '../components/ElCard.vue'
  import Table from './Table.vue'
  import { reactive, ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import DialogOfItem from '../components/workLoads/DialogOfItem.vue'
  import { deleteDaemonSetHandler, getDaemonSetListHandler } from '../../api/daemonSet.js'

  // 删除 DaemonSet
  const deleteItem = (row) => {
    // 删除提醒
    ElMessageBox.confirm('确认删除 DaemonSet :  ' + row.metadata.name, {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'warning',
    })
      .then(() => {
        deleteDaemonSetHandler(data.clusterId, data.nameSpace, row.metadata.name).then((res) => {
          if (res.data.status == 200) {
            ElMessage({
              type: 'success',
              message: res.data.message,
            })
            getList()
          }
        })
      })
      .catch(() => {
        return
      })
  }

  // 渲染表格数据
  const getList = () => {
    if (!data.clusterId || !data.nameSpace) {
      data.items = []
      return
    }
    getDaemonSetListHandler(data.clusterId, data.nameSpace).then((res) => {
      data.items = res.data.data.items || []
    })
  }

  // 从子组件 ELCard 中获取所需数据
  const getSelectValue = (selectValue) => {
    Object.assign(data, selectValue)
    getList()
  }
  const closeDialogOfItem = () => {
    createItemDialogVisible.value = false
    getList()
  }

  // 接受子组件table传递的参数
  const data = reactive({})
  const createItemDialogVisible = ref(false)
</script>

<template>
  <!-- 卡片主体: 挂载前会自动获取当前的选择，并通过getSelectValue赋值给data -->
  <ElCard
    title="DaemonSet 列表"
    :op-cluster="true"
    :op-ns="true"
    :op-search="true"
    :op-refresh="true"
    :op-create="true"
    @change="getSelectValue"
    @refresh="getList"
    @create-item="createItemDialogVisible = true"
  >
    <!-- 卡片 main 部分 table 数据 -->
    <template #mainData>
      <Table :table-data="data" @delete-item="deleteItem" @get-list="getList"></Table>
    </template>
  </ElCard>

  <!-- 创建 item 按钮 -->
  <DialogOfItem
    :open-dialog="createItemDialogVisible"
    resource-type="DaemonSet"
    action-method="create"
    @close-dialog="closeDialogOfItem"
    @get-list="getList"
  />
</template>
