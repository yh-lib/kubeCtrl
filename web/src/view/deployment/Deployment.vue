<script setup>
import ElCard from '../components/ElCard.vue';
import Table from './Table.vue';
import { reactive, ref } from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus'
import { getdeploymentListHandler,deletedeploymentHandler } from '../../api/deployment'
import DialogByCreateItem from '../components/DialogByCreateItem.vue'

// 删除 deployment
const deleteItem = (row) => {
    // 删除提醒
    ElMessageBox.confirm(
        '确认删除 deployment :  ' + row.metadata.name,
        {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
    .then(() => {
        deletedeploymentHandler(data.curClusterId,data.curNsName,row.metadata.name).then((res)=>{
            if (res.data.status == 200) {
                ElMessage({
                    type: 'success',
                    message: res.data.message,
                })
                getList()                
            }
        })
    })
    .catch(() => {
        return
    }) 
}

// 渲染表格数据
const getList = () => {
    if (!data.curClusterId || !data.curNsName) {
        data.items = []
        return
    }
    getdeploymentListHandler(data.curClusterId, data.curNsName).then((res) => {
        data.items = res.data.data.items || []
    })
}

// 从子组件 ELCard 中获取所需数据
const getSelectValue = (selectValue) =>{
    Object.assign(data, selectValue)
    getList()
}

// 接受子组件table传递的参数
const data = reactive({})
const createItemData = reactive({})
const createItemDialogVisible = ref(false)
</script>

<template>
    <!-- 卡片主体: 挂载前会自动获取当前的选择，并通过getSelectValue赋值给data -->
    <ElCard
        title="Deployment 列表"
        :op-cluster="true"
        :op-ns="true"
        :op-search="true"
        :op-refresh="true"
        :op-create="true"
        @change="getSelectValue"
        @refresh="getList"
        @create-item="createItemDialogVisible=true"
     >
        <!-- 卡片 main 部分 table 数据 -->
        <template #mainData>
            <Table :table-data="data" @delete-item="deleteItem"></Table>
        </template>
     </ElCard>

         <!-- 创建 item 按钮 -->
    <DialogByCreateItem 
        :open-dialog="createItemDialogVisible" 
        @close-dialog="createItemDialogVisible=false"
    >
    </DialogByCreateItem>
</template>    