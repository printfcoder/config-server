# CONFIG-SRV

```text
├── config-srv               
│   ├── README.md            
│   ├── broker               # 消息模块 提供处理或发送异步消息的能力
│   │   └── broker.go         
│   ├── config               # 引导及初始化服务的基础组件
│   │   ├── config.go        # 引导
│   │   ├── db.go            # 引导数据库相关
│   │   └── mysql.go         # 引导Mysql配置
│   ├── docs                 # 项目文档目录
│   │   └── config_db.sql    # 创建表语句
│   ├── facade               # 接口层
│   │   ├── entry.go         # 
│   │   ├── facade.go        # 
│   │   └── source.go        # 
│   ├── loader               # 配置加载器（占位，可能会删除掉），从Mysql等加载
│   │   └── loader.go        # 
│   ├── main.go              # 程序启动入口
│   ├── repository               # 
│   │   ├── README.md               # 
│   │   ├── mysql.go               # 
│   │   └── repository.go               # 
│   └── resources               # 
│       └── application.yml               # 
```