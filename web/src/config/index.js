// 放置项目的配置信息

// 后端接口配置
const BASE_URL = import.meta.env.VITE_BASE_URL
export const API_CONFIG = {
    // 登录管理
    loginApi: `${BASE_URL}/auth/login`,
    logoutApi: `${BASE_URL}/auth/logout`,
    // 用户管理
    userListApi: `${BASE_URL}/user/list`,
    userDeleteApi: `${BASE_URL}/user/delete`,
    userAddApi: `${BASE_URL}/user/add`,
    userUpdateApi: `${BASE_URL}/user/update`,
    // 集群管理
    clusterListApi: `${BASE_URL}/cluster/list`,
    clusterDeleteApi: `${BASE_URL}/cluster/delete`,
    clusterAddApi: `${BASE_URL}/cluster/add`,
    clusterUpdateApi: `${BASE_URL}/cluster/update`,
    clusterGetApi: `${BASE_URL}/cluster/get`,
    // 节点管理
    nodeListApi: `${BASE_URL}/node/list`,
    nodeGetApi: `${BASE_URL}/node/get`,
    nodeUpdateApi: `${BASE_URL}/node/update`,
    // 命名空间管理
    namespaceCreateApi: `${BASE_URL}/namespace/create`,
    namespaceDeleteApi: `${BASE_URL}/namespace/delete`,
    namespaceUpdateApi: `${BASE_URL}/namespace/update`,
    namespaceListApi: `${BASE_URL}/namespace/list`,
    namespaceGetApi: `${BASE_URL}/namespace/get`,
    // pod 管理
    podListApi: `${BASE_URL}/pod/list`,
    podGetApi: `${BASE_URL}/pod/get`,
    podDeleteApi: `${BASE_URL}/pod/delete`,
    // deployment 管理
    deploymentListApi: `${BASE_URL}/deployment/list`,
    deploymentGetApi: `${BASE_URL}/deployment/get`,
    deploymentDeleteApi: `${BASE_URL}/deployment/delete`,
    deploymentCreateApi: `${BASE_URL}/deployment/create`,
    // secret 管理
    secretListApi: `${BASE_URL}/secret/list`,
    secretGetApi: `${BASE_URL}/secret/get`,
    secretDeleteApi: `${BASE_URL}/secret/delete`,
}

export const CONFIG = {
    TOKEN_NAME: "Authorization"
}