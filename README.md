# 配置管理系统

Config-Server 是Go-Micro体系下的配置中心，与Go-Micro仅支持其它K-V与文件存储的配置方式不同，Config-Server提供集中化管理的平台，提供针对不同环境、集群、版本化、回滚等丰富的特性。

Config-Server 基于go-micro/config编写，可以无缝集成到任何Go-Micro服务中。

## 与原生go-micro/config的关系

1. 基于go-micro/config编写
2. 提供关系性数据库存储能力
3. 支持界面管理

## 架构设计

![](https://github.com/micro-in-cn/docs/blob/master/architecture-design/config-server/design.png)

## SRV

## ADMIN

