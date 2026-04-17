<script setup>
import { onBeforeMount, reactive, ref } from 'vue';
import { useWorkLoadData } from '../../../../store';
import { storeToRefs } from 'pinia';

// store from pinia
const store = useWorkLoadData()
const { workLoadItem } = storeToRefs(store)

const data = reactive({
    volumeItem:{
        'name':'',
        'hostPath':{
            'path':'',
            'type':''
        }
    },
    mountTypeOptions:[
        { value: '', label: '默认' },
        { value: 'Directory', label: '目录' },
        { value: 'DirectoryOrCreate', label: '目录，不存在则创建' },
        { value: 'File', label: '文件' },
        { value: 'FileOrCreate', label: '文件，不存在则创建' },
        { value: 'Socket', label: 'Unix 域套接字' },
        { value: 'CharDevice', label: '字符设备' },
        { value: 'BlockDevice', label: '块设备' }
    ]
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
    if (!value) {
        callback()
        return
    }

    const volumeList = workLoadItem.value?.item?.spec?.template?.spec?.volumes || []
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
        { required: true, message: '请输入存储卷名称', trigger: 'blur' },
        { validator: validateVolumeName, trigger: 'blur' },
    ],
    'hostPath.path': [
        { required: true, message: '请输入宿主机路径', trigger: 'blur' },
        { pattern: '^(?=/)', message: '宿主机路径必须以 / 开头', trigger: 'blur' },
    ]
})

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
        <!-- 名称 -->
        <el-form-item label="名称" prop="name" v-show="props.method">
            <el-input v-model="data.volumeItem.name" placeholder="请输入Volume名称"/>                    
        </el-form-item>
        <el-form-item label="宿主机路径" prop="hostPath.path">
            <el-input v-model="data.volumeItem.hostPath.path" placeholder="请输入宿主机路径"/>                    
        </el-form-item>
        <el-form-item label="挂载类型" prop="mountType">
            <el-select v-model="data.volumeItem.hostPath.type" placeholder="请选择存储卷类型" :empty-values="[null, undefined]">
            <el-option
                v-for="item in data.mountTypeOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            />
        </el-select>                    
        </el-form-item>
    </el-form>
</template>    