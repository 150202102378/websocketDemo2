## 文件目录结构：

## 配置文件：
- listenPort1: register端口
- listenPort2 = push端口
- reactPort1: docker上register映射接口
- reactPort2: docker上push映射端口
- postgresPort: docker上psql数据库映射端口

- serverIP: server ip地址
- connNum: 客户端连接数量

- [psql] : 数据库相关信息

## 待解决问题
1. ping pong 时间间隔设置（需要考虑push的时候占用时间情况）
2. 频道设置（设置频道数量待定,频道数量设置应该添加到配置文件中）

## 测试执行流程：
1. 修改配置文件的连接数，服务器地址，psql地址
2. 修改系统的文件操作数
3. make all（在docker上生成一个psql容器）
4. 在本机上直接运行服务器
5. make test（自行决定开几个docker）

## 如何与其它系统结合使用
**该框架在订阅和推送方面已经基本完善，与其它系统结合使用应注意以下两点：**
1. 信息的获取源 
2. 订阅的频道需要多个分类可以在 channelManager 中修改数量
3. 同一频道下有不同分类时，可以在 client 中添加字段，来区分该客户想订阅的类型
4. 注意ping、pong时间调整，传送字节大小调整




