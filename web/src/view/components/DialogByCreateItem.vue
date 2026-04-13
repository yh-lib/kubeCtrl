<script setup>
import { reactive, ref } from 'vue'
import ElCard from './ElCard.vue';
import { getSecretListHandler } from '../../api/secret';
import { createdeploymentHandler } from '../../api/deployment';
import { ElMessage } from 'element-plus';

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
  updatePlicyList:[{value:'RollingUpdate'},{value:'OnDelete'}],
  imagePullPolicyList: [{value:'Never'},{value:'IfNotPresent'},{value:'Always'}],
  switchAddService: false,
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

const dataSend = reactive({
  // 存储创建对象数据
  itemData: {
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
              "app": ""
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
            "imagePullSecrets": [  // ← 关键新增字段
                {
                  "name": ""
                }
            ]
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
  }, 
})

const createItem = () => {
  dataSend.itemData.clusterId = data.curClusterId
  dataSend.itemData.nameSpace = data.curNsName
  dataSend.itemData.name = data.curResourceName
  dataSend.itemData.item.metadata.name = data.curResourceName
  dataSend.itemData.item.metadata.namespace = data.curNsName
  dataSend.itemData.item.metadata.labels.app = data.curResourceName
  dataSend.itemData.item.spec.selector.matchLabels.app = data.curResourceName
  dataSend.itemData.item.spec.template.metadata.labels.app = data.curResourceName
  createdeploymentHandler(dataSend.itemData)

  console.log('创建资源：', dataSend.itemData)
}

const updatePoicySwitch = () => {
  dataSend.itemData.item.spec.strategy.type == 'RollingUpdate'?data.updatePoicySwitch=true:data.updatePoicySwitch=false
}

const hostNetworkSwitch = () => {
  if (dataSend.itemData.item.spec.template.spec.hostNetwork == true) {
    ElMessage({
      message: '提示: 已开启宿主机网络',
      type: 'warning',
    })
  }
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
                    <el-col :span="8" class="form-item">
                      <el-form-item label="名称" required>
                        <el-input placeholder="请输入资源名称" v-model="data.curResourceName"/>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <el-form-item label="副本数" >
                        <el-input placeholder="请输入副本数" v-model="dataSend.itemData.item.spec.replicas"/>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <el-form-item label="镜像地址">
                        <el-input placeholder="请输入镜像地址" v-model="dataSend.itemData.item.spec.template.spec.containers[0].image" />
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <el-form-item label="镜像拉取策略">
                        <el-select placeholder="请选择镜像拉取策略" v-model="dataSend.itemData.item.spec.template.spec.containers[0].imagePullPolicy">
                          <el-option 
                            v-for="item in data.imagePullPolicyList"
                            :key="item.value"
                            :label="item.value"
                            :value="item.value"
                          />
                        </el-select>
                      </el-form-item>
                    </el-col>                   
                     <el-col :span="8" class="form-item">
                      <el-form-item label="私有仓库密钥">
                        <el-select placeholder="请选择私有仓库密钥" v-model="dataSend.itemData.item.spec.template.spec.imagePullSecrets[0].name">
                          <el-option 
                            v-for="item in data.privateRepoSecretList"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                          />
                        </el-select>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <el-form-item label="DNS 策略">
                        <el-select placeholder="请选择 DNS 策略" v-model="dataSend.itemData.item.spec.template.spec.dnsPolicy">
                          <el-option 
                            v-for="item in data.dnsPolicyList"
                            :key="item.value"
                            :label="item.value"
                            :value="item.value"
                          />
                        </el-select>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <el-form-item label="更新策略">
                        <el-select placeholder="请选择更新策略" v-model="dataSend.itemData.item.spec.strategy.type" @change="updatePoicySwitch">
                          <el-option 
                            v-for="item in data.updatePlicyList"
                            :key="item.value"
                            :label="item.value"
                            :value="item.value"
                          />
                        </el-select>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <div style="display: flex;justify-content: space-between;" v-show="data.updatePoicySwitch">
                        <el-form-item label="最大不可用" label-width="100px">
                          <el-input placeholder="" style="width: 100px" v-model="dataSend.itemData.item.spec.strategy.rollingUpdate.maxUnavailable"/>
                        </el-form-item>
                        <el-form-item label="可超出最大值" label-width="100px">
                          <el-input placeholder="" style="width: 100px" v-model="dataSend.itemData.item.spec.strategy.rollingUpdate.maxSurge"/>
                        </el-form-item>
                      </div>
                    </el-col>              
                    <el-col :span="8" class="form-item">
                      <el-form-item label="使用宿主机网络">
                        <el-switch
                          v-model="dataSend.itemData.item.spec.template.spec.hostNetwork"
                          style="--el-switch-on-color: #13ce66;"
                          @change="hostNetworkSwitch"
                        />
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <el-form-item label="自动添加 Servcie">
                        <el-switch
                          v-model="data.switchAddService"
                        />
                      </el-form-item>
                    </el-col>
                  </el-row>
                <!-- 标签|注释|容忍 -->
                <el-tabs tab-position="left" style="height: 260px" type="border-card">
                    <!-- 标签 -->
                    <el-tab-pane label="标签">

                    </el-tab-pane>
                    <!-- 注释 -->
                    <el-tab-pane label="注释">

                    </el-tab-pane>
                    <!-- 容忍 -->
                    <el-tab-pane label="容忍">

                    </el-tab-pane>
                </el-tabs>
                </el-form>
              </template>                
            </ElCard>
        </el-tab-pane>
        <el-tab-pane label="存储卷配置" name="Volume">存储卷配置</el-tab-pane>
        <el-tab-pane label="调度配置" name="Schedule">调度配置</el-tab-pane>
        <el-tab-pane label="容器配置" name="Container">容器配置</el-tab-pane>
        <el-tab-pane label="初始化容器" name="InitContainer">初始化容器</el-tab-pane>
    </el-tabs>
    <el-button type="primary" size="large" style="margin-top: 20px;width: 90px;" @click="createItem">创建</el-button>
  </el-dialog>
</template>

<style scoped>
.form-item{
  margin-bottom: 15px;
}
</style>

