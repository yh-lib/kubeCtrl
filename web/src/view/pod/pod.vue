<script setup>
import { reactive } from 'vue';
import Elcard from '../components/ELCard.vue';
import Table from './Table.vue';
import { ElMessage, ElMessageBox } from 'element-plus'
import { getPodListHandler,deletePodHandler } from '../../api/pod'

// 删除 pod
const deleteItem = (row) => {
    // 删除提醒
    ElMessageBox.confirm(
        '确认删除 pod :  ' + row.metadata.name,
        {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
    .then(() => {
        deletePodHandler(data.curClusterId,data.curNsName,row.metadata.name).then((res)=>{
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
    getPodListHandler(data.curClusterId, data.curNsName).then((res) => {
        data.items = res.data.data.items || []
    })
}

// 从子组件 ELCard 中获取所需数据
const getSelectValue = (selectValue) =>{
    Object.assign(data, selectValue)
    getList()
}

// 接受子组件传递的参数
const data = reactive({})
</script>

<template>
    <!-- 卡片主体 -->
    <Elcard 
        title="pod 列表" 
        :op-cluster="true" 
        :op-ns="true" 
        :op-search="true" 
        :op-refresh="true"
        @change="getSelectValue"
        @refresh="getList"
    >
        <!-- 卡片 main 部分 table 数据 -->
        <template #table>
            <Table :table-data="data" @delete-item="deleteItem"></Table>
        </template>
    </Elcard> 
</template>