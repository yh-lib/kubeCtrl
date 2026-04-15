<script setup>
const props = defineProps({
    tolerationsList:{
        type: Array,
        default: []
    },
    tolerationsOperatorSelectOptions:{
        type: Array,
        default: []
    },
    tolerationsEffectSelectOptions:{
        type: Array,
        default: []
    }
})
const emit = defineEmits([
    'addToleration',
    'deleteToleration'
])
</script>

<template>
    <el-table :data="props.tolerationsList" style="width: 100%; height:100%">
        <!-- Key -->
        <el-table-column prop="key" label="Key">
            <template #default="scope">
                <el-input v-model="scope.row.key" placeholder="请输入容忍的Key"></el-input>
            </template>                        
        </el-table-column>
        <!-- Operator -->
        <el-table-column prop="operator" label="Operator">
            <template #default="scope">
            <el-select v-model="scope.row.operator" placeholder="请选择容忍操作" style="width: 240px">
                <el-option
                v-for="item in props.tolerationsOperatorSelectOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                />
                </el-select>
            </template>                        
        </el-table-column>
        <!-- Value -->
        <el-table-column prop="value" label="Value">
            <template #default="scope">
                <el-input v-model="scope.row.value" placeholder="请输入容忍的Value" v-if="scope.row.operator!='Exists'"></el-input>
            </template> 
        </el-table-column>
        <!-- Effect -->
        <el-table-column prop="effect" label="Effect">
            <template #default="scope">
            <el-select v-model="scope.row.effect" placeholder="请选择容忍的影响" style="width: 240px">
                <el-option
                v-for="item in props.tolerationsEffectSelectOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                />
                </el-select>
            </template> 
        </el-table-column>
        <!-- Operation -->
        <el-table-column width="200px" align="center">
            <template #header>
                <el-button type="primary" link style="font-size: 16px;margin-bottom: 10px;" @click="emit('addToleration')">添加</el-button>
            </template>
            <template #default="scope">
                <el-button type="danger" link style="font-size: 16px;margin-bottom: 10px;" @click="emit('deleteToleration',scope.$index)">删除</el-button>
            </template>                      
        </el-table-column>
    </el-table>
</template>    