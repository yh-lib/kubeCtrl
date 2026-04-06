<script setup>
import { computed } from 'vue'

const emit = defineEmits(['deleteItem'])

const getContainerStatus = (row) => {
    const statuses = row?.status?.containerStatuses || []
    const total = statuses.length
    const ready = statuses.filter((item) => item?.ready).length
    return `${ready}/${total}`
}

const getRestartCount = (row) => {
    const statuses = row?.status?.containerStatuses || []
    if (!statuses.length) {
        return 0
    }
    return Math.max(...statuses.map((item) => item?.restartCount || 0))
}

const getAgeText = (row) => {
    const creationTimestamp = row?.metadata?.creationTimestamp
    if (!creationTimestamp) {
        return '-'
    }

    const createdAt = new Date(creationTimestamp).getTime()
    if (Number.isNaN(createdAt)) {
        return '-'
    }

    const diffSeconds = Math.max(0, Math.floor((Date.now() - createdAt) / 1000))
    if (diffSeconds < 60) {
        return `${diffSeconds}秒钟前`
    }

    const diffMinutes = Math.floor(diffSeconds / 60)
    if (diffMinutes < 60) {
        return `${diffMinutes}分钟前`
    }

    const diffHours = Math.floor(diffMinutes / 60)
    if (diffHours < 24) {
        return `${diffHours}小时前`
    }

    const diffDays = Math.floor(diffHours / 24)
    return `${diffDays}天前`
}

const props = defineProps({
    tableData: {
        type: Object,
        default: () => ({
            items: [],
            item: {},
            curClusterId: '',
            curNsName: '',
            search: '',
        }),
    },
})

// 搜索后的 pod 列表数据
const filterTableData = computed(() =>
    (props.tableData.items || []).filter(
        (item) =>
            !props.tableData.search ||
            item.metadata.name.toLowerCase().includes(props.tableData.search.toLowerCase()) ||
            item.status.podIP.toLowerCase().includes(props.tableData.search.toLowerCase()) || 
            item.status.phase.toLowerCase().includes(props.tableData.search.toLowerCase()) || 
            item.spec.nodeName.includes(props.tableData.search.toLowerCase()) ||
            item.status.hostIP.toLowerCase().includes(props.tableData.search.toLowerCase())
    )
)

const getItem = (row) => {
    console.log('getItem', row)
    props.tableData.item = row
    // emit('getItem', row)
}
</script>

<template>
    <el-table :data="filterTableData" height="70vh" >
        <el-table-column label="Pod 名称" prop="metadata.name" width="300px">
            <template #default="scope">
                <el-button type="primary" link @click="getItem(scope.row)">{{ scope.row.metadata.name }}</el-button>
            </template>
        </el-table-column>
        <el-table-column label="Pod IP" prop="status.podIP" />
        <el-table-column label="创建时间" prop="metadata.creationTimestamp" />
        <el-table-column label="存活时间" prop="age">
            <template #default="scope">
                {{ getAgeText(scope.row) }}
            </template>
        </el-table-column>
        <el-table-column label="命名空间" prop="metadata.namespace" />
        <el-table-column label="状态" prop="status.phase" />
        <el-table-column label="重启次数" prop="restartCount">
            <template #default="scope">
                {{ getRestartCount(scope.row) }}
            </template>
        </el-table-column>
        <el-table-column label="容器状态" prop="containerStatus">
            <template #default="scope">
                {{ getContainerStatus(scope.row) }}
            </template>
        </el-table-column>
        <el-table-column label="宿主机名" prop="spec.nodeName" />
        <el-table-column label="宿主机IP" prop="status.hostIP" />
        <el-table-column label="操作" prop="operations">
            <template #default="scope">
                <el-button link type="danger" @click="emit('deleteItem',scope.row)">删除</el-button>
            </template>   
        </el-table-column>
    </el-table>
</template> 