## 简介
这是一个临时性修复Steam社区的小工具，就叫它AnotherSteamCommunityFix（简称ASCF）吧~
ASCF是由Go语言写的，所以它可以运行在几乎任何系统平台上！

ASCF会修改hosts文件，把steamcommunity.com域名指向本地（127.0.0.1），然后程序会监听本地的80和443端口。
当程序退出时（按Ctrl+C退出），它会把hosts恢复原样。

## 下载地址
* [Github Release](https://github.com/zyfworks/AnotherSteamCommunityFix/releases)
* [百度云网盘分流下载](https://pan.baidu.com/s/1nvBW8qP)

## 注意事项
* 访问Steam社区时必须保持该程序运行！
* 如果出现闪退，请使用管理员权限启动，并确保系统中没有其他程序占用80和443端口。
* 您可以通过运行参数`-ip=xxx.xxx.xxx.xxx`令程序强制使用这个IP地址来连接。
* 第一次使用前请先清空hosts文件中和steamcommunity.com相关的条目。
* 树莓派2及以上请使用ascf_Linux_ARMv7、树莓派1请使用ascf_linux_ARMv6、64位ARM平台请使用ascf_linux_ARMv8。

## Linux/macOS使用指南
1. 下载并解压缩
2. 打开终端（Terminal），进入到ascf程序目录：
   如ascf程序在 /User/Makazeu/Downloads/ascf_darwin_amd64/文件夹中，那么在终端中输入:
   ```cd /User/Makazeu/Downloads/ascf_darwin_amd64/```
3. 赋予程序可执行权限，在终端中输入命令：
   ```chmod +x ./ascf```
4. 使用root用户（管理员用户）运行程序，在终端中输入：
   ```sudo ./ascf```
   输入root用户密码后，看程序是否运行。
   因为程序涉及到hosts文件修改，需要高权限，所以你需要输入root密码
5. 若程序已经成功运行，此时就不要关闭终端窗口了，否则程序就会退出！试下Steam社区能否正常打开。
6. 一切都没问题后，在终端窗口中退出程序（按Ctrl+C），然后以后台的方式运行程序，输入
   ```nohup sudo ./ascf &```
7. 之后就可以关闭终端窗口了，此时程序在后台运行。现在steamcommunity.com可以打开咯！
