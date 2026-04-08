<script setup>
import { computed, ref } from 'vue'
import DialogByYaml from '../components/DialogByYaml.vue'
import { obj2yaml } from '../../utils/typeConv/type.conv.js'
import DialogHeaderLabel from '../components/DialogHeaderLabel.vue'

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
            item.metadata.name.toLowerCase().includes(props.tableData.search.toLowerCase())
    )
)
const curItem = ref('')
const itemByYaml = ref('')
const itemDetailDialog = ref(false)
const getItem = (row) => {
    curItem.value = row
    itemByYaml.value = obj2yaml(row)
    itemDetailDialog.value = true
}
</script>

<template>
    <el-table :data="filterTableData" height="70vh" >
        <el-table-column label="名称" prop="metadata.name" width="300px">
            <template #default="scope">
                <el-button type="primary" link @click="getItem(scope.row)">{{ scope.row.metadata.name }}</el-button>
            </template>
        </el-table-column>
        <el-table-column label="创建时间" prop="metadata.creationTimestamp" />
        <el-table-column label="存活时间" prop="age">
            <template #default="scope">
                {{ getAgeText(scope.row) }}
            </template>
        </el-table-column>
        <el-table-column label="命名空间" prop="metadata.namespace" />
        <el-table-column label="状态" prop="status" >
            <template #default="scope">
                <span v-if="scope.row.status.conditions[0].status == 'True'" style="color: green;">Available</span>
                <span v-else style="color: red;">UnAvailable</span>
            </template>
        </el-table-column>
        <el-table-column label="Pods" prop="containerStatus">
            <template #default="scope">
                {{ scope.row.status.availableReplicas || 0 }}/{{ scope.row.status.replicas }}
            </template>
        </el-table-column>
        <el-table-column label="操作" prop="operations">
            <template #default="scope">
                <el-button link type="danger" @click="emit('deleteItem',scope.row)">删除</el-button>
            </template>   
        </el-table-column>
    </el-table>
    
    <DialogByYaml
        :dialogVisible="itemDetailDialog"
        @closeDialog="itemDetailDialog=false"
        :item-by-yaml="itemByYaml"
    >
        <template #header>
            <DialogHeaderLabel 
                :cur-cluster-id="props.tableData.curClusterId"
                :cur-ns-name="curItem.metadata.namespace"
                :cur-resource-name="curItem.metadata.name"
                :cur-node-name="curItem.spec.nodeName"
                cur-resource-kind="Deployment"
            />
        </template>
    </DialogByYaml>
</template> 