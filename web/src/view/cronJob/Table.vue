<script setup>
  import { computed, reactive, ref } from 'vue'
  import DialogByYaml from '../components/DialogByYaml.vue'
  import { obj2yaml } from '../../utils/typeConv/type.conv.js'
  import DialogHeaderLabel from '../components/DialogHeaderLabel.vue'
  import DialogOfItem from '../components/workLoads/DialogOfItem.vue'
  import { useWorkLoadData } from '../../store/index.js'
  import { storeToRefs } from 'pinia'
  import { ElMessage } from 'element-plus'
  import { getHandler } from '../../api/generic.js'

  // store from pinia
  const store = useWorkLoadData()
  const { workLoadItem } = storeToRefs(store)

  const emit = defineEmits(['deleteItem', 'getList'])

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

  const props = defineProps(['tableData'])

  // 搜索后的 pod 列表数据
  const filterTableData = computed(() =>
    (props.tableData.items || []).filter(
      (item) =>
        !props.tableData.search ||
        item.metadata.name.toLowerCase().includes(props.tableData.search.toLowerCase())
    )
  )

  const mergeIfExists = (target, source) => {
    Object.keys(source || {}).forEach((key) => {
      if (!(key in target)) {
        return
      }

      const sourceValue = source[key]
      const targetValue = target[key]

      if (sourceValue === undefined) {
        return
      }

      if (Array.isArray(sourceValue)) {
        if (!Array.isArray(targetValue)) {
          return
        }

        target[key] = sourceValue
          .map((item, index) => {
            const currentTargetItem = targetValue[index]

            if (currentTargetItem === undefined) {
              return currentTargetItem
            }

            if (
              item &&
              typeof item === 'object' &&
              !Array.isArray(item) &&
              currentTargetItem &&
              typeof currentTargetItem === 'object' &&
              !Array.isArray(currentTargetItem)
            ) {
              return mergeIfExists(currentTargetItem, item)
            }

            return currentTargetItem === undefined ? currentTargetItem : item
          })
          .filter((item) => item !== undefined)
        return
      }

      if (
        sourceValue &&
        typeof sourceValue === 'object' &&
        !Array.isArray(sourceValue) &&
        targetValue &&
        typeof targetValue === 'object' &&
        !Array.isArray(targetValue)
      ) {
        mergeIfExists(targetValue, sourceValue)
        return
      }

      target[key] = sourceValue
    })

    return target
  }

  const getItem = (row) => {
    // 重置模板
    store.resetWorkLoadItem()
    // 模板赋值
    getHandler(
      props.tableData.clusterId,
      row.metadata.namespace,
      'cronJob',
      row.metadata.name
    ).then((res) => {
      if (res.data.status == 200) {
        workLoadItem.value.name = row.metadata.name
        mergeIfExists(workLoadItem.value.item, res.data.data.items)
        data.actionMethod = 'update'
        data.updateItemDialogVisible = true
      }
    })
  }

  const data = reactive({
    updateItemDialogVisible: false,
    actionMethod: '',
    suspendScheduleLoadingValue: false,
  })
  const closeDialogOfItem = () => {
    data.updateItemDialogVisible = false
    emit('getList')
  }
  const suspendScheduleFunc = () => {
    // data.suspendScheduleLoadingValue = true
    return new Promise((resolve, reject) => {
      // 异步操作成功时
      resolve('ture')
      // 异步操作失败时
      // reject('false')
    })
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
    <el-table-column label="创建时间" prop="metadata.creationTimestamp" />
    <el-table-column label="调度周期" prop="spec.schedule" />
    <el-table-column label="最后调度时间" prop="status.lastScheduleTime" />
    <el-table-column label="最后调度成功时间" prop="status.lastSuccessfulTime" />
    <el-table-column label="命名空间" prop="metadata.namespace" />
    <el-table-column label="并发调度策略" prop="spec.concurrencyPolicy" />
    <el-table-column label="暂停调度" prop="spec.suspend">
      <template #default="scope">
        <el-switch
          v-model="scope.row.spec.suspend"
          :before-change="suspendScheduleFunc"
          :loading="data.suspendScheduleLoadingValue"
        />
      </template>
    </el-table-column>
    <el-table-column label="操作" prop="operations">
      <template #default="scope">
        <el-button link type="warning" @click="getItem(scope.row)">编辑</el-button>
        <el-button link type="danger" @click="emit('deleteItem', scope.row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  <DialogOfItem
    :open-dialog="data.updateItemDialogVisible"
    @close-dialog="closeDialogOfItem"
    @get-list="emit('getList')"
    :actionMethod="data.actionMethod"
    resource-type="cronJob"
  />
</template>
