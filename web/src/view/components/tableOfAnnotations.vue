<script setup>
const props = defineProps({
    annotationsList:[],
})
const emit = defineEmits([
    'addAnnotation',
    'deleteAnnotation'
])
</script>

<template>
    <el-table :data="props.annotationsList" style="width: 100%; height:100%">
        <el-table-column prop="key" label="Key">
            <template #default="scope">
                <el-input v-model="scope.row.key" placeholder="请输入注释的 Key"></el-input>
            </template>                        
        </el-table-column>
        <el-table-column prop="value" label="Value">
            <template #default="scope">
                <el-input v-model="scope.row.value" placeholder="请输入注释的 Value"></el-input>
            </template> 
        </el-table-column>
        <el-table-column width="200px" align="center">
            <template #header>
                <el-button type="primary" link style="font-size: 16px;margin-bottom: 10px;" @click="emit('addAnnotation')">添加</el-button>
            </template>
            <template #default="scope">
                <el-button type="danger" link style="font-size: 16px;margin-bottom: 10px;" @click="emit('deleteAnnotation',scope.$index) ">删除</el-button>
            </template>                      
        </el-table-column>
    </el-table>
</template>    