<script lang="ts" setup>
import { onBeforeMount, onMounted, reactive, ref } from 'vue'
import { addUserHandler,updateUserHandler } from '../../api/user.js'
import { ElMessage } from 'element-plus'
// 定义用户属性
const data = reactive({
  userForm:{
    username:"",
    qq:"",
    address:"",
    id:""    
  }
})
// 重置表单
const userFormRef = ref()
const resetForm = () => {
    userFormRef.value.resetFields()
}
// 添加用户：提交表单
const submitForm = (userForm) => {
    userFormRef.value.validate((valid)=>{
      if (valid) {
        if (props.subMethod == 'create') {
          addUserHandler(userForm).then((Response)=>{
            ElMessage({
              message: response.data.message,
              type: 'success',
            })
          })         
        }else{
          updateUserHandler(userForm).then((Response)=>{
            ElMessage({
              message: response.data.message,
              type: 'success',
            })
            emit('refresh')
          })         
        }
      }else{
          ElMessage({
            message: "请完善表单内容",
            type: 'warning',
          })        
      }
    })
}
// 表单校验
const rules = reactive({
    username: [{ required: true, message: '请输入用户名', trigger: 'blur' },],
    qq: [{ required: true, message: '请输入qq号码', trigger: 'blur' }],
    id: [{ required: true, message: '请输入id', trigger: 'blur' }]
})
// 接受父组件的参数
const subMethod = ref('')
const subRow = reactive({})
const props = defineProps(['subMethod','subRow'])

// 解决单向数据流的问题，否则子组件将可以修改父组件的数据：原理是将浅拷贝修改为深度拷贝
onBeforeMount(()=>{
  // step1: 将父组件传递的对象转换成一个字符串
  //    JSON.stringify(props.subRow)
  // step2: 再将字符串转换成JSON
  data.userForm=JSON.parse(JSON.stringify(props.subRow))
})

// 更新用户后刷新用户列表
const emit = defineEmits(['refresh'])
</script>

<template>
  <el-form
    ref="userFormRef"
    :model="data.userForm"
    label-width="60px"
    center
    class="el-form"
    :rules="rules"
  >
    <!-- 表单 -->
    <el-form-item label="姓名"  prop="username">
      <el-input v-model="data.userForm.username" autocomplete="off" />
    </el-form-item>
    <el-form-item label="ID"  prop="id">
      <el-input v-model="data.userForm.id" autocomplete="off" />
    </el-form-item>
    <el-form-item label="QQ"  prop="qq">
      <el-input v-model="data.userForm.qq" autocomplete="off" />
    </el-form-item>
    <el-form-item label="地址"  prop="address">
      <el-input v-model="data.userForm.address" autocomplete="off" />
    </el-form-item>
    <!-- 按钮 -->
     <div class="button-group">
        <el-button type="primary" @click="submitForm(data.userForm)" > 
          {{ props.subMethod == 'create' ? '提交' : '更新' }}
        </el-button>   
        <el-button @click="resetForm()">
          {{ props.subMethod == 'create' ? '清空' : '重置' }}
        </el-button>
     </div>
  </el-form>
</template>

<style scoped>
.el-form .el-form-item{
    width: 70%;
    margin-top: 10px;
    margin-left: 50px;
}

.button-group {
  display: flex;             /* 启用 Flex 布局 */
  justify-content: center;  /* 让按钮组整体居中 */
  gap: 40px;                 /* 【关键】两个按钮之间的间距 */
  margin-top: 20px;          /* 与上方表单的间距 */
}
</style>