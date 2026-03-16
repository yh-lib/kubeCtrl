<script lang="ts" setup>
import { computed, reactive, ref, onMounted, onBeforeMount  } from 'vue'
import { getClusterListHandler as getListHandler,deleteClusterHandler as deleteHandler } from '../../api/cluster.js'
import { ElMessage, ElMessageBox} from 'element-plus'
import { CircleCheck, CircleClose } from '@element-plus/icons-vue'
import Add from './Add.vue'


// ++++++增

// add_2:打开添加集群的表单
const opDialog = ref(false)
const addItem = () => {
    method.value='create'
    data.itemForm={}
    opDialog.value = true
}
// add_4 定义表单初始数据
const data = reactive({
    tableData:[] as Cluster[],
    itemForm:{
        clusterId:"",
        clusterAlias:"",
        clusterVersion:"",
        clusterLocation:"",
        clusterStatus:"",
        clusterSize:""     
    }    
})
// add_7 监听refresh事件，刷新用户列表
const refreshList = () =>{
    opDialog.value = false
    getList()
}
// add_8 获取当前列表
const getList = () =>{
    loading.value = true
    getListHandler().then((response)=>{
        if (response.status === 200) {
            data.tableData = response.data.data; // 更新 tableData
            loading.value = false
        } else {
            console.error('获取列表失败:', response.msg);
        }
    })
}
// add_9 关闭dialog时刷新用户列表
const closeDialog = () =>{
    method.value == 'create' && getList()
}

// ++++++删

// delete_2 删除集群
const deleteRow = (row) => {
    // 删除提醒
    ElMessageBox.confirm(
        '确认删除集群：' + row.clusterId,
        {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
    .then(() => {
        deleteHandler(row.clusterId).then((response)=>{
            ElMessage({
                type: 'success',
                message: response.data.msg,
            })
            // 刷新列表
            getList()
        })
    })
    .catch(() => {
        return
    }) 
}

// ++++++改

// update_2 更新
const method = ref('')
const updateItem = (row) => {
    // 传递给操作参数给子组件
    method.value='update'
    // 传递当前用户数据给子组件
    data.itemForm = row
    // 打开表单弹窗
    opDialog.value = true
}
// update_3 将method.value、data.itemForm赋值给子组件
// update_4 在子组件中将值渲染到form表单中

// ++++++查

// select_1 子组件加载时自动获取列表
onBeforeMount(() => {
    getList();
})
// select_2 搜索功能
const search = ref('')
const filterTableData = computed(() =>
  data.tableData.filter(
    (data) =>
      !search.value ||
      data.clusterId.toLowerCase().includes(search.value.toLowerCase())
  )
)


// ++++++独立配置
// Main 标题
const titleName = "集群列表"
// Table 表头
const tableTtile = {
    f1: { prop: "clusterId", label: "ID" },
    f2: { prop: "clusterAlias", label: "别名" },
    f3: { prop: "clusterVersion", label: "版本" },
    f4: { prop: "clusterLocation", label: "位置" },
    f5: { prop: "clusterStatus", label: "状态" },
    f6: { prop: "clusterSize", label: "规模" }
}
// 集群信息接口
interface Cluster {
  clusterId: string
  clusterAlias: string
  clusterVersion: string
  clusterLocation: string
  clusterStatus: string
  clusterSize: number
}
// 加载图标默认关闭
const loading = ref(false)




</script>

<template>
    <el-card>
        <!-- add_1:添加按钮 -->
        <template #header>
            <div class="card-header">
                <span style="font-size: 24px;">{{ titleName }}</span>
                <el-button type="primary" @click="addItem">添加</el-button>
            </div>
        </template>
        <!-- none -->
        <el-table :data="filterTableData" style="width: 100%"  height="70vh" v-loading="loading">
            <el-table-column :label="tableTtile.f1.label"  :prop="tableTtile.f1.prop" />
            <el-table-column :label="tableTtile.f2.label"   :prop="tableTtile.f2.prop" />
            <el-table-column :label="tableTtile.f3.label" :prop="tableTtile.f3.prop" />
            <el-table-column :label="tableTtile.f4.label" :prop="tableTtile.f4.prop" />
            <el-table-column :label="tableTtile.f5.label" prop="" >
                <template #default="scope">
                    <el-icon v-if="scope.row.clusterStatus == 'true'" color="green"><CircleCheck /></el-icon>  
                    <el-icon v-else color="red"><CircleClose /></el-icon>  
                </template>
            </el-table-column>
            <el-table-column :label="tableTtile.f6.label" :prop="tableTtile.f6.prop" />
            <el-table-column align="right">
                <template #header>
                    <el-input v-model="search" size="small" placeholder="Search by clusterId" />
                </template>             
                <template #default="scope">
                    <!-- update_1 触发编辑操作 -->
                    <el-button :disabled="scope.row.clusterStatus != 'true'" size="small" @click="updateItem(scope.row)">
                        编辑
                    </el-button>
                    <!-- delete_1 触发删除动作 -->
                    <el-button
                        size="small"    
                        type="danger"
                        @click="deleteRow(scope.row)"
                    >
                        删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>  
        <!-- add_3:动态渲染表单抬头/add_5:打开表单-->
        <el-dialog 
            v-model="opDialog"
            :title="method == 'create' ? '添加集群' : '更新集群'"
            width="50vw"
            @close="closeDialog"
            destroy-on-close
        >
            <!-- add_6:添加集群表单 -->
            <!-- update3 将method.value、data.itemForm赋值给子组件 -->
            <Add :subMethod='method' :subRow="data.itemForm" @refresh="refreshList"/>
        </el-dialog>            
    </el-card>
</template>

<style scoped>
.card-header{
    display: flex;
    justify-content: space-between;
    align-items: center;

}
</style>