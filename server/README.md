## 传统部署

1. 拉取代码
```bash
git clone https://github.com/yh-lib/Ks.git
```

2. 编译
```bash
# 进入后端目录
cmd
cd server
```
```bash
# Windows 运行环境
set GOOS=windows
# 构建
go build -o ks-backend-windows-x86.exe main.go
```
```bash
# Linux 运行环境
set GOOS=linux
# 构建
go build -o ks-backend-linux-x86 main.go
```

3. 测试
```bash
# Windows 运行环境
setx KUBECONFIGPATH YourKubeConfig # 指定存放集群元数据的集群kubeconfig(持久化变量)
set KUBECONFIGPATH=YourKubeConfig # 当前会话变量
 .\ks-backend.exe # 启动后端服务
echo %USERNAME%   # 获取用户名
echo %PASSWORD%   # 获取密码  默认为Admin123
# 接口调用测试
curl -X POST ^
-H "Content-Type: application/json" ^
-d "{\"username\":\"Admin\",\"password\":\"Admin123\"}" ^
"127.0.0.1:8080/api/auth/login"
# 接口调用成功时的数据返回
{"status":200,"message":"登录成功","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IllIIiwiaXNzIjoiWUgiLCJzdWIiOiJZSCIsImV4cCI6MjQ5ODQwMDg0OSwibmJmIjoxNzc4NDAwODQ5LCJpYXQiOjE3Nzg0MDA4NDl9.BkXm8iapZ32co3lff5NbKkGhA9tSR_9a5QnfBm9fcYU"}}
```
```bash
# Linux 运行环境
echo 'export KUBECONFIGPATH="YourKubeConfig"' >> /etc/profile # 配置 kubeconfig 路径
echo 'export USERNAME="YourUserName"' >> /etc/profile  # 配置后端服务用户名
echo 'export PASSWORD="YourPassword"' >> /etc/profile  # 配置后端服务密码
source /etc/profile
chmod +x ks-backend
# 接口调用测试
curl -X POST \
-H "Content-Type: application/json" \
-d "{\"username\":\"Admin\",\"password\":\"Admin123\"}" \
"127.0.0.1:8080/api/auth/login"
# 接口调用成功时的数据返回
{"status":200,"message":"登录成功","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IllIIiwiaXNzIjoiWUgiLCJzdWIiOiJZSCIsImV4cCI6MjQ5ODQwMDg0OSwibmJmIjoxNzc4NDAwODQ5LCJpYXQiOjE3Nzg0MDA4NDl9.BkXm8iapZ32co3lff5NbKkGhA9tSR_9a5QnfBm9fcYU"}}
```
## 容器化部署

1. 准备主控集群 kubeconfig
```bash
# 提供你的kubeconfig
cp  YourKubeconfig server/config/baseKubeConfig
# 根据自己的需求 配置dockerfile
vim ./dockerfile
```
2. 制作镜像
```bash
docker build -t crpi-o5e9g8vha41iltg8.cn-hangzhou.personal.cr.aliyuncs.com/ks_required/ks-backend:v1.0 .
```
3. 推送至镜像仓库
```bash
docker push crpi-o5e9g8vha41iltg8.cn-hangzhou.personal.cr.aliyuncs.com/ks_required/ks-backend:v1.0
```
4. 启动服务
```bash
docker run \
--name ks-backend \
-e  "USERNAME=Admin" \      # 配置登录用户名
-e  "PASSWORD=Admin123" \   # 配置登录密码
-p  "10001:8080"        \   # 端口映射
crpi-o5e9g8vha41iltg8.cn-hangzhou.personal.cr.aliyuncs.com/ks_required/ks-backend:v1.0  # 构建的镜像
```
5. 测试
```bash
# 接口调用测试
curl -X POST ^
-H "Content-Type: application/json" ^
-d "{\"username\":\"Admin\",\"password\":\"Admin123\"}" ^
"127.0.0.1:10001/api/auth/login"
# 接口调用成功时的数据返回
curl -X POST ^
-H "Content-Type: application/json" ^
-d "{\"username\":\"Admin\",\"password\":\"Admin123\"}" ^
"127.0.0.1:10001/api/auth/login"
```


