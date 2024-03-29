# 配置管理系统

Config-Server 是Go-Micro体系下的配置中心，与Go-Micro仅支持其它K-V与文件存储的配置方式不同，Config-Server提供集中化管理的平台，提供针对不同集群、版本化、回滚等丰富的特性。

Config-Server 基于go-micro/config编写，可以无缝集成到任何Go-Micro服务中。

## 与原生go-micro/config的关系

1. 基于go-micro/config编写
2. 提供关系性数据库存储能力
3. 支持界面管理

## 借鉴Apollo

Config-Server参考了Apollo在业界成熟的设计方案，详见下方的服务架构与业务架构。感谢Apollo这款优秀的产品，我们在她的设计之上简化了表设计、部署流程，更重要的是面向Go-Micro风格的服务。打包即运行，不会有像Apollo那样的部署烦恼。

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
  - Cluster 集群 （A区，B区，C区）
    - Namespace 空间 （配置所属域）
      - Item 配置项
        - Key 配置名
        - Value 配置值

## 接口设计

### 格式协议

1. 结构：
 
```text
VERSION1:#{APP}/#{CLUSTER}/#{Namespace1}
```

2. 释意：

```text
- VERSION1 版本号，固定，目前仅支持“VERSION1”，在接口统一的情况下，用于解析不同版本的请求文本
  - APP 应用名，一次请求只能填一个
    - Cluster 应用部署的集群，一次请求只能填一个，Config-Server支持向同一服务不同集群提供服务
      - Namespace 配置所属域（空间），目前只能传一个，返回这些空间下的所有配置。
```

> Namespace 只支持传一个的原因是如果是多个，而改动往往只会有一个，此时不得不把所有未改动的Namespace一起返回，数据量会非常庞大。

3. 存储

表结构大致如下

id | namespace_id | key | value 
--- | --- | --- | --- 
1 | 2 | micro.book.name | BOOK-NAME 
2 | 2 | micro.book.price | 123 

有几个需要注意的地方：

**-- 强定义的结构**

为了方便当key-value转成其它格式如JSON、Properties、YAML等标准，不允许以下不规范的冲突出现：

id | namespace_id | key | value 
--- | --- | --- | --- 
1 | 2 | micro.book.name | BOOK-NAME 
2 | 2 | micro.book.price | 123 
2 | 2 | **micro.book** | 一本好书

**micro.book**有子key（name、price），它不能再被当成一个key，比如我们需要转换成JSON：

```json
{
  "micro": {
    "book": {
      "name": "BOOK-NAME",
      "price": "123"
    }
  }
}
```

格式非法时，无法转换。

### 配置服务

### 管理层服务

## CONFIG-SRV

## ADMIN-SRV

## FAQ

1. 如何解决客户端来自不同集群的注册问题？

答：不同集群使用注册中心可能不一样，导致客户端无法通过Go-Micro原生Selector查询到Config-Server。Config-Server的Source支持使用IP列表直连配置服务获取配置。
