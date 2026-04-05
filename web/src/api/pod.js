import { API_CONFIG } from "../config/index.js"
import request from "./axiosEncap.js"

// 获取 pod 列表
export const getPodListHandler = (clusterId, nameSpace) => {
    return request(API_CONFIG.podListApi, { clusterId, nameSpace }, 'get')
}

// 获取 pod 详情
// export const getpodHandler = (clusterId, name) => {
//     return request(API_CONFIG.podGetApi, { clusterId, name }, 'get')
// }

// 删除 pod
export const deletePodHandler = (clusterId, nameSpace, name) => {
    return request(API_CONFIG.podDeleteApi, { clusterId, nameSpace, name }, 'post')
}

// // 添加 pod
// export const createpodHandler = (clusterId, item) => {
//     return request(API_CONFIG.podCreateApi, { clusterId, item }, 'post')
// }


// // 更新 pod
// export const updatepodHandler = (podInfo) => {
//     return request(API_CONFIG.podUpdateApi, podInfo, 'post')
// }