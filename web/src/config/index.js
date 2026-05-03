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
  deploymentUpdateApi: `${BASE_URL}/deployment/update`,
  // statefulSet 管理
  statefulSetListApi: `${BASE_URL}/statefulSet/list`,
  statefulSetGetApi: `${BASE_URL}/statefulSet/get`,
  statefulSetDeleteApi: `${BASE_URL}/statefulSet/delete`,
  statefulSetCreateApi: `${BASE_URL}/statefulSet/create`,
  statefulSetUpdateApi: `${BASE_URL}/statefulSet/update`,
  // daemonSet 管理
  daemonSetListApi: `${BASE_URL}/daemonSet/list`,
  daemonSetGetApi: `${BASE_URL}/daemonSet/get`,
  daemonSetDeleteApi: `${BASE_URL}/daemonSet/delete`,
  daemonSetCreateApi: `${BASE_URL}/daemonSet/create`,
  daemonSetUpdateApi: `${BASE_URL}/daemonSet/update`,
  // secret 管理
  secretListApi: `${BASE_URL}/secret/list`,
  secretGetApi: `${BASE_URL}/secret/get`,
  secretDeleteApi: `${BASE_URL}/secret/delete`,
  // pvc 管理
  pvcListApi: `${BASE_URL}/pvc/list`,
  // configMap 管理
  configMapListApi: `${BASE_URL}/configMap/list`,
  configMapGetApi: `${BASE_URL}/configMap/get`,
  configMapDeleteApi: `${BASE_URL}/configMap/delete`,
  configMapCreateApi: `${BASE_URL}/configMap/create`,
  // service 管理
  serviceListApi: `${BASE_URL}/service/list`,
  serviceGetApi: `${BASE_URL}/service/get`,
  serviceDeleteApi: `${BASE_URL}/service/delete`,
  serviceCreateApi: `${BASE_URL}/service/create`,
  serviceUpdateApi: `${BASE_URL}/service/update`,
}

export const CONFIG = {
  TOKEN_NAME: 'Authorization',
}
