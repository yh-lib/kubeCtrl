import { API_CONFIG } from "../config/index.js"
import request from "./axiosEncap.js"

// 获取节点列表
export const getNodeListHandler = () => {
    return request(API_CONFIG.nodeListApi, {}, 'get', 10000)
}