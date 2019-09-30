
## [K8s ](https://www.kubernetes.org.cn)

## Kubernetes 节点
在Kubernets 系统架构图中，服务分为运行在工作节点上的服务和组成集群级别控制板的服务

Kubernetes 节点有运行应用容器必备的服务，而这些都是受Master的控制

每次个节点上当然都要运行 Docker。Docker 来负责所有具体的镜像下载和容器运行

Kubernetes 主要由以下几个核心组件组成：

* etcd 保存了整个集群的状态
* apiserver 提供了资源操作的唯一入口，并提供认证、授权、访问控制、API注册和发现等机制；
* controller manager 负责维护集群的状态，比如故障检测、自动拓展、滚动更新等
* sceduler 负责资源的调度，按照预定的调度策略将Pod调度到相应的机器上
* kubelet 负责维护容器的生命周期，同时也负责 Volume（CVI）和网络（CNI）的管理
* Container runtime 负责镜像管理以及Pod和容器的真正运行(CRI)
* kube-proxy 负责为Service提供cluster内部的服务发现和负载均衡

除了核心组件，还有一些推荐的Add-ons
* kube-dns负责为整个集群提供DNS服务
* Ingress Controller 为服务提供外网入口
* Heapster 提供资源监控
* Dashboard 提供GUI
* Federation 提供跨可用区的集群
* Fluentd-elasticsearch 提供集群日志采集、存储与查询

## 分层架构
Kubernetes 设计理念和功能其实就是一个类似 Linux 的分层架构

* 核心层：Kubernetes 最核心的功能，对外提供 API 构建高层的应用，对应提供插件式应用执行环境
* 应用层：部署（无状态应用、有状态应用、批处理任务、集群应用等）和路由（服务发现、DNS解析等）
* 管理层：系统度量（如基础设施、容器和网络的度量），自动化（如自动拓展、动态Provision等）以及策略管理（PBAC、Quota、PSP、NetworkPolicy等）
* 接口层：kubectl命令行工具、客户端SDK以及集群联邦
* 生态系统: 在接口层之上的庞大容器集群管理调度的生态系统，可以划分为两个范畴
    * Kubernetes外部：日志、监控、配置管理、CI、CD、Workflow、Faas、OTS应用、ChatOps等
    * Kubernetes内部：CRI、CNI、CVI、镜像仓库、Cloud Provider、集群自身的配置和管理等
    
## Kubelet
Kubelet 负责管理Pods和它们上面的容器、images镜像、volumes、etc

## Kube-proxy
每一个节点也运行一个简单的网络代理和负载均衡。正如Kubernetes API里卖弄定义的这些服务也可以在各种终端中以轮询的方式做一些简单的TCP和UDP传输

服务端点目前是通过DNS或者环境变量（Docker-links-compatible 和 Kubernetes{FOO}_SERVICE_HOST 及 {FOO}_SERVICE_PORT变量都支持）
这些变量由服务代理所管理的端口来解析

## Kubernetes 控制面板
Kubernetes 控制面板可以分为多个部分。目前它们都运行在一个master节点，然而为了达到高可性，这需要改变。不同部分一起协作提供一个统一的关于集群的视图

## etcd
所有master的持续状态都存在etcd的一个实例中。这可以很好地存储配置数据。因为有watch（观察者）的支持，各部件协调中的改变可以很快被察觉

## Kubernetes API Server
API服务提供Kubernetes API 的服务。这个服务试图通过把所有或者大部分的业务逻辑放到两个不同的部件中而使其具有CRUD特性。它主要处理REST操作，
在etcd中验证更新这些对象（并最终存储）

## Scheduler
调度器把未调度的pod通过 binding api 绑定到节点上。调度器是可插拔的，并且我们期待支持多集群的调度，未来甚至希望可以支持用户自定义的调度器。

## Kubernetes 控制管理服务器
所有其它的集群级别的功能目前都是由控制管理器所负责。例如，端点对象是被端点控制器来创建和更新。这些最终可以被分割成不同的部件来让它们独自的可插拔

## Kubernetes 设计理念
### Kubernetes 设计理念与分布式系统
分析和理解 Kubernetes 的设计理念可以使我们更深入地了解Kubernetes系统，更好地利用它管理分布式部署的云原生应用，另一方面也可以让我们
借鉴其在分布式系统设计方面的经验

### API 设计原则
对于云计算系统，系统API实际上处于系统设计的统领地位，正如本文前面所说，K8s集群系统每支持一项新功能，引入一项新技术，一定会新引入对应的API对象，
支持对该功能的管理操作，理解掌握的API，K8s系统API的设计有以下几条原则：

1. 所有API应该是声明式的。
2. API对象是彼此互补而且可组合的
3. 高层API以操作意图为基础设计
4. 低层API根据高层API的控制需要设计
5. 尽量避免简单封装，不要有在外部API无法显式知道的内部隐藏的机制。
6.  API操作复杂度与对象数量成正比
7.  API对象状态不能依赖于网络连接状态
8.  尽量避免让操作机制依赖于全局状态，因为在分布式系统中要保证全局状态的同步是非常困难的
    
### 控制器设计原则

* 控制逻辑应该只依赖于当前状态
* 假设任何错误的可能，并做容错处理
* 尽量避免复杂状态机，控制逻辑不要依赖无法监控的内部状态
* 假设任何操作都可能被任何操作对象拒绝，甚至被错误解析
* 每个模块都可以在出错后自动恢复
* 每个模块都可以在必要时优雅地降级服务

## Kubernetes 的核心技术概念和API对象
* Pod
* 复制控制器 （Replication Controller, RC）
* 副本集 （Replica Set, RS）
* 部署 （Deployment）
* 服务 （Service）
* 任务（Job）
* 后台支撑服务集（DaemonSet）
* 有状态服务集（PetSet）
* 集群联邦（Federation）
* 存储卷（Volume）
* 持久存储卷（persistent Volume, PV）和持久存储卷声明（Persistent Volume Claim, PVC）
* 节点（Node）
* 密钥对象（Secret）
* 用户账户（User Account）和服务账户（Service Account）
* 名字空间（Namespace）
* RBAC访问授权

从K8s的系统架构、技术概念和设计理念，我们可以看到K8s系统最核心的两个设计理念：一个是容错性，一个是易拓展性。容错性实际是保证
K8s系统稳定性和安全性的基础，易拓展性是保证K8s对变更友好，可以快速迭代增加新功能的基础

