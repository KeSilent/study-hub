# StudyHub

# 需求
因为个人做了一个培训班，所以需要一个记录学生课时的工具。因为需求很简单，只是简单的增加、扣除课时而已，在网上找的收费的感觉没必要，免费的又不适用，于是自己弄了个微信小程序。

# 架构
前端：uni-app
后端：gin-vue-admin

# 开发工具

Visual Studio Code、微信开发者工具、MySQL

# 目录结构

gin-vue-admin
├── server (后端服务)
└── web（后端页面）
└── wx（微信小程序）


# 实现内容
## 后端

### 微信的appId和secret配置

![](https://image.ypplog.cn/202408091149914.png)


### 登录页面
> 第一次运行，可以直接点击“前往初始化”，进行数据库和基础数据的初始化。

![](https://image.ypplog.cn/202408091132667.png)

### 首页

> 可以通过左侧菜单和右侧快捷按钮进入。

![](https://image.ypplog.cn/202408091134019.png)

### 组织机构

> 可以进行创建、编辑。

![](https://image.ypplog.cn/202408091137478.png)

### 课程管理

> 进行课程管理

> 可以对机构进行课程管理。

![](https://image.ypplog.cn/202408091146103.png)

### 用户管理

> 需要在“超级管理员”-》“用户管理”中对用户进行管理。

![](https://image.ypplog.cn/202408091147598.png)



## 小程序端
### 学生管理页面
![](https://image.ypplog.cn/202408221613506.png)