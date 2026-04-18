import { API_CONFIG } from "../config/index.js"
import request from "./axiosEncap.js"

// 获取 pvc 列表
export const getPvcListHandler = (clusterId, nameSpace) => {
    return request(API_CONFIG.pvcListApi, { clusterId, nameSpace }, 'get')
}

// // 获取 secret 详情
// export const getSecretHandler = (clusterId, nameSpace, name) => {
//     return request(API_CONFIG.secretGetApi, { clusterId, nameSpace, name }, 'get')
// }

// // 删除 secret
// export const deleteSecretHandler = (clusterId, nameSpace, name) => {
//     return request(API_CONFIG.secretDeleteApi, { clusterId, nameSpace, name }, 'post')
// }

// // 添加 secret
// export const createsecretHandler = (clusterId, item) => {
//     return request(API_CONFIG.secretCreateApi, { clusterId, item }, 'post')
// }


// // 更新 secret
// export const updatesecretHandler = (secretInfo) => {
//     return request(API_CONFIG.secretUpdateApi, secretInfo, 'post')
// }