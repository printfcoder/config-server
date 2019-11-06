# ADMIN

Admin 是管理配置中心的接口服务，它负责与Platform-Web协作，完成配置的增删改查，配置集群环境的管理，以及对人员权限的管理控制。

它基于go-micro/web编写。与srv进行RPC调用，完成配置的修改。