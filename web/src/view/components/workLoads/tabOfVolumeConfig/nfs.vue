<script setup>
import { onBeforeMount, reactive, readonly, ref } from 'vue';
import { useWorkLoadData } from '../../../../store';
import { storeToRefs } from 'pinia';

// store from pinia
const store = useWorkLoadData()
const { workLoadItem } = storeToRefs(store)

const data = reactive({
    volumeItem:{
        'name':'',
        'nfs':{
            'server':'',
            'path':'',
            'readOnly':false
        }
    },
    'readOnlyOptions':[true,false]
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
    'nfs.server':[
        { required: true, message: '请输入NFS服务器地址', trigger: 'blur' },
    ],
    'nfs.path': [
        { required: true, message: '请输入NFS路径地址', trigger: 'blur' },
        { pattern: '^(?=/)', message: 'NFS路径地址必须以 / 开头', trigger: 'blur' },
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
        <el-form-item label="NFS服务器地址" prop="nfs.server">
            <el-input v-model="data.volumeItem.nfs.server" placeholder="请输入NFS服务器地址"/>                    
        </el-form-item>
        <el-form-item label="NFS路径地址" prop="nfs.path">
            <el-input v-model="data.volumeItem.nfs.path" placeholder="请输入NFS路径地址"/>                    
        </el-form-item>
        <el-form-item label="是否只读" prop="nfs.readOnly">
            <el-select v-model="data.volumeItem.nfs.readOnly" placeholder="是否只读">
                <el-option
                    v-for="item in data.readOnlyOptions"
                    :key="item"
                    :label="item"
                    :value="item"
                />
            </el-select>                    
        </el-form-item>
    </el-form>
</template>    