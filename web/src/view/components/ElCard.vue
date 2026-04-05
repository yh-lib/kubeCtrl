<script setup>
import { ElSelect } from 'element-plus'
import { onBeforeMount, reactive } from 'vue';
import { getClusterListHandler} from '../../api/cluster.js'
import { getNamespaceListHandler } from '../../api/namespace.js';

const emit = defineEmits(['change','refresh'])

const data = reactive({
    // 集群信息
    curClusterId: '',
    clusterOptions: [],
    // namespace 信息
    curNsName: '',
    nsOptions: [],
    search: '',
})

const props = defineProps({
    title: String,
    opCluster: Boolean,
    opNs:Boolean,
    opSearch: Boolean,
    opRefresh: Boolean,
})

const syncToParent = () => {
    emit('change', data)
}

onBeforeMount(async () => {
    // 获取集群列表
    await getclusterOptions()
    // 获取select默认选择集群ID
    if (data.clusterOptions.length > 0) {
        data.curClusterId = data.clusterOptions[0].clusterId
    }
    // 获取 namespace 列表
    await getNsOptions()
    // 获取 namespace 默认选择 namespace
    if (data.nsOptions.length > 0) {
        data.curNsName = data.nsOptions[0].metadata.name
    }
    syncToParent()
})

const handleClusterChange = async () => {
    await getNsOptions()
    if (data.nsOptions.length > 0) {
        data.curNsName = data.nsOptions[0].metadata.name
    }
    syncToParent()
}

const handleNsChange = () => {
    syncToParent()
}

const handleSearchChange = () => {
    syncToParent()
}
// 获取 namespace 列表
const getNsOptions = async ()=>{
    await getNamespaceListHandler(data.curClusterId).then((res)=>{
        if (res.data.status === 200) {
            data.nsOptions = res.data.data.items || [];
        }
    })
}

// 获取集群列表
const getclusterOptions = async ()=>{
    await getClusterListHandler().then((res)=>{
        if (res.data.status === 200) {
            data.clusterOptions = res.data.data.items || [];
        }
    })     
}
</script>

<template>
  <!-- card -->
  <el-card>
    <!-- header -->
    <template #header>
      <div class="card-header">
        <!-- 标题 -->
        <div>
            <span style="font-size: 24px;">{{ props.title }}</span>
        </div>
        <div>
          <!-- 选择集群 -->
          <el-select v-model="data.curClusterId" placeholder="选择集群" style="width: 240px;" v-show="props.opCluster" filterable @change="handleClusterChange">
              <el-option
                  v-for="item in data.clusterOptions"
                  :key="item.clusterId"
                  :label="item.clusterAlias"
                  :value="item.clusterId"
              />
          </el-select>               
          <!-- 选择namespace -->
          <el-select v-model="data.curNsName" placeholder="选择命名空间" style="width: 240px;margin-left: 10px;" v-show="props.opNs" filterable @change="handleNsChange">
              <el-option
                  v-for="item in data.nsOptions"
                  :key="item.metadata.name"
                  :label="item.metadata.name"
                  :value="item.metadata.name"
              />
          </el-select>
          <!-- 搜索框 -->
          <el-input v-model="data.search" style="width: 240px;margin-left: 10px;" placeholder="搜索" v-show="props.opSearch"  @change="handleSearchChange"/>
          <!-- 创建按钮 -->
          <el-button type="primary" style="width: 105px; margin-left: 10px;" v-show="props.opRefresh" @click="emit('refresh')">刷新</el-button>
        </div>
      </div>
    </template>
    <!-- main -->
    <slot name="table">table 数据</slot>
  </el-card>
</template>

<style>
.card-header{
    display: flex;
    justify-content: space-between;
    align-items: center;
}
</style>