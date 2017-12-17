## 简介
这是一个临时性修复Steam社区的小工具，就叫它AnotherSteamCommunityFix（简称ASCF）吧~
ASCF是由Go语言写的，所以它可以运行在几乎任何系统平台上！

程序有2种运行模式：
1. 转发模式：把本地的HTTP请求重定向为HTTPS，并解析出正确的IP地址进行连接
2. 代理模式：把HTTP/HTTPS请求通过KCP协议发送到代理服务器，并由代理服务器转发

*程序默认、推荐转发模式*

ASCF会修改hosts文件，把steamcommunity.com域名指向本地（127.0.0.1），然后程序会监听本地的80和443端口。
当程序退出时（按Ctrl+C退出），它会把hosts恢复原样。

## 下载使用
下载地址：
* [在Github Release页面下载二进制文件](https://github.com/zyfworks/AnotherSteamCommunityFix/releases)
* [百度云网盘分流下载](https://pan.baidu.com/s/1nvBW8qP)

运行：
1. 转发模式：直接运行即可
2. 代理模式：运行参数为 `-mode=2`

当程序解析域名失败时，会使用备用IP进行连接，默认备用IP为`104.115.125.124`，您也可以手动指定这个备用IP：
```./ascf.exe -ip=xxx.xxx.xxx.xxx```

## 使用说明
* 访问Steam社区时必须保持该程序运行！
* 如果出现闪退，请使用管理员权限启动，并确保系统中没有其他程序占用80和443端口。
* 第一次使用前请先清空hosts文件中和steamcommunity.com相关的条目。
* 树莓派用户请使用linux_arm版！

## Linux使用
在Linux服务器上挂卡的朋友们，如果你用的是x64的Linux，那么下载linux_amd64版的，把可执行文件上传至服务器，使用超级用户运行：
```sudo ./ascf```
或是后台运行 
```sudo nohup ./ascf &```
之后steamcommunity.com便可访问。
