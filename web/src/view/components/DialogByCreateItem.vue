<script setup>
import { onBeforeMount, reactive, ref } from 'vue'
import ElCard from './ElCard.vue';
import { getSecretListHandler } from '../../api/secret';
import { createdeploymentHandler } from '../../api/deployment';
import { ElMessage } from 'element-plus';
import { list2obj, obj2list } from '../../utils/typeConv/type.conv';
import TableOfLabels from './tableOfLabels.vue';
import TableOfAnnotations from './tableOfAnnotations.vue';
import TableOfTolerations from './TableOfTolerations.vue';

const props = defineProps(['openDialog'])
const emit = defineEmits(['closeDialog'])
const activeName = ref('Basic')
const data = reactive({
  updatePoicySwitch: false,
  curClusterId:'',
  curNsName:'',
  curResourceName: '',
  privateRepoSecretList:[], // 存储私有仓库密钥列表
  dnsPolicyList:[{value:'Default'},{value:'ClusterFirst'},{value:'ClusterFirstWithHostNet'}], // 存储dns策略列表
  updatePlicyList:[{value:'RollingUpdate'},{value:'Recreate'}],
  imagePullPolicyList: [{value:'Never'},{value:'IfNotPresent'},{value:'Always'}],
  switchAddService: false,
  itemLabelsList: [],
  itemAnnotationsList: [],
  tolerationOperatorValue: 'Equal',
  tolerationEffectValue: 'NoSchedule',
  labelsAndAnnotationsSwtich: 'auto',
  // 容忍操作选项
  tolerationsOperatorSelectOptions : [
    {value: 'Equal', label: '等值'},
    {value: 'Exists',label: '存在'}
  ],
  // 容忍影响选项
  tolerationsEffectSelectOptions : [
    {label: "禁止调度", value: "NoSchedule"},
    {label: "驱逐", value: "NoExecute"},
    {label: "尽量不调度", value: "PreferNoSchedule"}
  ]
})

// 从子组件 ELCard 中获取所需数据
const getSelectValue = (selectValue) =>{
    Object.assign(data, selectValue)
    getSecretListHandler(data.curClusterId,data.curNsName).then((res)=>{
      if (res.data.status == 200) {
          data.privateRepoSecretList = []
          data.privateRepoSecretValue = ''
          res.data.data.items == null ||
          res.data.data.items.forEach(item => {
            item.data['.dockerconfigjson'] == null ||
            data.privateRepoSecretList.push({
              label:item.metadata.name,
              value:item.metadata.name,
            })
            console.log('获取到的数据有：',data.privateRepoSecretList)
          });                   
      }
  })
}

const formData = reactive({
  'clusterId':'',
  'nameSpace':'',
  'name':'',
  'item':{
    "kind": "Deployment",
    "apiVersion": "apps/v1",
    "metadata": {
      "name": "",
      "namespace": "",
      "labels": {
        "app": ""
      },
      "annotations": {
        "deployment.kubernetes.io/revision": "1",
      },
    },
    "spec": {
      "replicas": 3,
      "selector": {
        "matchLabels": {
          "app": ""
        }
      },
      "template": {
        "metadata": {
          "labels": {
            "app": "testApp"
          }
        },
        "spec": {
          "hostNetwork": false,
          "containers": [
            {
              "name": "container-1",
              "image": "",
              "resources": {},
              "imagePullPolicy": "IfNotPresent"
            }
          ],
          "restartPolicy": "Always",
          "terminationGracePeriodSeconds": 30,
          "dnsPolicy": "Default",
          "securityContext": {},
          "schedulerName": "default-scheduler",
          "tolerations":[
            {
              key:'node-role.kubernetes.io/master',
              operator: 'Exists',
              effect: 'NoSchedule'
            }
          ],
          "imagePullSecrets": [{}]
        }
      },
      "strategy": {
        "type": "RollingUpdate",
        "rollingUpdate": {
          "maxUnavailable": "25%",
          "maxSurge": "25%"
        }
      },
      "revisionHistoryLimit": 10,
      "progressDeadlineSeconds": 600
    }
  },
}
)

const createItem = () => {
  formData.clusterId = data.curClusterId
  formData.nameSpace = data.curNsName
  formData.name = data.curResourceName
  formData.item.metadata.name = data.curResourceName
  formData.item.metadata.namespace = data.curNsName
  formData.item.metadata.labels.app = data.curResourceName
  formData.item.spec.selector.matchLabels.app = data.curResourceName
  formData.item.spec.template.metadata.labels.app = data.curResourceName
  formData.item.metadata.labels = list2obj(data.itemLabelsList)
  formData.item.metadata.annotations = list2obj(data.itemAnnotationsList)
  if (formData.item.spec.template.spec.imagePullSecrets && formData.item.spec.template.spec.imagePullSecrets.length && formData.item.spec.template.spec.imagePullSecrets[0].name == ''){
    delete formData.item.spec.template.spec.imagePullSecrets
  }
  createdeploymentHandler(formData).then((res)=>{
      if (res.data.status === 200) {
        ElMessage({
                    message: res.data.message,
                    type: 'success',
        })
      }
  })
}

const updatePoicySwitchFunc = () => {
  const isRolling = formData.item.spec.strategy.type === 'RollingUpdate'
  data.updatePoicySwitch = isRolling
  if (isRolling) {
    formData.item.spec.strategy.rollingUpdate = {
      maxUnavailable: '25%',
      maxSurge: '25%'
    }
  } else {
    delete formData.item.spec.strategy.rollingUpdate
  }
}

const hostNetworkSwitch = () => {
  if (formData.item.spec.template.spec.hostNetwork == true) {
    ElMessage({
      message: '提示: 已开启宿主机网络',
      type: 'warning',
    })
  }
}

onBeforeMount(()=>{
  data.itemLabelsList = obj2list(formData.item.metadata.labels)
  // 同步初始的更新策略显示状态，防止首次打开时 v-show 与 dataSend 不一致
  data.updatePoicySwitch = formData.item.spec.strategy.type === 'RollingUpdate'
  // 确保 rollingUpdate 初始存在，避免模板访问 undefined
  if (data.updatePoicySwitch && !formData.item.spec.strategy.rollingUpdate) {
    formData.item.spec.strategy.rollingUpdate = {
      maxUnavailable: '25%',
      maxSurge: '25%'
    }
  }
  data.itemAnnotationsList = obj2list(formData.item.metadata.annotations)
})

// 标签列表项
const addLabelItem = () => {data.itemLabelsList.unshift({key:"",value:""})}
const deleteLabelItem = (index) => {
    data.itemLabelsList.splice(index,1)
}
// 注释列表项
const addAnnotationItem = () => {data.itemAnnotationsList.unshift({key:"",value:""})}
const deleteAnnotationItem = (index) => {
    data.itemAnnotationsList.splice(index,1)
}
// 容忍列表项
const addTolerationItem = () => {formData.item.spec.template.spec.tolerations.unshift({key:"",value:""})}
const deleteTolerationItem = (index) => {
    formData.item.spec.template.spec.tolerations.splice(index,1)
}
</script>

<template>
  <!-- 资源创建 dialog -->
  <el-dialog
    v-model="props.openDialog"
    width="1600px"
    style="height: 740px;"
    @close="emit('closeDialog')"
  >
    <el-tabs v-model="activeName">
        <!-- 标签：基本配置 -->
        <el-tab-pane label="基本配置" name="Basic">
            <ElCard :op-cluster="true" :op-ns="true" style="border-radius: 0px;width: 1560px;" @change="getSelectValue">
              <template #mainData>
                <!-- 基础信息 -->
                <el-form label-width="150px" label-position="left" style="height: 450px;width: 1490px;">
                  <el-row :gutter="100">
                    <!-- 名称 -->
                    <el-col :span="8" class="form-item">
                      <el-form-item label="名称" required>
                        <el-input placeholder="请输入资源名称" v-model="data.curResourceName"/>
                      </el-form-item>
                    </el-col>
                    <!-- 副本数 -->
                    <el-col :span="8" class="form-item">
                      <el-form-item label="副本数" >
                        <el-input placeholder="请输入副本数" v-model.number="formData.item.spec.replicas"/>
                      </el-form-item>
                    </el-col>
                    <!-- 镜像地址 -->
                    <el-col :span="8" class="form-item">
                      <el-form-item label="镜像地址">
                        <el-input placeholder="请输入镜像地址" v-model="formData.item.spec.template.spec.containers[0].image" />
                      </el-form-item>
                    </el-col>
                    <!-- 镜像拉取策略 -->
                    <el-col :span="8" class="form-item">
                      <el-form-item label="镜像拉取策略">
                        <el-select placeholder="请选择镜像拉取策略" v-model="formData.item.spec.template.spec.containers[0].imagePullPolicy">
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
                        v-if="formData.item.spec.template.spec.imagePullSecrets && formData.item.spec.template.spec.imagePullSecrets.length"
                        placeholder="请选择私有仓库密钥"
                        v-model="formData.item.spec.template.spec.imagePullSecrets[0].name"
                      >
                        <el-option 
                          v-for="item in data.privateRepoSecretList"
                          :key="item.value"
                          :label="item.label"
                          :value="item.value"
                        />
                      </el-select>
                      <el-select v-else disabled placeholder="无可用密钥">
                        <el-option :label="'无'" :value="''" />
                      </el-select>
                    </el-form-item>
                  </el-col>
                  <!-- DNS 策略 -->
                  <el-col :span="8" class="form-item">
                    <el-form-item label="DNS 策略">
                      <el-select placeholder="请选择 DNS 策略" v-model="formData.item.spec.template.spec.dnsPolicy">
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
                      <el-select placeholder="请选择更新策略" v-model="formData.item.spec.strategy.type" @change="updatePoicySwitchFunc">
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
                    <div style="display: flex;justify-content: space-between;" v-if="data.updatePoicySwitch">
                      <el-form-item label="最大不可用" label-width="100px">
                        <el-input placeholder="" style="width: 100px" v-model="formData.item.spec.strategy.rollingUpdate.maxUnavailable"/>
                      </el-form-item>
                      <el-form-item label="可超出最大值" label-width="100px">
                        <el-input placeholder="" style="width: 100px" v-model="formData.item.spec.strategy.rollingUpdate.maxSurge"/>
                      </el-form-item>
                    </div>
                  </el-col>          
                  <!-- 使用宿主机网络 -->
                  <el-col :span="8" class="form-item">
                    <el-form-item label="使用宿主机网络">
                      <el-switch
                        v-model="formData.item.spec.template.spec.hostNetwork"
                        style="--el-switch-on-color: #13ce66;"
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
                      <el-switch
                        v-model="data.switchAddService"
                      />
                    </el-form-item>
                  </el-col>
                </el-row>
                <!-- 标签|注释|容忍 tabs -->
                <el-tabs tab-position="left" style="height: 260px" type="border-card" class="no-border-input" v-if="data.labelsAndAnnotationsSwtich=='manual'">
                    <!-- 控制器标签标签 -->
                    <el-tab-pane label="控制器标签">
                        <TableOfLabels
                          :label-list="data.itemLabelsList"
                          @add-label="addLabelItem"
                          @delete-label="deleteLabelItem"
                        />
                    </el-tab-pane>
                    <!-- 控制器注释 -->
                    <el-tab-pane label="控制器注释">
                        <TableOfAnnotations 
                          :annotations-list="data.itemAnnotationsList"
                          @add-annotation="addAnnotationItem"
                          @delete-annotation="deleteAnnotationItem"
                        />
                    </el-tab-pane>
                    <!-- 容忍 -->
                    <el-tab-pane label="容忍">
                        <TableOfTolerations 
                          :tolerations-list="formData.item.spec.template.spec.tolerations"
                          :tolerations-operator-select-options="data.tolerationsOperatorSelectOptions"
                          :tolerations-effect-select-options="data.tolerationsEffectSelectOptions"
                          @add-toleration="addTolerationItem"
                          @delete-toleration="deleteTolerationItem"                      
                        />
                    </el-tab-pane>
                </el-tabs>
                </el-form>
              </template>                
            </ElCard>
        </el-tab-pane>
        <!-- <el-tab-pane label="调度配置" name="Schedule">
          
        </el-tab-pane> -->
        <!-- <el-tab-pane label="存储卷配置" name="Volume">存储卷配置</el-tab-pane>        
        <el-tab-pane label="容器配置" name="Container">容器配置</el-tab-pane>
        <el-tab-pane label="初始化容器" name="InitContainer">初始化容器</el-tab-pane> -->
    </el-tabs>
    <el-button type="primary" size="large" style="margin-top: 20px;width: 90px;" @click="createItem">创建</el-button>
  </el-dialog>
</template>

<style scoped>
.form-item{
  margin-bottom: 15px;
}
:deep(.no-border-input .el-input__wrapper){
    box-shadow: none;
    border: none;
}
</style>

