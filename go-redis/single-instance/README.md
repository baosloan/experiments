# 单节点客户端

## 安装

```shell
go get github.com/redis/go-redis/v9
```

## 连接redis服务器

连接到 Redis 服务器示例：
```go
import "github.com/redis/go-redis/v9"

rdb := redis.NewClient(&redis.Options{
	Addr:	  "localhost:6379",
	Password: "", // 没有密码，默认值
	DB:		  0,  // 默认DB 0
})
```
同时也支持另外一种常见的连接字符串:
```go
opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
if err != nil {
	panic(err)
}

rdb := redis.NewClient(opt)
```
redis.Options参数详解：
```go
type Options struct {
    // 连接网络类型，如: tcp、udp、unix等方式
    // 如果为空默认tcp
    Network string

    // redis服务器地址，ip:port格式，比如：192.168.1.100:6379
    // 默认为 :6379
    Addr string

    // ClientName 是对网络连接设置一个名字，使用 "CLIENT LIST" 命令
    // 可以查看redis服务器当前的网络连接列表
    // 如果设置了ClientName，go-redis对每个连接调用 `CLIENT SETNAME ClientName` 命令
    // 查看: https://redis.io/commands/client-setname/
    // 默认为空，不设置客户端名称
    ClientName string

    // 如果你想自定义连接网络的方式，可以自定义 `Dialer` 方法，
    // 如果不指定，将使用默认的方式进行网络连接 `redis.NewDialer`
    Dialer func(ctx context.Context, network, addr string) (net.Conn, error)

    // 建立了新连接时调用此函数
    // 默认为nil
    OnConnect func(ctx context.Context, cn *Conn) error

    // 当redis服务器版本在6.0以上时，作为ACL认证信息配合密码一起使用，
    // ACL是redis 6.0以上版本提供的认证功能，6.0以下版本仅支持密码认证。
    // 默认为空，不进行认证。
    Username string

    // 当redis服务器版本在6.0以上时，作为ACL认证信息配合密码一起使用，
    // 当redis服务器版本在6.0以下时，仅作为密码认证。
    // ACL是redis 6.0以上版本提供的认证功能，6.0以下版本仅支持密码认证。
    // 默认为空，不进行认证。
    Password string

    // 允许动态设置用户名和密码，go-redis在进行网络连接时会获取用户名和密码，
    // 这对一些认证鉴权有时效性的系统来说很有用，比如一些云服务商提供认证信息有效期为12小时。
    // 默认为nil
    CredentialsProvider func() (username string, password string)

    // redis DB 数据库，默认为0
    DB int

    // 命令最大重试次数， 默认为3
    MaxRetries int

    // 每次重试最小间隔时间
    // 默认 8 * time.Millisecond (8毫秒) ，设置-1为禁用
    MinRetryBackoff time.Duration

    // 每次重试最大间隔时间
    // 默认 512 * time.Millisecond (512毫秒) ，设置-1为禁用
    MaxRetryBackoff time.Duration

    // 建立新网络连接时的超时时间
    // 默认5秒
    DialTimeout time.Duration

    // 从网络连接中读取数据超时时间，可能的值：
    //  0 - 默认值，3秒
    // -1 - 无超时，无限期的阻塞
    // -2 - 不进行超时设置，不调用 SetReadDeadline 方法
    ReadTimeout time.Duration

    // 把数据写入网络连接的超时时间，可能的值：
    //  0 - 默认值，3秒
    // -1 - 无超时，无限期的阻塞
    // -2 - 不进行超时设置，不调用 SetWriteDeadline 方法
    WriteTimeout time.Duration

    // 是否使用context.Context的上下文截止时间，
    // 有些情况下，context.Context的超时可能带来问题。
    // 默认不使用
    ContextTimeoutEnabled bool

    // 连接池的类型，有 LIFO 和 FIFO 两种模式，
    // PoolFIFO 为 false 时使用 LIFO 模式，为 true 使用 FIFO 模式。
    // 当一个连接使用完毕时会把连接归还给连接池，连接池会把连接放入队尾，
    // LIFO 模式时，每次取空闲连接会从"队尾"取，就是刚放入队尾的空闲连接，
    // 也就是说 LIFO 每次使用的都是热连接，连接池有机会关闭"队头"的长期空闲连接，
    // 并且从概率上，刚放入的热连接健康状态会更好；
    // 而 FIFO 模式则相反，每次取空闲连接会从"队头"取，相比较于 LIFO 模式，
    // 会使整个连接池的连接使用更加平均，有点类似于负载均衡寻轮模式，会循环的使用
    // 连接池的所有连接，如果你使用 go-redis 当做代理让后端 redis 节点负载更平均的话，
    // FIFO 模式对你很有用。
    // 如果你不确定使用什么模式，请保持默认 PoolFIFO = false
    PoolFIFO bool

    // 连接池最大连接数量，注意：这里不包括 pub/sub，pub/sub 将使用独立的网络连接
    // 默认为 10 * runtime.GOMAXPROCS
    PoolSize int

    // PoolTimeout 代表如果连接池所有连接都在使用中，等待获取连接时间，超时将返回错误
    // 默认是 1秒+ReadTimeout
    PoolTimeout time.Duration

    // 连接池保持的最小空闲连接数，它受到PoolSize的限制
    // 默认为0，不保持
    MinIdleConns int

    // 连接池保持的最大空闲连接数，多余的空闲连接将被关闭
    // 默认为0，不限制
    MaxIdleConns int

    // ConnMaxIdleTime 是最大空闲时间，超过这个时间将被关闭。
    // 如果 ConnMaxIdleTime <= 0，则连接不会因为空闲而被关闭。
    // 默认值是30分钟，-1禁用
    ConnMaxIdleTime time.Duration

    // ConnMaxLifetime 是一个连接的生存时间，
    // 和 ConnMaxIdleTime 不同，ConnMaxLifetime 表示连接最大的存活时间
    // 如果 ConnMaxLifetime <= 0，则连接不会有使用时间限制
    // 默认值为0，代表连接没有时间限制
    ConnMaxLifetime time.Duration

    // 如果你的redis服务器需要TLS访问，可以在这里配置TLS证书等信息
    // 如果配置了证书信息，go-redis将使用TLS发起连接，
    // 如果你自定义了 `Dialer` 方法，你需要自己实现网络连接
    TLSConfig *tls.Config

    // 限流器的配置，参照 `Limiter` 接口
    Limiter Limiter

    // 设置启用在副本节点只读查询，默认为false不启用
    // 参照：https://redis.io/commands/readonly
    readOnly bool
}
```

### 使用TLS

你需要手动设置 tls.Config

```go
rdb := redis.NewClient(&redis.Options{
	TLSConfig: &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "you domain",
		//Certificates: []tls.Certificate{cert}
	},
})
```

如果你使用的是域名连接，且遇到了类似 x509: cannot validate certificate for xxx.xxx.xxx.xxx because it doesn't contain any IP SANs的错误，应该在 ServerName 中指定你的域名：#1710

```go
rdb := redis.NewClient(&redis.Options{
	TLSConfig: &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "你的域名",
	},
})
```

### SSH方式

使用SSH协议连接：
```go
sshConfig := &ssh.ClientConfig{
	User:			 "root",
	Auth:			 []ssh.AuthMethod{ssh.Password("password")},
	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	Timeout:		 15 * time.Second,
}

sshClient, err := ssh.Dial("tcp", "remoteIP:22", sshConfig)
if err != nil {
	panic(err)
}

rdb := redis.NewClient(&redis.Options{
	Addr: net.JoinHostPort("127.0.0.1", "6379"),
	Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
		return sshClient.Dial(network, addr)
	},
	// SSH不支持超时设置，在这里禁用
	ReadTimeout:  -1,
	WriteTimeout: -1,
})
```

### dial tcp: i/o timeout

当你遇到 dial tcp: i/o timeout 错误时，表示 go-redis 无法连接 Redis 服务器，比如 redis 服务器没有正常运行或监听了其他端口，以及可能被防火墙拦截等。你可以使用一些网络命令排查问题，例如 telnet:

```shell
telnet localhost 6379
Trying 127.0.0.1...
telnet: Unable to connect to remote host: Connection refused
```
如果你使用 Docker、Kubernetes、Istio、Service Mesh、Sidecar 方式运行，应该确保服务在容器完全可用后启动，你可以通过健康检查、Readiness Gate、Istio holdApplicationUntilProxyStarts等。

## Context上下文

go-redis 支持 Context，你可以使用它控制 超时 或者传递一些数据, 也可以 监控 go-redis 性能。

```go
ctx := context.Background()
```

## 执行Redis命令

执行Redis命令：

```go
val, err := rdb.Get(ctx, "key").Result()
fmt.Println(val)
```
你也可以分别访问值和错误：

```go
get := rdb.Get(ctx, "key")
fmt.Println(get.Val(), get.Err())
```

## 执行尚不支持的命令

可以使用 Do() 方法执行尚不支持或者任意命令:

```go
val, err := rdb.Do(ctx, "get", "key").Result()
if err != nil {
	if err == redis.Nil {
		fmt.Println("key does not exists")
		return
	}
	panic(err)
}
fmt.Println(val.(string))
```

Do() 方法返回 Cmd 类型，你可以使用它获取你想要的类型：

```go
// Text is a shortcut for get.Val().(string) with proper error handling.
val, err := rdb.Do(ctx, "get", "key").Text()
fmt.Println(val, err)
```

方法列表:

```go
s, err := cmd.Text()
flag, err := cmd.Bool()

num, err := cmd.Int()
num, err := cmd.Int64()
num, err := cmd.Uint64()
num, err := cmd.Float32()
num, err := cmd.Float64()

ss, err := cmd.StringSlice()
ns, err := cmd.Int64Slice()
ns, err := cmd.Uint64Slice()
fs, err := cmd.Float32Slice()
fs, err := cmd.Float64Slice()
bs, err := cmd.BoolSlice()
```

## redis.Nil

redis.Nil 是一种特殊的错误，严格意义上来说它并不是错误，而是代表一种状态，例如你使用 Get 命令获取 key 的值，当 key 不存在时，返回 redis.Nil。在其他比如 BLPOP 、 ZSCORE 也有类似的响应，你需要区分错误：

```go
val, err := rdb.Get(ctx, "key").Result()
switch {
case err == redis.Nil:
	fmt.Println("key不存在")
case err != nil:
	fmt.Println("错误", err)
case val == "":
	fmt.Println("值是空字符串")
}
```

## Conn

redis.Conn 是从连接池中取出的单个连接，除非你有特殊的需要，否则尽量不要使用它。你可以使用它向 redis 发送任何数据并读取 redis 的响应，当你使用完毕时，应该把它返回给 go-redis，否则连接池会永远丢失一个连接。

```go
cn := rdb.Conn(ctx)
defer cn.Close()

if err := cn.ClientSetName(ctx, "myclient").Err(); err != nil {
	panic(err)
}

name, err := cn.ClientGetName(ctx).Result()
if err != nil {
	panic(err)
}
fmt.Println("client name", name)
```