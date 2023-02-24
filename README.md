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


