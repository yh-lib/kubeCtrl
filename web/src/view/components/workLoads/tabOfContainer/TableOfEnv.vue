<script setup>
  // store from pinia
  import { reactive, ref } from 'vue'
  import { useWorkLoadData } from '../../../../store'
  import { storeToRefs } from 'pinia'
  import { getConfigMapListHandler } from '../../../../api/configMap'
  import { getSecretListHandler } from '../../../../api/secret'

  // store from pinia
  const store = useWorkLoadData()
  const { workLoadItem } = storeToRefs(store)

  // data from parent cmt
  const props = defineProps(['containerItem'])

  const data = reactive({
    // 页面状态
    activeName: 'manual',
    dialogVisible: false,
    radioValue: 'configMap',
    // 临时数据
    envFromData: {
      prefix: '',
      configMapRef: {},
      secretRef: {},
    },
    configMapList: [],
    secretList: [],
  })

  const envFromDataDefault = {
    prefix: '',
    configMapRef: {},
    secretRef: {},
  }
  const resetEnvFromData = () => {
    data.envFromData = envFromDataDefault
  }
  // 获取 configMap、secret 列表
  const getConfigMapList = () => {
    getConfigMapListHandler(workLoadItem.value.clusterId, workLoadItem.value.nameSpace).then(
      (res) => {
        if (res.data.status === 200) {
          data.configMapList = res.data.data.items
        }
        console.log(data.configMapList)
      }
    )
  }
  const getSecretList = () => {
    getSecretListHandler(workLoadItem.value.clusterId, workLoadItem.value.nameSpace).then((res) => {
      if (res.data.status === 200) {
        data.secretList = res.data.data.items
      }
      console.log(data.secretList)
    })
  }

  const addEnvFrom = () => {
    data.radioValue == 'configMap'
      ? (data.envFromData.secretRef = {})
      : (data.envFromData.configMapRef = {})

    props.containerItem.envFrom.push(JSON.parse(JSON.stringify(data.envFromData)))
  }

  const isConfigMapRow = (row) => {
    return !!row?.configMapRef?.name
  }

  const getEnvFromName = (row) => {
    return isConfigMapRow(row) ? row.configMapRef.name : row.secretRef?.name
  }

  const updateEnvFromName = (row, value) => {
    if (isConfigMapRow(row)) {
      row.configMapRef.name = value
      return
    }

    row.secretRef.name = value
  }
</script>

<template>
  <el-tabs v-model="data.activeName">
    <el-tab-pane label="手动生成" name="manual">
      <el-table :data="props.containerItem.env" style="width: 100%; height: 440px" border>
        <!-- Key -->
        <el-table-column prop="name" label="变量名">
          <template #default="scope">
            <el-input v-model="scope.row.name" placeholder="请输入变量名"></el-input>
          </template>
        </el-table-column>
        <!-- Value -->
        <el-table-column prop="value" label="变量值">
          <template #default="scope">
            <el-input v-model="scope.row.value" placeholder="请输入变量值"></el-input>
          </template>
        </el-table-column>
        <!-- Operation -->
        <el-table-column width="200px" align="center">
          <template #header>
            <el-button
              type="primary"
              link
              style="font-size: 16px; margin-bottom: 10px"
              @click="props.containerItem.env.push({ name: '', value: '' })"
              >添加</el-button
            >
          </template>
          <template #default="scope">
            <el-button
              type="danger"
              link
              style="font-size: 16px; margin-bottom: 10px"
              @click="props.containerItem.env.splice(scope.$index, 1)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </el-tab-pane>
    <el-tab-pane label="批量生成" name="auto">
      <el-table :data="props.containerItem.envFrom" style="width: 100%; height: 440px" border>
        <el-table-column prop="prefix" label="变量前缀">
          <template #default="scope">
            <el-input v-model="scope.row.prefix" placeholder="请输入变量前缀"></el-input>
          </template>
        </el-table-column>
        <el-table-column prop="type" label="来源类型">
          <template #default="scope">
            {{ isConfigMapRow(scope.row) ? 'configMap' : 'secret' }}
          </template>
        </el-table-column>
        <el-table-column prop="from" label="变量来源">
          <template #default="scope">
            <el-select
              :model-value="getEnvFromName(scope.row)"
              @update:model-value="updateEnvFromName(scope.row, $event)"
              placeholder="请选择变量来源"
              style="width: 240px"
              @visible-change="isConfigMapRow(scope.row) ? getConfigMapList() : getSecretList()"
            >
              <el-option
                v-for="item in isConfigMapRow(scope.row) ? data.configMapList : data.secretList"
                :key="item.metadata.name"
                :label="item.metadata.name"
                :value="item.metadata.name"
              />
            </el-select>
          </template>
        </el-table-column>
        <el-table-column width="200px" align="center">
          <template #header>
            <el-button
              type="primary"
              link
              style="font-size: 16px; margin-bottom: 10px"
              @click="data.dialogVisible = true"
              >添加</el-button
            >
          </template>
          <template #default="scope">
            <el-button
              type="danger"
              link
              style="font-size: 16px; margin-bottom: 10px"
              @click="props.containerItem.envFrom.splice(scope.$index, 1)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>
    </el-tab-pane>
  </el-tabs>
  <!-- 批量生成 dialog -->
  <el-dialog
    v-model="data.dialogVisible"
    title="批量生成环境变量"
    width="600px"
    style="margin-top: 460px; height: 300px"
    destroy-on-close
    :close-on-click-modal="false"
  >
    <!-- 单选按钮 -->
    <el-radio-group
      v-model="data.radioValue"
      style="display: flex; justify-content: left; margin-left: 40px"
      @change="resetEnvFromData"
    >
      <el-radio value="configMap" size="large">configMap</el-radio>
      <el-radio value="secret" size="large">secret</el-radio>
    </el-radio-group>
    <!-- configMap -->
    <div v-if="data.radioValue == 'configMap'" style="display: flex; margin: 20px 40px 0px 40px">
      <el-form :model="data.envFromData" label-width="100px">
        <el-form-item label="变量名前缀">
          <el-input
            v-model="data.envFromData.prefix"
            placeholder="请输入变量名前缀"
            style="width: 400px"
          />
        </el-form-item>
        <el-form-item label="configMap">
          <el-select
            v-model="data.envFromData.configMapRef.name"
            placeholder="请选择configMap"
            @visible-change="getConfigMapList"
          >
            <el-option
              v-for="item in data.configMapList"
              :key="item.metadata.name"
              :label="item.metadata.name"
              :value="item.metadata.name"
            />
          </el-select>
        </el-form-item>
      </el-form>
    </div>

    <!-- secret -->
    <div v-if="data.radioValue == 'secret'" style="display: flex; margin: 20px 40px 0px 40px">
      <el-form :model="data.envFromData" label-width="100px">
        <el-form-item label="变量名前缀">
          <el-input
            v-model="data.envFromData.prefix"
            placeholder="请输入变量名前缀"
            style="width: 400px"
          />
        </el-form-item>
        <el-form-item label="secret">
          <el-select
            v-model="data.envFromData.secretRef.name"
            placeholder="请选择secret"
            @visible-change="getSecretList"
          >
            <el-option
              v-for="item in data.secretList"
              :key="item.metadata.name"
              :label="item.metadata.name"
              :value="item.metadata.name"
            />
          </el-select>
        </el-form-item>
      </el-form>
    </div>
    <el-button type="primary" style="width: 100px; margin-top: 20px" @click="addEnvFrom">
      添加
    </el-button>
  </el-dialog>
</template>
