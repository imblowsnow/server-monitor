

## 启动脚本
```shell
# github原始脚本
curl https://raw.githubusercontent.com/imblowsnow/server-monitor/master/deploy-server.sh | bash -s -- 22251

# 加速脚本
curl https://cdn.jsdelivr.net/gh/imblowsnow/server-monitor@master/deploy-server.sh | bash -s -- 22251
```

## 测试编译
```shell
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build 
```
## 预览
![image](https://github.com/imblowsnow/server-monitor/assets/74449531/270497a4-4cd3-463a-816e-40fed98a0033)
![image](https://github.com/imblowsnow/server-monitor/assets/74449531/63b5d99e-0976-4b69-86cc-cffe9296371a)


## TODO
- [x] 编写后台API / 首页API
- [x] 编写视图页面
- [ ] 对接服务器终端
- [ ] 通知渠道
- [ ] 上线/离线通知
