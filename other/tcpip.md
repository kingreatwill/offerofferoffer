
# TCP/IP

## KEEPALIVE
https://blog.csdn.net/lanyang123456/article/details/90578453
### 为什么应用层需要heart beat/心跳包？
TCP keepalive已经很牛逼了，但为什么还会提到应用层的心跳呢？

个人认知的原因包括两个：
- TCP keepalive处于传输层，由操作系统负责，能够判断进程存在，网络通畅，但无法判断进程阻塞或死锁等问题。
- 客户端与服务器之间有四层代理或负载均衡，即在传输层之上的代理，只有传输层以上的数据才被转发，例如socks5等

所以，基于以上原因，有时候还是需要应用程序自己去设计心跳规则的。
可以服务端负责周期发送心跳包，检测客户端，也可以客户端负责发送心跳包，或者服服务端和客户端同时发送心跳包。

可以根据具体的应用场景进行设计。