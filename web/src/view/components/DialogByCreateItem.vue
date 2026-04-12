<script setup>
import { reactive, ref } from 'vue'
import ElCard from './ElCard.vue';
import { getSecretListHandler } from '../../api/secret';

const props = defineProps(['openDialog'])
const emit = defineEmits(['closeDialog'])
const activeName = ref('Basic')
const data = reactive({
  privateRepoSecretList:[], // 存储私有仓库密钥列表
  dnsPolicyList:[{value:'Default'},{value:'ClusterFirst'},{value:'ClusterFirstWithHostNet'}], // 存储dns策略列表
  updatePlicyList:[{value:'RollingUpdate'},{value:'OnDelete'}],
  privateRepoSecretValue : '',  // 存储当前选择的私有仓库secret
  dnsPolicyValue : 'Default',  // 存储当前选择的DNS策略
  updatePlicyValue : 'RollingUpdate',  // 存储当前选择的更新策略
  maxUnavailable: '25%',
  maxSurge: '25%',
  switchHostNetwork: false,
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

const handleClick = () => {
  console.log("点击创建deployment按钮")
}
</script>

<template>
  <!-- 资源创建 dialog -->
  <el-dialog
    v-model="props.openDialog"
    width="1600px"
    style="height: 700px;"
    @close="emit('closeDialog')"
  >
    <el-tabs v-model="activeName" @tab-click="handleClick">
        <!-- 标签：基本配置 -->
        <el-tab-pane label="基本配置" name="Basic">
            <ElCard :op-cluster="true" :op-ns="true" style="border-radius: 0px;width: 1560px;" @change="getSelectValue">
              <template #mainData>
                <el-form label-width="150px" label-position="left" style="height: 450px;width: 1460px;">
                  <el-row :gutter="100">
                    <!-- 第一行 -->
                    <el-col :span="8" class="form-item">
                      <el-form-item label="名称" required>
                        <el-input placeholder="请输入资源名称"/>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <el-form-item label="副本数" >
                        <el-input placeholder="请输入副本数"/>
                      </el-form-item>
                    </el-col>
                    <el-col :span="8" class="form-item">
                      <el-form-item label="私有仓库密钥">
                        <el-select placeholder="请选择私有仓库密钥" v-model="data.privateRepoSecretValue">
                          <el-option 
                            v-for="item in data.privateRepoSecretList"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                          />
                        </el-select>
                      </el-form-item>
                    </el-col>
                    <!-- 第二行 -->
                    <el-col :span="8" class="form-item">
                      <el-form-item label="DNS 策略">
                        <el-select placeholder="请选择 DNS 策略" v-model="data.dnsPolicyValue">
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
                        <el-select placeholder="请选择更新策略" v-model="data.updatePlicyValue">
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
                      <div style="display: flex;justify-content: space-between;">
                        <el-form-item label="最大不可用" label-width="100px">
                          <el-input placeholder="" style="width: 100px" v-model="data.maxUnavailable"/>
                        </el-form-item>
                        <el-form-item label="可超出最大值" label-width="100px">
                          <el-input placeholder="" style="width: 100px" v-model="data.maxUnavailable"/>
                        </el-form-item>
                      </div>
                      
                    </el-col>
                    <!-- 第三行 -->
                    <el-col :span="8" class="form-item">
                      <el-form-item label="使用宿主机网络">
                        <el-switch
                          v-model="data.switchHostNetwork"
                          style="--el-switch-on-color: #13ce66;"
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
                </el-form>
              </template>                
            </ElCard>
        </el-tab-pane>
        <el-tab-pane label="存储卷配置" name="Volume">存储卷配置</el-tab-pane>
        <el-tab-pane label="调度配置" name="Schedule">调度配置</el-tab-pane>
        <el-tab-pane label="容器配置" name="Container">容器配置</el-tab-pane>
        <el-tab-pane label="初始化容器" name="InitContainer">初始化容器</el-tab-pane>
    </el-tabs>
  </el-dialog>
</template>

<style scoped>
.form-item{
  margin-bottom: 15px;
}
</style>

