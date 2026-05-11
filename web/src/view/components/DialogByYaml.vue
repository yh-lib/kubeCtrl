<script setup>
  import YamlEdit from './YamlEdit.vue'

  const props = defineProps({
    dialogVisible: Boolean,
    itemByYaml: String,
  })

  const emit = defineEmits(['closeDialog'])
</script>

<template>
  <el-dialog
    v-model="props.dialogVisible"
    title="配置编辑"
    width="1600px"
    style="height: 840px"
    @close="emit('closeDialog')"
    destroy-on-close
    :close-on-click-modal="false"
  >
    <slot name="header"></slot>
    <yaml-edit :code="props.itemByYaml"></yaml-edit>
  </el-dialog>
</template>
