<script lang="ts" setup>
  import { computed, reactive, ref, onBeforeMount } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import ElCard from '../components/ElCard.vue'
  import { deleteHandler, getHandler, getListHandler } from '../../api/generic'
  import Table from './table.vue'
  import Create from './create.vue'

  const opDialog = ref(false)
  const createItem = () => {
    method.value = 'create'
    data.itemForm = {}
    opDialog.value = true
  }

  const data = reactive({
    tableData: [] as Cluster[],
    itemForm: {
      clusterId: '',
      clusterAlias: '',
      clusterVersion: '',
      clusterLocation: '',
      clusterStatus: '',
      clusterSize: '',
      clusterNodeNum: '',
      clusterPodNum: '',
      kubeconfig: '',
    },
  })

  const refreshList = () => {
    opDialog.value = false
    getList()
  }

  const getList = () => {
    getListHandler('', '', 'cluster', true).then((response) => {
      if (response.data.status === 200) {
        data.tableData = response.data.data.items
      }
    })
  }

  const closeDialog = () => {
    method.value == 'create' && getList()
  }

  const deleteItem = (row) => {
    console.log('删除集群 res:', row.clusterId)
    // 删除提醒
    ElMessageBox.confirm('确认删除集群：' + row.clusterId, {
      confirmButtonText: '确认',
      cancelButtonText: '取消',
      type: 'warning',
    })
      .then(() => {
        deleteHandler(row.clusterId, 'kc', 'cluster', row.clusterId).then((res) => {
          if (res.data.status == 200) {
            ElMessage({
              type: 'success',
              message: res.data.message,
            })
            // 刷新列表
            getList()
          }
        })
      })
      .catch(() => {
        return
      })
  }

  const method = ref('')
  const getItem = (row) => {
    // 获取集群详情
    getHandler(row.clusterId, '', 'cluster').then((response) => {
      data.itemForm = response.data.data.item
      // 注意下面这两步：如果放在axios外面，则赋值可能失败，因为是异步运行
      // 传递给操作参数给子组件
      method.value = 'update'
      // 打开表单弹窗
      opDialog.value = true
    })
  }
  onBeforeMount(() => {
    getList()
  })
  // 搜索功能
  const search = ref('')
  const tableData = computed(() =>
    data.tableData.filter(
      (data) => !search.value || data.clusterId.toLowerCase().includes(search.value.toLowerCase())
    )
  )

  const handleHeaderChange = (headerData) => {
    search.value = headerData.search || ''
  }

  // 集群信息接口
  interface Cluster {
    clusterId: string
    clusterAlias: string
    clusterVersion: string
    clusterLocation: string
    clusterStatus: string
    clusterSize: number
    clusterNodeNum: number
    clusterPodNum: number
  }
</script>

<template>
  <ElCard
    title="集群列表"
    :op-search="true"
    :op-create="true"
    :op-refresh="true"
    @change="handleHeaderChange"
    @create-item="createItem"
    @refresh="getList"
  >
    <template #mainData>
      <Table :table-data="tableData" @get-item="getItem" @delete-item="deleteItem"></Table>
      <el-dialog
        v-model="opDialog"
        :title="method == 'create' ? '添加集群' : '更新集群'"
        width="1000px"
        @close="closeDialog"
        destroy-on-close
        :close-on-click-modal="false"
      >
        <Create :subMethod="method" :subRow="data.itemForm" @refresh="refreshList" />
      </el-dialog>
    </template>
  </ElCard>
</template>

<style scoped></style>
