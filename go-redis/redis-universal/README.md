# Go Redis Universal客户端

UniversalClient 并不是一个客户端，而是对 Client 、 ClusterClient 、 FailoverClient 客户端的包装。

根据不同的选项，客户端的类型如下：

如果指定了 MasterName 选项，则返回 FailoverClient 哨兵客户端。
如果 Addrs 是 2 个以上的地址，则返回 ClusterClient 集群客户端。
其他情况，返回 Client 单节点客户端。

示例如下：

```go
// *redis.Client.
rdb := NewUniversalClient(&redis.UniversalOptions{
    Addrs: []string{":6379"},
})

// *redis.ClusterClient.
rdb := NewUniversalClient(&redis.UniversalOptions{
    Addrs: []string{":6379", ":6380"},
})

// *redis.FailoverClient.
rdb := NewUniversalClient(&redis.UniversalOptions{
    Addrs: []string{":6379"},
    MasterName: "mymaster",
})
```

更多设置请参照redis.UniversalOptions:
部分配置项继承自 Options，请查看 Options 说明。

```go
type UniversalOptions struct {
    // 单个主机或集群配置
	// 例如：[]string{"192.168.1.10:6379"}
	Addrs []string

	// ClientName 和 `Options` 相同，会对每个Node节点的每个网络连接配置
	ClientName string

    // 设置 DB, 只针对 `Redis Client` 和 `Failover Client`
	DB int

    // 下面的配置项，和 `Options`、`Sentinel` 基本一致，请参照 `Options` 的说明

	Dialer    func(ctx context.Context, network, addr string) (net.Conn, error)
	OnConnect func(ctx context.Context, cn *Conn) error

	Username         string
	Password         string
	SentinelUsername string
	SentinelPassword string

	MaxRetries      int
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration

	DialTimeout           time.Duration
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ContextTimeoutEnabled bool

	PoolFIFO bool

    // 连接池配置项，是针对一个节点的设置，而不是所有节点
    // 例如你的集群有15个redis节点， `PoolSize` 代表和每个节点的连接数量
    // 最终最大连接数为 PoolSize * 15节点数量
	PoolSize        int
	PoolTimeout     time.Duration
	MinIdleConns    int
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration

	TLSConfig *tls.Config

	// 集群配置项

	MaxRedirects   int
	ReadOnly       bool
	RouteByLatency bool
	RouteRandomly  bool

    // 哨兵 Master Name，仅适用于 `Failover Client`
	MasterName string
}
```