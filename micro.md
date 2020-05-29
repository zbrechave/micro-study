## micro基础

### API，Web，SRV服务之间的区别

<img src="/Users/davve/Library/Application Support/typora-user-images/image-20200529114835393.png" alt="image-20200529114835393" style="zoom:50%;" />

### API服务

micro api 默认把**go.micro.api**作为API网关服务的命名空间。

### Web服务

micro web 默认把**go.micro.web**作为服务的命名空间。

### SRV服务

SRV 服务是RPC服务的基础，也就是你常写的服务类型。我们一般把它称作RPC或后后端服务，因为它作为后台架构的一部分，不应该暴露在最外层。默认情况下，我们使用**go.micro.srv**作为它的命名空间，或者你可以使用像**com.example.srv**这样的名字。

### Micro Api

#### Handler

Handler负责持有并管理HTTP请求路由。默认的handler使用从注册中心获取的端口元数据来决定指向服务的路由，如果路由不匹配，就会回退到使用”rpc” hander。

API有如下请求handler：

- --handler=api 处理http请求，通过RPC来完全控制http的请求/响应。

  Content-Type: 支持任何类型

  Body: 支持任何格式

  Forward Format: 转发格式，[api.Request](https://github.com/micro/go-api/blob/master/proto/api.proto#L11)/[api.Response](https://github.com/micro/go-api/blob/master/proto/api.proto#L21)

  Path: 请求路径，`/[service]/[method]`

  Resolver: 请求解析器，路径会被解析成服务与方法

  Configure: 配置，在启动时指定`--handler=api`或在启动命令前指定环境变量`MICRO_API_HANDLER=api`

- --handler=rpc 处理json及protobuf格式的POST请求，并转向RPC。

  Content-Type: `application/json` or `application/protobuf`

  Body: JSON 或者 Protobuf

  Forward Format: **json-rpc**或者**proto-rpc**，与`Content-Type`有关

  Path: `/[service]/[method]`

  Resolver: 请求解析器，路径会被解析成服务与方法

  Configure: 配置，在启动时指定`--handler=rpc`或在启动命令前指定环境变量`MICRO_API_HANDLER=rpc`

  如果没有设置时，RPC Handler就是**默认**的handler，

- --handler=proxy 处理http请求并转向反向代理。

  Content-Type: 支持任何类型

  Body: 支持任何格式

  Forward Format: HTTP反向代理

  Path: `/[service]`

  Resolver: 请求解析器，路径会被解析成服务名

  Configure: 配置，在启动时指定`--handler=proxy`或在启动命令前指定环境变量`MICRO_API_HANDLER=proxy`

  REST can be implemented behind the API as microservices

- --handler=web 包含web socket的http反向代理

  Content-Type: 支持任何类型

  Body: 支持任何格式

  Forward Format: HTTP反向代理，包括web socket

  Path: `/[service]`

  Resolver: 请求解析器，路径会被解析成服务名

  Configure: 配置，在启动时指定`--handler=web`或在启动命令前指定环境变量`MICRO_API_HANDLER=web`

#### RPC Endpoint

**/rpc**端点允许绕过主handler，然后与任何服务直接会话。

- `service` - 指定服务名
- `method` - 指定方法名
- `request` - 请求body体
- `address` - 可选，指定特定的目标主机地址

```rust
curl -d 'service=go.micro.srv.greeter' \
     -d 'method=Say.Hello' \
     -d 'request={"name": "Bob"}' \
     http://localhost:8080/rpc
```

### Resolver

解析器，Micro使用命名空间与HTTP请求路径来动态路由到具体的服务。

API命名的空间是`go.micro.api`。可以通过指令`--namespace`或者环境变量`MICRO_NAMESPACE=`设置命名空间。

#### RPC Resolver

RPC解析器示例中的RPC服务有名称与方法，分别是`go.micro.api.greeter`，`Greeter.Hello`。

URL会被解析成以下几部分：

| 路径             | 服务                 | 方法    |
| ---------------- | -------------------- | ------- |
| /foo/bar         | go.micro.api.foo     | Foo.Bar |
| /foo/bar/baz     | go.micro.api.foo     | Bar.Baz |
| /foo/bar/baz/cat | go.micro.api.foo.bar | Baz.Cat |

带版本号的API URL也可以很容易定位到具体的服务：

| Path            | Service             | Method  |
| --------------- | ------------------- | ------- |
| /foo/bar        | go.micro.api.foo    | Foo.Bar |
| /v1/foo/bar     | go.micro.api.v1.foo | Foo.Bar |
| /v1/foo/bar/baz | go.micro.api.v1.foo | Bar.Baz |
| /v2/foo/bar     | go.micro.api.v2.foo | Foo.Bar |
| /v2/foo/bar/baz | go.micro.api.v2.foo | Bar.Baz |

#### Proxy Resolver

代理解析器只处理服务名，所以处理方案和RPC解析器有点不太一样。

URL会被解析成以下几部分：

| 路径           | 服务                 | 方法           |
| -------------- | -------------------- | -------------- |
| /foo           | go.micro.api.foo     | /foo           |
| /foo/bar       | go.micro.api.foo     | /foo/bar       |
| /greeter       | go.micro.api.greeter | /greeter       |
| /greeter/:name | go.micro.api.greeter | /greeter/:name |