import { API_CONFIG } from "../config/index.js"
import request from "./axiosEncap.js"

// 获取集群列表
export const getClusterListHandler = () => {
    return request(API_CONFIG.clusterListApi, {}, 'get', 10000)
}

// 删除集群
export const deleteClusterHandler = (id) => {
    return request(API_CONFIG.clusterDeleteApi, {id}, 'post', 10000)
}

// 添加集群
export const addClusterHandler = (clusterInfo) => {
    return request(API_CONFIG.clusterAddApi, clusterInfo, 'post', 10000)
}


// 更新集群
export const updateClusterHandler = (clusterInfo) => {
    return request(API_CONFIG.clusterUpdateApi, clusterInfo, 'post', 10000)
}