<script setup>
import { ref } from 'vue';
// 导入 codemirror 编辑器主体
import { Codemirror } from 'vue-codemirror';
// 导入编辑器支持的语言：语法高亮等功能
import { yaml } from '@codemirror/lang-yaml'
const extensions = [yaml()]
// 定义存放代码的变量
const code = ref('')
// 
</script>

<template>
    <Codemirror 
        v-model="code"
        placeholder="在此输入yaml数据..."
        :style="{ height: '600px', border: '1px solid #dcdfe6', textAlign:'left'}"
        :autofocus="true"
        :indent-with-tab="true"
        :tab-size="2"
        :extensions="extensions" 
        @ready="console.log('yaml 编辑器初始化已完成')"
        @change="console.log('change', '编辑器内容发生变化')"
        @focus="console.log('focus', '编辑器获得焦点')"
        @blur="console.log('blur', '编辑器失去焦点')"
    />
</template>