<script lang="ts" setup>
import { computed, reactive, ref, onMounted, onBeforeMount  } from 'vue'
import { getUserListHandler as getListHandler,deleteUserHandler as deleteHandler } from '../../api/user.js'
import { User } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import Add from './Add.vue'

// config
const titleName = "用户列表"
interface User {
  username: string
  address: string
  id: number
  qq: string
}


const search = ref('')

const filterTableData = computed(() =>
  data.tableData.filter(
    (data) =>
      !search.value ||
      data.username.toLowerCase().includes(search.value.toLowerCase())
  )
)

const handleEdit = (index: number, row: User) => {
  console.log(index, row)
}
const handleDelete = (index: number, row: User) => {
  console.log(index, row)
}

const data = reactive({
    tableData:[] as User[],
    userForm:{
        username:"",
        qq:"",
        address:"",
        id:""        
    }    
})

// getList
const getUserList = () =>{
    loading.value = true
    getListHandler().then((response)=>{
        if (response.status === 200) {
            data.tableData = response.data.data; // 更新 tableData
            loading.value = false
        } else {
            console.error('获取用户列表失败:', response.msg);
        }
    })
}


// 加载图标
const loading = ref(false)

// 删除用户
const deleteUser = (row) => {
    // 删除提醒
    ElMessageBox.confirm(
        '确认删除用户：' + row.username,
        {
            confirmButtonText: '确认',
            cancelButtonText: '取消',
            type: 'warning',
        }
    )
    .then(() => {
        deleteHandler(row.id).then((response)=>{
            ElMessage({
                type: 'success',
                message: response.data.message,
            })
            // 刷新用户列表
            getUserList()
        })
    })
    .catch(() => {
        return
    }) 
}

// 添加用户
const userDialog = ref(false)

const addUser = () => {
    method.value='create'
    data.userForm={}
    userDialog.value = true
}
// 关闭弹窗时是否刷新用户列表
const closeDialog = () =>{
    method.value == 'create' && getUserList()
}
// 更新用户
const method = ref('')
const updateUser = (row) => {
    // 传递给子组件的操作参数
    method.value='update'
    // 传递当前用户数据给子组件
    data.userForm = row
    // 打开表单弹窗
    userDialog.value = true
}
// 更新用户时刷新用户列表
const updateUserOperation = () =>{
    userDialog.value = false
    getUserList()
}
// 组件加载时自动调用 getUserList 方法
onBeforeMount(() => {
    getUserList();
})
</script>

<template>
    <el-card>
        <!-- 添加用户的表单 -->
        <el-dialog 
            v-model="userDialog"
            :title="method == 'create' ? '添加用户' : '更新用户'"
            width="500px"
            @close="closeDialog"
            destroy-on-close
        >
            <!-- 添加用户的表单组件 -->
            <Add :subMethod='method' :subRow="data.userForm" @refresh="updateUserOperation"/>
        </el-dialog>
        <template #header>
            <div class="card-header">
                <span style="font-size: 24px;">{{ titleName }}</span>
                <el-button type="primary" @click="addUser">添加</el-button>
            </div>
        </template>
        <el-table :data="filterTableData" style="width: 100%"  height="70vh" v-loading="loading">
            <el-table-column label="用户名"  prop="username" />
            <el-table-column label="ID"   prop="id" />
            <el-table-column label="QQ号" prop="qq" />
            <el-table-column label="地址" prop="address" />
            <el-table-column align="right">
            <template #header>
                <el-input v-model="search" size="small" placeholder="Type to search" />
            </template>             
            <template #default="scope">
                <el-button size="small" @click="updateUser(scope.row)">
                    编辑
                </el-button>
                <el-button
                    size="small"    
                    type="danger"
                    @click="deleteUser(scope.row)"
                >
                    删除
                </el-button>
            </template>
            </el-table-column>
        </el-table>      
    </el-card>
</template>

<style scoped>
.card-header{
    display: flex;
    justify-content: space-between;
    align-items: center;

}
</style>