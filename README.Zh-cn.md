# rs-api


## ✨ 特性

- 遵循 RESTful API 设计规范

- 基于 GIN WEB API 框架，提供了丰富的中间件支持（用户认证、跨域、访问日志、追踪ID等）

- 基于Casbin的 RBAC 访问控制模型

- JWT 认证

- 支持 Swagger 文档(基于swaggo)

- 基于 GORM 的数据库存储，可扩展多种类型数据库

- 配置文件简单的模型映射，快速能够得到想要的配置

- 代码生成工具

- 表单构建工具

- 多指令模式

- 多租户的支持

- TODO: 单元测试

### 环境要求

go 1.18

node版本: v14.16.0

npm版本: 6.14.11

### 开发目录创建

```bash

# 创建开发目录
mkdir goadmin
cd goadmin
```

### 获取代码

> 重点注意：两个项目必须放在同一文件夹下；

```bash
# 获取后端代码
git clone https://github.com/go-admin-team/go-admin.git

# 获取前端代码
git clone https://github.com/go-admin-team/go-admin-ui.git

```

### 启动说明

#### 服务端启动说明

```bash
# 进入 go-admin 后端项目
cd ./go-admin

# 更新整理依赖
go mod tidy

# 编译项目
go build

# 修改配置 
# 文件路径  go-admin/config/settings.yml
vi ./config/settings.yml

# 1. 配置文件中修改数据库信息 
# 注意: settings.database 下对应的配置数据
# 2. 确认log路径
```

:::tip ⚠️注意 在windows环境如果没有安装中CGO，会出现这个问题；

```bash
E:\go-admin>go build
# github.com/mattn/go-sqlite3
cgo: exec /missing-cc: exec: "/missing-cc": file does not exist
```

or

```bash
D:\Code\go-admin>go build
# github.com/mattn/go-sqlite3
cgo: exec gcc: exec: "gcc": executable file not found in %PATH%
```

[解决cgo问题进入](https://doc.go-admin.dev/zh-CN/guide/faq#cgo-%E7%9A%84%E9%97%AE%E9%A2%98)

:::

#### 初始化数据库，以及服务启动

``` bash
# 首次配置需要初始化数据库资源信息
# macOS or linux 下使用
$ ./go-admin migrate -c config/settings.dev.yml

# ⚠️注意:windows 下使用
$ go-admin.exe migrate -c config/settings.dev.yml


# 启动项目，也可以用IDE进行调试
# macOS or linux 下使用
$ ./go-admin server -c config/settings.yml


# ⚠️注意:windows 下使用
$ go-admin.exe server -c config/settings.yml
```

#### sys_api 表的数据如何添加

在项目启动时，使用`-a true` 系统会自动添加缺少的接口数据
```bash
./go-admin server -c config/settings.yml -a true
```

#### 使用docker 编译启动

```shell
# 编译镜像
docker build -t go-admin .

# 启动容器，第一个go-admin是容器名字，第二个go-admin是镜像名称
# -v 映射配置文件 本地路径：容器路径
docker run --name go-admin -p 8000:8000 -v /config/settings.yml:/config/settings.yml -d go-admin-server
```

#### 文档生成

```bash
go generate
```

#### 交叉编译

```bash
# windows
env GOOS=windows GOARCH=amd64 go build main.go

# or
# linux
env GOOS=linux GOARCH=amd64 go build main.go
```

### UI交互端启动说明

```bash
# 安装依赖
npm install

# 建议不要直接使用 cnpm 安装依赖，会有各种诡异的 bug。可以通过如下操作解决 npm 下载速度慢的问题
npm install --registry=https://registry.npmmirror.com

# 启动服务
npm run dev
```



