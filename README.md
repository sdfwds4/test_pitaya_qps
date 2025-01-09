# test_pitaya_qps
测试  go nrpc 服务器的 qps


## 回显测试

启动gate
```
server.exe -frontend=true -type=gate
```

启动game
```
server.exe -frontend=false -type=gam
```

测试工具
```
go install github.com/topfreegames/pitaya/v3
```
```
pitaya repl
>>> connect 127.0.0.1:3250
>>> request gate.connector.getsessiondata
>>> request game.room.getsessiondata
```



## 回显测试

- 测试代码：
  - [server.go](server.go)
  - [client.go](client.go)

