Docker 相关的几个核心概念，
使用 Linux NameSpace 进行网络、进程空间、命名空间等资源的隔离，
使用 Cgroups 技术对资源的占用、使用量进行限制，
使用 AUFS 等存储驱动来实现分层结构、增量更新等能力。

但 Docker 技术的原理远不止这些，但理解上述内容，已经能够很好的帮助我们大致理解 Docker 的运行原理。


## Docker 的资源隔离：NameSpace
1. Linux NameSpace
多个 Docker 容器之间需要保持相互隔离，来模拟 “独立环境”，而 Docker 正式利用了 Linux NameSpace 来实现这一能力，Linux Namespace 是 Linux 提供的一种机制，可以实现不同资源的隔离。

Linux 提供的主要的 NameSpace
- Mount NameSpace- 用于隔离文件系统的挂载点
- UTS NameSpace- 用于隔离 HostName 和 DomianName
- IPC NameSpace- 用于隔离进程间通信
- PID NameSpace- 用于隔离进程 ID
- Network NameSpace- 用于隔离网络
- User NameSpace- 用于隔离用户和用户组 UID/GID

2. Linux NameSpace API
在已经有了 Linux NameSpace 概念下，怎么来利用这些 NameSpace 呢？其实我们知道，Docker 容器本身就是一个进程，那么对容器的 NameSpace 操作，显然也就是对进程的 API 操作。Linux 主要提供了以下对 NameSpace 的操作接口

Linux NameSpace API

- clone - 创建一个新进程
- setns - 允许进程重新加入一个已经存在的 NameSpace
- unshare - 将调用进程移动到新的 NameSpace

## Docker 的资源限制：cgroups
1、cgroups

其实在调用了 Linux NameSpace 的情况下，还有一个问题没有解决： 那就是资源的限制 。

虽然实现了资源的隔离，但是 Docker 的本质还是一个进程，在多个 Docker 进程的情况下，如果其中一个进程就占满了所有的 CPU 与内存，其他进程处于忙等而导致服务无响应，这是难以想象的。因此除了 NameSpace 隔离，还需要通过另外一种技术来限制进程资源使用大小情况：cgroups(control groups)。

2、subsystem

Linux 有以下的 cgroups subsystem，用于资源控制

- memory - 内存限制
- hugetlb - huge pages 使用量
- cpu - CPU 使用率
- cpuacct - CPU 使用率
- cpuset - 绑定 cgroups 到指定 CPUs 和 NUMA 节点
- innodb_lock_wait_timeout - block 设备的 IO 速度
- net_cls - 网络接口设置优先级
- devices - mknode 访问设备权限
- freezer - suspend 和 restore cgroups 进程
- perf_event - 性能监控
- pids - 限制子树 cgroups 总进程数
  
3、cgroups 控制资源
cgroups 用来将进程统一分组，并用于控制进程的内存、CPU 以及网络等资源使用。cgroups 会将系统进程组成成独立的树，树的节点是进程组，每棵树与一个或者多个 subsystem 关联。subsystem 的作用就是对组进行操作。

## Docker 存储驱动：Union File Systems
### 1、AUFS

其实在有了 NameSpace 和 cgroups 的情况下，对于 Docker 项目的初始化可以简单抽象为：

1. 启动 Namespace 配置
2. 设置 cgroups 参数，对资源进行限制
3. 切换进程的根目录
   
但是还有另外一个问题， 是否每次打包、升级镜像都要重新走一遍整个初始化流程 ？这显然是不合理的。

Docker 中最典型的存储驱动就是 AUFS（Advanced Multi-layered unification filesytem），可以将 AUFS 想象为一个可以 “栈式叠加” 的文件系统，AUFS 允许在一个基础的文件系统的上，“增量式” 的增加文件。

AUFS 支持将不同目录挂载到同一个目录下，这种挂载对用户来说是透明的。通常，AUFS 最上层是可读写层，而最底层是只读层，每一层都是一个普通的文件系统。

技术干货分享：Docker 核心技术与基本原理
Docker 的镜像分层，正式通过 AUFS 通过分层结构、挂载等方式实现。在用户制作镜像的每一步，Docker 都会生成一个层，也就是增量 rootfs。

rootfs (root file system) ，根文件系统，本质上就是一个 Linux 系统中基本的文件目录组织结构，包含的就是典型 Linux 系统中的 /dev, /proc, /bin, /etc 等标准目录和文件。

同时，除了只读层，可读写层，Docker 在 AUFS 中有一个单独生成的内部层， init 层 ，用于存放 /etc/hosts 、 /etc/resolv.conf 等相关信息。

### 2、AUFS 的读写操作

读操作
当需要 写入 一个文件时，不存在时则在可读写层新建一个，否则一直从向下寻找。

写操作
当 删除 一个文件中，如果文件仅存在可读写层，则直接删除这个文件。

但是又有一个问题，如果删除的是只读层的文件呢？所以在这种情况下，会先删除可读写层中的备份，之后通过创建一个 whiteout 文件来标记文件不存在，这其实是一种 “遮挡”，只读层文件却不会被真正的删除，但是表现上确是已经被 “删除” 了。

当 新建 文件时，由于 whiteout 的 存在，所以需要先检查 whiteout 是否存在，存在的情况下，需要先删除再创建。

AUFS 只是 Docker 存储驱动的其中一种，在有些场景下并不是最优的选择，但都是属于 Union File System，主要是基于 “写时复制” 以及 “用时配置” 两种方式，但它能够有效帮助我们理解 Docker 的分层结构以及原理。其他的 Docker 存储驱动还有 OverlayFS、Devicemapper、Btrfs、ZFS 等，这里不再赘述。

## 参考
[Docker 核心技术与基本原理](https://www.toutiao.com/i6720007968827392520)