# Go Redis Sentinel客户端

## 服务器客户端

连接到 哨兵模式 管理的服务器：

```go
rdb := redis.NewFailoverClient(&redis.FailoverOptions{
    MasterName:    "master-name",
    SentinelAddrs: []string{":9126", ":9127", ":9128"},
})
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