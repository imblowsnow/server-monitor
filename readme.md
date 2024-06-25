

## 启动脚本
```shell
curl https://raw.githubusercontent.com/imblowsnow/server-monitor/master/deploy-server.sh | bash -s -- 22251
curl https://cdn.jsdelivr.net/gh/imblowsnow/server-monitor@master/deploy-server.sh | bash -s -- 22251
```

## 测试编译
```shell
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build 
```


## TODO
- [ ] 编写后台API / 首页API
- [ ] 编写视图页面
- [ ] 对接服务器终端
- [ ] 通知渠道
- [ ] 上线/离线通知
