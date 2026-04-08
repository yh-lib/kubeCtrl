
<script setup>
const props = defineProps(['curClusterId','curNodeName','curNsName','curResourceKind','curResourceName'])
</script>

<!-- dialog header -->
 <template>
    <div style="display: flex; justify-content: flex-start;">
        <el-tag type="primary" effect="plain" style="margin-right: 10px;">集群: {{ curClusterId }}</el-tag>
        <el-tag type="primary" effect="plain" style="margin-right: 10px;">命名空间: {{ curNsName }}</el-tag>
        <el-tag type="primary" effect="plain" style="margin-right: 10px;">资源类型: {{ curResourceKind }}</el-tag>
        <el-tag type="primary" effect="plain" style="margin-right: 10px;">资源名称: {{ curResourceName }}</el-tag>
        <el-tag type="primary" effect="plain" style="margin-right: 10px;" v-if="curNodeName">节点: {{ curNodeName }}</el-tag>
    </div>
 </template>
             
        
