# v1版本说明
1. 支持zmq-server接收消息。（已完成）
2. 支持集群自动发现和转发消息(已完成)
3. 支持webSocket转发消息,支持ssl协议（已完成）
   * 建立连接后先发送缓存消息。
   * 完成后再发送实时消息
4. 完成mysql存储。(已完成)
5. 完成消息分析和结果回调。（已完成）
6. 支持http接收消息。


# v2版本说明
1. 支持接收docker日志，并选举接收节点。
   选举策略：
      * etcd中已有对应关系，检测节点是否正常，如果正常返回该节点。
      * 对于新的serviceID根据监控数据选举闲节点。(以每分钟日志量+20倍服务量算标志，标志最小为最优)
      * 应用关闭时，删除etcd中的对应数据。（region-api完成）。
   选举请求：
      * 请求由dockerd为容器创建logger时完成。
      * dockerd完成链接状态检测工作，若分配的服务端状态异常。重新请求心得节点。   
2. 支持集群间各类消息通信。
   * 操作日志消息。（及时）
   * docker日志监控数据消息。（用于选举）
   * 容器状态检测停止消息。（用于应用启动和停止操作处理节点不在同一节点）
3. docker日志文件存储。
   * 文件缓存写入。（每128条写入或每1分钟写入）
   * 分布式文件存储。
   * 历史日志文件压缩。


# v3版本说明 

1. 支持监控数据websocket消息转发服务。
2. 消息输入zmq sub订阅数据源。
3. 进行数据切割，区分主题。
3. 消息输出以不同的主题进行区分和订阅。
 
