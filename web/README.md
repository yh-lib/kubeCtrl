## 传统部署

1. 编译

```bash
npm run build
```

2. 部署代码

```bash
cp -rp dist /usr/share/nginx/html
```

3. 修改 nginx.conf 配置

```bash
vim nginx.conf
```

4. 启动服务

```bash
systemctl restart nginx
```

# 容器化部署

3. docker build . -t image_tag
4. docker run -it
