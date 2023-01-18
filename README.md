# my-world-server
my-world服务器代码，基于go-leaf+mongodb+mysql开发
## 概述
登陆服进入游戏，网关服监听用户状态，必要时可强踢/强拉，游戏服负责开放地图部分玩法逻辑（如任务系统），副本服负责子游戏部分逻辑（如战斗系统），聊天服进行系统通知/邮件/聊天。

## 服务器结构
* 登陆服
* 网关服
* 游戏服
* 副本服
* 聊天服

## 项目配置
* config.yaml

## 产品配置
* role.xml
* npc.xml
* goods.xml
* shop.xml