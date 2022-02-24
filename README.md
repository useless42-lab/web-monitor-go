# 柠檬监控网页后端

GO版本 1.17


## 项目安装
```
cp .env.example .env
cp plan.json.example plan.json // 采用面包多作为支付系统
```

env 部分配置信息
```
ROUTE_PORT=8080 // 路由端口
NODE=1          // snowflake节点
AUTH_API=       // 自建登录系统配置 // 后续修改
KEY=            // 自建登录系统配置 // 后续修改
MBD_TOKEN=      // 面包多开发者密钥
```

