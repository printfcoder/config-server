# 配置管理系统

Config-Server 是Go-Micro体系下的配置中心，与Go-Micro仅支持其它K-V与文件存储的配置方式不同，Config-Server提供集中化管理的平台，提供针对不同环境、集群、版本化、回滚等丰富的特性。

Config-Server 基于go-micro/config编写，可以无缝集成到任何Go-Micro服务中。

## 与原生go-micro/config的关系

1. 基于go-micro/config编写
2. 提供关系性数据库存储能力
3. 支持界面管理

## 架构设计

### 服务架构

![](https://github.com/micro-in-cn/docs/blob/master/architecture-design/config-server/design.png)

- Clients 任意服务或客户端，调用[config-srv](./config-srv)获取配置
- Platform-Web
- [config-srv](./config-srv) 配置服务 
- [admin-srv](./admin-srv) 管理层服务
- [platform-web](https://github.com/micro-in-cn/platform-web) UI接入层

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

### 配置服务

### 管理层服务

## CONFIG-SRV

## ADMIN-SRV