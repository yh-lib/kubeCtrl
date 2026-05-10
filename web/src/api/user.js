import { API_CONFIG } from '../config/index.js'
import request from './axiosEncap.js'

// 获取用户列表
export const getUserListHandler = () => {
  return request(API_CONFIG.userListApi, {}, 'get')
}

// 删除用户
export const deleteUserHandler = (id) => {
  return request(API_CONFIG.userDeleteApi, { id }, 'post')
}

// 添加用户
export const addUserHandler = (userForm) => {
  return request(API_CONFIG.userAddApi, userForm, 'post')
}

// 更新用户
export const updateUserHandler = (userForm) => {
  return request(API_CONFIG.userUpdateApi, userForm, 'post')
}
