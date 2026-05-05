<script setup>
  import { ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { list2obj } from '../../../utils/typeConv/type.conv'
  import { useWorkLoadData } from '../../../store'
  import { storeToRefs } from 'pinia'
  import TabOfBasicConfig from './TabOfBasicConfig/TabOfBasicConfig.vue'
  import TabOfScheduleConfig from './tabOfScheduleConfig/TabOfScheduleConfig.vue'
  import DialogByYaml from '../DialogByYaml.vue'
  import { obj2yaml } from '../../../utils/typeConv/type.conv'
  import YamlEdit from '../YamlEdit.vue'
  import TabOfVolumeConfig from './tabOfVolumeConfig/tabOfVolumeConfig.vue'
  import TabOfContainer from './tabOfContainer/TabOfContainer.vue'
  import Deployment from '../../deployment/Deployment.vue'
  import { createHandler, updateHandler } from '../../../api/generic'

  const props = defineProps(['openDialog', 'actionMethod', 'resourceType'])
  const emit = defineEmits(['closeDialog', 'getList'])
  const activeName = ref('Basic')

  // 创建ref来获取子组件实例
  const basicRef = ref(null)
  const scheduleRef = ref(null)
  const volumeRef = ref(null)

  // store from pinia
  const store = useWorkLoadData()
  const { workLoadItem } = storeToRefs(store)

  const handleClose = () => {
    store.resetWorkLoadItem()
    activeName.value = 'Basic'
    itemOfYaml.value = ''
    emit('closeDialog')
  }

  // 创建调度管理资源
  const createItem = () => {
    // 同步子组件数据至模板
    syncToWorkLoadItem()
    // console.log(workLoadItem.value)
    const itemForm = getPostData(workLoadItem.value)
    if (props.resourceType == 'cronJob') {
      delete itemForm.item.spec.template
    }
    // 调用后端接口 创建 资源
    createHandler(props.resourceType, itemForm).then((res) => {
      if (res.data.status === 200) {
        ElMessage({
          message: res.data.message,
          type: 'success',
        })
        emit('getList')
        handleClose()
      }
    })
  }
  // 更新调度管理资源
  const updateItem = () => {
    // 生成模板数据
    syncToWorkLoadItem()
    const itemForm = JSON.parse(JSON.stringify(workLoadItem.value))
    // 移除空值字段
    removeEmptyFieldsDeep(itemForm)
    // 调用后端接口
    updateHandler(props.resourceType, itemForm).then((res) => {
      if (res.data.status === 200) {
        ElMessage({
          message: res.data.message,
          type: 'success',
        })
        emit('getList')
        handleClose()
      }
    })
  }
  // Yaml
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
  const getPostData = (obj) => {
    const postData = JSON.parse(JSON.stringify(obj))
    removeEmptyFieldsDeep(postData)
    return postData
  }
  const itemOfYaml = ref('')
  const getItemOfYaml = (tab) => {
    if (tab.paneName !== 'Yaml') return
    // 同步数据至模板
    syncToWorkLoadItem()
    // 转换模板数据为yaml
    if (props.resourceType == 'cronJob') {
      const sourceData = JSON.parse(JSON.stringify(workLoadItem.value.item))
      delete sourceData.spec.template
      itemOfYaml.value = obj2yaml(getPostData(sourceData))
    } else {
      itemOfYaml.value = obj2yaml(getPostData(workLoadItem.value.item))
    }
  }
  // 同步子组件数据至模板
  const syncToWorkLoadItem = () => {
    // 基本配置组件：  标签 注释
    if (basicRef.value.data.labelsAndAnnotationsSwtich == 'auto') {
      // 自动生成
      workLoadItem.value.item.metadata.labels = {}
      workLoadItem.value.item.metadata.labels.app = workLoadItem.value.item.metadata.name
      workLoadItem.value.item.spec.selector.matchLabels.app = workLoadItem.value.item.metadata.name
      workLoadItem.value.item.spec.template.metadata.labels.app =
        workLoadItem.value.item.metadata.name
    } else {
      // 手动配置标签、注释
      // 控制器
      workLoadItem.value.item.metadata.labels = list2obj(basicRef.value.data.controllerLabelsList)
      workLoadItem.value.item.metadata.annotations = list2obj(
        basicRef.value.data.controllerAnnotationsList
      )
      // 选择器
      workLoadItem.value.item.spec.selector.matchLabels = list2obj(
        basicRef.value.data.podLabelsList
      )
      // pod
      workLoadItem.value.item.spec.template.metadata.labels = list2obj(
        basicRef.value.data.podLabelsList
      )
    }
    // 调度组件数据
    workLoadItem.value.item.spec.template.spec.nodeSelector = list2obj(
      scheduleRef.value.data.nodeLabelsList
    )
    // 修正 emptyDir medium 数据
    workLoadItem.value.item.spec.template.spec.volumes.forEach((item) => {
      item?.emptyDir?.medium == 'Disk' && delete item.emptyDir.medium
    })
    // cronJob 数据修正
    if (props.resourceType == 'cronJob') {
      workLoadItem.value.item.spec.jobTemplate.spec.template = workLoadItem.value.item.spec.template
    }
  }
</script>

<template>
  <!-- 资源创建 dialog -->
  <el-dialog
    v-model="props.openDialog"
    width="1600px"
    style="height: 740px"
    @close="handleClose"
    destroy-on-close
  >
    <el-tabs v-model="activeName" @tab-click="getItemOfYaml">
      <el-tab-pane label="基本配置" name="Basic">
        <TabOfBasicConfig
          ref="basicRef"
          :actionMethod="props.actionMethod"
          :resource-type="props.resourceType"
        />
      </el-tab-pane>
      <el-tab-pane label="调度配置" name="Schedule">
        <TabOfScheduleConfig ref="scheduleRef" />
      </el-tab-pane>
      <el-tab-pane label="存储配置" name="Volume">
        <TabOfVolumeConfig />
      </el-tab-pane>
      <el-tab-pane label="容器配置" name="Container">
        <TabOfContainer container-type="container" />
      </el-tab-pane>
      <el-tab-pane label="初始化容器" name="InitContainer">
        <TabOfContainer container-type="init" />
      </el-tab-pane>
      <el-tab-pane label="Yaml" name="Yaml">
        <YamlEdit :code="itemOfYaml" style="height: 562px" />
      </el-tab-pane>
    </el-tabs>
    <el-button
      type="primary"
      size="large"
      style="margin-top: 20px; width: 90px"
      @click="createItem"
      v-if="props.actionMethod == 'create'"
    >
      创建
    </el-button>
    <el-button
      type="primary"
      size="large"
      style="margin-top: 20px; width: 90px"
      @click="updateItem"
      v-if="props.actionMethod == 'update'"
    >
      更新
    </el-button>
  </el-dialog>
</template>

<style scoped>
  .form-item {
    margin-bottom: 15px;
  }
  :deep(.no-border-input .el-input__wrapper) {
    box-shadow: none;
    border: none;
  }
</style>
