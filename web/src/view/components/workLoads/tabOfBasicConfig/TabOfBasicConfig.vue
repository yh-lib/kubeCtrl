<script setup>
  import ElCard from '../../ElCard.vue'
  import { useWorkLoadData } from '../../../../store'
  import { storeToRefs } from 'pinia'
  import { onBeforeMount, reactive } from 'vue'
  import TableOfKeyValue from '../TableOfKeyValue.vue'
  import { getSecretListHandler } from '../../../../api/secret'
  import { ElMessage } from 'element-plus'
  import Deployment from '../../../deployment/Deployment.vue'
  import { getServiceListHandler } from '../../../../api/service'
  import Daemonset from '../../../daemonset/Daemonset.vue'
  import { obj2list } from '../../../../utils/typeConv/type.conv'

  // store from pinia
  const store = useWorkLoadData()
  const { workLoadItem } = storeToRefs(store)

  const props = defineProps(['actionMethod', 'resourceType'])
  // 当前页面所需数据
  const data = reactive({
    // table 数据
    controllerLabelsList: [],
    controllerAnnotationsList: [],
    // switch 状态
    switchAddService: false,
    labelsAndAnnotationsSwtich: '',
    // 选项
    dnsPolicyList: [
      { value: 'Default' },
      { value: 'ClusterFirst' },
      { value: 'ClusterFirstWithHostNet' },
    ],
    imagePullPolicyList: [{ value: 'Never' }, { value: 'IfNotPresent' }, { value: 'Always' }],
    privateRepoSecretList: [],
    svcList: [],
    bindSvcValue: '',
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
    if (props.resourceType == 'Deployment') {
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
    if (props.resourceType == 'StatefulSet') {
      const isRolling = workLoadItem.value.item.spec.updateStrategy.type === 'RollingUpdate'
      if (isRolling) {
        workLoadItem.value.item.spec.updateStrategy.rollingUpdate = {
          partition: 0,
          maxUnavailable: '25%',
        }
      } else {
        delete workLoadItem.value.item.spec.updateStrategy.rollingUpdate
      }
    }
    if (props.resourceType == 'DaemonSet') {
      const isRolling = workLoadItem.value.item.spec.updateStrategy.type === 'RollingUpdate'
      if (isRolling) {
        workLoadItem.value.item.spec.updateStrategy.rollingUpdate = {
          maxUnavailable: '25%',
        }
      } else {
        delete workLoadItem.value.item.spec.updateStrategy.rollingUpdate
      }
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
    console.log('ELCARD数据切换')
    // 集群 or NS 切换后 数据重置
    data.privateRepoSecretList = []
    data.svcList = []
    workLoadItem.value.item.spec.template.spec.imagePullSecrets[0].name = ''
    workLoadItem.value.item.spec.serviceName = ''

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

  const getSvcList = () => {
    data.svcList = []
    getServiceListHandler(workLoadItem.value.clusterId, workLoadItem.value.nameSpace).then(
      (res) => {
        if (res.data.status == 200) {
          res.data.data.items == null ||
            res.data.data.items.forEach((item) => {
              item.spec.clusterIP != 'None' ||
                data.svcList.push({
                  label: item.metadata.name,
                  value: item.metadata.name,
                })
            })
        }
      }
    )
  }

  const bindSvcValueChanged = () => {
    if (data.bindSvcValue == 'autoCreateSvc') {
      if (workLoadItem.value.item.metadata.name == '') {
        data.bindSvcValue = ''
        ElMessage.error('请先输入资源名称')
      }
      workLoadItem.value.item.spec.serviceName = workLoadItem.value.item.metadata.name
    } else {
      workLoadItem.value.item.spec.serviceName = ''
      getSvcList()
    }
  }

  onBeforeMount(() => {
    workLoadItem.value.item.kind = props.resourceType
    data.labelsAndAnnotationsSwtich = ''
    if (props.actionMethod == 'create') {
      data.labelsAndAnnotationsSwtich = 'auto'
    }
    if (props.actionMethod == 'update') {
      data.labelsAndAnnotationsSwtich = 'manual'
      data.controllerLabelsList = obj2list(workLoadItem.value.item.metadata.labels)
      data.controllerAnnotationsList = obj2list(workLoadItem.value.item.metadata.annotations)
    }
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
          <el-col
            :span="8"
            class="form-item"
            v-if="props.resourceType == 'Deployment' || props.resourceType == 'StatefulSet'"
          >
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
              <!-- 资源类型为deployment时 -->
              <el-select
                placeholder="请选择更新策略"
                v-model="workLoadItem.item.spec.strategy.type"
                @change="updatePoicySwitchFunc"
                v-if="props.resourceType == 'Deployment'"
              >
                <el-option
                  v-for="item in ['RollingUpdate', 'Recreate']"
                  :key="item"
                  :label="item"
                  :value="item"
                />
              </el-select>
              <!-- 资源类型为statefulset、daemonSet时 -->
              <el-select
                placeholder="请选择更新策略"
                v-model="workLoadItem.item.spec.updateStrategy.type"
                @change="updatePoicySwitchFunc"
                v-if="props.resourceType == 'StatefulSet' || props.resourceType == 'DaemonSet'"
              >
                <el-option
                  v-for="item in ['RollingUpdate', 'OnDelete']"
                  :key="item"
                  :label="item"
                  :value="item"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <!-- 更新策略 参数 -->
          <el-col
            :span="8"
            class="form-item"
            v-if="
              workLoadItem.item.spec.strategy.type == 'RollingUpdate' ||
              workLoadItem.item.spec.updateStrategy.type == 'RollingUpdate'
            "
          >
            <!-- deployment -->
            <div
              style="display: flex; justify-content: space-between"
              v-if="props.resourceType == 'Deployment'"
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
            <!-- statefulSet -->
            <div
              style="display: flex; justify-content: space-between"
              v-if="props.resourceType == 'StatefulSet'"
            >
              <el-form-item label="分区序号" label-width="100px">
                <el-input
                  placeholder=""
                  style="width: 100px"
                  v-model="workLoadItem.item.spec.updateStrategy.rollingUpdate.partition"
                />
              </el-form-item>
              <el-form-item label="最大不可用" label-width="100px">
                <el-input
                  placeholder=""
                  style="width: 100px"
                  v-model="workLoadItem.item.spec.updateStrategy.rollingUpdate.maxUnavailable"
                />
              </el-form-item>
            </div>
            <!-- DaemonSet -->
            <div
              style="display: flex; justify-content: space-between"
              v-if="props.resourceType == 'DaemonSet'"
            >
              <el-form-item label="最大不可用" label-width="150px">
                <el-input
                  placeholder="请输入最大不可用百分比"
                  style="width: 280px"
                  v-model="workLoadItem.item.spec.updateStrategy.rollingUpdate.maxUnavailable"
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
            <el-form-item label="自动添加 Servcie" v-if="props.resourceType == 'Deployment'">
              <el-switch v-model="data.switchAddService" />
            </el-form-item>
            <el-form-item label="绑定 Servcie" v-if="props.resourceType == 'StatefulSet'">
              <el-radio-group v-model="data.bindSvcValue" @change="bindSvcValueChanged">
                <el-radio value="autoCreateSvc">自动生成</el-radio>
                <el-radio value="manualSelectSvc">手动选择</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label-width="0px" v-if="data.bindSvcValue == 'manualSelectSvc'">
              <el-select
                v-model="workLoadItem.item.spec.serviceName"
                placeholder="请选择要绑定的无头服务"
                style="width: 500px"
                @visible-change="getSvcList"
              >
                <el-option
                  v-for="item in data.svcList"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
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
