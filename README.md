# todo 小商店kernel

```config
go 1.18
java openjdk 17.0.6 2023-01-17 LTS
ndk version:21.4.7075529
gcc version:8.1.0
pprof https://graphviz.org/download/
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

```text
├─.githooks
├─.vscode
├─assert
│  └─ectype
├─cmd
│  ├─api
│  └─mobile
├─internal
│  ├─container
│  ├─contextx
│  ├─ctrl
│  ├─entity
│  ├─Infrastructure
│  ├─middleware
│  ├─model
│  ├─router
│  ├─server
│  └─service
├─logs
├─output
├─pkg
│  ├─bmobx
│  ├─constx
│  ├─datatypesx
│  ├─db
│  ├─httpx
│  ├─i18n
│  │  ├─entry
│  │  ├─errorx
│  │  └─module
│  ├─loggerx
│  ├─processx
│  ├─recoverx
│  └─utils
├─script
└─test
```

1. assert 存放静态资源
infrastructure 
基础结构层：用于存放不带事务的基础处理，加上基础逻辑校验

service
服务层：携带事务，用于组合infrastructure逻辑，组合成基础api接口逻辑

ctrl
接口层：用于校验参数

