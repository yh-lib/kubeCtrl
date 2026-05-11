<script setup lang="ts">
  import { reactive, ref } from 'vue'
  import ElCard from '../components/ElCard.vue'
  import { createHandler, deleteHandler, getListHandler } from '../../api/generic'
  import Table from './Table.vue'
  import { ElMessage, ElMessageBox } from 'element-plus'

  const itemForm = reactive({
    clusterId: '',
    nameSpace: '',
    name: '',
    item: {
      kind: 'Namespace',
      apiVersion: 'v1',
      metadata: {
        name: '',
      },
    },
  })
  // 渲染表格数据
  const getList = () => {
    if (!itemForm.clusterId) {
      data.items = []
      return
    }
    getListHandler(itemForm.clusterId, '', 'namespace').then((res) => {
      data.items = res.data.data.items || []
    })
  }
  // 从子组件 ELCard 中获取所需数据
  const getSelectValue = (selectValue) => {
    Object.assign(itemForm, selectValue)
    Object.assign(data, selectValue)
    getList()
  }
  const closeDialog = () => {
    createItemDialogVisible.value = false
    getList()
  }
  // 接受子组件table传递的参数
  const data = reactive({
    tableData: [],
  })
  const createItemDialogVisible = ref(false)

  // 删除 namespace
  const deleteItem = (resourceName) => {
    console.log('执行删除逻辑', resourceName)
    // 删除提醒
    ElMessageBox.confirm('确认删除 namespace :  ' + resourceName, {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'warning',
    })
      .then(() => {
        deleteHandler(itemForm.clusterId, itemForm.nameSpace, 'namespace', resourceName).then(
          (res) => {
            if (res.data.status == 200) {
              ElMessage({
                type: 'success',
                message: res.data.message,
              })
              getList()
            }
          }
        )
      })
      .catch(() => {
        return
      })
  }
  // 创建namespace
  const createItemConfirm = () => {
    itemForm.item.metadata.name = itemForm.name
    createHandler('namespace', itemForm).then((res) => {
      if (res.data.status === 200) {
        ElMessage({
          message: res.data.message,
          type: 'success',
        })
        data.createDialog = false
        getList()
      }
    })
  }
</script>

<template>
  <ElCard
    title="Namespace 列表"
    :op-cluster="true"
    :op-search="true"
    :op-refresh="true"
    :op-create="true"
    @change="getSelectValue"
    @refresh="getList"
    @create-item="createItemDialogVisible = true"
  >
    <template #mainData>
      <Table :table-data="data" @delete-item="deleteItem"></Table>
    </template>
  </ElCard>
  <!-- 创建 namespace 窗口 -->
  <el-dialog
    v-model="createItemDialogVisible"
    width="580px"
    style="height: 210px"
    @close="closeDialog"
    destroy-on-close
    :close-on-click-modal="false"
  >
    <!-- dialog header -->
    <template #header>
      <span style="font-size: 18px">创建命名空间</span>
    </template>
    <!-- dialog middle -->
    <div>
      <el-input
        v-model="itemForm.name"
        placeholder="请输入命名空间名称"
        style="width: 400px; margin-top: 10px"
        size="large"
      />
    </div>
    <div>
      <el-button type="primary" @click="createItemConfirm" style="width: 105px; margin-top: 30px"
        >确认</el-button
      >
    </div>
  </el-dialog>
</template>
