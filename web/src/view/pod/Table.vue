<script setup>
  import { computed, ref } from 'vue'
  import DialogByYaml from '../components/DialogByYaml.vue'
  import { obj2yaml } from '../../utils/typeConv/type.conv.js'
  import DialogHeaderLabel from '../components/DialogHeaderLabel.vue'
  import { getHandler } from '../../api/generic.js'

  const emit = defineEmits(['deleteItem'])

  const getContainerStatus = (row) => {
    const statuses = row?.status?.containerStatuses || []
    const total = statuses.length
    const ready = statuses.filter((item) => item?.ready).length
    return `${ready}/${total}`
  }

  const getRestartCount = (row) => {
    const statuses = row?.status?.containerStatuses || []
    if (!statuses.length) {
      return 0
    }
    return Math.max(...statuses.map((item) => item?.restartCount || 0))
  }

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

  const props = defineProps({
    tableData: {
      type: Object,
      default: () => ({
        items: [],
        item: {},
        curClusterId: '',
        curNsName: '',
        search: '',
      }),
    },
  })

  const removeEmptyFieldsDeep = (obj) => {
    Object.keys(obj).forEach((key) => {
      const value = obj[key]

      if (Array.isArray(value)) {
        if (value.length === 0) {
          delete obj[key]
          return
        }

        value.forEach((item) => {
          if (item && typeof item === 'object') {
            removeEmptyFieldsDeep(item)
          }
        })

        obj[key] = value.filter((item) => {
          if (!item || typeof item !== 'object' || Array.isArray(item)) {
            return true
          }

          return Object.keys(item).length > 0
        })

        if (obj[key].length === 0) {
          delete obj[key]
        }

        return
      }

      if (value && typeof value === 'object') {
        removeEmptyFieldsDeep(value)

        if (Object.keys(value).length === 0) {
          delete obj[key]
        }
        return
      }

      if (value === '' || value === null || value === undefined) {
        delete obj[key]
      }
    })

    return obj
  }
  // 搜索后的 pod 列表数据
  const filterTableData = computed(() =>
    (props.tableData.items || []).filter(
      (item) =>
        !props.tableData.search ||
        item.metadata.name.toLowerCase().includes(props.tableData.search.toLowerCase()) ||
        item.status.podIP.toLowerCase().includes(props.tableData.search.toLowerCase()) ||
        item.status.phase.toLowerCase().includes(props.tableData.search.toLowerCase()) ||
        item.spec.nodeName.includes(props.tableData.search.toLowerCase()) ||
        item.status.hostIP.toLowerCase().includes(props.tableData.search.toLowerCase())
    )
  )
  const curItem = ref('')
  const itemByYaml = ref('')
  const itemDetailDialog = ref(false)
  const getItem = (row) => {
    getHandler(props.tableData.clusterId, row.metadata.namespace, 'pod', row.metadata.name).then(
      (res) => {
        if (res.data.status == 200) {
          curItem.value = removeEmptyFieldsDeep(res.data.data.items)
          itemByYaml.value = obj2yaml(curItem.value)
          itemDetailDialog.value = true
        }
      }
    )
  }
</script>

<template>
  <el-table :data="filterTableData" height="1010px">
    <el-table-column label="名称" prop="metadata.name" width="300px">
      <template #default="scope">
        <el-button type="primary" link @click="getItem(scope.row)">{{
          scope.row.metadata.name
        }}</el-button>
      </template>
    </el-table-column>
    <el-table-column label="Pod IP" prop="status.podIP" />
    <el-table-column label="创建时间" prop="metadata.creationTimestamp" />
    <el-table-column label="存活时间" prop="age">
      <template #default="scope">
        {{ getAgeText(scope.row) }}
      </template>
    </el-table-column>
    <el-table-column label="命名空间" prop="metadata.namespace" />
    <el-table-column label="状态" prop="status.phase" />
    <el-table-column label="重启次数" prop="restartCount">
      <template #default="scope">
        {{ getRestartCount(scope.row) }}
      </template>
    </el-table-column>
    <el-table-column label="容器状态" prop="containerStatus">
      <template #default="scope">
        {{ getContainerStatus(scope.row) }}
      </template>
    </el-table-column>
    <el-table-column label="宿主机名" prop="spec.nodeName" />
    <el-table-column label="宿主机IP" prop="status.hostIP" />
    <el-table-column label="操作" prop="operations">
      <template #default="scope">
        <el-button link type="danger" @click="emit('deleteItem', scope.row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>

  <DialogByYaml
    :dialogVisible="itemDetailDialog"
    @closeDialog="itemDetailDialog = false"
    :item-by-yaml="itemByYaml"
  >
    <template #header>
      <DialogHeaderLabel
        :cur-cluster-id="props.tableData.curClusterId"
        :cur-ns-name="curItem.metadata.namespace"
        :cur-resource-name="curItem.metadata.name"
        :cur-node-name="curItem.spec.nodeName"
        cur-resource-kind="Pod"
      />
    </template>
  </DialogByYaml>
</template>
