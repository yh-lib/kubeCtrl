<script setup>
  const props = defineProps(['containerItem'])
  const addTableRow = () => {
    props.containerItem.ports.push({
      name: '',
      containerPort: '',
      protocol: '',
    })
  }
  const deleteTableRow = (index) => {
    props.containerItem.ports.splice(index, 1)
  }
</script>

<template>
  <el-table :data="props.containerItem.ports" style="height: 500px; width: 100%">
    <!-- Key -->
    <el-table-column prop="name" label="端口名称">
      <template #default="scope">
        <el-input v-model="scope.row.name" placeholder="请输入端口名称"></el-input>
      </template>
    </el-table-column>
    <!-- Value -->
    <el-table-column prop="containerPort" label="容器端口">
      <template #default="scope">
        <el-input
          v-model.number="scope.row.containerPort"
          placeholder="请输入标签的 Value"
        ></el-input>
      </template>
    </el-table-column>
    <el-table-column prop="protocol" label="端口协议">
      <template #default="scope">
        <el-select v-model="scope.row.protocol" placeholder="请选择端口协议">
          <el-option v-for="item in ['TCP', 'UDP']" :key="item" :label="item" :value="item" />
        </el-select>
      </template>
    </el-table-column>
    <!-- Operation -->
    <el-table-column width="200px" align="center">
      <template #header>
        <el-button
          type="primary"
          link
          style="font-size: 16px; margin-bottom: 10px"
          @click="addTableRow"
          >添加</el-button
        >
      </template>
      <template #default="scope">
        <el-button
          type="danger"
          link
          style="font-size: 16px; margin-bottom: 10px"
          @click="deleteTableRow(scope.$index)"
          >删除</el-button
        >
      </template>
    </el-table-column>
  </el-table>
</template>
