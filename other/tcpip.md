
# TCP/IP

## KEEPALIVE

在双方长时间未通讯时，如何得知对方还活着？如何得知这个TCP连接是健康且具有通讯能力的？

TCP的保活机制就是用来解决此类问题，这个机制我们也可以称作：keepalive。
保活机制默认是关闭的，TCP连接的任何一方都可打开此功能。有三个主要配置参数用来控制保活功能。

三个参数保活时间：tcp_keepalive_time、探测时间间隔：tcp_keepalive_intvl、探测循环次数：tcp_keepalive_probes。

如果在一段时间（保活时间：tcp_keepalive_time）内此连接都不活跃，开启保活功能的一端会向对端发送一个保活探测报文。

- 若对端正常存活，且连接有效，对端必然能收到探测报文并进行响应。此时，发送端收到响应报文则证明TCP连接正常，重置保活时间计数器即可。
- 若由于网络原因或其他原因导致，发送端无法正常收到保活探测报文的响应。那么在一定探测时间间隔（tcp_keepalive_intvl）后，将继续发送保活探测报文。直到收到对端的响应，或者达到配置的探测循环次数上限（tcp_keepalive_probes）都没有收到对端响应，这时对端会被认为不可达，TCP连接随存在但已失效，需要将连接做中断处理。

在探测过程中，`对方主机`会处于以下四种状态之一：

状态 | 处理
---|---
对方主机 | TCP连接正常，将保活计时器重置
对方主机已崩溃，包括：已关闭或者正在重启 | TCP连接不正常，经过指定次数的探测依然没有得到响应，则断开连接
对方主机崩溃并且已经重启 | 重启后原连接已失效，对方由于不认识探测报文，会响应重置报文段，请求端将连接断开
对方主机仍在工作，由于某些原因不可达（如：网络）| TCP连接不正常，经过指定次数的探测依然没有得到响应，则断开连接



这三个参数，在linux上可以在/proc/sys/net/ipv4/路径下找到，或者通过sysctl -a | grep keepalive命令查看当前内核运行参数。

```
[root@vm01 ~]# cd /proc/sys/net/ipv4
[root@vm01 ipv4]# pwd
/proc/sys/net/ipv4
[root@vm01 ipv4]# cat /proc/sys/net/ipv4/tcp_keepalive_time
7200
[root@vm01 ipv4]# cat /proc/sys/net/ipv4/tcp_keepalive_probes
9
[root@vm01 ipv4]# cat /proc/sys/net/ipv4/tcp_keepalive_intvl
75
[root@vm01 ipv4]# sysctl -a | grep keepalive
net.ipv4.tcp_keepalive_time = 7200
net.ipv4.tcp_keepalive_probes = 9
net.ipv4.tcp_keepalive_intvl = 75
```
- 保活时间（tcp_keepalive_time）默认：7200秒
- 保活时间间隔（tcp_keepalive_intvl）默认：75秒
- 探测循环次数（tcp_keepalive_probes）默认：9次


### 为什么应用层需要heart beat/心跳包？
TCP keepalive已经很牛逼了，但为什么还会提到应用层的心跳呢？

个人认知的原因包括两个：
- TCP keepalive处于传输层，由操作系统负责，能够判断进程存在，网络通畅，但无法判断进程阻塞或死锁等问题。
- 客户端与服务器之间有四层代理或负载均衡，即在传输层之上的代理，只有传输层以上的数据才被转发，例如socks5等

所以，基于以上原因，有时候还是需要应用程序自己去设计心跳规则的。
可以服务端负责周期发送心跳包，检测客户端，也可以客户端负责发送心跳包，或者服服务端和客户端同时发送心跳包。

可以根据具体的应用场景进行设计。

> 客户端不仅要设计`心跳包`，也要考虑`重连`

### 参考
[TCP keepalive的详解(解惑)](https://blog.csdn.net/lanyang123456/article/details/90578453)