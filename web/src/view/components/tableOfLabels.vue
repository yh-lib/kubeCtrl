<script setup>
const props = defineProps({
    labelList:{
        type: Array,
        default: []
    },
})
const emit = defineEmits([
    'addLabel',
    'deleteLabel'
])
</script>

<template>
    <el-table :data="props.labelList" style="width: 100%; height:100%">
        <!-- Key -->
        <el-table-column prop="key" label="Key">
            <template #default="scope">
                <el-input v-model="scope.row.key" placeholder="请输入标签的 Key"></el-input>
            </template>                        
        </el-table-column>
        <!-- Value -->
        <el-table-column prop="value" label="Value">
            <template #default="scope">
                <el-input v-model="scope.row.value" placeholder="请输入标签的 Value"></el-input>
            </template> 
        </el-table-column>
        <!-- Operation -->
        <el-table-column width="200px" align="center">
            <template #header>
                <el-button type="primary" link style="font-size: 16px;margin-bottom: 10px;" @click="emit('addLabel')">添加</el-button>
            </template>
            <template #default="scope">
                <el-button type="danger" link style="font-size: 16px;margin-bottom: 10px;" @click="emit('deleteLabel',scope.$index)">删除</el-button>
            </template>                      
        </el-table-column>
    </el-table> 
</template>    