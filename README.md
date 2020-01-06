- 配置文件：
1. listenPort: service监听接口
2. reactPort: service所在docker映射接口
3. serviceIP: service所在docker地址

- 运行：
1. make all后进入service所在docker
2. 运行指令./service
3. make all后进入client所在docker
4. 运行指令./client
 
- 测试工具运行指令：
    ./bin/jmeter -n -t processGroup.jmx -l log.jtl



* 待解决问题
1. client收到connect: connection timed out, (应该是没有心跳机制，所以60秒后断开) (read: connection reset by peer)
2. client中一个dial可以创建多少个连接？ 
3. 广播与心跳冲突？上锁