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
  import CronJob from '../../../cronJob/cronJob.vue'

  // store from pinia
  const store = useWorkLoadData()
  const { workLoadItem } = storeToRefs(store)

  const props = defineProps(['actionMethod', 'resourceType'])
  // 当前页面所需数据
  const data = reactive({
    // table 数据
    controllerLabelsList: [],
    controllerAnnotationsList: [],
    podLabelsList: [],
    podAnnotationsList: [],
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
  // Pod 标签
  const addPodLabelItem = () => {
    data.podLabelsList.unshift({ key: '', value: '' })
  }
  const deletePodLabelItem = (index) => {
    data.podLabelsList.splice(index, 1)
  }
  // Pod 注释
  const addPodAnnotationItem = () => {
    data.controllerAnnotationsList.unshift({ key: '', value: '' })
  }
  const deletePodAnnotationItem = (index) => {
    data.controllerAnnotationsList.splice(index, 1)
  }
  // 切换更新策略
  const updatePoicySwitchFunc = () => {
    if (props.resourceType == 'deployment') {
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
    if (props.resourceType == 'statefulSet') {
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
    if (props.resourceType == 'daemonSet') {
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
            item.data == undefined ||
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
    // kind
    workLoadItem.value.item.kind = props.resourceType
    // apiversion
    switch (props.resourceType) {
      case 'deployment':
        workLoadItem.value.item.apiVersion = 'apps/v1'
        break
      case 'statefulSet':
        workLoadItem.value.item.apiVersion = 'apps/v1'
        break
      case 'daemonSet':
        workLoadItem.value.item.apiVersion = 'apps/v1'
        break
      case 'cronJob':
        workLoadItem.value.item.apiVersion = 'batch/v1'
        if (props.actionMethod == 'update') {
          workLoadItem.value.item.spec.template =
            workLoadItem.value.item.spec.jobTemplate.spec.template
        }
        break
      default:
        break
    }
    data.labelsAndAnnotationsSwtich = ''
    if (props.actionMethod == 'create') {
      data.labelsAndAnnotationsSwtich = 'auto'
    }
    if (props.actionMethod == 'update') {
      data.labelsAndAnnotationsSwtich = 'manual'
      // 控制器标签、注释
      data.controllerLabelsList = obj2list(workLoadItem.value.item.metadata.labels)
      data.controllerAnnotationsList = obj2list(workLoadItem.value.item.metadata.annotations)
      // Pod 标签、注释
      data.podLabelsList = obj2list(workLoadItem.value.item.spec.template.metadata.labels)
      data.podAnnotationsList = obj2list(workLoadItem.value.item.spec.template.metadata.annotations)
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
          <el-col :span="8">
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
            v-if="props.resourceType == 'deployment' || props.resourceType == 'statefulSet'"
          >
            <el-form-item label="副本数">
              <el-input
                placeholder="请输入副本数"
                v-model.number="workLoadItem.item.spec.replicas"
              />
            </el-form-item>
          </el-col>
          <!-- 镜像地址 -->
          <el-col :span="8">
            <el-form-item label="镜像地址" required>
              <el-input
                placeholder="请输入镜像地址"
                v-model="workLoadItem.item.spec.template.spec.containers[0].image"
              />
            </el-form-item>
          </el-col>
          <!-- 镜像拉取策略 -->
          <el-col :span="8">
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
          <!-- 重启策略 -->
          <el-col :span="8" v-if="props.resourceType == 'cronJob'">
            <el-form-item label="重启策略" required>
              <el-select
                placeholder="请选择重启策略"
                v-model="workLoadItem.item.spec.template.spec.restartPolicy"
              >
                <el-option
                  v-for="item in [
                    { label: '无论容器退出状态如何，总是重启', value: 'Always' },
                    { label: '只有容器退出状态非 0 时才重启', value: 'OnFailure' },
                    { label: '从不重启', value: 'Never' },
                  ]"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <!-- 私有仓库密钥 -->
          <el-col :span="8">
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
          <el-col :span="8">
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
          <!-- 调度周期 -->
          <el-col :span="8" v-if="props.resourceType == 'cronJob'">
            <el-form-item label="调度周期" required>
              <el-input placeholder="请配置调度周期" v-model="workLoadItem.item.spec.schedule" />
            </el-form-item>
          </el-col>
          <!-- 并发策略 -->
          <el-col :span="8" v-if="props.resourceType == 'cronJob'">
            <el-form-item label="并发策略">
              <el-select
                placeholder="请选择并发策略"
                v-model="workLoadItem.item.spec.concurrencyPolicy"
                @change="
                  workLoadItem.item.spec.concurrencyPolicy != 'Allow' &&
                    (workLoadItem.item.spec.parallelism = null)
                "
              >
                <el-option
                  v-for="item in [
                    { label: '允许同时运行多个任务实例', value: 'Allow' },
                    { label: '禁止同时运行多个任务实例', value: 'Forbid' },
                    { label: '用新任务替换正在运行的任务', value: 'Replace' },
                  ]"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>
          </el-col>
          <!-- Job -->
          <el-col :span="8" v-if="workLoadItem.item.spec.concurrencyPolicy == 'Allow'">
            <el-form-item label="Job启动pod并行数量">
              <el-input-number
                v-model="workLoadItem.item.spec.jobTemplate.spec.parallelism"
                :min="1"
              />
            </el-form-item>
          </el-col>
          <!-- 并发数量 -->
          <el-col :span="8" v-if="workLoadItem.item.spec.concurrencyPolicy == 'Allow'">
            <el-form-item label="Job成功阈值">
              <el-input-number
                v-model="workLoadItem.item.spec.jobTemplate.spec.completions"
                :min="1"
              />
            </el-form-item>
          </el-col>
          <!-- 保留成功的 Job 数量 -->
          <el-col :span="8" v-if="props.resourceType == 'cronJob'">
            <el-form-item label="保留成功的 Job 数量">
              <el-input-number
                :min="1"
                v-model="workLoadItem.item.spec.successfulJobsHistoryLimit"
              />
            </el-form-item>
          </el-col>
          <!-- 保留失败的 Job 数量 -->
          <el-col :span="8" v-if="props.resourceType == 'cronJob'">
            <el-form-item label="保留失败的 Job 数量">
              <el-input-number :min="1" v-model="workLoadItem.item.spec.failedJobsHistoryLimit" />
            </el-form-item>
          </el-col>
          <!-- 暂停调度 -->
          <el-col :span="8" v-if="props.resourceType == 'cronJob'">
            <el-form-item label="暂停调度">
              <el-switch v-model="workLoadItem.item.spec.suspend" />
            </el-form-item>
          </el-col>
          <!-- 更新策略 -->
          <el-col :span="8" v-if="props.resourceType != 'CronJob'">
            <el-form-item label="更新策略">
              <!-- 资源类型为deployment时 -->
              <el-select
                placeholder="请选择更新策略"
                v-model="workLoadItem.item.spec.strategy.type"
                @change="updatePoicySwitchFunc"
                v-if="props.resourceType == 'deployment'"
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
                v-if="props.resourceType == 'statefulSet' || props.resourceType == 'daemonSet'"
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
            v-if="
              workLoadItem.item.spec.strategy.type == 'RollingUpdate' ||
              workLoadItem.item.spec.updateStrategy.type == 'RollingUpdate'
            "
          >
            <!-- deployment -->
            <div
              style="display: flex; justify-content: space-between"
              v-if="props.resourceType == 'deployment'"
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
              v-if="props.resourceType == 'statefulSet'"
            >
              <el-form-item label="分区序号" label-width="100px">
                <el-input
                  placeholder=""
                  style="width: 100px"
                  v-model.number="workLoadItem.item.spec.updateStrategy.rollingUpdate.partition"
                />
              </el-form-item>
              <el-form-item label="最大不可用" label-width="100px">
                <el-input
                  placeholder=""
                  style="width: 100px"
                  v-model.number="
                    workLoadItem.item.spec.updateStrategy.rollingUpdate.maxUnavailable
                  "
                />
              </el-form-item>
            </div>
            <!-- DaemonSet -->
            <div
              style="display: flex; justify-content: space-between"
              v-if="props.resourceType == 'daemonSet'"
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
          <el-col :span="8" v-if="props.resourceType != 'cronJob'">
            <el-form-item label="使用宿主机网络">
              <el-switch
                v-model="workLoadItem.item.spec.template.spec.hostNetwork"
                style="--el-switch-on-color: #13ce66"
                @change="hostNetworkSwitch"
              />
            </el-form-item>
          </el-col>
          <!-- 标签及注释 -->
          <el-col :span="8">
            <el-form-item label="标签及注释">
              <el-radio-group v-model="data.labelsAndAnnotationsSwtich">
                <el-radio value="auto" size="large">自动生成</el-radio>
                <el-radio value="manual" size="large">手动配置</el-radio>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <!-- 自动添加 Servcie -->
          <el-col :span="8">
            <el-form-item label="自动添加 Servcie" v-if="props.resourceType == 'deployment'">
              <el-switch v-model="data.switchAddService" />
            </el-form-item>
            <el-form-item label="绑定 Servcie" v-if="props.resourceType == 'statefulSet'">
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
          <!-- 控制器标签 -->
          <el-tab-pane label="控制器标签">
            <TableOfKeyValue
              :table-list="data.controllerLabelsList"
              @add-table-row="addControllerLabelItem"
              @delete-table-row="deleteControllerLabelItem"
            />
          </el-tab-pane>
          <!-- 控制器注释 -->
          <el-tab-pane label="控制器注释">
            <TableOfKeyValue
              :table-list="data.controllerAnnotationsList"
              @add-table-row="addControllerAnnotationItem"
              @delete-table-row="deleteControllerAnnotationItem"
            />
          </el-tab-pane>
          <!-- pod标签 -->
          <el-tab-pane label="Pod 标签">
            <TableOfKeyValue
              :table-list="data.podLabelsList"
              @add-table-row="addPodLabelItem"
              @delete-table-row="deletePodLabelItem"
            />
          </el-tab-pane>
          <!-- pod注释 -->
          <el-tab-pane label="Pod 注释">
            <TableOfKeyValue
              :table-list="data.podAnnotationsList"
              @add-table-row="addPodAnnotationItem"
              @delete-table-row="deletePodAnnotationItem"
            />
          </el-tab-pane>
        </el-tabs>
      </el-form>
    </template>
  </ElCard>
</template>
