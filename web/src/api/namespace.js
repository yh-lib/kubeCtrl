import { API_CONFIG } from "../config/index.js"
import request from "./axiosEncap.js"

// 获取集群列表
export const getNamespaceListHandler = (clusterId) => {
    return request(API_CONFIG.namespaceListApi, { clusterId: clusterId }, 'get', 10000)
}

// 获取集群详情
export const getNamespaceHandler = (clusterId, name) => {
    return request(API_CONFIG.namespaceGetApi, { clusterId, name }, 'get', 10000)
}

// 删除集群
export const deleteNamespaceHandler = (clusterId, name) => {
    return request(API_CONFIG.namespaceDeleteApi, { clusterId, name }, 'get', 10000)
}

// 添加集群
export const createNamespaceHandler = (clusterId, item) => {
    return request(API_CONFIG.namespaceCreateApi, { clusterId, item }, 'post', 10000)
}


// 更新集群
export const updateNamespaceHandler = (NamespaceInfo) => {
    return request(API_CONFIG.namespaceUpdateApi, NamespaceInfo, 'post', 10000)
}