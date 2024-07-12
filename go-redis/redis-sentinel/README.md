# Go Redis Sentinel客户端

## 服务器客户端

连接到 哨兵模式 管理的服务器：

```go
rdb := redis.NewFailoverClient(&redis.FailoverOptions{
    MasterName:    "master-name",
    SentinelAddrs: []string{":9126", ":9127", ":9128"},
})
```
更多配置项，请参照 redis.FailoverOptions:
部分配置项继承自 Options，请查看 Options 说明。

```go
type FailoverOptions struct {
	// sentinel master节点名称
	MasterName string

	// 哨兵节点地址列表
	// 示例:[]string{"192.168.1.10:6379", "192.168.1.11:6379"}
	SentinelAddrs []string

	// ClientName 和 `Options` 相同，会对每个Node节点的每个网络连接配置
	ClientName string

	// 用于ACL认证的用户名
	SentinelUsername string

	// Sentinel中 `requirepass<password>` 的密码配置
	// 如果同时提供了 `SentinelUsername` ，则启用ACL认证
	SentinelPassword string

	// 把只读命令发送到响应最快的节点，
	// 仅限于 `Failover Cluster Client`
	RouteByLatency bool

    // 把只读命令随机到一个节点
	// 仅限于 `Failover Cluster Client`
    RouteRandomly bool

	// 把所有命令发送到发送到只读节点
	ReplicaOnly bool

	// 当所有副本节点都无法连接时，尝试使用与Sentinel已断开连接的副本
	UseDisconnectedReplicas bool

	// 下面的配置项，和 `Options` 基本一致，请参照 `Options` 的说明

	Dialer    func(ctx context.Context, network, addr string) (net.Conn, error)
	OnConnect func(ctx context.Context, cn *Conn) error

	Username string
	Password string
	DB       int

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
}
```

从 go-redis v8 版本开始，你可以尝试使用 NewFailoverClusterClient 把只读命令路由到从节点，请注意， NewFailoverClusterClient 借助了 Cluster Client 实现，不支持 DB 选项（只能操作 DB 0）：

```go
import "github.com/redis/go-redis/v9"

rdb := redis.NewFailoverClusterClient(&redis.FailoverOptions{
    MasterName:    "master-name",
    SentinelAddrs: []string{":9126", ":9127", ":9128"},

    // 你可以选择把只读命令路由到最近的节点，或者随机节点，二选一
    // RouteByLatency: true,
    // RouteRandomly: true,
})
```

## 哨兵服务器客户端

请注意，哨兵客户端本身用于连接哨兵服务器，你可以从哨兵上获取管理的 redis 服务器信息：

```go
import "github.com/redis/go-redis/v9"

sentinel := redis.NewSentinelClient(&redis.Options{
    Addr: ":9126",
})

addr, err := sentinel.GetMasterAddrByName(ctx, "master-name").Result()
```