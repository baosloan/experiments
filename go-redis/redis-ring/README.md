# Go Redis Ring客户端

## 介绍

Ring 分片客户端，是采用了一致性 HASH 算法在多个 redis 服务器之间分发 key，每个节点承担一部分 key 的存储。

Ring 客户端会监控每个节点的健康状况，并从 Ring 中移除掉宕机的节点，当节点恢复时，会再加入到 Ring 中。这样实现了可用性和容错性，但节点和节点之间没有一致性，仅仅是通过多个节点分摊流量的方式来处理更多的请求。如果你更注重一致性、分区、安全性，请使用 Redis Cluster。

go-redis 默认使用 Rendezvous Hash 算法，你也可以通过设置 RingOptions.NewConsistentHash 自定义一致性 HASH 算法，更多 Ring 客户端设置请参照 redis.RingOptions。

## 开始使用

创建一个由三个节点组成的 Ring 客户端，更多设置请参照 redis.RingOptions:

```go
import "github.com/redis/go-redis/v9"

rdb := redis.NewRing(&redis.RingOptions{
    Addrs: map[string]string{
        // shardName => host:port
        "shard1": "localhost:7000",
        "shard2": "localhost:7001",
        "shard3": "localhost:7002",
    },
})
```

更多设置请参照 redis.RingOptions:
部分配置项继承自 Options，请查看 Options 说明。

```go
// RingOptions are used to configure a ring client and should be
// passed to NewRing.
type RingOptions struct {
	// redis服务器地址
	// 示例："one" => "192.168.1.10:6379", "two" => "192.168.1.11:6379"
	Addrs map[string]string

	// New集群节点 `*redis.Client` 的对象，
	// go-redis 默认使用 `redis.NewClient(opt)` 方法
	NewClient func(opt *Options) *Client

    // ClientName 和 `Options` 相同，会对每个Node节点的每个网络连接配置
	ClientName string

	// 节点健康检查的时间间隔，默认500毫秒
	// 如果连续3次检查失败，认为节点宕机
	HeartbeatFrequency time.Duration

	// 设置自定义的一致性hash算法，ring会在多个节点之间通过hash算法分布key
	// 参考: https://medium.com/@dgryski/consistent-hashing-algorithmic-tradeoffs-ef6b8e2fcae8
	NewConsistentHash func(shards []string) ConsistentHash

    // 下面的配置项，和 `Options` 基本一致，请参照 `Options` 的说明

	Dialer    func(ctx context.Context, network, addr string) (net.Conn, error)
	OnConnect func(ctx context.Context, cn *Conn) error

	Username string
	Password string
	DB       int

	MaxRetries      int
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration

	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PoolFIFO bool

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
	Limiter   Limiter
}
```

你可以像其他客户端一样执行命令：

```go
if err := rdb.Set(ctx, "foo", "bar", 0).Err(); err != nil {
    panic(err)
}
```

## 遍历每个节点：

```go
err := rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
    return shard.Ping(ctx).Err()
})
if err != nil {
    panic(err)
}
```

## 节点选项配置

你可以手动设置连接节点，例如设置用户名和密码：

```go
rdb := redis.NewRing(&redis.RingOptions{
    NewClient: func(opt *redis.Options) *redis.NewClient {
        user, pass := userPassForAddr(opt.Addr)
        opt.Username = user
        opt.Password = pass

        return redis.NewClient(opt)
    },
})
```

## 自定义Hash算法

go-redis 默认使用 Rendezvous Hash 算法将 Key 分布到多个节点上，你可以更改为其他 Hash 算法：

```go
import "github.com/golang/groupcache/consistenthash"

ring := redis.NewRing(&redis.RingOptions{
    NewConsistentHash: func() {
        return consistenthash.New(100, crc32.ChecksumIEEE)
    },
})
```
