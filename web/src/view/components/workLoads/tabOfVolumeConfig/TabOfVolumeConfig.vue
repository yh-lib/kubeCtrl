<script setup>
import { computed, reactive, ref } from 'vue';
import { useWorkLoadData } from '../../../../store';
import { storeToRefs } from 'pinia';
// 导入存储卷类型
import configMap from './configMap.vue';
import secret from './secret.vue';
import emptyDir from './emptyDir.vue';
import hostPath from './hostPath.vue';
import pvc from './pvc.vue';
import nfs from './nfs.vue';

// store from pinia
const store = useWorkLoadData()
const { workLoadItem } = storeToRefs(store)
const data = reactive({
    dialogOfAddVolumeVisible: false,
    volumeType: '',
})

defineExpose({
    data,
})

// 添加存储卷
const addTableRow = () => {
    data.dialogOfAddVolumeVisible = true
}
// 选项列表
const volumeTypeOptions = [
    'configMap',
    'secret',
    'emptyDir',
    'hostPath',
    'pvc',
    'nfs'
]
// 组件列表
const volumeTypeCommpents = {
    configMap,
    secret,
    emptyDir,
    hostPath,
    pvc,
    nfs,
}
// 添加Volume至模板文件
const addVolume = async () => {
    if (childDataRef.value?.validate) {
        const valid = await childDataRef.value.validate().catch(() => false)
        if (!valid) return
    }
    const newVolumeItem = JSON.parse(JSON.stringify(childDataRef.value.data.volumeItem))
    workLoadItem.value.item.spec.template.spec.volumes.push(newVolumeItem)
}
const deleteTableRow = (index) => {
    workLoadItem.value.item.spec.template.spec.volumes.splice(index,1)
}

const childDataRef = ref(null)

// 获取volume类型
const getVolumeType = (row) => {
    if (!row) return ''

    if (row.configMap) return 'configMap'
    if (row.secret) return 'secret'
    if (row.emptyDir) return 'emptyDir'
    if (row.hostPath) return 'hostPath'
    if (row.persistentVolumeClaim) return 'pvc'
    if (row.nfs) return 'nfs'

    return ''
}
</script>

<template>
    <el-table :data="workLoadItem.item.spec.template.spec.volumes" style="width: 100%; height:560px" class="no-border-input">
        <!-- 名称 -->
        <el-table-column prop="name" label="名称" width="400px">
            <template #default="scope">
                <el-input v-model="scope.row.name"></el-input>
            </template>
        </el-table-column>
        <!-- 类型 -->
        <el-table-column prop="type" label="类型" width="200px">
            <template #default="scope">
                {{ getVolumeType(scope.row) }}
            </template>
        </el-table-column>
         <!-- Volume 配置 -->
        <el-table-column prop="config" label="Volume 配置">   
            <template #default="scope">
                <component 
                :is="volumeTypeCommpents[getVolumeType(scope.row)]"
                :parentVolumeItem="scope.row"
                style="margin-top: 20px;"
                />
            </template>            
        </el-table-column>
        <!-- Operation -->
        <el-table-column width="200px" align="center">
            <template #header>
                <el-button type="primary" link style="font-size: 16px;margin-bottom: 10px;" @click="addTableRow">添加</el-button>
            </template>
            <template #default="scope">
                <el-button type="danger" link style="font-size: 16px;margin-bottom: 10px;" @click="deleteTableRow(scope.$index)">删除</el-button>
            </template>                      
        </el-table-column>
    </el-table>
    <!-- Dialog 添加存储卷 -->
    <el-dialog
        v-model="data.dialogOfAddVolumeVisible"
        width="800px"
        title="添加存储卷"
        style="margin-top: 350px;"
        destroy-on-close
    >
        <div style="display: flex;justify-content: left;padding: 0 70px;margin-top: 10px;">
            <!-- Select 选择存储卷类型 -->
            <span style="width: 120px;margin-top: 10px;">存储卷类型</span>
            <el-select v-model="data.volumeType" placeholder="请选择存储卷类型" size="large" style="width: 470px;margin-left: 20px;">
                <el-option
                    v-for="item in volumeTypeOptions"
                    :key="item"
                    :label="item"
                    :value="item"
                />
            </el-select>
        </div>       
        <!-- 动态选择组件 -->
        <component 
        :is="volumeTypeCommpents[data.volumeType]" 
        style="margin-top: 26px;padding:0 90px;" 
        ref="childDataRef"
        method="add"
        />
        <!-- 创建按钮 -->
         <template #footer>
            <div style="display: flex;justify-content: center;">
                <el-button type="primary" style="width: 100px;height: 40px;margin-bottom: 10px;font-size: 14px;" @click="addVolume">
                    添加
                </el-button>
            </div>
         </template>
    </el-dialog>
</template>