<script lang="ts" setup>
  import { computed, reactive, onBeforeMount } from 'vue'
  import { ElSelect, ElMessage } from 'element-plus'
  import { list2obj, obj2list } from '../../utils/typeConv/type.conv.js'
  import request from '../../api/axiosEncap.js'
  import { API_CONFIG } from '../../config/index.js'
  import ElCard from '../components/ElCard.vue'
  import { getHandler, getListHandler } from '../../api/generic.js'

  // 需要的数据变量
  const data = reactive({
    items: [], // 后端返回的 list
    item: {}, // 后端返回的 obj
    search: '', // 接收 header 搜索框中数据
    curHostName: '',
    // 集群选择器
    clusterOptions: [], // 集群选择 options
    curClusterId: '', // 当前选择的集群id
    defaultClusterId: 'in-cluster', // 默认选择的集群id
    // 表格
    opDialog: false, // 配置、主机名 对话框
    nodeLabels: [], // 后端返回的label
    nodeTaints: [],
    curTabName: '',
    taintOptions: [
      // 污点选项
      { label: '禁止调度', value: 'NoSchedule' },
      { label: '驱逐', value: 'NoExecute' },
      { label: '尽量不调度', value: 'PreferNoSchedule' },
    ],
    // axios
    requestItem: {
      clusterId: '',
      nameSpace: '',
      name: '',
      item: {},
    },
    // 列表页面
    swSchedule: false,
    swExeCute: false,
  })

  // // // 子组件加载前自动获取数据
  onBeforeMount(async () => {
    await getclusterOptions() // 获取集群列表
    data.curClusterId = data.defaultClusterId // 获取基础设施集群的Node列表
    getList()
  })

  // 更新按钮实现逻辑
  const updateItem = () => {
    if (data.curTabName == 'nodeLabel') {
      // 构建后端所需数据
      data.requestItem.clusterId = data.curClusterId
      data.requestItem.nameSpace = ''
      data.requestItem.name = data.curHostName
      data.requestItem.item = data.item
      data.requestItem.item.metadata.labels = list2obj(data.nodeLabels)
      // 调用后端API
      request(API_CONFIG.nodeUpdateApi, data.requestItem, 'post').then((res) => {
        if (res.data.status === 200) {
          ElMessage({
            message: res.data.message,
            type: 'success',
          })
          data.opDialog = false
        } else {
          ElMessage.error({ message: res.data.message })
        }
      })
    } else if (data.curTabName == 'nodeTaint') {
      // 构建后端所需数据
      data.requestItem.clusterId = data.curClusterId
      data.requestItem.nameSpace = ''
      data.requestItem.name = data.curHostName
      data.requestItem.item = data.item
      data.requestItem.item.spec.taints = data.nodeTaints
      // 调用后端API
      request(API_CONFIG.nodeUpdateApi, data.requestItem, 'post').then((res) => {
        if (res.data.status === 200) {
          ElMessage({
            message: res.data.message,
            type: 'success',
          })
          data.opDialog = false
        } else {
          ElMessage.error({ message: res.data.message })
        }
      })
    }
  }

  // 获取节点角色（纯 JS 版本）
  const getNodeRole = (labels = {}) => {
    const keyList = Object.keys(labels || {})
    if (!keyList.length) {
      return 'Worker'
    }
    const controlPlaneKeys = [
      'node-role.kubernetes.io/control-plane',
      'node-role.kubernetes.io/controlplane',
      'node-role.kubernetes.io/master',
    ]
    return controlPlaneKeys.some((key) => keyList.includes(key)) ? 'Master' : 'Worker'
  }

  // 添加列表项
  const addLabelItem = () => {
    data.nodeLabels.unshift({ key: '', value: '' })
  }
  // 添加污点项（为空时先初始化数组）
  const addTaintItem = () => {
    if (!Array.isArray(data.nodeTaints)) {
      data.nodeTaints = []
    }
    data.nodeTaints.unshift({ key: '', value: '', effect: '' })
  }

  // 删除列表项
  const deleteLabelItem = (index) => {
    data.nodeLabels.splice(index, 1)
  }
  // 删除污点项
  const deleteTaintItem = (index) => {
    data.nodeTaints.splice(index, 1)
  }

  // 获取节点标签
  const getLabel = () => {
    // console.log("获取节点调试数据:",data.item.spec.taints)
    // 获取标签
    data.nodeLabels = obj2list(data.item.metadata.labels)
  }

  // 节点配置
  const getItem = (row) => {
    data.curTabName = 'nodeDetail'
    data.curHostName = row.metadata.name
    data.nodeTaints = Array.isArray(row.spec?.taints) ? [...row.spec.taints] : []
    // 获取节点详情
    getHandler(data.curClusterId, '', 'node', data.curHostName).then((res) => {
      data.item = res.data.data.items
      // console.log("获取节点污点：：：:",data.nodeTaints)
      data.opDialog = true
    })
  }

  // 关闭dialog时刷新用户列表
  const closeDialog = () => {
    getList()
  }

  // 内存地址单位转换
  const memoryKi2MB = (memory) => {
    if (memory === null || memory === undefined) return ''
    const match = String(memory)
      .trim()
      .match(/^(\d+(?:\.\d+)?)Ki$/i)
    return match ? Math.floor(Number(match[1]) / 1000) : ''
  }

  // 搜索后的 node 列表数据
  const filterTableData = computed(() =>
    data.items.filter(
      (item) =>
        !data.search ||
        item.metadata.name.toLowerCase().includes(data.search.toLowerCase()) ||
        item.status.addresses[0].address.toLowerCase().includes(data.search.toLowerCase())
    )
  )

  // 获取当前集群Node列表
  const getList = () => {
    if (!data.clusterId) {
      data.items = []
      return
    }
    getListHandler(data.clusterId, '', 'node').then((res) => {
      data.items = res.data.data.items || []
    })
  }

  // 获取集群列表
  const getclusterOptions = async () => {
    await getListHandler('', '', 'cluster').then((res) => {
      if (res.data.status === 200) {
        data.clusterOptions = res.data.data.items
        console.log('获取集群列表:::', data.clusterOptions)
      } else {
        console.error('获取集群列表失败:', res.data.message)
      }
    })
  }

  const getSelectValue = (selectValue) => {
    Object.assign(data, selectValue)
    getList()
  }
</script>

<template>
  <ElCard
    title="节点列表"
    :op-cluster="true"
    :op-search="true"
    :op-refresh="true"
    @change="getSelectValue"
    @refresh="getList"
  >
    <template #mainData>
      <el-table :data="filterTableData" style="width: 100%" height="1010px">
        <el-table-column label="主机名" prop="hostName">
          <template #default="scope">
            <el-button link type="primary" @click="getItem(scope.row)">
              {{ scope.row.status.addresses[1].address }}
            </el-button>
          </template>
        </el-table-column>
        <el-table-column label="IP地址" prop="ipAddress">
          <template #default="scope">
            {{ scope.row.status.addresses[0].address }}
          </template>
        </el-table-column>
        <el-table-column label="规格" prop="capacity">
          <template #default="scope">
            {{ scope.row.status.capacity.cpu }}C |
            {{ memoryKi2MB(scope.row.status.capacity.memory) }}MB
          </template>
        </el-table-column>
        <el-table-column label="角色" prop="role">
          <template #default="scope">
            {{ getNodeRole(scope.row.metadata?.labels || {}) }}
          </template>
        </el-table-column>
        <el-table-column label="状态" prop="status">
          <template #default="scope">
            <span
              v-if="
                scope.row.status.conditions[scope.row.status.conditions.length - 1].status == 'True'
              "
              style="color: green"
              >Ready</span
            >
            <span v-else style="color: red">NotReady</span>
          </template>
        </el-table-column>
        <el-table-column label="禁止调度" prop="schedule">
          <template #default="scope">
            <el-switch v-model="data.swSchedule" />
          </template>
        </el-table-column>
        <el-table-column label="驱逐" prop="eviction">
          <template #default="scope">
            <el-switch v-model="data.swExeCute" />
          </template>
        </el-table-column>
        <el-table-column label="操作" prop="edit">
          <template #default="scope">
            <el-button link type="primary" @click="getItem(scope.row)">配置</el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- dialog -->
      <el-dialog
        v-model="data.opDialog"
        width="1600px"
        style="height: 620px"
        @close="closeDialog"
        destroy-on-close
        :close-on-click-modal="false"
      >
        <!-- dialog header -->
        <template #header>
          <div style="display: flex; justify-content: flex-start">
            <el-tag type="primary" effect="plain" style="margin-right: 10px"
              >集群: {{ data.curClusterId || '-' }}</el-tag
            >
            <el-tag type="primary" effect="plain" style="margin-right: 10px"
              >节点: {{ data.item?.metadata?.name || '-' }}</el-tag
            >
            <span style="font-size: 18px; margin-left: 500px">节点配置</span>
          </div>
        </template>
        <!-- dialog middle -->
        <el-tabs v-model="data.curTabName" class="demo-tabs" @tab-click="getLabel">
          <!-- 标签   详情 -->
          <el-tab-pane label="详情" name="nodeDetail">
            <el-descriptions :column="2" size="large" border style="height: 440px">
              <el-descriptions-item label="主机名">{{ data.curHostName }}</el-descriptions-item>
              <el-descriptions-item label="IP地址">{{
                data.item.status.addresses[0].address
              }}</el-descriptions-item>
              <el-descriptions-item label="角色">{{
                getNodeRole(data.item?.metadata?.labels || {})
              }}</el-descriptions-item>

              <el-descriptions-item label="节点状态">
                <el-tag
                  :type="
                    data.item.status.conditions[data.item.status.conditions.length - 1].status ==
                    'True'
                      ? 'success'
                      : 'danger'
                  "
                >
                  {{ data.item.status.conditions[data.item.status.conditions.length - 1].reason }}
                </el-tag>
              </el-descriptions-item>
              <el-descriptions-item label="Pods"
                ><el-button link type="primary">{{
                  data.item.status.capacity.pods
                }}</el-button></el-descriptions-item
              >
              <el-descriptions-item label="操作系统">{{
                data.item.status.nodeInfo.osImage
              }}</el-descriptions-item>

              <el-descriptions-item label="Runtime">{{
                data.item.status.nodeInfo.containerRuntimeVersion
              }}</el-descriptions-item>
              <el-descriptions-item label="CPU架构">{{
                data.item.status.nodeInfo.architecture
              }}</el-descriptions-item>
              <el-descriptions-item label="组件版本">{{
                data.item.status.nodeInfo.kubeletVersion
              }}</el-descriptions-item>

              <el-descriptions-item label="系统内核">{{
                data.item.status.nodeInfo.kernelVersion
              }}</el-descriptions-item>
            </el-descriptions>
          </el-tab-pane>
          <!-- 标签   标签 -->
          <el-tab-pane label="标签" name="nodeLabel" class="no-border-input">
            <el-table :data="data.nodeLabels" style="width: 100%; height: 440px">
              <el-table-column prop="key" label="Key">
                <template #default="scope">
                  <el-input v-model="scope.row.key" placeholder="请输入标签的Key"></el-input>
                </template>
              </el-table-column>
              <el-table-column prop="value" label="Value">
                <template #default="scope">
                  <el-input v-model="scope.row.value" placeholder="请输入标签的Value"></el-input>
                </template>
              </el-table-column>
              <el-table-column>
                <template #header>
                  <el-button
                    type="primary"
                    link
                    style="font-size: 16px; margin-bottom: 10px"
                    @click="addLabelItem"
                    >添加</el-button
                  >
                </template>
                <template #default="scope">
                  <el-button
                    type="danger"
                    link
                    style="font-size: 16px; margin-bottom: 10px"
                    @click="deleteLabelItem(scope.$index)"
                    >删除</el-button
                  >
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>
          <!-- 标签   污点 -->
          <el-tab-pane label="污点" name="nodeTaint" class="no-border-input">
            <el-table :data="data.nodeTaints" style="width: 100%; height: 440px">
              <el-table-column prop="key" label="Key">
                <template #default="scope">
                  <el-input v-model="scope.row.key" placeholder="请输入污点的Key"></el-input>
                </template>
              </el-table-column>
              <el-table-column prop="value" label="Value">
                <template #default="scope">
                  <el-input v-model="scope.row.value" placeholder="请输入污点的Value"></el-input>
                </template>
              </el-table-column>
              <el-table-column prop="effect" label="Effect">
                <template #default="scope">
                  <!-- <el-input v-model="scope.row.effect" placeholder="请输入污点的Effect"></el-input> -->
                  <el-select v-model="scope.row.effect" placeholder="Select" style="width: 240px">
                    <el-option
                      v-for="item in data.taintOptions"
                      :key="item.value"
                      :label="item.label"
                      :value="item.value"
                    />
                  </el-select>
                </template>
              </el-table-column>
              <el-table-column>
                <template #header>
                  <el-button
                    type="primary"
                    link
                    style="font-size: 16px; margin-bottom: 10px"
                    @click="addTaintItem"
                    >添加</el-button
                  >
                </template>
                <template #default="scope">
                  <el-button
                    type="danger"
                    link
                    style="font-size: 16px; margin-bottom: 10px"
                    @click="deleteTaintItem(scope.$index)"
                    >删除</el-button
                  >
                </template>
              </el-table-column>
            </el-table>
          </el-tab-pane>
        </el-tabs>
        <template #footer>
          <div style="display: flex; justify-content: center">
            <el-button v-if="data.curTabName !== 'nodeDetail'" type="primary" @click="updateItem"
              >更新</el-button
            >
          </div>
        </template>
      </el-dialog>
    </template>
  </ElCard>
</template>

<style lang="less" scoped>
  :deep(.no-border-input .el-input__wrapper) {
    box-shadow: none;
    border: none;
  }
</style>
