<script setup>
import { onBeforeMount, reactive, readonly, ref } from 'vue';
import { useWorkLoadData } from '../../../../store';
import { storeToRefs } from 'pinia';
import { getSecretListHandler } from '../../../../api/secret.js';

// store from pinia
const store = useWorkLoadData()
const { workLoadItem } = storeToRefs(store)

const data = reactive({
    volumeItem:{
        'name':'',              
        'secret':{
            'secretName': '',
            'defaultMode': null,
            'optional': null,  
            'items': [],        
        }
    },
    'Options':[
        {'label':'是','value':true},
        {'label':'否','value':false}
    ],
    secretList:[]
})

const validate = () => ruleFormRef.value?.validate()

// 暴露 data 给父组件
defineExpose({
    data,
    validate
})

// 接收父组件参数
const props = defineProps({
    method: String,
    parentVolumeItem: {
        type: Object,
        default: {},
    }
})

// 数据渲染
onBeforeMount(()=>{
    if (props.method != 'add') {
        data.volumeItem = props.parentVolumeItem
    }
})

// 表单校验：存储卷名称不能重复
const validateVolumeName = (rule, value, callback) => {
    // 如果没有输入，视为必填不通过，返回对应错误信息
    if (!value) {
        callback(new Error('请输入存储卷名称'))
        return
    }

    // 从当前 workload 中取出 volumes 列表，未取到时用空数组兜底
    const volumeList = workLoadItem.value?.item?.spec?.template?.spec?.volumes || []

    // 查找是否存在与当前输入同名且不是正在编辑对象的 volume
    const duplicated = volumeList.some(item => item !== data.volumeItem && item?.name === value)
    
    if (duplicated) {
        callback(new Error('存储卷名称不能重复'))
        return
    }
    callback()
}

// 表单校验
const ruleFormRef = ref()
const rules = reactive({
    name: [
        { validator: validateVolumeName, trigger: 'blur' },
    ],
    items: {
        key: [
            { required: true, message: '请输入secret Key', trigger: 'blur' },  
        ],
        path: [
            { required: true, message: '请输入挂载路径', trigger: 'blur' },
            // { pattern: '^(?=/)', message: '挂载路径地址必须以 / 开头', trigger: 'blur' },
        ],
        mode: [
            { type: 'number', min: 0, max: 777, message: '权限需在 000-777 之间', trigger: ['blur','change'] }
        ],
    },
})

// 获取 secretList
const getSecretList = () => {   
    getSecretListHandler(workLoadItem.value.clusterId,workLoadItem.value.nameSpace).then((res)=>{
        if (res.data.status === 200) {
            data.secretList = res.data.data.items
        }
        console.log(data.secretList)
    })
}

// 添加自定义挂载路径
const addMountPath = () => {
    data.volumeItem.secret.items.push({
        'key':'',
        'path':'',
        'mode':''
    })
}
const deleteMountPath = (index) => {
    data.volumeItem.secret.items.splice(index,1)
}

// 权限 number 转换成前后端都能识别的十进制
const oct2Dec = (row) => {
    if (typeof(row) != 'object'){
        data.volumeItem.secret.defaultMode = parseInt(data.volumeItem.secret.defaultMode,8)
    }else{
        row.mode = parseInt(row.mode,8)
    }
}
</script>

<template>
    <el-form 
        label-width="120px" 
        label-position="left" 
        size="large"
        :rules="rules"
        ref="ruleFormRef"
        :model="data.volumeItem"           
    >
        <el-tabs>
            <el-tab-pane label="基础配置">
                <el-form-item label="名称" prop="name" v-show="props.method">
                    <el-input v-model="data.volumeItem.name" placeholder="请输入Volume名称"/>                    
                </el-form-item>
                <el-form-item label="secret" prop="secret.secretName">
                    <el-select v-model="data.volumeItem.secret.secretName" placeholder="请选择要挂载的secret" @visible-change="getSecretList">
                        <el-option
                            v-for="item in data.secretList"
                            :key="item.metadata.name"
                            :label="item.metadata.name"
                            :value="item.metadata.name"
                        />
                    </el-select>                    
                </el-form-item>
                <el-form-item label="默认挂载权限" prop="secret.defaultMode">
                    <el-input v-model.number="data.volumeItem.secret.defaultMode" placeholder="请输入默认挂载权限" @change="oct2Dec"/>                    
                </el-form-item>
                <el-form-item label="可选" prop="secret.optional">
                    <el-select v-model="data.volumeItem.secret.optional" placeholder="secret 缺失时是否允许 Pod 启动">
                        <el-option
                            v-for="item in data.Options"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
            </el-tab-pane>
            <el-tab-pane label="自定义挂载路径" class="no-border-input">
                <el-table :data="data.volumeItem.secret.items">
                    <el-table-column prop="key" label="Key" width="200px">
                        <template #default="scope">
                            <el-form-item label="" label-width="0px" :prop="('secret.items.'+scope.$index+'.key')" :rules="rules.items.key">
                                <el-input v-model="scope.row.key" placeholder="请输入挂载的Key"/> 
                            </el-form-item>                            
                        </template>                        
                    </el-table-column>
                    <el-table-column prop="path" label="挂载路径">
                        <template #default="scope">
                            <el-form-item label="" label-width="0px" :prop="('secret.items.'+scope.$index+'.path')" :rules="rules.items.path">
                                <el-input v-model="scope.row.path" placeholder="请输入挂载路径"/> 
                            </el-form-item>
                        </template>  
                    </el-table-column>
                    <el-table-column prop="mode" label="挂载权限" width="200px">
                        <template #default="scope">
                            <el-form-item label="" label-width="0px" :prop="('secret.items.'+scope.$index+'.mode')" :rules="rules.items.mode">
                                <el-input v-model.number="scope.row.mode" placeholder="请输入挂载权限" @change="oct2Dec(scope.row)"/>
                            </el-form-item>                            
                        </template>                        
                    </el-table-column>
                    <el-table-column prop="operation" align="center" width="100px">
                        <template #header>
                            <el-button type="primary" link size="large" style="margin-bottom: 5px;" @click="addMountPath">添加</el-button>
                        </template>
                        <template #default="scope">
                            <el-button type="danger" link size="large" @click="deleteMountPath(scope.$index)">删除</el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-tab-pane>
        </el-tabs>
    </el-form>
</template>    