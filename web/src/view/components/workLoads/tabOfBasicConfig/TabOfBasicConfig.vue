<script setup>
  import ElCard from '../../ElCard.vue'
  import { useWorkLoadData } from '../../../../store'
  import { storeToRefs } from 'pinia'
  import { reactive } from 'vue'
  import TableOfKeyValue from '../TableOfKeyValue.vue'
  import { getSecretListHandler } from '../../../../api/secret'
  import { ElMessage } from 'element-plus'

  // store from pinia
  const store = useWorkLoadData()
  const { workLoadItem } = storeToRefs(store)

  const props = defineProps(['actionMethod'])
  // 当前页面所需数据
  const data = reactive({
    // table 数据
    controllerLabelsList: [],
    controllerAnnotationsList: [],
    // switch 状态
    switchAddService: false,
    labelsAndAnnotationsSwtich: 'auto',
    // 选项
    dnsPolicyList: [
      { value: 'Default' },
      { value: 'ClusterFirst' },
      { value: 'ClusterFirstWithHostNet' },
    ],
    updatePlicyList: [{ value: 'RollingUpdate' }, { value: 'Recreate' }],
    imagePullPolicyList: [{ value: 'Never' }, { value: 'IfNotPresent' }, { value: 'Always' }],
    privateRepoSecretList: [],
  })
  // 控制器标签
  const addControllerLabelItem = () => {
    data.controllerLabelsList.unshift({ key: '', value: '' })
  }
  const deleteControllerLabelItem = (index) => {
    data.controllerLabelsList.splice(index, 1)
  }
  // 控制器注释
  const addControllerAnnotationItem = () => {
    data.controllerAnnotationsList.unshift({ key: '', value: '' })
  }
  const deleteControllerAnnotationItem = (index) => {
    data.controllerAnnotationsList.splice(index, 1)
  }
  // 切换更新策略
  const updatePoicySwitchFunc = () => {
    const isRolling = workLoadItem.value.item.spec.strategy.type === 'RollingUpdate'
    if (isRolling) {
      workLoadItem.value.item.spec.strategy.rollingUpdate = {
        maxUnavailable: '25%',
        maxSurge: '25%',
      }
    } else {
      delete workLoadItem.value.item.spec.strategy.rollingUpdate
    }
  }
  // 切换主机网络
  const hostNetworkSwitch = () => {
    if (workLoadItem.value.item.spec.template.spec.hostNetwork == true) {
      ElMessage({
        message: '提示: 已开启宿主机网络',
        type: 'warning',
      })
    }
  }
  // 从子组件 ELCard 中获取所需数据
  const getSelectValue = (selectValue) => {
    data.privateRepoSecretList = []
    delete workLoadItem.value.item.spec.template.spec.imagePullSecrets[0].name
    Object.assign(workLoadItem.value, selectValue)
    getSecretListHandler(workLoadItem.value.clusterId, workLoadItem.value.nameSpace).then((res) => {
      if (res.data.status == 200) {
        res.data.data.items == null ||
          res.data.data.items.forEach((item) => {
            item.data['.dockerconfigjson'] == null ||
              data.privateRepoSecretList.push({
                label: item.metadata.name,
                value: item.metadata.name,
              })
          })
      }
    })
  }

  const syncToWorkLoadItemName = () => {
    workLoadItem.value.name = workLoadItem.value.item.metadata.name
  }

  // 暴露 data 给父组件
  defineExpose({
    data,
  })
</script>

<template>
  <ElCard
    :op-cluster="true"
    :op-ns="true"
    style="border-radius: 0px; width: 1560px; height: 550px"
    @change="getSelectValue"
    :actionMethod="props.actionMethod"
  >
    <template #mainData>
      <!-- 基础信息 -->
      <el-form label-width="150px" label-position="left" style="height: 450px; width: 1490px">
        <el-row :gutter="100">
          <!-- 名称 -->
          <el-col :span="8" class="form-item">
            <el-form-item label="名称" required>
              <el-input
                placeholder="请输入资源名称"
                v-model="workLoadItem.item.metadata.name"
                @change="syncToWorkLoadItemName"
                :disabled="props.actionMethod == 'update'"
              />
            </el-form-item>
          </el-col>
          <!-- 副本数 -->
          <el-col :span="8" class="form-item">
            <el-form-item label="副本数">
              <el-input
                placeholder="请输入副本数"
                v-model.number="workLoadItem.item.spec.replicas"
              />
            </el-form-item>
          </el-col>
          <!-- 镜像地址 -->
          <el-col :span="8" class="form-item">
            <el-form-item label="镜像地址" required>
              <el-input
                placeholder="请输入镜像地址"
                v-model="workLoadItem.item.spec.template.spec.containers[0].image"
              />
            </el-form-item>
          </el-col>
          <!-- 镜像拉取策略 -->
          <el-col :span="8" class="form-item">
            <el-form-item label="镜像拉取策略">
              <el-select
                placeholder="请选择镜像拉取策略"
                v-model="workLoadItem.item.spec.template.spec.containers[0].imagePullPolicy"
              >
                <el-option
                  v-for="item in data.imagePullPolicyList"
                  :key="item.value"
                  :label="item.value"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <!-- 私有仓库密钥 -->
          <el-col :span="8" class="form-item">
            <el-form-item label="私有仓库密钥">
              <el-select
                placeholder="请选择私有仓库密钥"
                v-model="workLoadItem.item.spec.template.spec.imagePullSecrets[0].name"
              >
                <el-option
                  v-for="item in data.privateRepoSecretList"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <!-- DNS 策略 -->
          <el-col :span="8" class="form-item">
            <el-form-item label="DNS 策略">
              <el-select
                placeholder="请选择 DNS 策略"
                v-model="workLoadItem.item.spec.template.spec.dnsPolicy"
              >
                <el-option
                  v-for="item in data.dnsPolicyList"
                  :key="item.value"
                  :label="item.value"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <!-- 更新策略 -->
          <el-col :span="8" class="form-item">
            <el-form-item label="更新策略">
              <el-select
                placeholder="请选择更新策略"
                v-model="workLoadItem.item.spec.strategy.type"
                @change="updatePoicySwitchFunc"
              >
                <el-option
                  v-for="item in data.updatePlicyList"
                  :key="item.value"
                  :label="item.value"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <!-- 更新策略 参数 -->
          <el-col :span="8" class="form-item">
            <div
              style="display: flex; justify-content: space-between"
              v-if="workLoadItem.item.spec.strategy.type == 'RollingUpdate'"
            >
              <el-form-item label="最大不可用" label-width="100px">
                <el-input
                  placeholder=""
                  style="width: 100px"
                  v-model="workLoadItem.item.spec.strategy.rollingUpdate.maxUnavailable"
                />
              </el-form-item>
              <el-form-item label="可超出最大值" label-width="100px">
                <el-input
                  placeholder=""
                  style="width: 100px"
                  v-model="workLoadItem.item.spec.strategy.rollingUpdate.maxSurge"
                />
              </el-form-item>
            </div>
          </el-col>
          <!-- 使用宿主机网络 -->
          <el-col :span="8" class="form-item">
            <el-form-item label="使用宿主机网络">
              <el-switch
                v-model="workLoadItem.item.spec.template.spec.hostNetwork"
                style="--el-switch-on-color: #13ce66"
                @change="hostNetworkSwitch"
              />
            </el-form-item>
          </el-col>
          <!-- 标签及注释 -->
          <el-col :span="8" class="form-item">
            <el-form-item label="标签及注释">
              <el-radio-group v-model="data.labelsAndAnnotationsSwtich">
                <el-radio value="auto" size="large">自动生成</el-radio>
                <el-radio value="manual" size="large">手动配置</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <!-- 自动添加 Servcie -->
          <el-col :span="8" class="form-item">
            <el-form-item label="自动添加 Servcie">
              <el-switch v-model="data.switchAddService" />
            </el-form-item>
          </el-col>
        </el-row>
        <!-- 标签|注释|容忍 tabs -->
        <el-tabs
          tab-position="left"
          style="height: 250px"
          type="border-card"
          class="no-border-input"
          v-if="data.labelsAndAnnotationsSwtich == 'manual'"
        >
          <!-- 标签 -->
          <el-tab-pane label="标签">
            <TableOfKeyValue
              :table-list="data.controllerLabelsList"
              @add-table-row="addControllerLabelItem"
              @delete-table-row="deleteControllerLabelItem"
            />
          </el-tab-pane>
          <!-- 注释 -->
          <el-tab-pane label="注释">
            <TableOfKeyValue
              :table-list="data.controllerAnnotationsList"
              @add-table-row="addControllerAnnotationItem"
              @delete-table-row="deleteControllerAnnotationItem"
            />
          </el-tab-pane>
        </el-tabs>
      </el-form>
    </template>
  </ElCard>
</template>
