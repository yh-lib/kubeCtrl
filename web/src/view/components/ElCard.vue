<script setup>
  import { ElSelect } from 'element-plus'
  import { Refresh } from '@element-plus/icons-vue'
  import { onBeforeMount, reactive } from 'vue'
  // store from pinia
  import { useWorkLoadData } from '../../store'
  import { storeToRefs } from 'pinia'
  import { getListHandler } from '../../api/generic.js'

  // store from pinia
  const store = useWorkLoadData()
  const { workLoadItem } = storeToRefs(store)

  const emit = defineEmits(['change', 'refresh', 'createItem'])

  const data = reactive({
    // 集群信息
    clusterId: '',
    clusterOptions: [],
    // namespace 信息
    nameSpace: '',
    nsOptions: [],
    search: '',
  })

  // 选择组件元素
  const props = defineProps({
    title: String,
    opCluster: Boolean,
    opNs: Boolean,
    opSearch: Boolean,
    opRefresh: Boolean,
    opCreate: Boolean,
    actionMethod: String,
  })

  const syncToParent = () => {
    emit('change', data)
  }

  onBeforeMount(async () => {
    // 获取集群列表
    await getclusterOptions()
    // 获取select默认选择集群ID
    if (data.clusterOptions.length > 0) {
      data.clusterId = data.clusterOptions[0].clusterId
    }
    // 获取 namespace 列表
    await getNsOptions()
    // 获取 namespace 默认选择 namespace
    if (data.nsOptions.length > 0) {
      data.nameSpace = ''
    }
    syncToParent()
  })

  const handleClusterChange = () => {
    syncToParent()
  }

  const handleNsChange = () => {
    syncToParent()
  }

  const handleSearchChange = () => {
    syncToParent()
  }
  // 获取 namespace 列表
  const getNsOptions = async () => {
    data.nameSpace = ''
    await getListHandler(data.clusterId, '', 'namespace').then((res) => {
      if (res.data.status === 200) {
        data.nsOptions = res.data.data.items || []
      }
    })
  }

  // 获取集群列表
  const getclusterOptions = async () => {
    await getListHandler('', '', 'cluster').then((res) => {
      if (res.data.status === 200) {
        data.clusterOptions = res.data.data.items || []
      }
    })
  }
</script>

<template>
  <el-card class="ks-card">
    <template #header>
      <div class="card-header">
        <div class="card-title-wrap">
          <span class="card-title">{{ props.title }}</span>
        </div>
        <div class="card-controls">
          <el-select
            v-model="data.clusterId"
            placeholder="选择集群"
            class="control-select"
            v-show="props.opCluster"
            filterable
            @change="handleClusterChange"
            :disabled="props.actionMethod == 'update'"
          >
            <el-option
              v-for="item in data.clusterOptions"
              :key="item.clusterId"
              :label="item.clusterAlias"
              :value="item.clusterId"
            />
          </el-select>
          <el-select
            v-model="data.nameSpace"
            placeholder="All namespaces"
            class="control-select"
            v-show="props.opNs"
            filterable
            @change="handleNsChange"
            :disabled="props.actionMethod == 'update'"
          >
            <el-option key="All namespaces" label="All namespaces" value="" />
            <el-option
              v-for="item in data.nsOptions"
              :key="item.metadata.name"
              :label="item.metadata.name"
              :value="item.metadata.name"
            />
          </el-select>
          <el-input
            v-model="data.search"
            class="control-search"
            placeholder="搜索"
            v-show="props.opSearch"
            @change="handleSearchChange"
          />
          <el-button
            style="border-radius: 12px"
            :icon="Refresh"
            v-show="props.opRefresh"
            @click="emit('refresh')"
            >刷新</el-button
          >
        </div>
      </div>
    </template>
    <slot name="mainData">无数据</slot>
    <!-- 创建或删除按钮 -->
    <el-button
      type="primary"
      style="
        margin-top: -100px;
        z-index: 1;
        position: absolute;
        right: 60px;
        font-size: 26px;
        width: 60px;
        height: 60px;
      "
      circle
      @click="emit('createItem')"
      v-show="props.opCreate"
    >
      +
    </el-button>
  </el-card>
</template>

<style scoped>
  .ks-card {
    width: 100%;
    height: 1150px;
    display: block;
    border: 1px solid #e6eaf0;
    border-radius: 24px;
    box-shadow: 0 12px 28px rgba(15, 23, 42, 0.04);
    margin-top: 15px;
  }

  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 16px;
  }

  .card-title {
    font-size: 24px;
    font-weight: 700;
    color: #111827;
  }

  .card-controls {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    gap: 10px;
    flex-wrap: wrap;
  }

  .control-select,
  .control-search {
    width: 240px;
  }
</style>
