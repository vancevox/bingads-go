# BingAds Go API 客户端

这是一个用于访问Microsoft Advertising API（原BingAds API）的Go语言客户端库。该库提供了简单直观的方式来与Microsoft Advertising API进行交互。

## 功能特性

- 支持SOAP协议与Microsoft Advertising API通信
- 自动处理认证和请求头
- 简洁的API接口设计
- 完整的XML序列化和反序列化支持
- 丰富的数据模型

## 目前支持的服务

- 共享列表服务(SharedListService)
  - 获取共享列表项(GetListItemsBySharedList)
  - 获取共享实体(GetSharedEntities)
  - 获取共享实体关联(GetSharedEntityAssociationsBySharedEntityIds)
  - 添加列表项到共享列表(AddListItemsToSharedList)
  - 从共享列表删除列表项(DeleteListItemsFromSharedList)

## 快速开始

```go
package main

import (
    "fmt"
    
    "github.com/vancevox/bingads-go/campaignManagement/models"
    "github.com/vancevox/bingads-go/campaignManagement/service"
    "github.com/vancevox/bingads-go/config"
)

func main() {
    // 创建客户端
    client := service.NewClient(&config.Config{
        Auth: &config.AuthConfig{
            DeveloperToken:      "你的开发者令牌",
            AuthenticationToken: "你的认证令牌",
            CustomerID:          "你的客户ID",
            CustomerAccountID:   "你的客户账户ID",
        },
        API: config.DefaultConfig(),
    })
    
    // 启用调试模式查看请求和响应
    client.Config.API.Debug = true
    
    // 获取共享列表服务
    sharedListService := client.SharedListService()
    
    // 调用API方法，例如获取共享实体
    entities, err := sharedListService.GetSharedEntities(
        models.SharedEntityTypePlacementExclusionList, 
        models.EntityScopeCustomer,
    )
    
    if err != nil {
        fmt.Printf("错误: %v\n", err)
        return
    }
    
    fmt.Printf("共享实体列表: %+v\n", entities)
}
```

## 配置选项

可以通过`config.Config`结构体配置客户端：

```go
config := &config.Config{
    Auth: &config.AuthConfig{
        DeveloperToken:      "开发者令牌",
        AuthenticationToken: "认证令牌",
        CustomerID:          "客户ID",
        CustomerAccountID:   "客户账户ID",
    },
    API: &config.APIConfig{
        Debug:            false, // 设置为true可以查看请求和响应详情
        Timeout:          30,    // 请求超时时间（秒）
        RetryCount:       3,     // 重试次数
        RetryWaitTimeSec: 5,     // 重试等待时间（秒）
    },
}
```

## 测试

运行单元测试：

```bash
go test -v ./test/unit
```
