### 目录结构：

web     前端代码
web/README.md   前端部署文档

server  后端代码
server/README.md   后端部署文档

chart   ks项目chart包文件

### 兼容性

kubernetes 1.23
kubernetes 1.35


### 一键安装

```bash
# 根据需求  编辑 values.yaml
vim chart/values.yaml
# 安装 release
helm install ks -n ks .
```

### 依赖关系

1. 前端服务依赖后端服务（否则无法正常登录验证）
2. 后端服务依赖k8s集群状态（否则无法获取METADATA_NAMSPACE中集群的kubeconfig，后端服务将无法启动）