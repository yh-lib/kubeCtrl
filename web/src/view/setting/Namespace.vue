<script lang="ts" setup>
import { computed, reactive, ref, onBeforeMount, effect  } from 'vue'
import { getNamespaceListHandler as getListHandler,getNamespaceHandler as getHandler,deleteNamespaceHandler as deleteHandler,createNamespaceHandler as createHandler } from '../../api/namespace.js'
import { getClusterListHandler} from '../../api/cluster.js'
import { ElSelect,ElMessage,ElMessageBox } from 'element-plus'
import { list2obj, obj2list } from '../../utils/typeConv/type.conv.js'
import request from '../../api/axiosEncap.js'
import { API_CONFIG } from '../../config/index.js'

// 需要的数据变量
const data = reactive({
    items: [],  // 后端返回的 list
    item: {
        "kind": "Namespace",
        "apiVersion": "v1",
        "metadata": {
            "name":"",
        }     
    }, // 后端返回的 obj
    search: "", // 接收 header 搜索框中数据
    curHostName: "",
    // 集群选择器
    clusterOptions:[],  // 集群选择 options
    curClusterId: "",   // 当前选择的集群id
    defaultClusterId: "in-cluster", // 默认选择的集群id
    // 表格
    opDialog: false,    // 配置、主机名 对话框
    createDialog: false,    // 创建namespace
    nodeLabels: [],     // 后端返回的label
    nodeTaints: [],
    curTabName: "",
    // axios
    requestItem: {
        clusterId:"",
        nameSpace:"",
        name:"",
        item:{},
    },
    // 命名空间
    newNsName: "",
    namespaceItem:{
        "clusterId":"",
        "item":{
            "kind": "Namespace",
            "apiVersion": "v1",
            "metadata": {
                "name": ""            
             }
        }
    }
})

// // // 子组件加载前自动获取数据
onBeforeMount(async () => {
    await getclusterOptions()   // 获取集群列表
    data.curClusterId=data.defaultClusterId // 获取基础设施集群的Node列表
    getList();
})

// 创建namespace
const createItemConfirm =() =>{
    data.item.metadata.name = data.newNsName  
    createHandler(data.curClusterId,data.item).then((res)=>{
        if (res.data.status === 200) {
            ElMessage({
                message: res.data.message,
                type: 'success',
            })
            data.createDialog = false
            getList();  
        }
})
}

// 创建namespace
const createItem = () =>{
    data.createDialog = "true"
}

// 删除namespace
const deleteItem = (name) => {
    console.log("删除函数")
    ElMessageBox.confirm(
        '确认删除命名空间：' + name,
        {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
    .then(()=>{
        deleteHandler(data.curClusterId,name).then((res)=>{
                if (res.data.status === 200) {
                    ElMessage({
                        message: res.data.message,
                        type: 'success',
                    })
                    getList();    
                }     
        })                  
    })
    .catch(() => {
        return
    })            
}

// 更新按钮实现逻辑
const updateItem = () => {
    if (data.curTabName == 'nodeLabel'){
        // 构建后端所需数据
        data.requestItem.clusterId = data.curClusterId
        data.requestItem.nameSpace = ""
        data.requestItem.name = data.curHostName
        data.requestItem.item = data.item
        data.requestItem.item.metadata.labels = list2obj(data.nodeLabels)
        // 调用后端API
        request(API_CONFIG.nodeUpdateApi,data.requestItem,'post').then((res)=>{
            if (res.data.status === 200) {
                ElMessage({
                    message: res.data.message,
                    type: 'success',
                })
                data.opDialog = false
            } else {
                ElMessage.error({message: res.data.message})

            }

        })
    }
    else if (data.curTabName == 'nodeTaint') {
        // 构建后端所需数据
        data.requestItem.clusterId = data.curClusterId
        data.requestItem.nameSpace = ""
        data.requestItem.name = data.curHostName
        data.requestItem.item = data.item
        data.requestItem.item.spec.taints = data.nodeTaints
        // 调用后端API
        request(API_CONFIG.nodeUpdateApi,data.requestItem,'post').then((res)=>{
            if (res.data.status === 200) {
                ElMessage({
                    message: res.data.message,
                    type: 'success',
                })
                data.opDialog = false
            } else {
                ElMessage.error({message: res.data.message})

            }

        })
    }
}

// 获取节点角色（纯 JS 版本）
const getNodeRole = (labels = {}) => {
    const keyList = Object.keys(labels || {})
    if (!keyList.length) {
        return 'Worker'
    }
    const controlPlaneKeys = [
        'node-role.kubernetes.io/control-plane',
        'node-role.kubernetes.io/controlplane',
        'node-role.kubernetes.io/master',
    ]
    return controlPlaneKeys.some((key) => keyList.includes(key)) ? 'Master' : 'Worker'
}

// 添加列表项
const addLabelItem = () => {data.nodeLabels.unshift({key:"",value:""})}
// 添加污点项（为空时先初始化数组）
const addTaintItem = () => {
    if (!Array.isArray(data.nodeTaints)) {
        data.nodeTaints = []
    }
    data.nodeTaints.unshift({ key: "", value: "", effect: "" })
}

// 删除列表项
const deleteLabelItem = (index) => {
    data.nodeLabels.splice(index,1)
}
// 删除污点项
const deleteTaintItem = (index) => {
    data.nodeTaints.splice(index,1)
}
    
// 获取节点标签
const getLabel = () => {
    // console.log("获取节点调试数据:",data.item.spec.taints)
    // 获取标签
    data.nodeLabels = obj2list(data.item.metadata.labels)
}
// 关闭dialog时刷新用户列表
const closeDialog = () =>{
    getList()
}

// 内存地址单位转换
const memoryKi2MB = (memory) => {
    if (memory === null || memory === undefined) return ''
    const match = String(memory).trim().match(/^(\d+(?:\.\d+)?)Ki$/i)
    return match ? Math.floor(Number(match[1]) / 1000) : ''
}

// 搜索后的 namespace 列表数据
const filterTableData = computed(() =>
  data.items.filter(
    (item) =>
      !data.search ||
      item.metadata.name.toLowerCase().includes(data.search.toLowerCase())
  )
)

// 获取当前集群Namespace列表
const getList = () =>{
    getListHandler(data.curClusterId).then((res)=>{
        data.items = res.data.data.items;
    })
}

// 获取集群列表
const getclusterOptions = async ()=>{
    await getClusterListHandler().then((res)=>{
        if (res.data.status === 200) {
            data.clusterOptions = res.data.data.items;
            console.log('获取集群列表:::', data.clusterOptions);
        } else {
            console.error('获取集群列表失败:', res.data.message);
        }
    })
}
</script>

<template>
    <!-- card -->
    <el-card>
        <!-- card header -->
        <template #header>
            <div class="card-header">
                <div>
                    <span style="font-size: 24px;">命名空间列表</span>
                </div>
                <div>
                    <el-select v-model="data.curClusterId" placeholder="选择集群" style="width: 240px;margin-right: 10px;" @change="getList(data.curClusterId)">
                        <el-option
                            v-for="item in data.clusterOptions"
                            :key="item.clusterId"
                            :label="item.clusterId"
                            :value="item.clusterId"
                            :disabled="item.disabled"
                        />     
                    </el-select>               
                    <el-input v-model="data.search" style="width: 240px" placeholder="搜索命名空间"/>    
                    <el-button type="primary" @click="createItem()" style="width: 105px;margin-left: 10px;">创建</el-button> 
                </div>
            </div>
        </template>
        <!-- card middle -->
        <el-table :data="filterTableData" style="width: 100%;"  height="70vh">
            <el-table-column label="名称" prop="metadata.name"/>            
            <el-table-column label="创建时间" prop="metadata.creationTimestamp"/>
            <el-table-column label="状态" prop="status.phase"/>
            <el-table-column label="操作" prop="operation" >     
                <template #default="scope">
                    <el-button link type="danger" @click="deleteItem(scope.row.metadata.name)">删除</el-button>
                </template>                                             
            </el-table-column>
        </el-table> 
        <!-- 创建namespace -->
        <el-dialog 
            v-model="data.createDialog"
            width="580px"
            style="height: 210px;"
            @close="closeDialog"
            destroy-on-close
        >
            <!-- dialog header -->
            <template #header>                
                <span style="font-size: 18px;">创建命名空间</span>
            </template>
            <!-- dialog middle -->
             <div>
                <el-input v-model="data.newNsName" placeholder="请输入命名空间名称" style="width: 400px; margin-top: 10px;" size="large"/> 
             </div>
            <div>
                <el-button type="primary" @click="createItemConfirm" style="width: 105px; margin-top: 30px;">确认</el-button>
            </div>
            
        </el-dialog>               
    </el-card>
           
</template>

<style lang="less" scoped>
.card-header{
    display: flex;
    justify-content: space-between;
    align-items: center;
}

:deep(.no-border-input .el-input__wrapper){
    box-shadow: none;
    border: none;
}
</style>
