import { API_CONFIG } from "../config/index.js"
import request from "./axiosEncap.js"

// 获取 deployment 列表
export const getdeploymentListHandler = (clusterId, nameSpace) => {
    return request(API_CONFIG.deploymentListApi, { clusterId, nameSpace }, 'get')
}

// 获取 deployment 详情
export const getdeploymentHandler = (clusterId, nameSpace, name) => {
    return request(API_CONFIG.deploymentGetApi, { clusterId, nameSpace, name }, 'get')
}

// 删除 deployment
export const deletedeploymentHandler = (clusterId, nameSpace, name) => {
    return request(API_CONFIG.deploymentDeleteApi, { clusterId, nameSpace, name }, 'post')
}

// 创建 deployment
export const createdeploymentHandler = (itemForm) => {
    return request(API_CONFIG.deploymentCreateApi, itemForm, 'post')
}


// // 更新 deployment
// export const updatedeploymentHandler = (deploymentInfo) => {
//     return request(API_CONFIG.deploymentUpdateApi, deploymentInfo, 'post')
// }