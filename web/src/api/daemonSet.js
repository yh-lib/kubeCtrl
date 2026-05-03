import { API_CONFIG } from '../config/index.js'
import request from './axiosEncap.js'

// 获取 daemonSet 列表
export const getDaemonSetListHandler = (clusterId, nameSpace) => {
  return request(API_CONFIG.daemonSetListApi, { clusterId, nameSpace }, 'get')
}

// 获取 daemonSet 详情
export const getDaemonSetHandler = (clusterId, nameSpace, name) => {
  return request(API_CONFIG.daemonSetGetApi, { clusterId, nameSpace, name }, 'get')
}

// 删除 daemonSet
export const deleteDaemonSetHandler = (clusterId, nameSpace, name) => {
  return request(API_CONFIG.daemonSetDeleteApi, { clusterId, nameSpace, name }, 'post')
}

// 创建 daemonSet
export const createDaemonSetHandler = (itemForm) => {
  return request(API_CONFIG.daemonSetCreateApi, itemForm, 'post')
}

// 更新 daemonSet
export const updateDaemonSetHandler = (itemFrom) => {
  return request(API_CONFIG.daemonSetUpdateApi, itemFrom, 'post')
}
