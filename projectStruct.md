###目录结构
addmgr 				- 地址管理
blockchain 			- 实现比特币区块处理和链选择规则
blockchain / fullblocktests 	- 提供一组用于测试共识验证规则的块测试
btcec 				- 实现对比特币脚本所需的椭圆曲线密码函数的支持
btcjson 			- 为底层的JSON-RPC命令和返回值提供一个扩展的API
chaincfg/chainhash 		- 提供通用的散列类型和相关函数，允许抽象特定的散列算法。
connmgr 			- 链接管理器
database 			- 为比特币区块链提供数据库接口,实现了ffldb
mempool 			- 未发掘的比特币交易池。
mining 				- pow挖矿
netsync 			- 同步管理器，同步区块链和交易
peer 				- 创建和管理比特币对等网络，及上行下行数据包的处理
rpcclient 			- 实现一个强大且易于使用的支持Websocket的比特币JSON-RPC客户端
txscript 			- 实现比特币交易脚本语言
wire 				- 实现比特币网络协议


###工程root目录下文件：
├── btcd.go  - btcd程序主入口，程序初始化，并启动btcd全节点
├── config.go  -负责从配置文件解释相关的参数
├── log.go  - 初始化相关日志管理器
├── params.go - 定义不同类型bitcoin网络参数(mainnet, 以及其他测试网络 )
├── rpcadapters.go - rpc节点连接，同步管理
├── rpcserver.go - rpc服务器实现
├── rpcserverhelp.go - rpc服务器相关帮助
├── rpcwebsocket.go - 实现rpc websocket协议
├── server.go - bitcoin服务器，主要处理跟bitcoin节点的通信
├── service_windows.go - window系统下，bitcion服务实现
├── signal.go - 监听系统中断信号
├── signalsigterm.go - 定义要监听信息(SIGINT:中断信号, SIGTERM:终止程序信号)
├── upgrade.go - 修改0.2.0之前版本db数据库路径
├── upnp.go - 实现NAT-PMP功能
├── version.go - 版本信息

###命令行指令


Application Options:
  -V, --version               打印版本号
  
  -C, --configfile=           指定配置文件
  
  -b, --datadir=              指定区块数据存储目录
  
      --logdir=               指定日志输出目录
      
  -a, --addpeer=              启动时，指定一个节点并连接
  
      --connect=              启动时，只连接指定的节点。
      
      --nolisten             禁止进入的连接，注意： 如果使用了--connect 或 --proxy选项，监听会自动停止，即便使用了--listen选项也会停止监听。 
                              
      --listen=              添加监听端口(默认是8333， 测试网络是：18333) 
      
      --maxpeers=             最大的进出连接数(默认125)
      
      --nobanning             禁止屏蔽恶意节点功能
      
      --banduration=          屏蔽恶意节点时长.  有效时间单位 {s, m, h}.  最小1秒 (默认：24h0m0s)
    
      --banthreshold=         触发断开连接，屏蔽恶意节点的最大ban值(100)
      
      --whitelist=            免屏蔽白名单(eg. 192.168.1.0/24 or ::1)
      
  -u, --rpcuser=              登录RPC服务器管理员用户名
  
  -P, --rpcpass=              登录RPC服务器管理员密码
  
      --rpclimituser=         登录RPC服务器普通用户名
      
      --rpclimitpass=         登录RPC服务器普通用户密码
      
      --rpclisten=            设置rpc服务器监听端口 (默认主网: 8334, 测试网: 18334)
      
      --rpccert=              rpc证书
      
      --rpckey=               rpc证书密钥
      
      --rpcmaxclients=        rpc服务器最大客服端连接个数 (10)
      
      --rpcmaxwebsockets=     rpc服务器最大websocket连接个数 (25)
      
      --rpcmaxconcurrentreqs=  最大rpc并发请求个数 (20)
      
      --rpcquirks            开启JSON-RPC quirk。注意： 不鼓励使用，除非为了解决一些互操作问题时，才开启。主要是兼容JSON-RPC 1.0 跟 JSON-RPC 2.0 中 ，对notification id为null处理。
                              
      --norpc                关闭内置RPC服务器。 注意：如果没有指定rpcuser/rpcpass或rpclimiteduser/rpclimitedpass时，RPC服务器是不会开启。
                              
      --notls                RPC服务器停止使用TLS。注意：这个控制选项只有在RPC服务器绑定到本机才有效。 
      
      --nodnsseed           不使用DNS获取其他节点信息。
      
      --externalip=         添加外部监听ip地址。如果运行btcd的主机在NAT后，那么这个地址就是外网的地址。 
      
      --proxy=               SOCKS5代理服务器地址 (eg. 127.0.0.1:9050)
      
      --proxyuser=           登录SOCKS5代理服务器的用户名
      
      --proxypass=           登录SOCKS5代理服务器用户的密码
      
      --onion=               用于连接到onion暗网服务的SOCKS5服务器地址 (eg. 127.0.0.1:9050)
      
      --onionuser=            登录oion SOCKS5服务器的用户名
      
      --onionpass=            登录oion SOCKS5服务器的用户的密码
      
      --noonion               关闭连接onion服务。
      
      --torisolation         为每个连接使用随机密码来隔离Tor流。
      
      --testnet              使用测试网络
      
      --regtest              使用回归测试网络
      
      --simnet               使用模拟测试网络
      
      --addcheckpoint=       添加检查点: '<块高度>:<块哈希>'，用在优化下载旧区块，避免从旧区块分叉
      
      --nocheckpoints       关闭内置的检查点。不要关闭除非你自己清楚这样做的目的。  
      
      --dbtype=             存储区块链所用的数据库类型 (默认：ffldb)
      
      --profile=            指定开启HTTP服务器性能检测的端口。 注意端口必须在1024 ~ 65536区间。 
      
      --cpuprofile=        指定服务器CPU性能检测输出文件
      
  -d, --debuglevel=       为所用子系统指定log等级{trace, debug, info, warn, error, critical} ，你也可以通过<子系统>=<等级>, <子系统>=<等级>,... 来为单个子系统设置log等级。 使用 show 来列出可用的子系统。(btcd -d show )
          
      --upnp              使用UPnP来发现路由设备实现NAT
      
      --minrelaytxfee=    最少交易费用按来算BTC/kB (默认： 1e-05) 
      
      --limitfreerelay=    免交易费的交易个数，每分钟多少k字节算(默认：15kb/分钟) 
                              
      --norelaypriority   对无手续费，低手续费交易进行中转时，不要求高优先级
      
      --maxorphantx=         内存中孤块中包含最大的交易数 (100)
      
      --generate             使用cpu挖矿
      
      --miningaddr=          添加接受区块奖励地址，如果开启generate选项，则必须至少指定一个地址。 
                              
      --blockminsize=        创建区块时，最小区块大小。 
      
      --blockmaxsize=       创建区块时，最大区块大小(默认750000Bytes) 
      
      --blockminweight=     创建区块时，区块最小权重
      
      --blockmaxweight=    创建区块时，区块最大权重(默认：3000000)   
      
      --blockprioritysize=    创建区块时，对于高级优先级/低手续费交易，区块大小（默认： 50000Bytes）
      
      --uacomment=          在用户代理上留言 - 具体请查看BIP 14.
      
      --nopeerbloomfilters    关闭bloom滤波器
      
      --nocfilters            关闭committed过滤(CF)
      
      --dropcfindex         在程序启动时，删掉数据库中用于committed过滤(CF)支持的索引，然后程序退出。  
                              
      --sigcachemaxsize=      签名证书缓存条数最大值（默认100000）   
      --blocksonly            不从远程节点接受交易
      
      --txindex               维护一个完整hash-based交易索引。 这样可以通过getrawtransaction RPC让交易数据可用。
                              
      --droptxindex          在程序启动时，删掉来自数据库中的hash-based 交易索引，然后退出。 
      
      --addrindex             维护一个address-based交易索引，这样可以让searchrawtransactions RPC可用。 
      
      --dropaddrindex        在程序启动后，删掉来自数据库中的address-based交易索引，然后程序退出。 
      
      --relaynonstd          中转非标准交易，忽略当前bitcoin网络默认设置。 
      
      --rejectnonstd         拒绝非标准交易，忽略当前bitcoin网络默认设置。


