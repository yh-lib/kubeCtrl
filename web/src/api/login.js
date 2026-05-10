import { API_CONFIG, CONFIG } from '../config/index.js'
import router from '../router/index.js'
import request from './axiosEncap.js'
import { ElMessage, ElMessageBox } from 'element-plus'

export const login = (username, password) => {
  request(API_CONFIG.loginApi, { username, password }, 'post').then((response) => {
    console.log('response:', response)
    if (response.data.status === 200) {
      // 将后端返回的token保存到本地
      const token = response.data.data.token
      window.localStorage.setItem(CONFIG.TOKEN_NAME, token)
      // // 提示登录成功
      ElMessage({
        message: response.data.message,
        type: 'success',
      })
      router.replace('/cluster/dashboard')
    }
  })
}

export const logout = () => {
  // 退出确认提示框
  ElMessageBox.confirm('确定要离开了吗？', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    type: 'warning',
  })
    .then(() => {
      // 退出逻辑
      request(API_CONFIG.logoutApi, {}, 'get').then((response) => {
        if (response.data.status === 200) {
          // 后端已经完成退出逻辑
          // 移除本地token
          window.localStorage.removeItem(CONFIG.TOKEN_NAME)
          // 提示退出成功
          ElMessage({
            message: response.data.message,
            type: 'success',
          })
          router.replace('/login')
        } else {
          ElMessage.error('后端退出逻辑未完成')
        }
      })
    })
    .catch(() => {
      return
    })
}
