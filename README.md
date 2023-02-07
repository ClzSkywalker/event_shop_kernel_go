# todo 小商店kernel

```config
go 1.18
java openjdk 17.0.6 2023-01-17 LTS
ndk version:21.4.7075529
gcc version:8.1.0
```

```cmd
go build golang.org/x/mobile/cmd/gomobile
gomobile init
```

```issue
问题一：
Could not create task ':shared_preferences_android:generateDebugUnitTestConfig'.
this and base files have different roots: D:\project\flutter\event_shop_app\build\shared_preferences_android and C:\Users\Administrator\AppData\Local\Pub\Cache\hosted\pub.flutter-io.cn\shared_preferences_android-2.0.15\android
解决1：
flutter clean
flutter pub get
解决2：
把当前的项目目录 放到 C盘上去执行
```

# 文件目录
```
├─.githooks
├─.vscode
├─assert
├─cmd --后端调试启动入口
├─internal 
│  ├─container -- 全局对象容器
│  ├─ctrl -- curd 接口
│  ├─middleware -- 中间器
│  ├─model 
│  ├─router -- 路由管理
│  ├─server -- 启动服务相关入口
│  └─service 
├─kernel -- 第三方平台编译入口
├─logs -- 日志输出
├─output -- 打包结果入口
├─pkg
│  ├─db -- db 初始化模块
│  ├─jwt
│  ├─logger -- 日志管理
│  └─utils -- 工具类
└─script -- 脚本
```
