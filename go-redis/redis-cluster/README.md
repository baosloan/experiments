# Go Redis Cluster客户端

go-redis 支持 Redis Cluster 客户端，如下面示例，redis.ClusterClient 表示集群对象，对集群内每个 redis 节点使用 redis.Client 对象进行通信，每个 redis.Client 会拥有单独的连接池。

连接到 redis 集群示例：
```go
import "github.com/redis/go-redis/v9"

rdb := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
})

```

更多配置参数，请参照 redis.ClusterOptions:
部分配置项继承自 Options，请查看 Options 说明。

```go
type ClusterOptions struct {
    // redis集群的列表地址
	// 例如：[]string{"192.168.1.10:6379", "192.168.1.11:6379"}
    Addrs []string

    // ClientName 和 `Options` 相同，会对集群每个Node节点的每个网络连接配置
    ClientName string

    // New集群节点 `*redis.Client` 的对象，
	// go-redis 默认使用 `redis.NewClient(opt)` 方法
    NewClient func(opt *Options) *Client

    // 同 `Options`
    MaxRedirects int

    // 启用从节点处理只读命令，go-redis会把只读命令发给从节点(如果有从节点)
	// 默认不启用
    ReadOnly bool

	// 把只读命令发送到响应最快的节点，自动启用 `ReadOnly` 选项
    RouteByLatency bool

    // 把只读命令随机到一个节点，自动启用 `ReadOnly` 选项
    RouteRandomly bool

	// 返回redis集群Slot信息的函数，go-redis默认将获取redis-cluster的配置信息
	// 如果你是自建redis集群在节点直接操作读写，需要自己配置Slot信息
	// 可以使用 `Cluster.ReloadState` 手动加载集群配置信息
    ClusterSlots func(context.Context) ([]ClusterSlot, error)

    // 下面的配置项，和 `Options` 基本一致，请参照 `Options` 的说明

    Dialer func(ctx context.Context, network, addr string) (net.Conn, error)

    OnConnect func(ctx context.Context, cn *Conn) error

    Username string
    Password string

    MaxRetries      int
    MinRetryBackoff time.Duration
    MaxRetryBackoff time.Duration

    DialTimeout           time.Duration
    ReadTimeout           time.Duration
    WriteTimeout          time.Duration
    ContextTimeoutEnabled bool

    PoolFIFO        bool

	// 连接池配置项，是针对集群中的一个节点，而不是整个集群
	// 例如你的集群有15个redis节点， `PoolSize` 代表和每个节点的连接数量
	// 最终最大连接数为 PoolSize * 15节点数量
    PoolSize        int
    PoolTimeout     time.Duration
    MinIdleConns    int
    MaxIdleConns    int
    ConnMaxIdleTime time.Duration
    ConnMaxLifetime time.Duration

    TLSConfig *tls.Config
}
```

遍历每个节点：

```go
err := rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
    return shard.Ping(ctx).Err()
})
if err != nil {
    panic(err)
}
```

只遍历主节点请使用： ForEachMaster， 只遍历从节点请使用： ForEachSlave

你也可以自定义的设置每个节点的初始化:

```go
rdb := redis.NewClusterClient(&redis.ClusterOptions{
    NewClient: func(opt *redis.Options) *redis.NewClient {
        user, pass := userPassForAddr(opt.Addr)
        opt.Username = user
        opt.Password = pass

        return redis.NewClient(opt)
    },
})
```