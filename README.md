# 极简网上银行项目

>Go小白的第一个Go语言项目

此项目分为两个模块：一个用于程序的核心逻辑，另一个用于通过 Web API 公开逻辑。

## 定义功能和要求

网上银行系统将：

- 允许客户创建帐户。
- 允许客户取款。
- 允许客户将资金转到其他帐户。
- 提供包含客户数据和最终余额的对账单。
- 通过终结点公开一个 Web API，用于输出对账单。

## 创建初始项目文件

$GOPATH 目录中创建以下文件结构：

```
simple-online-bank/
    bankcore/
        go.mod
        bank.go
    bankapi/
        go.mod
        main.go
```