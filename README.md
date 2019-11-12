# 配置管理系统

Config-Server 是Go-Micro体系下的配置中心，与Go-Micro仅支持其它K-V与文件存储的配置方式不同，Config-Server提供集中化管理的平台，提供针对不同环境、集群、版本化、回滚等丰富的特性。

Config-Server 基于go-micro/config编写，可以无缝集成到任何Go-Micro服务中。

## 与原生go-micro/config的关系

1. 基于go-micro/config编写
2. 提供关系性数据库存储能力
3. 支持界面管理

## 借鉴Apollo

Config-Server参考了Apollo在业界成熟的设计方案，详见下方的服务架构与业务架构。感谢Apollo这款优秀的产品，我们在她的设计之上简化了表设计、发布流程，更重要的是面向Go-Micro风格的服务。

## 架构设计

### 服务架构

![](https://github.com/micro-in-cn/docs/blob/master/architecture-design/config-server/design.png)

- Clients 任意服务或客户端，调用[config-srv](./config-srv)获取配置
- [platform-web](https://github.com/micro-in-cn/platform-web) UI接入层，集成配置管理界面，与[admin-srv](./admin-srv)交互
- [config-srv](./config-srv) 配置服务，向微服务提供接口下发配置，不提供改动接口，只向[admin-srv](./admin-srv)提供CRUD接口用于管理配置数据。
- [admin-srv](./admin-srv) 管理层服务

### 业务架构

![](https://github.com/micro-in-cn/docs/blob/master/architecture-design/config-server/business-desgin.png)

- App 应用
  - Env 环境（开发、测试、压测、集成、产线）
    - Cluster 集群 （A区，B区，C区）
      - Namespace 空间 （配置所属域）
        - Item 配置项
          - Key 配置名
          - Value 配置值

## 接口设计

### 格式协议

结构：
 
```text
VERSION1:#{APP}/#{ENV}/#{CLUSTER}/#{Namespace1},#{Namespace2},#{Namespace3}
```

释意：

```text
- VERSION1 版本号，固定，目前仅支持“VERSION1”，在接口统一的情况下，用于解析不同版本的请求文本
  - APP 应用名，一次请求只能填一个
    - Env 应用部署环境，一次请求只能填一个
      - Cluster 应用部署的集群，一次请求只能填一个，Config-Server支持向同一服务不同集群提供服务
        - Namespace 配置所属域（空间），可以传任意多个，返回这些空间下的所有配置
```

### 配置服务

### 管理层服务

## CONFIG-SRV

## ADMIN-SRV

## FAQ

1. 如何解决客户端来自不同集群的注册问题？

答：不同集群使用注册中心可能不一样，导致客户端无法通过Go-Micro原生Selector查询到Config-Server。Config-Server的Source支持使用IP列表直连配置服务获取配置。